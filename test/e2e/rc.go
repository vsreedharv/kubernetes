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

package e2e

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/controller/replication"
	"k8s.io/kubernetes/test/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = framework.KubeDescribe("ReplicationController", func() {
	f := framework.NewDefaultFramework("replication-controller")

	It("should serve a basic image on each replica with a public image [Conformance]", func() {
		ServeImageOrFail(f, "basic", "gcr.io/google_containers/serve_hostname:v1.4")
	})

	It("should serve a basic image on each replica with a private image", func() {
		// requires private images
		framework.SkipUnlessProviderIs("gce", "gke")

		ServeImageOrFail(f, "private", "gcr.io/k8s-authenticated-test/serve_hostname:v1.4")
	})

	It("should surface a failure condition on a common issue like exceeded quota", func() {
		rcConditionCheck(f)
	})

	It("should adopt matching pods on creation", func() {
		testRCAdoptMatchingOrphans(f)
	})

	It("should release no longer matching pods", func() {
		testRCReleaseControlledNotMatching(f)
	})
})

func newRC(rsName string, replicas int32, rcPodLabels map[string]string, imageName string, image string) *v1.ReplicationController {
	zero := int64(0)
	return &v1.ReplicationController{
		ObjectMeta: metav1.ObjectMeta{
			Name: rsName,
		},
		Spec: v1.ReplicationControllerSpec{
			Replicas: func(i int32) *int32 { return &i }(replicas),
			Template: &v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: rcPodLabels,
				},
				Spec: v1.PodSpec{
					TerminationGracePeriodSeconds: &zero,
					Containers: []v1.Container{
						{
							Name:  imageName,
							Image: image,
						},
					},
				},
			},
		},
	}
}

func newRCWithSelector(rcName string, replicas int32, rcPodLabels map[string]string, imageName string, image string, nameLabelSelector string) *v1.ReplicationController {
	rcObj := newRC(rcName, replicas, rcPodLabels, imageName, image)

	rcObj.Spec.Selector = map[string]string{"name": nameLabelSelector}

	return rcObj
}

// A basic test to check the deployment of an image using
// a replication controller. The image serves its hostname
// which is checked for each replica.
func ServeImageOrFail(f *framework.Framework, test string, image string) {
	name := "my-hostname-" + test + "-" + string(uuid.NewUUID())
	replicas := int32(2)

	// Create a replication controller for a service
	// that serves its hostname.
	// The source for the Docker containter kubernetes/serve_hostname is
	// in contrib/for-demos/serve_hostname
	By(fmt.Sprintf("Creating replication controller %s", name))
	controller, err := f.ClientSet.Core().ReplicationControllers(f.Namespace.Name).Create(&v1.ReplicationController{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.ReplicationControllerSpec{
			Replicas: func(i int32) *int32 { return &i }(replicas),
			Selector: map[string]string{
				"name": name,
			},
			Template: &v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"name": name},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  name,
							Image: image,
							Ports: []v1.ContainerPort{{ContainerPort: 9376}},
						},
					},
				},
			},
		},
	})
	Expect(err).NotTo(HaveOccurred())
	// Cleanup the replication controller when we are done.
	defer func() {
		// Resize the replication controller to zero to get rid of pods.
		if err := framework.DeleteRCAndPods(f.ClientSet, f.InternalClientset, f.Namespace.Name, controller.Name); err != nil {
			framework.Logf("Failed to cleanup replication controller %v: %v.", controller.Name, err)
		}
	}()

	// List the pods, making sure we observe all the replicas.
	label := labels.SelectorFromSet(labels.Set(map[string]string{"name": name}))

	pods, err := framework.PodsCreated(f.ClientSet, f.Namespace.Name, name, replicas)

	By("Ensuring each pod is running")

	// Wait for the pods to enter the running state. Waiting loops until the pods
	// are running so non-running pods cause a timeout for this test.
	for _, pod := range pods.Items {
		if pod.DeletionTimestamp != nil {
			continue
		}
		err = f.WaitForPodRunning(pod.Name)
		Expect(err).NotTo(HaveOccurred())
	}

	// Verify that something is listening.
	By("Trying to dial each unique pod")
	retryTimeout := 2 * time.Minute
	retryInterval := 5 * time.Second
	err = wait.Poll(retryInterval, retryTimeout, framework.PodProxyResponseChecker(f.ClientSet, f.Namespace.Name, label, name, true, pods).CheckAllResponses)
	if err != nil {
		framework.Failf("Did not get expected responses within the timeout period of %.2f seconds.", retryTimeout.Seconds())
	}
}

