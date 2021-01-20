/*
Copyright 2021 The Kubernetes Authors.

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

package scheduler

import (
	"fmt"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/kubernetes/pkg/scheduler"
	schedapi "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/serviceaffinity"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
	testutils "k8s.io/kubernetes/test/integration/util"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

func TestServiceAffinityEnqueue(t *testing.T) {
	profile := schedapi.KubeSchedulerProfile{
		SchedulerName: v1.DefaultSchedulerName,
		Plugins: &schedapi.Plugins{
			PreFilter: &schedapi.PluginSet{
				Enabled: []schedapi.Plugin{
					{Name: serviceaffinity.Name},
				},
			},
			Filter: &schedapi.PluginSet{
				Enabled: []schedapi.Plugin{
					{Name: serviceaffinity.Name},
				},
			},
		},
		PluginConfig: []schedapi.PluginConfig{
			{
				Name: serviceaffinity.Name,
				Args: &schedapi.ServiceAffinityArgs{
					AffinityLabels: []string{"hostname"},
				},
			},
		},
	}
	// Use zero backoff seconds to bypass backoffQ.
	testCtx := testutils.InitTestSchedulerWithOptions(
		t,
		testutils.InitTestMaster(t, "serviceaffinity-enqueue", nil),
		nil,
		scheduler.WithProfiles(profile),
		scheduler.WithPodInitialBackoffSeconds(0),
		scheduler.WithPodMaxBackoffSeconds(0),
	)
	testutils.SyncInformerFactory(testCtx)
	// It's intended to not start the scheduler's queue, and hence to
	// keep flushing logic away. We will pop and schedule the Pods manually later.
	defer testutils.CleanupTest(t, testCtx)

	cs, ns, ctx := testCtx.ClientSet, testCtx.NS.Name, testCtx.Ctx
	// Create two Nodes.
	for i := 1; i <= 2; i++ {
		nodeName := fmt.Sprintf("node%d", i)
		capacity := map[v1.ResourceName]string{v1.ResourcePods: "1"}
		node := st.MakeNode().Name(nodeName).Label("hostname", nodeName).Capacity(capacity).Obj()
		if _, err := cs.CoreV1().Nodes().Create(ctx, node, metav1.CreateOptions{}); err != nil {
			t.Fatalf("Failed to create Node %q: %v", nodeName, err)
		}
	}

	// Create a Service.
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "svc",
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{{
				Port: int32(80),
			}},
			Selector: map[string]string{
				"foo": "bar",
			},
		},
	}
	if _, err := cs.CoreV1().Services(ns).Create(ctx, svc, metav1.CreateOptions{}); err != nil {
		t.Fatalf("Failed to create Service %q: %v", svc.Name, err)
	}

	// Create two Pods.
	pause := imageutils.GetPauseImageName()
	for i := 1; i <= 2; i++ {
		podName := fmt.Sprintf("pod%d", i)
		pod := st.MakePod().Namespace(ns).Name(podName).Label("foo", "bar").Container(pause).Obj()
		// Make Pod1 an existing Pod.
		if i == 1 {
			pod.Spec.NodeName = fmt.Sprintf("node%d", i)
		}
		if _, err := cs.CoreV1().Pods(ns).Create(ctx, pod, metav1.CreateOptions{}); err != nil {
			t.Fatalf("Failed to create Pod %q: %v", pod.Name, err)
		}
	}

	// Wait for pod2 to be present in the scheduling queue.
	if err := wait.Poll(time.Millisecond*200, wait.ForeverTestTimeout, func() (bool, error) {
		return len(testCtx.Scheduler.SchedulingQueue.PendingPods()) == 1, nil
	}); err != nil {
		t.Fatal(err)
	}

	// Pop one Pod. It should be schedulable.
	podInfo := testCtx.Scheduler.NextPod()
	fwk, ok := testCtx.Scheduler.Profiles[podInfo.Pod.Spec.SchedulerName]
	if !ok {
		t.Fatalf("Cannot find the profile for Pod %v", podInfo.Pod.Name)
	}
	// Schedule the Pod manually.
	_, fitError := testCtx.Scheduler.Algorithm.Schedule(ctx, fwk, framework.NewCycleState(), podInfo.Pod)
	// The fitError is expected to be:
	// 0/2 nodes are available: 1 Too many pods, 1 node(s) didn't match service affinity
	if fitError == nil {
		t.Fatalf("Expect Pod %v to fail at scheduling.", podInfo.Pod.Name)
	}
	testCtx.Scheduler.Error(podInfo, fitError)

	// Scheduling cycle is incremented from 0 to 1 after NextPod() is called, so
	// pass a number larger than 1 to move Pod to unschedulableQ.
	testCtx.Scheduler.SchedulingQueue.AddUnschedulableIfNotPresent(podInfo, 10)

	// Spawn a Service event.
	// We expect this event to trigger moving the test Pod from unschedulableQ to activeQ.
	if err := cs.CoreV1().Services(ns).Delete(ctx, "svc", metav1.DeleteOptions{}); err != nil {
		t.Fatalf("Failed to delete service 'svc': %v", err)
	}

	// Now we should be able to pop the Pod from activeQ again.
	podInfo = testCtx.Scheduler.NextPod()
	if podInfo.Attempts != 2 {
		t.Errorf("Expected the Pod to be attempted 2 times, but got %v", podInfo.Attempts)
	}
}
