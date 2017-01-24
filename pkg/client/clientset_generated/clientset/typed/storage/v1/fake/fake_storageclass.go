/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	v1 "k8s.io/kubernetes/pkg/apis/storage/v1"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

// FakeStorageClasses implements StorageClassInterface
type FakeStorageClasses struct {
	Fake *FakeStorageV1
}

var storageclassesResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"}

func (c *FakeStorageClasses) Create(storageClass *v1.StorageClass) (result *v1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootCreateAction(storageclassesResource, storageClass), &v1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StorageClass), err
}

func (c *FakeStorageClasses) Update(storageClass *v1.StorageClass) (result *v1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateAction(storageclassesResource, storageClass), &v1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StorageClass), err
}

func (c *FakeStorageClasses) Delete(name string, options *api_v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewRootDeleteAction(storageclassesResource, name), &v1.StorageClass{})
	return err
}

func (c *FakeStorageClasses) DeleteCollection(options *api_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	action := core.NewRootDeleteCollectionAction(storageclassesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1.StorageClassList{})
	return err
}

func (c *FakeStorageClasses) Get(name string, options meta_v1.GetOptions) (result *v1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootGetAction(storageclassesResource, name), &v1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StorageClass), err
}

func (c *FakeStorageClasses) List(opts meta_v1.ListOptions) (result *v1.StorageClassList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootListAction(storageclassesResource, opts), &v1.StorageClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := core.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.StorageClassList{}
	for _, item := range obj.(*v1.StorageClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageClasses.
func (c *FakeStorageClasses) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewRootWatchAction(storageclassesResource, opts))
}

// Patch applies the patch and returns the patched storageClass.
func (c *FakeStorageClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.StorageClass, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootPatchSubresourceAction(storageclassesResource, name, data, subresources...), &v1.StorageClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.StorageClass), err
}
