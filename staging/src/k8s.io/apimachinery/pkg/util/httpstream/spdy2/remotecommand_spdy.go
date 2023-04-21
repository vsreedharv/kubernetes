/*
Copyright 2015 The Kubernetes Authors.

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

package spdy2

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"k8s.io/apimachinery/pkg/util/httpstream"
	"k8s.io/apimachinery/pkg/util/remotecommand"
	"k8s.io/klog/v2"
)

// spdyStreamExecutor handles transporting standard shell streams over an httpstream connection.
type spdyStreamExecutor struct {
	upgrader  Upgrader
	transport http.RoundTripper

	method    string
	url       *url.URL
	protocols []string
}

// NewSPDYExecutorForTransports connects to the provided server using the given transport,
// upgrades the response using the given upgrader to multiplexed bidirectional streams.
func NewSPDYExecutorForTransports(transport http.RoundTripper, upgrader Upgrader, method string, url *url.URL) (Executor, error) {
	return NewSPDYExecutorForProtocols(
		transport, upgrader, method, url,
		remotecommand.StreamProtocolV4Name,
		remotecommand.StreamProtocolV3Name,
		remotecommand.StreamProtocolV2Name,
		remotecommand.StreamProtocolV1Name,
	)
}

// NewSPDYExecutorForProtocols connects to the provided server and upgrades the connection to
// multiplexed bidirectional streams using only the provided protocols. Exposed for testing, most
// callers should use NewSPDYExecutor or NewSPDYExecutorForTransports.
func NewSPDYExecutorForProtocols(transport http.RoundTripper, upgrader Upgrader, method string, url *url.URL, protocols ...string) (Executor, error) {
	return &spdyStreamExecutor{
		upgrader:  upgrader,
		transport: transport,
		method:    method,
		url:       url,
		protocols: protocols,
	}, nil
}

// Stream opens a protocol streamer to the server and streams until a client closes
// the connection or the server disconnects.
func (e *spdyStreamExecutor) Stream(options StreamOptions) error {
	return e.StreamWithContext(context.Background(), options)
}

// newConnectionAndStream creates a new SPDY connection and a stream protocol handler upon it.
func (e *spdyStreamExecutor) newConnectionAndStream(ctx context.Context, options StreamOptions) (httpstream.Connection, streamProtocolHandler, error) {
	req, err := http.NewRequestWithContext(ctx, e.method, e.url.String(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating request: %v", err)
	}

	conn, protocol, err := Negotiate(
		e.upgrader,
		&http.Client{Transport: e.transport},
		req,
		e.protocols...,
	)
	if err != nil {
		return nil, nil, err
	}

	var streamer streamProtocolHandler

	switch protocol {
	case remotecommand.StreamProtocolV4Name:
		streamer = newStreamProtocolV4(options)
	case remotecommand.StreamProtocolV3Name:
		streamer = newStreamProtocolV3(options)
	case remotecommand.StreamProtocolV2Name:
		streamer = newStreamProtocolV2(options)
	case "":
		klog.V(4).Infof("The server did not negotiate a streaming protocol version. Falling back to %s", remotecommand.StreamProtocolV1Name)
		fallthrough
	case remotecommand.StreamProtocolV1Name:
		streamer = newStreamProtocolV1(options)
	}

	return conn, streamer, nil
}

// StreamWithContext opens a protocol streamer to the server and streams until a client closes
// the connection or the server disconnects or the context is done.
func (e *spdyStreamExecutor) StreamWithContext(ctx context.Context, options StreamOptions) error {
	conn, streamer, err := e.newConnectionAndStream(ctx, options)
	if err != nil {
		return err
	}
	defer conn.Close()

	panicChan := make(chan any, 1)
	errorChan := make(chan error, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		errorChan <- streamer.stream(spdyStreamCreator{Connection: conn})
	}()

	select {
	case p := <-panicChan:
		panic(p)
	case err := <-errorChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

type spdyStreamCreator struct {
	httpstream.Connection
}

func (spdyStreamCreator) Run() {
	// nothing to do
}
