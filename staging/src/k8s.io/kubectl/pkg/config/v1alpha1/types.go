/*
Copyright 2024 The Kubernetes Authors.

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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Preferences stores elements of KubeRC configuration file
type Preferences struct {
	metav1.TypeMeta `json:",inline"`

	// overrides allows changing default flag values of commands.
	// This is especially useful, when user doesn't want to explicitly
	// set flags each time.
	// +listType=atomic
	Overrides []CommandOverride `json:"overrides"`

	// aliases allows defining command aliases for existing kubectl commands, with optional default flag values.
	// If the alias name collides with a built-in command, built-in command always takes precedence.
	// kubectl [ALIAS NAME] [USER_FLAGS] [USER_EXPLICIT_ARGS] expands to
	// kubectl [COMMAND] # built-in command alias points to
	//         [USER_FLAGS]
	//         [KUBERC_FLAGS] # rest of the flags that are not passed by user in [USER_FLAGS]
	//         [USER_EXPLICIT_ARGS]
	//         [KUBERC_ALIAS_ARGUMENTS]
	// e.g.
	// - name: runx
	//   command: run
	//   flags:
	//   - name: image
	//     default: nginx
	//   args:
	//   - --
	//   - custom-arg1
	// For example, if user invokes "kubectl runx test-pod" command,
	// this will be expanded to "kubectl run --image=nginx test-pod -- custom-arg1"
	// +listType=atomic
	Aliases []AliasOverride `json:"aliases"`
}

// AliasOverride stores the alias definitions.
type AliasOverride struct {
	// Name is the name of alias that can only include alphabetical characters
	// If the alias name conflicts with the built-in command,
	// built-in command will be used.
	Name string `json:"name"`
	// Command is the single or set of commands to execute, such as "set env" or "create"
	Command string `json:"command"`
	// Arguments is allocated for the arguments such as resource names, etc.
	// These arguments are appended after the explicitly defined arguments.
	// +listType=atomic
	Args []string `json:"args,omitempty"`
	// Flag is allocated to store the flag definitions of alias.
	// Flag only modifies the default value of the flag and if
	// user explicitly passes a value, explicit one is used.
	// +listType=atomic
	Flags []CommandOverrideFlag `json:"flags,omitempty"`
}

// CommandOverride stores the commands and their associated flag's
// default values.
type CommandOverride struct {
	// Command refers to a command whose flag's default value is changed.
	Command string `json:"command"`
	// Flags is a list of flags storing different default values.
	// +listType=atomic
	Flags []CommandOverrideFlag `json:"flags"`
}

// CommandOverrideFlag stores the name and the specified default
// value of the flag.
type CommandOverrideFlag struct {
	// Flags name without dashes as prefix.
	Name string `json:"name"`

	// In a string format of a default value. It will be parsed
	// by kubectl to the compatible value of the flag.
	Default string `json:"default"`
}
