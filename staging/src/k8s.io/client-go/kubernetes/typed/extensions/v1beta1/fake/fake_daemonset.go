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

	v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	extensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
	testing "k8s.io/client-go/testing"
)

// FakeDaemonSets implements DaemonSetInterface
type FakeDaemonSets struct {
	Fake *FakeExtensionsV1beta1
	ns   string
}

var daemonsetsResource = v1beta1.SchemeGroupVersion.WithResource("daemonsets")

var daemonsetsKind = v1beta1.SchemeGroupVersion.WithKind("DaemonSet")

// Get takes name of the daemonSet, and returns the corresponding daemonSet object, and an error if there is any.
func (c *FakeDaemonSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(daemonsetsResource, c.ns, name), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

// List takes label and field selectors, and returns the list of DaemonSets that match those selectors.
func (c *FakeDaemonSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DaemonSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(daemonsetsResource, daemonsetsKind, c.ns, opts), &v1beta1.DaemonSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DaemonSetList{ListMeta: obj.(*v1beta1.DaemonSetList).ListMeta}
	for _, item := range obj.(*v1beta1.DaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daemonSets.
func (c *FakeDaemonSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(daemonsetsResource, c.ns, opts))

}

// Create takes the representation of a daemonSet and creates it.  Returns the server's representation of the daemonSet, and an error, if there is any.
func (c *FakeDaemonSets) Create(ctx context.Context, daemonSet *v1beta1.DaemonSet, opts v1.CreateOptions) (result *v1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(daemonsetsResource, c.ns, daemonSet), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

// Update takes the representation of a daemonSet and updates it. Returns the server's representation of the daemonSet, and an error, if there is any.
func (c *FakeDaemonSets) Update(ctx context.Context, daemonSet *v1beta1.DaemonSet, opts v1.UpdateOptions) (result *v1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(daemonsetsResource, c.ns, daemonSet), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDaemonSets) UpdateStatus(ctx context.Context, daemonSet *v1beta1.DaemonSet, opts v1.UpdateOptions) (*v1beta1.DaemonSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(daemonsetsResource, "status", c.ns, daemonSet), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

// Delete takes name of the daemonSet and deletes it. Returns an error if one occurs.
func (c *FakeDaemonSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(daemonsetsResource, c.ns, name, opts), &v1beta1.DaemonSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDaemonSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(daemonsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DaemonSetList{})
	return err
}

// Patch applies the patch and returns the patched daemonSet.
func (c *FakeDaemonSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(daemonsetsResource, c.ns, name, pt, data, subresources...), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied daemonSet.
func (c *FakeDaemonSets) Apply(ctx context.Context, daemonSet *extensionsv1beta1.DaemonSetApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.DaemonSet, err error) {
	if daemonSet == nil {
		return nil, fmt.Errorf("daemonSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(daemonSet)
	if err != nil {
		return nil, err
	}

	manager := "default-test-manager"
	if m := opts.FieldManager; m != "" {
		manager = m
	}

	name := daemonSet.Name
	if name == nil {
		return nil, fmt.Errorf("daemonSet.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewApplySubresourceAction(daemonsetsResource, c.ns, *name, data, manager, opts.Force), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDaemonSets) ApplyStatus(ctx context.Context, daemonSet *extensionsv1beta1.DaemonSetApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.DaemonSet, err error) {
	if daemonSet == nil {
		return nil, fmt.Errorf("daemonSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(daemonSet)
	if err != nil {
		return nil, err
	}

	manager := "default-test-manager"
	if m := opts.FieldManager; m != "" {
		manager = m
	}

	name := daemonSet.Name
	if name == nil {
		return nil, fmt.Errorf("daemonSet.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewApplySubresourceAction(daemonsetsResource, c.ns, *name, data, manager, opts.Force, "status"), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}
