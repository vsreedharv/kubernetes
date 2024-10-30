/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package factory

import (
	"context"
	"errors"
	"fmt"
	"net"

	pb "go.etcd.io/etcd/api/v3/etcdserverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"k8s.io/apiserver/pkg/apis/example"

	"testing"
	"time"

	"k8s.io/apimachinery/pkg/api/apitesting"
	"k8s.io/apimachinery/pkg/runtime/schema"
	examplev1 "k8s.io/apiserver/pkg/apis/example/v1"
	"k8s.io/apiserver/pkg/storage/storagebackend"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func Test_atomicLastError(t *testing.T) {
	aError := &atomicLastError{err: fmt.Errorf("initial error")}
	// no timestamp is always updated
	aError.Store(errors.New("updated error"), time.Time{})
	err := aError.Load()
	if err.Error() != "updated error" {
		t.Fatalf("Expected: \"updated error\" got: %s", err.Error())
	}
	// update to current time
	now := time.Now()
	aError.Store(errors.New("now error"), now)
	err = aError.Load()
	if err.Error() != "now error" {
		t.Fatalf("Expected: \"now error\" got: %s", err.Error())
	}
	// no update to past time
	past := now.Add(-5 * time.Second)
	aError.Store(errors.New("past error"), past)
	err = aError.Load()
	if err.Error() != "now error" {
		t.Fatalf("Expected: \"now error\" got: %s", err.Error())
	}
}

func TestClientWithGrpcHealthcheck(t *testing.T) {

	etcdMock1 := startEtcdMockServer(t)
	defer etcdMock1.Stop(t)
	etcdMock2 := startEtcdMockServer(t)
	defer etcdMock2.Stop(t)

	cfg := storagebackend.Config{
		Type: storagebackend.StorageTypeETCD3,
		Transport: storagebackend.TransportConfig{
			ServerList:            []string{etcdMock1.Listener.Addr().String(), etcdMock2.Listener.Addr().String()},
			EnableGrpcHealthcheck: true,
		},
		Codec: apitesting.TestCodec(codecs, examplev1.SchemeGroupVersion),
	}

	storage, destroyFunc, err := newETCD3Storage(*cfg.ForResource(schema.GroupResource{Resource: "pods"}), nil, nil, "")
	defer destroyFunc()
	if err != nil {
		t.Fatal(err)
	}

	maxCreates := 10
	for range maxCreates {
		err = storage.Create(context.TODO(), "/abc", &example.Pod{}, nil, 0)
		if err != nil {
			t.Fatalf("Create failed: %v", err)
		}
	}

	if (etcdMock1.MockKVServer.TxnCounter != maxCreates/2) && (etcdMock2.MockKVServer.TxnCounter != maxCreates/2) {
		t.Fatalf("Etcd client round robin balancer should balance equally. Got %d vs %d",
			etcdMock1.MockKVServer.TxnCounter, etcdMock2.MockKVServer.TxnCounter)
	}

	etcdMock1.HealthServer.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)

	for range maxCreates {
		err = storage.Create(context.TODO(), "/abc", &example.Pod{}, nil, 0)
		if err != nil {
			t.Fatalf("Create failed: %v", err)
		}
	}

	if etcdMock1.MockKVServer.TxnCounter > maxCreates/2+1 {
		t.Fatalf("Etcd client grpc healthcheck isn't working. Max 1 extra request is expected. Got %d vs %d",
			etcdMock1.MockKVServer.TxnCounter, etcdMock2.MockKVServer.TxnCounter)
	}

	if etcdMock2.MockKVServer.TxnCounter < maxCreates {
		t.Fatalf("Etcd client grpc healthcheck isn't working. Most of request should go to second server. Got %d vs %d",
			etcdMock1.MockKVServer.TxnCounter, etcdMock2.MockKVServer.TxnCounter)
	}
}

type mockEtcdServer struct {
	Listener     net.Listener
	GrpcServer   *grpc.Server
	MockKVServer *mockKVServer
	HealthServer *health.Server
}

func startEtcdMockServer(t *testing.T) *mockEtcdServer {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	svr := grpc.NewServer()
	kv := &mockKVServer{}
	hs := health.NewServer()

	pb.RegisterKVServer(svr, kv)
	healthpb.RegisterHealthServer(svr, hs)

	go func() {
		_ = svr.Serve(lis)
	}()

	return &mockEtcdServer{
		Listener:     lis,
		GrpcServer:   svr,
		MockKVServer: kv,
		HealthServer: hs,
	}
}

func (m *mockEtcdServer) Stop(t *testing.T) {
	err := m.Listener.Close()
	if err != nil {
		t.Fatal(err)
	}
	m.GrpcServer.Stop()
}

type mockKVServer struct {
	TxnCounter int
}

func (m *mockKVServer) Range(context.Context, *pb.RangeRequest) (*pb.RangeResponse, error) {
	return &pb.RangeResponse{}, nil
}

func (m *mockKVServer) Put(context.Context, *pb.PutRequest) (*pb.PutResponse, error) {
	return &pb.PutResponse{}, nil
}

func (m *mockKVServer) DeleteRange(context.Context, *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	return &pb.DeleteRangeResponse{}, nil
}

func (m *mockKVServer) Txn(context.Context, *pb.TxnRequest) (*pb.TxnResponse, error) {
	m.TxnCounter += 1
	// minimal response to make create call succeed
	return &pb.TxnResponse{Succeeded: true, Header: &pb.ResponseHeader{Revision: 0}}, nil
}

func (m *mockKVServer) Compact(context.Context, *pb.CompactionRequest) (*pb.CompactionResponse, error) {
	return &pb.CompactionResponse{}, nil
}
