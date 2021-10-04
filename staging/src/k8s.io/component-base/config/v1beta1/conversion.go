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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/component-base/config"
)

// Important! The public back-and-forth conversion functions for the types in this generic
// package with ComponentConfig types need to be manually exposed like this in order for
// other packages that reference this package to be able to call these conversion functions
// in an autogenerated manner.
// TODO: Fix the bug in conversion-gen so it automatically discovers these Convert_* functions
// in autogenerated code as well.

func Convert_v1beta1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *config.ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(in, out, s)
}

func Convert_config_ClientConnectionConfiguration_To_v1beta1_ClientConnectionConfiguration(in *config.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_config_ClientConnectionConfiguration_To_v1beta1_ClientConnectionConfiguration(in, out, s)
}

func Convert_v1beta1_DebuggingConfiguration_To_config_DebuggingConfiguration(in *DebuggingConfiguration, out *config.DebuggingConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_DebuggingConfiguration_To_config_DebuggingConfiguration(in, out, s)
}

func Convert_config_DebuggingConfiguration_To_v1beta1_DebuggingConfiguration(in *config.DebuggingConfiguration, out *DebuggingConfiguration, s conversion.Scope) error {
	return autoConvert_config_DebuggingConfiguration_To_v1beta1_DebuggingConfiguration(in, out, s)
}

func Convert_v1beta1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *config.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(in, out, s)
}

func Convert_config_LeaderElectionConfiguration_To_v1beta1_LeaderElectionConfiguration(in *config.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_config_LeaderElectionConfiguration_To_v1beta1_LeaderElectionConfiguration(in, out, s)
}

func Convert_v1beta1_LoggingConfiguration_To_config_LoggingConfiguration(in *LoggingConfiguration, out *config.LoggingConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_LoggingConfiguration_To_config_LoggingConfiguration(in, out, s)
}

func Convert_config_LoggingConfiguration_To_v1beta1_LoggingConfiguration(in *config.LoggingConfiguration, out *LoggingConfiguration, s conversion.Scope) error {
	return autoConvert_config_LoggingConfiguration_To_v1beta1_LoggingConfiguration(in, out, s)
}
