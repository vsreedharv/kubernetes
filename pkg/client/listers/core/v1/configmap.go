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

// This file was automatically generated by lister-gen with arguments: --input-dirs=[k8s.io/kubernetes/pkg/api,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/certificates/v1beta1,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/policy/v1alpha1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/rbac/v1beta1,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/apis/storage/v1,k8s.io/kubernetes/pkg/apis/storage/v1beta1]

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	api "k8s.io/kubernetes/pkg/api"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/cache"
)

// ConfigMapLister helps list ConfigMaps.
type ConfigMapLister interface {
	// List lists all ConfigMaps in the indexer.
	List(selector labels.Selector) (ret []*v1.ConfigMap, err error)
	// ConfigMaps returns an object that can list and get ConfigMaps.
	ConfigMaps(namespace string) ConfigMapNamespaceLister
	ConfigMapListerExpansion
}

// configMapLister implements the ConfigMapLister interface.
type configMapLister struct {
	indexer cache.Indexer
}

// NewConfigMapLister returns a new ConfigMapLister.
func NewConfigMapLister(indexer cache.Indexer) ConfigMapLister {
	return &configMapLister{indexer: indexer}
}

// List lists all ConfigMaps in the indexer.
func (s *configMapLister) List(selector labels.Selector) (ret []*v1.ConfigMap, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ConfigMap))
	})
	return ret, err
}

// ConfigMaps returns an object that can list and get ConfigMaps.
func (s *configMapLister) ConfigMaps(namespace string) ConfigMapNamespaceLister {
	return configMapNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigMapNamespaceLister helps list and get ConfigMaps.
type ConfigMapNamespaceLister interface {
	// List lists all ConfigMaps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ConfigMap, err error)
	// Get retrieves the ConfigMap from the indexer for a given namespace and name.
	Get(name string) (*v1.ConfigMap, error)
	ConfigMapNamespaceListerExpansion
}

// configMapNamespaceLister implements the ConfigMapNamespaceLister
// interface.
type configMapNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigMaps in the indexer for a given namespace.
func (s configMapNamespaceLister) List(selector labels.Selector) (ret []*v1.ConfigMap, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ConfigMap))
	})
	return ret, err
}

// Get retrieves the ConfigMap from the indexer for a given namespace and name.
func (s configMapNamespaceLister) Get(name string) (*v1.ConfigMap, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("configmap"), name)
	}
	return obj.(*v1.ConfigMap), nil
}
