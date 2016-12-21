// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	v1 "k8s.io/kubernetes/pkg/api/v1"
	meta_v1 "k8s.io/kubernetes/pkg/apis/meta/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	intstr "k8s.io/kubernetes/pkg/util/intstr"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_APIVersion, InType: reflect.TypeOf(&APIVersion{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CPUTargetUtilization, InType: reflect.TypeOf(&CPUTargetUtilization{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CustomMetricCurrentStatus, InType: reflect.TypeOf(&CustomMetricCurrentStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CustomMetricCurrentStatusList, InType: reflect.TypeOf(&CustomMetricCurrentStatusList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CustomMetricTarget, InType: reflect.TypeOf(&CustomMetricTarget{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CustomMetricTargetList, InType: reflect.TypeOf(&CustomMetricTargetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSet, InType: reflect.TypeOf(&DaemonSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSetList, InType: reflect.TypeOf(&DaemonSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSetSpec, InType: reflect.TypeOf(&DaemonSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DaemonSetStatus, InType: reflect.TypeOf(&DaemonSetStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Deployment, InType: reflect.TypeOf(&Deployment{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentCondition, InType: reflect.TypeOf(&DeploymentCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentList, InType: reflect.TypeOf(&DeploymentList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentRollback, InType: reflect.TypeOf(&DeploymentRollback{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentSpec, InType: reflect.TypeOf(&DeploymentSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentStatus, InType: reflect.TypeOf(&DeploymentStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentStrategy, InType: reflect.TypeOf(&DeploymentStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_FSGroupStrategyOptions, InType: reflect.TypeOf(&FSGroupStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HTTPIngressPath, InType: reflect.TypeOf(&HTTPIngressPath{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HTTPIngressRuleValue, InType: reflect.TypeOf(&HTTPIngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HorizontalPodAutoscaler, InType: reflect.TypeOf(&HorizontalPodAutoscaler{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HorizontalPodAutoscalerList, InType: reflect.TypeOf(&HorizontalPodAutoscalerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HorizontalPodAutoscalerSpec, InType: reflect.TypeOf(&HorizontalPodAutoscalerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HorizontalPodAutoscalerStatus, InType: reflect.TypeOf(&HorizontalPodAutoscalerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_HostPortRange, InType: reflect.TypeOf(&HostPortRange{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IDRange, InType: reflect.TypeOf(&IDRange{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Ingress, InType: reflect.TypeOf(&Ingress{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressBackend, InType: reflect.TypeOf(&IngressBackend{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressList, InType: reflect.TypeOf(&IngressList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressRule, InType: reflect.TypeOf(&IngressRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressRuleValue, InType: reflect.TypeOf(&IngressRuleValue{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressSpec, InType: reflect.TypeOf(&IngressSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressStatus, InType: reflect.TypeOf(&IngressStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_IngressTLS, InType: reflect.TypeOf(&IngressTLS{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicy, InType: reflect.TypeOf(&NetworkPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicyIngressRule, InType: reflect.TypeOf(&NetworkPolicyIngressRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicyList, InType: reflect.TypeOf(&NetworkPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicyPeer, InType: reflect.TypeOf(&NetworkPolicyPeer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicyPort, InType: reflect.TypeOf(&NetworkPolicyPort{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_NetworkPolicySpec, InType: reflect.TypeOf(&NetworkPolicySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_PodSecurityPolicy, InType: reflect.TypeOf(&PodSecurityPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_PodSecurityPolicyList, InType: reflect.TypeOf(&PodSecurityPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_PodSecurityPolicySpec, InType: reflect.TypeOf(&PodSecurityPolicySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSet, InType: reflect.TypeOf(&ReplicaSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSetCondition, InType: reflect.TypeOf(&ReplicaSetCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSetList, InType: reflect.TypeOf(&ReplicaSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSetSpec, InType: reflect.TypeOf(&ReplicaSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicaSetStatus, InType: reflect.TypeOf(&ReplicaSetStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ReplicationControllerDummy, InType: reflect.TypeOf(&ReplicationControllerDummy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RollbackConfig, InType: reflect.TypeOf(&RollbackConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RollingUpdateDeployment, InType: reflect.TypeOf(&RollingUpdateDeployment{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RunAsUserStrategyOptions, InType: reflect.TypeOf(&RunAsUserStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SELinuxStrategyOptions, InType: reflect.TypeOf(&SELinuxStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Scale, InType: reflect.TypeOf(&Scale{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ScaleSpec, InType: reflect.TypeOf(&ScaleSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ScaleStatus, InType: reflect.TypeOf(&ScaleStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SubresourceReference, InType: reflect.TypeOf(&SubresourceReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_SupplementalGroupsStrategyOptions, InType: reflect.TypeOf(&SupplementalGroupsStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ThirdPartyResource, InType: reflect.TypeOf(&ThirdPartyResource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ThirdPartyResourceData, InType: reflect.TypeOf(&ThirdPartyResourceData{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ThirdPartyResourceDataList, InType: reflect.TypeOf(&ThirdPartyResourceDataList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ThirdPartyResourceList, InType: reflect.TypeOf(&ThirdPartyResourceList{})},
	)
}

