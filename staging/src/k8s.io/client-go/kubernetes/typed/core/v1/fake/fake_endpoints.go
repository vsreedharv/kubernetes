/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	testing "k8s.io/client-go/testing"
)

// FakeEndpoints implements EndpointsInterface
type FakeEndpoints struct {
	Fake *FakeCoreV1
	ns   string
}

var endpointsResource = v1.SchemeGroupVersion.WithResource("endpoints")

var endpointsKind = v1.SchemeGroupVersion.WithKind("Endpoints")

// Get takes name of the endpoints, and returns the corresponding endpoints object, and an error if there is any.
func (c *FakeEndpoints) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointsResource, c.ns, name), &v1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoints), err
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *FakeEndpoints) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EndpointsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointsResource, endpointsKind, c.ns, opts), &v1.EndpointsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EndpointsList{ListMeta: obj.(*v1.EndpointsList).ListMeta}
	for _, item := range obj.(*v1.EndpointsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *FakeEndpoints) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointsResource, c.ns, opts))

}

// Create takes the representation of a endpoints and creates it.  Returns the server's representation of the endpoints, and an error, if there is any.
func (c *FakeEndpoints) Create(ctx context.Context, endpoints *v1.Endpoints, opts metav1.CreateOptions) (result *v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointsResource, c.ns, endpoints), &v1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoints), err
}

// Update takes the representation of a endpoints and updates it. Returns the server's representation of the endpoints, and an error, if there is any.
func (c *FakeEndpoints) Update(ctx context.Context, endpoints *v1.Endpoints, opts metav1.UpdateOptions) (result *v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointsResource, c.ns, endpoints), &v1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoints), err
}

// Delete takes name of the endpoints and deletes it. Returns an error if one occurs.
func (c *FakeEndpoints) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(endpointsResource, c.ns, name, opts), &v1.Endpoints{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpoints) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.EndpointsList{})
	return err
}

// Patch applies the patch and returns the patched endpoints.
func (c *FakeEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, name, pt, data, subresources...), &v1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoints), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied endpoints.
func (c *FakeEndpoints) Apply(ctx context.Context, endpoints *corev1.EndpointsApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Endpoints, err error) {
	if endpoints == nil {
		return nil, fmt.Errorf("endpoints provided to Apply must not be nil")
	}
	data, err := json.Marshal(endpoints)
	if err != nil {
		return nil, err
	}

	manager := "default-test-manager"
	if m := opts.FieldManager; m != "" {
		manager = m
	}

	name := endpoints.Name
	if name == nil {
		return nil, fmt.Errorf("endpoints.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewApplySubresourceAction(endpointsResource, c.ns, *name, data, manager, opts.Force), &v1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoints), err
}
