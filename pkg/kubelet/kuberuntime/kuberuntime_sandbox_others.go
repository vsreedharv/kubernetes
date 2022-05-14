//go:build !windows
// +build !windows

/*
Copyright 2022 The Kubernetes Authors.

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

package kuberuntime

import (
	v1 "k8s.io/api/core/v1"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func addNonWindowsRelatedContext(lc *runtimeapi.LinuxPodSandboxConfig, sc *v1.PodSecurityContext) {
	if sc.RunAsUser != nil {
		lc.SecurityContext.RunAsUser = &runtimeapi.Int64Value{Value: int64(*sc.RunAsUser)}
	}
	if sc.RunAsGroup != nil {
		lc.SecurityContext.RunAsGroup = &runtimeapi.Int64Value{Value: int64(*sc.RunAsGroup)}
	}

	if sc.FSGroup != nil {
		lc.SecurityContext.SupplementalGroups = append(lc.SecurityContext.SupplementalGroups, int64(*sc.FSGroup))
	}
	if sc.SELinuxOptions != nil {
		lc.SecurityContext.SelinuxOptions = &runtimeapi.SELinuxOption{
			User:  sc.SELinuxOptions.User,
			Role:  sc.SELinuxOptions.Role,
			Type:  sc.SELinuxOptions.Type,
			Level: sc.SELinuxOptions.Level,
		}
	}
}

func getPodSandboxWindowsConfig(m *kubeGenericRuntimeManager, pod *v1.Pod, podSandBoxConfig *runtimeapi.PodSandboxConfig) (*runtimeapi.WindowsPodSandboxConfig, error) {
	return nil, nil
}