func DeepCopy_v1beta1_APIVersion(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIVersion)
		out := out.(*APIVersion)
		out.Name = in.Name
		return nil
	}
}

func DeepCopy_v1beta1_CPUTargetUtilization(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CPUTargetUtilization)
		out := out.(*CPUTargetUtilization)
		out.TargetPercentage = in.TargetPercentage
		return nil
	}
}

func DeepCopy_v1beta1_CustomMetricCurrentStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomMetricCurrentStatus)
		out := out.(*CustomMetricCurrentStatus)
		out.Name = in.Name
		out.CurrentValue = in.CurrentValue.DeepCopy()
		return nil
	}
}

func DeepCopy_v1beta1_CustomMetricCurrentStatusList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomMetricCurrentStatusList)
		out := out.(*CustomMetricCurrentStatusList)
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CustomMetricCurrentStatus, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_CustomMetricCurrentStatus(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_CustomMetricTarget(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomMetricTarget)
		out := out.(*CustomMetricTarget)
		out.Name = in.Name
		out.TargetValue = in.TargetValue.DeepCopy()
		return nil
	}
}

func DeepCopy_v1beta1_CustomMetricTargetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomMetricTargetList)
		out := out.(*CustomMetricTargetList)
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CustomMetricTarget, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_CustomMetricTarget(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_DaemonSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSet)
		out := out.(*DaemonSet)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DaemonSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_v1beta1_DaemonSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSetList)
		out := out.(*DaemonSetList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]DaemonSet, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_DaemonSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_DaemonSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSetSpec)
		out := out.(*DaemonSetSpec)
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(meta_v1.LabelSelector)
			if err := meta_v1.DeepCopy_v1_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if err := v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_DaemonSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DaemonSetStatus)
		out := out.(*DaemonSetStatus)
		out.CurrentNumberScheduled = in.CurrentNumberScheduled
		out.NumberMisscheduled = in.NumberMisscheduled
		out.DesiredNumberScheduled = in.DesiredNumberScheduled
		out.NumberReady = in.NumberReady
		return nil
	}
}

