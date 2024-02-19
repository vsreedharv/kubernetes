/*
Copyright 2020 The Kubernetes Authors.

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

package topologymanager

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/kubelet/cm/admission"
	"k8s.io/kubernetes/pkg/kubelet/cm/containermap"
	"k8s.io/kubernetes/pkg/kubelet/events"
	"k8s.io/kubernetes/pkg/kubelet/lifecycle"
)

const (
	// containerTopologyScope specifies the TopologyManagerScope per container.
	containerTopologyScope = "container"
	// podTopologyScope specifies the TopologyManagerScope per pod.
	podTopologyScope = "pod"
	// noneTopologyScope specifies the TopologyManagerScope when topologyPolicyName is none.
	noneTopologyScope = "none"
)

type podTopologyHints map[string]map[string]TopologyHint

// Scope interface for Topology Manager
type Scope interface {
	Name() string
	GetPolicy() Policy
	Admit(pod *v1.Pod) lifecycle.PodAdmitResult
	// RegisterProvider adds a hint provider to manager to indicate the hint provider
	// wants to be consulted with when making topology hints, and is authoritative about
	// the current resource allocation.
	RegisterProvider(ra ResourceAllocator)
	// AddContainer adds pod to Manager for tracking
	AddContainer(pod *v1.Pod, container *v1.Container, containerID string)
	// RemoveContainer removes pod from Manager tracking
	RemoveContainer(containerID string) error
	// Store is the interface for storing pod topology hints
	Store
}

type scope struct {
	recorder record.EventRecorder
	mutex    sync.Mutex
	name     string
	// Mapping of a Pods mapping of Containers and their TopologyHints
	// Indexed by PodUID to ContainerName
	podTopologyHints podTopologyHints
	// The list of components registered with the Manager
	providers []ResourceAllocator
	// Topology Manager Policy
	policy Policy
	// Mapping of (PodUid, ContainerName) to ContainerID for Adding/Removing Pods from PodTopologyHints mapping
	podMap containermap.ContainerMap
}

func (s *scope) Name() string {
	return s.name
}

func (s *scope) getTopologyHints(podUID string, containerName string) TopologyHint {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.podTopologyHints[podUID][containerName]
}

func (s *scope) setTopologyHints(podUID string, containerName string, th TopologyHint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.podTopologyHints[podUID] == nil {
		s.podTopologyHints[podUID] = make(map[string]TopologyHint)
	}
	s.podTopologyHints[podUID][containerName] = th
}

func (s *scope) GetAffinity(podUID string, containerName string) TopologyHint {
	return s.getTopologyHints(podUID, containerName)
}

func (s *scope) GetPolicy() Policy {
	return s.policy
}

func (s *scope) RegisterProvider(ra ResourceAllocator) {
	s.providers = append(s.providers, ra)
}

// It would be better to implement this function in topologymanager instead of scope
// but topologymanager do not track mapping anymore
func (s *scope) AddContainer(pod *v1.Pod, container *v1.Container, containerID string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.podMap.Add(string(pod.UID), container.Name, containerID)
}

// It would be better to implement this function in topologymanager instead of scope
// but topologymanager do not track mapping anymore
func (s *scope) RemoveContainer(containerID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	klog.InfoS("RemoveContainer", "containerID", containerID)
	// Get the podUID and containerName associated with the containerID to be removed and remove it
	podUIDString, containerName, err := s.podMap.GetContainerRef(containerID)
	if err != nil {
		return nil
	}
	s.podMap.RemoveByContainerID(containerID)

	// In cases where a container has been restarted, it's possible that the same podUID and
	// containerName are already associated with a *different* containerID now. Only remove
	// the TopologyHints associated with that podUID and containerName if this is not true
	if _, err := s.podMap.GetContainerID(podUIDString, containerName); err != nil {
		delete(s.podTopologyHints[podUIDString], containerName)
		if len(s.podTopologyHints[podUIDString]) == 0 {
			delete(s.podTopologyHints, podUIDString)
		}
	}

	return nil
}

type allocationMap map[string][]string

func (am allocationMap) Add(containerName string, resources []string) {
	for _, resource := range resources {
		am[resource] = append(am[resource], containerName)
	}
}

func (am allocationMap) String() string {
	if len(am) == 0 {
		return "none"
	}

	resources := make([]string, 0, len(am))
	for resource := range am {
		resources = append(resources, resource)
	}
	sort.Strings(resources)

	items := []string{}
	for _, resource := range resources {
		contNames := am[resource]
		items = append(items, resource+fmt.Sprintf(": containers=%d", len(contNames)))
	}
	return strings.Join(items, "; ")
}

func (s *scope) resourceAllocationSuccessEvent(pod *v1.Pod, allocs allocationMap) {
	s.recorder.Event(pod, v1.EventTypeNormal, events.AllocatedResources, allocs.String())
}

func (s *scope) admitPolicyNone(pod *v1.Pod) lifecycle.PodAdmitResult {
	allocs := make(allocationMap)
	for _, container := range append(pod.Spec.InitContainers, pod.Spec.Containers...) {
		resources, err := s.allocateAlignedResources(pod, &container)
		if err != nil {
			err = admission.ResourceAllocationFailureEvent(s.recorder, pod, container.Name, err)
			return admission.GetPodAdmitResult(err)
		}
		allocs.Add(container.Name, resources)
	}
	s.resourceAllocationSuccessEvent(pod, allocs)
	return admission.GetPodAdmitResult(nil)
}

// It would be better to implement this function in topologymanager instead of scope
// but topologymanager do not track providers anymore
func (s *scope) allocateAlignedResources(pod *v1.Pod, container *v1.Container) ([]string, error) {
	allocated := []string{}
	for _, provider := range s.providers {
		err := provider.Allocate(pod, container)
		if err != nil {
			return allocated, err
		}
		res := provider.GetExclusiveResources(pod, container)
		allocated = append(allocated, res...)
	}
	return allocated, nil
}
