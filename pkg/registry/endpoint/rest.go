/*
Copyright 2014 Google Inc. All rights reserved.

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

package endpoint

import (
	"fmt"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/errors"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/apiserver"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/registry"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/watch"
)

// REST adapts endpoints into apiserver's RESTStorage model.
type REST struct {
	registry Registry
	store    registry.Store
}

// NewREST returns a new apiserver.RESTStorage implementation for endpoints
func NewREST(registry Registry) *REST {
	return &REST{
		registry: registry,
	}
}

func NewRESTv2(store registry.Store) *REST {
	return &REST{
		store: store,
	}
}

// Get satisfies the RESTStorage interface.
func (rs *REST) Get(ctx api.Context, id string) (runtime.Object, error) {
	if rs.registry != nil {
		// Old way.
		return rs.registry.GetEndpoints(ctx, id)
	}
	// New way.
	endpoints := &api.Endpoints{}
	if err := rs.store.Get(ctx, id, endpoints); err != nil {
		return nil, err
	}
	return endpoints, nil
}

// List satisfies the RESTStorage interface.
func (rs *REST) List(ctx api.Context, label, field labels.Selector) (runtime.Object, error) {
	if !label.Empty() || !field.Empty() {
		return nil, errors.NewBadRequest("label/field selectors are not supported on endpoints")
	}
	if rs.registry != nil {
		// Old way.
		return rs.registry.ListEndpoints(ctx)
	}
	// New way.
	list := &api.EndpointsList{}
	err := rs.store.List(ctx, list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Watch returns Endpoint events via a watch.Interface.
// It implements apiserver.ResourceWatcher.
func (rs *REST) Watch(ctx api.Context, label, field labels.Selector, resourceVersion string) (watch.Interface, error) {
	// Old way.
	return rs.registry.WatchEndpoints(ctx, label, field, resourceVersion)
}

// Create satisfies the RESTStorage interface.
func (rs *REST) Create(ctx api.Context, obj runtime.Object) (<-chan apiserver.RESTResult, error) {
	// Old way.
	endpoints, ok := obj.(*api.Endpoints)
	if !ok {
		return nil, fmt.Errorf("not an endpoints: %#v", obj)
	}
	if len(endpoints.Name) == 0 {
		return nil, fmt.Errorf("id is required: %#v", obj)
	}
	if !api.ValidNamespace(ctx, &endpoints.ObjectMeta) {
		return nil, errors.NewConflict("endpoints", endpoints.Namespace, fmt.Errorf("Endpoints.Namespace does not match the provided context"))
	}
	api.FillObjectMetaSystemFields(ctx, &endpoints.ObjectMeta)

	return apiserver.MakeAsync(func() (runtime.Object, error) {
		err := rs.registry.UpdateEndpoints(ctx, endpoints)
		if err != nil {
			return nil, err
		}
		return rs.registry.GetEndpoints(ctx, endpoints.Name)
	}), nil
}

// Update satisfies the RESTStorage interface.
func (rs *REST) Update(ctx api.Context, obj runtime.Object) (<-chan apiserver.RESTResult, error) {
	endpoints, ok := obj.(*api.Endpoints)
	if !ok {
		return nil, fmt.Errorf("not an endpoints: %#v", obj)
	}
	if rs.registry != nil {
		// Old way.
		return apiserver.MakeAsync(func() (runtime.Object, error) {
			err := rs.registry.UpdateEndpoints(ctx, endpoints)
			if err != nil {
				return nil, err
			}
			return rs.registry.GetEndpoints(ctx, endpoints.Name)
		}), nil
	}
	// New way.
	return apiserver.MakeAsync(func() (runtime.Object, error) {
		if err := rs.store.Update(ctx, endpoints.Name, endpoints); err != nil {
			return nil, err
		}
		if err := rs.store.Get(ctx, endpoints.Name, endpoints); err != nil {
			return nil, err
		}
		return endpoints, nil
	}), nil
}

// Delete satisfies the RESTStorage interface but is unimplemented.
func (rs *REST) Delete(ctx api.Context, id string) (<-chan apiserver.RESTResult, error) {
	return nil, errors.NewBadRequest("Endpoints are read-only")
}

// New implements the RESTStorage interface.
func (rs REST) New() runtime.Object {
	return &api.Endpoints{}
}