func DeepCopy_v1beta1_Deployment(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Deployment)
		out := out.(*Deployment)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DeploymentSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DeploymentStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentCondition)
		out := out.(*DeploymentCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.LastUpdateTime = in.LastUpdateTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		out.Reason = in.Reason
		out.Message = in.Message
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentList)
		out := out.(*DeploymentList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Deployment, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_Deployment(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentRollback(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentRollback)
		out := out.(*DeploymentRollback)
		out.TypeMeta = in.TypeMeta
		out.Name = in.Name
		if in.UpdatedAnnotations != nil {
			in, out := &in.UpdatedAnnotations, &out.UpdatedAnnotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.UpdatedAnnotations = nil
		}
		out.RollbackTo = in.RollbackTo
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentSpec)
		out := out.(*DeploymentSpec)
		if in.Replicas != nil {
			in, out := &in.Replicas, &out.Replicas
			*out = new(int32)
			**out = **in
		} else {
			out.Replicas = nil
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(meta_v1.LabelSelector)
			if err := meta_v1.DeepCopy_v1_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if err := v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DeploymentStrategy(&in.Strategy, &out.Strategy, c); err != nil {
			return err
		}
		out.MinReadySeconds = in.MinReadySeconds
		if in.RevisionHistoryLimit != nil {
			in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
			*out = new(int32)
			**out = **in
		} else {
			out.RevisionHistoryLimit = nil
		}
		out.Paused = in.Paused
		if in.RollbackTo != nil {
			in, out := &in.RollbackTo, &out.RollbackTo
			*out = new(RollbackConfig)
			**out = **in
		} else {
			out.RollbackTo = nil
		}
		if in.ProgressDeadlineSeconds != nil {
			in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
			*out = new(int32)
			**out = **in
		} else {
			out.ProgressDeadlineSeconds = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentStatus)
		out := out.(*DeploymentStatus)
		out.ObservedGeneration = in.ObservedGeneration
		out.Replicas = in.Replicas
		out.UpdatedReplicas = in.UpdatedReplicas
		out.ReadyReplicas = in.ReadyReplicas
		out.AvailableReplicas = in.AvailableReplicas
		out.UnavailableReplicas = in.UnavailableReplicas
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]DeploymentCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_DeploymentCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentStrategy)
		out := out.(*DeploymentStrategy)
		out.Type = in.Type
		if in.RollingUpdate != nil {
			in, out := &in.RollingUpdate, &out.RollingUpdate
			*out = new(RollingUpdateDeployment)
			if err := DeepCopy_v1beta1_RollingUpdateDeployment(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.RollingUpdate = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_FSGroupStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FSGroupStrategyOptions)
		out := out.(*FSGroupStrategyOptions)
		out.Rule = in.Rule
		if in.Ranges != nil {
			in, out := &in.Ranges, &out.Ranges
			*out = make([]IDRange, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Ranges = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_HTTPIngressPath(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HTTPIngressPath)
		out := out.(*HTTPIngressPath)
		out.Path = in.Path
		out.Backend = in.Backend
		return nil
	}
}

func DeepCopy_v1beta1_HTTPIngressRuleValue(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HTTPIngressRuleValue)
		out := out.(*HTTPIngressRuleValue)
		if in.Paths != nil {
			in, out := &in.Paths, &out.Paths
			*out = make([]HTTPIngressPath, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Paths = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_HorizontalPodAutoscaler(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscaler)
		out := out.(*HorizontalPodAutoscaler)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_HorizontalPodAutoscalerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerList)
		out := out.(*HorizontalPodAutoscalerList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]HorizontalPodAutoscaler, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_HorizontalPodAutoscaler(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_HorizontalPodAutoscalerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerSpec)
		out := out.(*HorizontalPodAutoscalerSpec)
		out.ScaleRef = in.ScaleRef
		if in.MinReplicas != nil {
			in, out := &in.MinReplicas, &out.MinReplicas
			*out = new(int32)
			**out = **in
		} else {
			out.MinReplicas = nil
		}
		out.MaxReplicas = in.MaxReplicas
		if in.CPUUtilization != nil {
			in, out := &in.CPUUtilization, &out.CPUUtilization
			*out = new(CPUTargetUtilization)
			**out = **in
		} else {
			out.CPUUtilization = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_HorizontalPodAutoscalerStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HorizontalPodAutoscalerStatus)
		out := out.(*HorizontalPodAutoscalerStatus)
		if in.ObservedGeneration != nil {
			in, out := &in.ObservedGeneration, &out.ObservedGeneration
			*out = new(int64)
			**out = **in
		} else {
			out.ObservedGeneration = nil
		}
		if in.LastScaleTime != nil {
			in, out := &in.LastScaleTime, &out.LastScaleTime
			*out = new(meta_v1.Time)
			**out = (*in).DeepCopy()
		} else {
			out.LastScaleTime = nil
		}
		out.CurrentReplicas = in.CurrentReplicas
		out.DesiredReplicas = in.DesiredReplicas
		if in.CurrentCPUUtilizationPercentage != nil {
			in, out := &in.CurrentCPUUtilizationPercentage, &out.CurrentCPUUtilizationPercentage
			*out = new(int32)
			**out = **in
		} else {
			out.CurrentCPUUtilizationPercentage = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_HostPortRange(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HostPortRange)
		out := out.(*HostPortRange)
		out.Min = in.Min
		out.Max = in.Max
		return nil
	}
}

func DeepCopy_v1beta1_IDRange(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IDRange)
		out := out.(*IDRange)
		out.Min = in.Min
		out.Max = in.Max
		return nil
	}
}

