/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-deep-copies.sh.

package metrics

import (
	time "time"

	api "k8s.io/kubernetes/pkg/api"
	resource "k8s.io/kubernetes/pkg/api/resource"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
	inf "speter.net/go/exp/math/dec/inf"
)

func deepCopy_resource_Quantity(in resource.Quantity, out *resource.Quantity, c *conversion.Cloner) error {
	if in.Amount != nil {
		if newVal, err := c.DeepCopy(in.Amount); err != nil {
			return err
		} else {
			out.Amount = newVal.(*inf.Dec)
		}
	} else {
		out.Amount = nil
	}
	out.Format = in.Format
	return nil
}

func deepCopy_unversioned_ListMeta(in unversioned.ListMeta, out *unversioned.ListMeta, c *conversion.Cloner) error {
	out.SelfLink = in.SelfLink
	out.ResourceVersion = in.ResourceVersion
	return nil
}

func deepCopy_unversioned_Time(in unversioned.Time, out *unversioned.Time, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.Time); err != nil {
		return err
	} else {
		out.Time = newVal.(time.Time)
	}
	return nil
}

func deepCopy_unversioned_TypeMeta(in unversioned.TypeMeta, out *unversioned.TypeMeta, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	return nil
}

func deepCopy_metrics_AggregateSample(in AggregateSample, out *AggregateSample, c *conversion.Cloner) error {
	if err := deepCopy_metrics_Sample(in.Sample, &out.Sample, c); err != nil {
		return err
	}
	if in.CPU != nil {
		out.CPU = new(CPUMetrics)
		if err := deepCopy_metrics_CPUMetrics(*in.CPU, out.CPU, c); err != nil {
			return err
		}
	} else {
		out.CPU = nil
	}
	if in.Memory != nil {
		out.Memory = new(MemoryMetrics)
		if err := deepCopy_metrics_MemoryMetrics(*in.Memory, out.Memory, c); err != nil {
			return err
		}
	} else {
		out.Memory = nil
	}
	if in.Network != nil {
		out.Network = new(NetworkMetrics)
		if err := deepCopy_metrics_NetworkMetrics(*in.Network, out.Network, c); err != nil {
			return err
		}
	} else {
		out.Network = nil
	}
	if in.Filesystem != nil {
		out.Filesystem = make([]FilesystemMetrics, len(in.Filesystem))
		for i := range in.Filesystem {
			if err := deepCopy_metrics_FilesystemMetrics(in.Filesystem[i], &out.Filesystem[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Filesystem = nil
	}
	return nil
}

func deepCopy_metrics_CPUMetrics(in CPUMetrics, out *CPUMetrics, c *conversion.Cloner) error {
	if in.TotalCores != nil {
		out.TotalCores = new(resource.Quantity)
		if err := deepCopy_resource_Quantity(*in.TotalCores, out.TotalCores, c); err != nil {
			return err
		}
	} else {
		out.TotalCores = nil
	}
	return nil
}

func deepCopy_metrics_ContainerSample(in ContainerSample, out *ContainerSample, c *conversion.Cloner) error {
	if err := deepCopy_metrics_Sample(in.Sample, &out.Sample, c); err != nil {
		return err
	}
	if in.CPU != nil {
		out.CPU = new(CPUMetrics)
		if err := deepCopy_metrics_CPUMetrics(*in.CPU, out.CPU, c); err != nil {
			return err
		}
	} else {
		out.CPU = nil
	}
	if in.Memory != nil {
		out.Memory = new(MemoryMetrics)
		if err := deepCopy_metrics_MemoryMetrics(*in.Memory, out.Memory, c); err != nil {
			return err
		}
	} else {
		out.Memory = nil
	}
	if in.Filesystem != nil {
		out.Filesystem = make([]FilesystemMetrics, len(in.Filesystem))
		for i := range in.Filesystem {
			if err := deepCopy_metrics_FilesystemMetrics(in.Filesystem[i], &out.Filesystem[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Filesystem = nil
	}
	return nil
}

func deepCopy_metrics_FilesystemMetrics(in FilesystemMetrics, out *FilesystemMetrics, c *conversion.Cloner) error {
	out.Device = in.Device
	if in.UsageBytes != nil {
		out.UsageBytes = new(resource.Quantity)
		if err := deepCopy_resource_Quantity(*in.UsageBytes, out.UsageBytes, c); err != nil {
			return err
		}
	} else {
		out.UsageBytes = nil
	}
	if in.LimitBytes != nil {
		out.LimitBytes = new(resource.Quantity)
		if err := deepCopy_resource_Quantity(*in.LimitBytes, out.LimitBytes, c); err != nil {
			return err
		}
	} else {
		out.LimitBytes = nil
	}
	return nil
}

func deepCopy_metrics_MemoryMetrics(in MemoryMetrics, out *MemoryMetrics, c *conversion.Cloner) error {
	if in.TotalBytes != nil {
		out.TotalBytes = new(resource.Quantity)
		if err := deepCopy_resource_Quantity(*in.TotalBytes, out.TotalBytes, c); err != nil {
			return err
		}
	} else {
		out.TotalBytes = nil
	}
	if in.UsageBytes != nil {
		out.UsageBytes = new(resource.Quantity)
		if err := deepCopy_resource_Quantity(*in.UsageBytes, out.UsageBytes, c); err != nil {
			return err
		}
	} else {
		out.UsageBytes = nil
	}
	if in.PageFaults != nil {
		out.PageFaults = new(int64)
		*out.PageFaults = *in.PageFaults
	} else {
		out.PageFaults = nil
	}
	if in.MajorPageFaults != nil {
		out.MajorPageFaults = new(int64)
		*out.MajorPageFaults = *in.MajorPageFaults
	} else {
		out.MajorPageFaults = nil
	}
	return nil
}

func deepCopy_metrics_MetricsMeta(in MetricsMeta, out *MetricsMeta, c *conversion.Cloner) error {
	out.SelfLink = in.SelfLink
	return nil
}

func deepCopy_metrics_NetworkMetrics(in NetworkMetrics, out *NetworkMetrics, c *conversion.Cloner) error {
	if in.RxBytes != nil {
		out.RxBytes = new(resource.Quantity)
		if err := deepCopy_resource_Quantity(*in.RxBytes, out.RxBytes, c); err != nil {
			return err
		}
	} else {
		out.RxBytes = nil
	}
	if in.RxErrors != nil {
		out.RxErrors = new(int64)
		*out.RxErrors = *in.RxErrors
	} else {
		out.RxErrors = nil
	}
	if in.TxBytes != nil {
		out.TxBytes = new(resource.Quantity)
		if err := deepCopy_resource_Quantity(*in.TxBytes, out.TxBytes, c); err != nil {
			return err
		}
	} else {
		out.TxBytes = nil
	}
	if in.TxErrors != nil {
		out.TxErrors = new(int64)
		*out.TxErrors = *in.TxErrors
	} else {
		out.TxErrors = nil
	}
	return nil
}

func deepCopy_metrics_NonLocalObjectReference(in NonLocalObjectReference, out *NonLocalObjectReference, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.UID = in.UID
	return nil
}

func deepCopy_metrics_PodSample(in PodSample, out *PodSample, c *conversion.Cloner) error {
	if err := deepCopy_metrics_Sample(in.Sample, &out.Sample, c); err != nil {
		return err
	}
	if in.Network != nil {
		out.Network = new(NetworkMetrics)
		if err := deepCopy_metrics_NetworkMetrics(*in.Network, out.Network, c); err != nil {
			return err
		}
	} else {
		out.Network = nil
	}
	return nil
}

func deepCopy_metrics_RawContainerMetrics(in RawContainerMetrics, out *RawContainerMetrics, c *conversion.Cloner) error {
	out.Name = in.Name
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Samples != nil {
		out.Samples = make([]ContainerSample, len(in.Samples))
		for i := range in.Samples {
			if err := deepCopy_metrics_ContainerSample(in.Samples[i], &out.Samples[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Samples = nil
	}
	return nil
}

func deepCopy_metrics_RawMetricsOptions(in RawMetricsOptions, out *RawMetricsOptions, c *conversion.Cloner) error {
	out.MaxSamples = in.MaxSamples
	return nil
}

func deepCopy_metrics_RawNodeMetrics(in RawNodeMetrics, out *RawNodeMetrics, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	out.NodeName = in.NodeName
	if in.Total != nil {
		out.Total = make([]AggregateSample, len(in.Total))
		for i := range in.Total {
			if err := deepCopy_metrics_AggregateSample(in.Total[i], &out.Total[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Total = nil
	}
	if in.SystemContainers != nil {
		out.SystemContainers = make([]RawContainerMetrics, len(in.SystemContainers))
		for i := range in.SystemContainers {
			if err := deepCopy_metrics_RawContainerMetrics(in.SystemContainers[i], &out.SystemContainers[i], c); err != nil {
				return err
			}
		}
	} else {
		out.SystemContainers = nil
	}
	return nil
}

func deepCopy_metrics_RawNodeMetricsList(in RawNodeMetricsList, out *RawNodeMetricsList, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]RawNodeMetrics, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_metrics_RawNodeMetrics(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_metrics_RawPodMetrics(in RawPodMetrics, out *RawPodMetrics, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if err := deepCopy_metrics_NonLocalObjectReference(in.PodRef, &out.PodRef, c); err != nil {
		return err
	}
	if in.Containers != nil {
		out.Containers = make([]RawContainerMetrics, len(in.Containers))
		for i := range in.Containers {
			if err := deepCopy_metrics_RawContainerMetrics(in.Containers[i], &out.Containers[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Containers = nil
	}
	if in.Samples != nil {
		out.Samples = make([]PodSample, len(in.Samples))
		for i := range in.Samples {
			if err := deepCopy_metrics_PodSample(in.Samples[i], &out.Samples[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Samples = nil
	}
	return nil
}

func deepCopy_metrics_RawPodMetricsList(in RawPodMetricsList, out *RawPodMetricsList, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := deepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		out.Items = make([]RawPodMetrics, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_metrics_RawPodMetrics(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_metrics_Sample(in Sample, out *Sample, c *conversion.Cloner) error {
	if err := deepCopy_unversioned_Time(in.SampleTime, &out.SampleTime, c); err != nil {
		return err
	}
	return nil
}

func init() {
	err := api.Scheme.AddGeneratedDeepCopyFuncs(
		deepCopy_resource_Quantity,
		deepCopy_unversioned_ListMeta,
		deepCopy_unversioned_Time,
		deepCopy_unversioned_TypeMeta,
		deepCopy_metrics_AggregateSample,
		deepCopy_metrics_CPUMetrics,
		deepCopy_metrics_ContainerSample,
		deepCopy_metrics_FilesystemMetrics,
		deepCopy_metrics_MemoryMetrics,
		deepCopy_metrics_MetricsMeta,
		deepCopy_metrics_NetworkMetrics,
		deepCopy_metrics_NonLocalObjectReference,
		deepCopy_metrics_PodSample,
		deepCopy_metrics_RawContainerMetrics,
		deepCopy_metrics_RawMetricsOptions,
		deepCopy_metrics_RawNodeMetrics,
		deepCopy_metrics_RawNodeMetricsList,
		deepCopy_metrics_RawPodMetrics,
		deepCopy_metrics_RawPodMetricsList,
		deepCopy_metrics_Sample,
	)
	if err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}
