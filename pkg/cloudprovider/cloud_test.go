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

package cloudprovider

import (
	"testing"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetLoadBalancerName(t *testing.T) {
	s := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-service",
			Namespace: "test-namespace",
			UID:       "abc-def-123-456",
		},
	}

	expected := "aabcdef123456"
	name := GetLoadBalancerName(s)
	if name != expected {
		t.Errorf("name[%s] != %s", name, expected)
	}
}

func TestGetLoadBalancerNameAnnotationLoadBalancerName(t *testing.T) {
	// Test without a value
	s := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-service",
			Namespace: "test-namespace",
			UID:       "abc-def-123-456",
			Annotations: map[string]string{
				ServiceAnnotationLoadBalancerName: "",
			},
		},
	}

	expected := "aabcdef123456"
	name := GetLoadBalancerName(s)
	if name != expected {
		t.Errorf("name[%s] != %s", name, expected)
	}

	// Test without a small value
	s = &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-service",
			Namespace: "test-namespace",
			UID:       "abc-def-123-456",
			Annotations: map[string]string{
				ServiceAnnotationLoadBalancerName: "abc-def",
			},
		},
	}

	expected = "k8s-abc-def-abcdef123456"
	name = GetLoadBalancerName(s)
	if name != expected {
		t.Errorf("name[%s] != %s", name, expected)
	}
	if len(name) > 32 {
		t.Error("name length is larger than 32 characters")
	}

	// Test without a large value
	s = &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-service",
			Namespace: "test-namespace",
			UID:       "abc-def-123-456",
			Annotations: map[string]string{
				ServiceAnnotationLoadBalancerName: "123456-123456-123456",
			},
		},
	}

	expected = "k8s-123456-123456-12-abcdef12345"
	name = GetLoadBalancerName(s)
	if name != expected {
		t.Errorf("name[%s] != %s", name, expected)
	}
	if len(name) > 32 {
		t.Error("name length is larger than 32 characters")
	}

	// Test with non-alphanumeric ASCII characters
	s = &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-service",
			Namespace: "test-namespace",
			UID:       "abc-def-123-456",
			Annotations: map[string]string{
				ServiceAnnotationLoadBalancerName: "  .!世界-123A",
			},
		},
	}

	expected = "k8s-zzzzzz-123A-abcdef123456"
	name = GetLoadBalancerName(s)
	if name != expected {
		t.Errorf("name[%s] != %s", name, expected)
	}
	if len(name) > 32 {
		t.Error("name length is larger than 32 characters")
	}

}

func TestGetLoadBalancerNameAnnotationLoadBalancerAutoGeneratedName(t *testing.T) {
	// Test with a small value
	s := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "one",
			Namespace: "project",
			UID:       "abc-def-123-456",
			Annotations: map[string]string{
				ServiceAnnotationLoadBalancerAutoGenerateName: "",
			},
		},
	}

	expected := "k8s-project-one-abcdef123456"
	name := GetLoadBalancerName(s)
	if name != expected {
		t.Errorf("name[%s] != %s", name, expected)
	}
	if len(name) > 32 {
		t.Error("name length is larger than 32 characters")
	}

	// Test with a large value
	s = &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mylargeproject",
			Namespace: "large-namespace",
			UID:       "abc-def-123-456",
			Annotations: map[string]string{
				ServiceAnnotationLoadBalancerAutoGenerateName: "",
			},
		},
	}

	expected = "k8s-large-na-mylargep-abcdef1234"
	name = GetLoadBalancerName(s)
	if name != expected {
		t.Errorf("name[%s] != %s", name, expected)
	}
	if len(name) > 32 {
		t.Error("name length is larger than 32 characters")
	}
}