func DeepCopy_v1beta1_Ingress(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Ingress)
		out := out.(*Ingress)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_IngressSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_IngressStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_IngressBackend(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressBackend)
		out := out.(*IngressBackend)
		out.ServiceName = in.ServiceName
		out.ServicePort = in.ServicePort
		return nil
	}
}

func DeepCopy_v1beta1_IngressList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressList)
		out := out.(*IngressList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Ingress, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_Ingress(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_IngressRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressRule)
		out := out.(*IngressRule)
		out.Host = in.Host
		if err := DeepCopy_v1beta1_IngressRuleValue(&in.IngressRuleValue, &out.IngressRuleValue, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_IngressRuleValue(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressRuleValue)
		out := out.(*IngressRuleValue)
		if in.HTTP != nil {
			in, out := &in.HTTP, &out.HTTP
			*out = new(HTTPIngressRuleValue)
			if err := DeepCopy_v1beta1_HTTPIngressRuleValue(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.HTTP = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_IngressSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressSpec)
		out := out.(*IngressSpec)
		if in.Backend != nil {
			in, out := &in.Backend, &out.Backend
			*out = new(IngressBackend)
			**out = **in
		} else {
			out.Backend = nil
		}
		if in.TLS != nil {
			in, out := &in.TLS, &out.TLS
			*out = make([]IngressTLS, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_IngressTLS(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.TLS = nil
		}
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]IngressRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_IngressRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Rules = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_IngressStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressStatus)
		out := out.(*IngressStatus)
		if err := v1.DeepCopy_v1_LoadBalancerStatus(&in.LoadBalancer, &out.LoadBalancer, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_IngressTLS(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IngressTLS)
		out := out.(*IngressTLS)
		if in.Hosts != nil {
			in, out := &in.Hosts, &out.Hosts
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Hosts = nil
		}
		out.SecretName = in.SecretName
		return nil
	}
}

func DeepCopy_v1beta1_NetworkPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicy)
		out := out.(*NetworkPolicy)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_NetworkPolicySpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_NetworkPolicyIngressRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyIngressRule)
		out := out.(*NetworkPolicyIngressRule)
		if in.Ports != nil {
			in, out := &in.Ports, &out.Ports
			*out = make([]NetworkPolicyPort, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_NetworkPolicyPort(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Ports = nil
		}
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = make([]NetworkPolicyPeer, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_NetworkPolicyPeer(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.From = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_NetworkPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyList)
		out := out.(*NetworkPolicyList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]NetworkPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_NetworkPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_NetworkPolicyPeer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyPeer)
		out := out.(*NetworkPolicyPeer)
		if in.PodSelector != nil {
			in, out := &in.PodSelector, &out.PodSelector
			*out = new(meta_v1.LabelSelector)
			if err := meta_v1.DeepCopy_v1_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.PodSelector = nil
		}
		if in.NamespaceSelector != nil {
			in, out := &in.NamespaceSelector, &out.NamespaceSelector
			*out = new(meta_v1.LabelSelector)
			if err := meta_v1.DeepCopy_v1_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.NamespaceSelector = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_NetworkPolicyPort(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicyPort)
		out := out.(*NetworkPolicyPort)
		if in.Protocol != nil {
			in, out := &in.Protocol, &out.Protocol
			*out = new(v1.Protocol)
			**out = **in
		} else {
			out.Protocol = nil
		}
		if in.Port != nil {
			in, out := &in.Port, &out.Port
			*out = new(intstr.IntOrString)
			**out = **in
		} else {
			out.Port = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_NetworkPolicySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetworkPolicySpec)
		out := out.(*NetworkPolicySpec)
		if err := meta_v1.DeepCopy_v1_LabelSelector(&in.PodSelector, &out.PodSelector, c); err != nil {
			return err
		}
		if in.Ingress != nil {
			in, out := &in.Ingress, &out.Ingress
			*out = make([]NetworkPolicyIngressRule, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_NetworkPolicyIngressRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Ingress = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_PodSecurityPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicy)
		out := out.(*PodSecurityPolicy)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_PodSecurityPolicySpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_PodSecurityPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicyList)
		out := out.(*PodSecurityPolicyList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PodSecurityPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_PodSecurityPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_PodSecurityPolicySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodSecurityPolicySpec)
		out := out.(*PodSecurityPolicySpec)
		out.Privileged = in.Privileged
		if in.DefaultAddCapabilities != nil {
			in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
			*out = make([]v1.Capability, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.DefaultAddCapabilities = nil
		}
		if in.RequiredDropCapabilities != nil {
			in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
			*out = make([]v1.Capability, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.RequiredDropCapabilities = nil
		}
		if in.AllowedCapabilities != nil {
			in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
			*out = make([]v1.Capability, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.AllowedCapabilities = nil
		}
		if in.Volumes != nil {
			in, out := &in.Volumes, &out.Volumes
			*out = make([]FSType, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Volumes = nil
		}
		out.HostNetwork = in.HostNetwork
		if in.HostPorts != nil {
			in, out := &in.HostPorts, &out.HostPorts
			*out = make([]HostPortRange, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.HostPorts = nil
		}
		out.HostPID = in.HostPID
		out.HostIPC = in.HostIPC
		if err := DeepCopy_v1beta1_SELinuxStrategyOptions(&in.SELinux, &out.SELinux, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_RunAsUserStrategyOptions(&in.RunAsUser, &out.RunAsUser, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_SupplementalGroupsStrategyOptions(&in.SupplementalGroups, &out.SupplementalGroups, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_FSGroupStrategyOptions(&in.FSGroup, &out.FSGroup, c); err != nil {
			return err
		}
		out.ReadOnlyRootFilesystem = in.ReadOnlyRootFilesystem
		return nil
	}
}

func DeepCopy_v1beta1_ReplicaSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSet)
		out := out.(*ReplicaSet)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_ReplicaSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_ReplicaSetStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_ReplicaSetCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSetCondition)
		out := out.(*ReplicaSetCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		out.Reason = in.Reason
		out.Message = in.Message
		return nil
	}
}

func DeepCopy_v1beta1_ReplicaSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSetList)
		out := out.(*ReplicaSetList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ReplicaSet, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ReplicaSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_ReplicaSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSetSpec)
		out := out.(*ReplicaSetSpec)
		if in.Replicas != nil {
			in, out := &in.Replicas, &out.Replicas
			*out = new(int32)
			**out = **in
		} else {
			out.Replicas = nil
		}
		out.MinReadySeconds = in.MinReadySeconds
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(meta_v1.LabelSelector)
			if err := meta_v1.DeepCopy_v1_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if err := v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_ReplicaSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicaSetStatus)
		out := out.(*ReplicaSetStatus)
		out.Replicas = in.Replicas
		out.FullyLabeledReplicas = in.FullyLabeledReplicas
		out.ReadyReplicas = in.ReadyReplicas
		out.AvailableReplicas = in.AvailableReplicas
		out.ObservedGeneration = in.ObservedGeneration
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]ReplicaSetCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ReplicaSetCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_ReplicationControllerDummy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ReplicationControllerDummy)
		out := out.(*ReplicationControllerDummy)
		out.TypeMeta = in.TypeMeta
		return nil
	}
}

