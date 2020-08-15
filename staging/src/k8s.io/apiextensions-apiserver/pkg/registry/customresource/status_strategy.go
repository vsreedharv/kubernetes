/*
Copyright 2018 The Kubernetes Authors.

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

package customresource

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

type statusStrategy struct {
	customResourceStrategy
}

func NewStatusStrategy(strategy customResourceStrategy) statusStrategy {
	return statusStrategy{strategy}
}

// GetResetFields returns the set of fields that get reset by the strategy
// and should not be modified by the user.
func (a statusStrategy) GetResetFields() map[fieldpath.APIVersion]*fieldpath.Set {
	fields := map[fieldpath.APIVersion]*fieldpath.Set{
		"apiextensions.k8s.io/v1": fieldpath.NewSet(
			// TODO: this is not correct, we might need a way to specify an inverse,
			// like ignore everything but status.
			fieldpath.MakePathOrDie("spec"),
		),
		"apiextensions.k8s.io/v1beta1": fieldpath.NewSet(
			// TODO: this is not correct, we might need a way to specify an inverse,
			// like ignore everything but status.
			fieldpath.MakePathOrDie("spec"),
		),
	}

	return fields
}

func (a statusStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	// update is only allowed to set status
	newCustomResourceObject := obj.(*unstructured.Unstructured)
	newCustomResource := newCustomResourceObject.UnstructuredContent()
	status, ok := newCustomResource["status"]

	// managedFields must be preserved since it's been modified to
	// track changed fields in the status update.
	managedFields := newCustomResourceObject.GetManagedFields()

	// copy old object into new object
	oldCustomResourceObject := old.(*unstructured.Unstructured)
	// overridding the resourceVersion in metadata is safe here, we have already checked that
	// new object and old object have the same resourceVersion.
	*newCustomResourceObject = *oldCustomResourceObject.DeepCopy()

	// set status
	newCustomResourceObject.SetManagedFields(managedFields)
	newCustomResource = newCustomResourceObject.UnstructuredContent()
	if ok {
		newCustomResource["status"] = status
	} else {
		delete(newCustomResource, "status")
	}
}

// ValidateUpdate is the default update validation for an end user updating status.
func (a statusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return a.customResourceStrategy.validator.ValidateStatusUpdate(ctx, obj, old, a.scale)
}
