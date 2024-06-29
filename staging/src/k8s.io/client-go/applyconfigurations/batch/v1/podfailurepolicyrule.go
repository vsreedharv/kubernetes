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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	batchv1 "k8s.io/api/batch/v1"
)

// PodFailurePolicyRuleApplyConfiguration represents a declarative configuration of the PodFailurePolicyRule type for use
// with apply.
type PodFailurePolicyRuleApplyConfiguration struct {
	Action          *batchv1.PodFailurePolicyAction                            `json:"action,omitempty"`
	OnExitCodes     *PodFailurePolicyOnExitCodesRequirementApplyConfiguration  `json:"onExitCodes,omitempty"`
	OnPodConditions []PodFailurePolicyOnPodConditionsPatternApplyConfiguration `json:"onPodConditions,omitempty"`
	Name            *string                                                    `json:"name,omitempty"`
}

// PodFailurePolicyRuleApplyConfiguration constructs a declarative configuration of the PodFailurePolicyRule type for use with
// apply.
func PodFailurePolicyRule() *PodFailurePolicyRuleApplyConfiguration {
	return &PodFailurePolicyRuleApplyConfiguration{}
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *PodFailurePolicyRuleApplyConfiguration) WithAction(value batchv1.PodFailurePolicyAction) *PodFailurePolicyRuleApplyConfiguration {
	b.Action = &value
	return b
}

// WithOnExitCodes sets the OnExitCodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OnExitCodes field is set to the value of the last call.
func (b *PodFailurePolicyRuleApplyConfiguration) WithOnExitCodes(value *PodFailurePolicyOnExitCodesRequirementApplyConfiguration) *PodFailurePolicyRuleApplyConfiguration {
	b.OnExitCodes = value
	return b
}

// WithOnPodConditions adds the given value to the OnPodConditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OnPodConditions field.
func (b *PodFailurePolicyRuleApplyConfiguration) WithOnPodConditions(values ...*PodFailurePolicyOnPodConditionsPatternApplyConfiguration) *PodFailurePolicyRuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOnPodConditions")
		}
		b.OnPodConditions = append(b.OnPodConditions, *values[i])
	}
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PodFailurePolicyRuleApplyConfiguration) WithName(value string) *PodFailurePolicyRuleApplyConfiguration {
	b.Name = &value
	return b
}