func DeepCopy_v1beta1_RollbackConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RollbackConfig)
		out := out.(*RollbackConfig)
		out.Revision = in.Revision
		return nil
	}
}

func DeepCopy_v1beta1_RollingUpdateDeployment(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RollingUpdateDeployment)
		out := out.(*RollingUpdateDeployment)
		if in.MaxUnavailable != nil {
			in, out := &in.MaxUnavailable, &out.MaxUnavailable
			*out = new(intstr.IntOrString)
			**out = **in
		} else {
			out.MaxUnavailable = nil
		}
		if in.MaxSurge != nil {
			in, out := &in.MaxSurge, &out.MaxSurge
			*out = new(intstr.IntOrString)
			**out = **in
		} else {
			out.MaxSurge = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_RunAsUserStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RunAsUserStrategyOptions)
		out := out.(*RunAsUserStrategyOptions)
		out.Rule = in.Rule
		if in.Ranges != nil {
			in, out := &in.Ranges, &out.Ranges
			*out = make([]IDRange, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Ranges = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_SELinuxStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SELinuxStrategyOptions)
		out := out.(*SELinuxStrategyOptions)
		out.Rule = in.Rule
		if in.SELinuxOptions != nil {
			in, out := &in.SELinuxOptions, &out.SELinuxOptions
			*out = new(v1.SELinuxOptions)
			**out = **in
		} else {
			out.SELinuxOptions = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_Scale(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Scale)
		out := out.(*Scale)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Spec = in.Spec
		if err := DeepCopy_v1beta1_ScaleStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_ScaleSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleSpec)
		out := out.(*ScaleSpec)
		out.Replicas = in.Replicas
		return nil
	}
}