// 1. Create a quota restricting pods in the current namespace to 2.
// 2. Create a replication controller that wants to run 3 pods.
// 3. Check replication controller conditions for a ReplicaFailure condition.
// 4. Relax quota or scale down the controller and observe the condition is gone.
func rcConditionCheck(f *framework.Framework) {
	c := f.ClientSet
	namespace := f.Namespace.Name
	name := "condition-test"

	By(fmt.Sprintf("Creating quota %q that allows only two pods to run in the current namespace", name))
	quota := newPodQuota(name, "2")
	_, err := c.Core().ResourceQuotas(namespace).Create(quota)
	Expect(err).NotTo(HaveOccurred())

	err = wait.PollImmediate(1*time.Second, 1*time.Minute, func() (bool, error) {
		quota, err = c.Core().ResourceQuotas(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		podQuota := quota.Status.Hard[v1.ResourcePods]
		quantity := resource.MustParse("2")
		return (&podQuota).Cmp(quantity) == 0, nil
	})
	if err == wait.ErrWaitTimeout {
		err = fmt.Errorf("resource quota %q never synced", name)
	}
	Expect(err).NotTo(HaveOccurred())

	By(fmt.Sprintf("Creating rc %q that asks for more than the allowed pod quota", name))
	rc := newRC(name, 3, map[string]string{"name": name}, nginxImageName, nginxImage)
	rc, err = c.Core().ReplicationControllers(namespace).Create(rc)
	Expect(err).NotTo(HaveOccurred())

	By(fmt.Sprintf("Checking rc %q has the desired failure condition set", name))
	generation := rc.Generation
	conditions := rc.Status.Conditions
	err = wait.PollImmediate(1*time.Second, 1*time.Minute, func() (bool, error) {
		rc, err = c.Core().ReplicationControllers(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		if generation > rc.Status.ObservedGeneration {
			return false, nil
		}
		conditions = rc.Status.Conditions

		cond := replication.GetCondition(rc.Status, v1.ReplicationControllerReplicaFailure)
		return cond != nil, nil
	})
	if err == wait.ErrWaitTimeout {
		err = fmt.Errorf("rc manager never added the failure condition for rc %q: %#v", name, conditions)
	}
	Expect(err).NotTo(HaveOccurred())

	By(fmt.Sprintf("Scaling down rc %q to satisfy pod quota", name))
	rc, err = framework.UpdateReplicationControllerWithRetries(c, namespace, name, func(update *v1.ReplicationController) {
		x := int32(2)
		update.Spec.Replicas = &x
	})
	Expect(err).NotTo(HaveOccurred())

	By(fmt.Sprintf("Checking rc %q has no failure condition set", name))
	generation = rc.Generation
	conditions = rc.Status.Conditions
	err = wait.PollImmediate(1*time.Second, 1*time.Minute, func() (bool, error) {
		rc, err = c.Core().ReplicationControllers(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		if generation > rc.Status.ObservedGeneration {
			return false, nil
		}
		conditions = rc.Status.Conditions

		cond := replication.GetCondition(rc.Status, v1.ReplicationControllerReplicaFailure)
		return cond == nil, nil
	})
	if err == wait.ErrWaitTimeout {
		err = fmt.Errorf("rc manager never removed the failure condition for rc %q: %#v", name, conditions)
	}
	Expect(err).NotTo(HaveOccurred())
}

func testRCAdoptMatchingOrphans(f *framework.Framework) {
	name := "pod-adoption"
	By(fmt.Sprintf("Given a Pod with a 'name' label %s is created", name))
	p := f.PodClient().CreateSync(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"name": name,
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  name,
					Image: nginxImageName,
				},
			},
		},
	})

	By("When a replication controller with a matching selector is created")
	replicas := int32(2)
	rcSt := newRCWithSelector(name, replicas, map[string]string{"name": name}, name, nginxImageName, name)
	rc, err := f.ClientSet.Core().ReplicationControllers(f.Namespace.Name).Create(rcSt)
	Expect(err).NotTo(HaveOccurred())
	// Cleanup the ReplicationController when we are done.
	defer func() {
		if err := framework.DeleteRCAndPods(f.ClientSet, f.InternalClientset, f.Namespace.Name, rc.Name); err != nil {
			framework.Logf("Failed to cleanup ReplicationController %v: %v.", rc.Name, err)
		}
	}()

	By("Then the orphan pod is adopted")
	err = wait.PollImmediate(1*time.Second, 1*time.Minute, func() (bool, error) {
		p2, err := f.ClientSet.Core().Pods(f.Namespace.Name).Get(p.Name, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		for _, owner := range p2.OwnerReferences {
			if *owner.Controller && owner.UID == rc.UID {
				// pod adopted
				return true, nil
			}
		}
		// pod still not adopted
		return false, nil
	})
	Expect(err).NotTo(HaveOccurred())
}

func testRCReleaseControlledNotMatching(f *framework.Framework) {
	name := "pod-release"
	By("Given a ReplicationController is created")
	replicas := int32(2)
	rcSt := newRCWithSelector(name, replicas, map[string]string{"name": name}, name, nginxImageName, name)
	rc, err := f.ClientSet.Core().ReplicationControllers(f.Namespace.Name).Create(rcSt)
	Expect(err).NotTo(HaveOccurred())
	// Cleanup the rc when we are done.
	defer func() {
		if err := framework.DeleteRCAndPods(f.ClientSet, f.InternalClientset, f.Namespace.Name, rc.Name); err != nil {
			framework.Logf("Failed to cleanup ReplicationController %v: %v.", rc.Name, err)
		}
	}()

	By("When the matched label of one of its pods change")
	pods, err := framework.PodsCreated(f.ClientSet, f.Namespace.Name, rc.Name, replicas)
	Expect(err).NotTo(HaveOccurred())

	p := pods.Items[0]
	podClient := f.ClientSet.Core().Pods(f.Namespace.Name)
	patch := []byte("{\"metadata\":{\"labels\":{\"name\":\"not-matching-name\"}}}")
	_, err = podClient.Patch(p.Name, types.StrategicMergePatchType, patch)
	Expect(err).NotTo(HaveOccurred())

	By("Then the pod is released")
	err = wait.PollImmediate(1*time.Second, 1*time.Minute, func() (bool, error) {
		p2, err := podClient.Get(p.Name, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		for _, owner := range p2.OwnerReferences {
			if *owner.Controller && owner.UID == rc.UID {
				// pod still belonging to the replication controller
				return false, nil
			}
		}
		// pod already released
		return true, nil
	})
	Expect(err).NotTo(HaveOccurred())
}