func DeepCopy_v1beta1_ScaleStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleStatus)
		out := out.(*ScaleStatus)
		out.Replicas = in.Replicas
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Selector = nil
		}
		out.TargetSelector = in.TargetSelector
		return nil
	}
}

func DeepCopy_v1beta1_SubresourceReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubresourceReference)
		out := out.(*SubresourceReference)
		out.Kind = in.Kind
		out.Name = in.Name
		out.APIVersion = in.APIVersion
		out.Subresource = in.Subresource
		return nil
	}
}

func DeepCopy_v1beta1_SupplementalGroupsStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SupplementalGroupsStrategyOptions)
		out := out.(*SupplementalGroupsStrategyOptions)
		out.Rule = in.Rule
		if in.Ranges != nil {
			in, out := &in.Ranges, &out.Ranges
			*out = make([]IDRange, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Ranges = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_ThirdPartyResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ThirdPartyResource)
		out := out.(*ThirdPartyResource)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Description = in.Description
		if in.Versions != nil {
			in, out := &in.Versions, &out.Versions
			*out = make([]APIVersion, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Versions = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_ThirdPartyResourceData(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ThirdPartyResourceData)
		out := out.(*ThirdPartyResourceData)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.Data != nil {
			in, out := &in.Data, &out.Data
			*out = make([]byte, len(*in))
			copy(*out, *in)
		} else {
			out.Data = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_ThirdPartyResourceDataList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ThirdPartyResourceDataList)
		out := out.(*ThirdPartyResourceDataList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ThirdPartyResourceData, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ThirdPartyResourceData(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_ThirdPartyResourceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ThirdPartyResourceList)
		out := out.(*ThirdPartyResourceList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ThirdPartyResource, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_ThirdPartyResource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}
