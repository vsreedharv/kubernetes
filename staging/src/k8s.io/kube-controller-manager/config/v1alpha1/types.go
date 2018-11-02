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

package v1alpha1

import (
	apimachineryconfigv1alpha1 "k8s.io/apimachinery/pkg/apis/config/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiserverconfigv1alpha1 "k8s.io/apiserver/pkg/apis/config/v1alpha1"
)

// PersistentVolumeRecyclerConfiguration contains elements describing persistent volume plugins.
type PersistentVolumeRecyclerConfiguration struct {
	// MaximumRetry is number of retries the PV recycler will execute on failure to recycle
	// PV.
	MaximumRetry int32 `json:"-"`
	// MinimumTimeoutNFS is the minimum ActiveDeadlineSeconds to use for an NFS Recycler
	// pod.
	MinimumTimeoutNFS int32 `json:"-"`
	// PodTemplateFilePathNFS is the file path to a pod definition used as a template for
	// NFS persistent volume recycling
	PodTemplateFilePathNFS string `json:"-"`
	// IncrementTimeoutNFS is the increment of time added per Gi to ActiveDeadlineSeconds
	// for an NFS scrubber pod.
	IncrementTimeoutNFS int32 `json:"-"`
	// PodTemplateFilePathHostPath is the file path to a pod definition used as a template for
	// HostPath persistent volume recycling. This is for development and testing only and
	// will not work in a multi-node cluster.
	PodTemplateFilePathHostPath string `json:"-"`
	// MinimumTimeoutHostPath is the minimum ActiveDeadlineSeconds to use for a HostPath
	// Recycler pod.  This is for development and testing only and will not work in a multi-node
	// cluster.
	MinimumTimeoutHostPath int32 `json:"-"`
	// IncrementTimeoutHostPath is the increment of time added per Gi to ActiveDeadlineSeconds
	// for a HostPath scrubber pod.  This is for development and testing only and will not work
	// in a multi-node cluster.
	IncrementTimeoutHostPath int32 `json:"-"`
}

// VolumeConfiguration contains *all* enumerated flags meant to configure all volume
// plugins. From this config, the controller-manager binary will create many instances of
// volume.VolumeConfig, each containing only the configuration needed for that plugin which
// are then passed to the appropriate plugin. The ControllerManager binary is the only part
// of the code which knows what plugins are supported and which flags correspond to each plugin.
type VolumeConfiguration struct {
	// EnableHostPathProvisioning enables HostPath PV provisioning when running without a
	// cloud provider. This allows testing and development of provisioning features. HostPath
	// provisioning is not supported in any way, won't work in a multi-node cluster, and
	// should not be used for anything other than testing or development.
	EnableHostPathProvisioning *bool `json:"-"`
	// EnableDynamicProvisioning enables the provisioning of volumes when running within an environment
	// that supports dynamic provisioning. Defaults to true.
	EnableDynamicProvisioning *bool `json:"-"`
	// PersistentVolumeRecyclerConfiguration holds configuration for persistent volume plugins.
	PersistentVolumeRecyclerConfiguration PersistentVolumeRecyclerConfiguration `json:"-"`
	// VolumePluginDir is the full path of the directory in which the flex
	// volume plugin should search for additional third party volume plugins
	FlexVolumePluginDir string `json:"-"`
}

// GroupResource describes an group resource.
type GroupResource struct {
	// Group is the group portion of the GroupResource.
	Group string `json:"-"`
	// Resource is the resource portion of the GroupResource.
	Resource string `json:"-"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeControllerManagerConfiguration contains elements describing kube-controller manager.
type KubeControllerManagerConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// Generic holds configuration for a generic controller-manager
	Generic GenericControllerManagerConfiguration `json:"-"`
	// KubeCloudShared holds configuration for shared related features
	// both in cloud controller manager and kube-controller manager.
	KubeCloudShared KubeCloudSharedConfiguration `json:"-"`
	// AttachDetachControllerConfiguration holds configuration for
	// AttachDetachController related features.
	AttachDetachController AttachDetachControllerConfiguration `json:"-"`
	// CSRSigningControllerConfiguration holds configuration for
	// CSRSigningController related features.
	CSRSigningController CSRSigningControllerConfiguration `json:"-"`
	// DaemonSetControllerConfiguration holds configuration for DaemonSetController
	// related features.
	DaemonSetController DaemonSetControllerConfiguration `json:"-"`
	// DeploymentControllerConfiguration holds configuration for
	// DeploymentController related features.
	DeploymentController DeploymentControllerConfiguration `json:"-"`
	// DeprecatedControllerConfiguration holds configuration for some deprecated
	// features.
	DeprecatedController DeprecatedControllerConfiguration `json:"-"`
	// EndpointControllerConfiguration holds configuration for EndpointController
	// related features.
	EndpointController EndpointControllerConfiguration `json:"-"`
	// GarbageCollectorControllerConfiguration holds configuration for
	// GarbageCollectorController related features.
	GarbageCollectorController GarbageCollectorControllerConfiguration `json:"-"`
	// HPAControllerConfiguration holds configuration for HPAController related features.
	HPAController HPAControllerConfiguration `json:"-"`
	// JobControllerConfiguration holds configuration for JobController related features.
	JobController JobControllerConfiguration `json:"-"`
	// NamespaceControllerConfiguration holds configuration for NamespaceController
	// related features.
	NamespaceController NamespaceControllerConfiguration `json:"-"`
	// NodeIPAMControllerConfiguration holds configuration for NodeIPAMController
	// related features.
	NodeIPAMController NodeIPAMControllerConfiguration `json:"-"`
	// NodeLifecycleControllerConfiguration holds configuration for
	// NodeLifecycleController related features.
	NodeLifecycleController NodeLifecycleControllerConfiguration `json:"-"`
	// PersistentVolumeBinderControllerConfiguration holds configuration for
	// PersistentVolumeBinderController related features.
	PersistentVolumeBinderController PersistentVolumeBinderControllerConfiguration `json:"-"`
	// PodGCControllerConfiguration holds configuration for PodGCController
	// related features.
	PodGCController PodGCControllerConfiguration `json:"-"`
	// ReplicaSetControllerConfiguration holds configuration for ReplicaSet related features.
	ReplicaSetController ReplicaSetControllerConfiguration `json:"-"`
	// ReplicationControllerConfiguration holds configuration for
	// ReplicationController related features.
	ReplicationController ReplicationControllerConfiguration `json:"-"`
	// ResourceQuotaControllerConfiguration holds configuration for
	// ResourceQuotaController related features.
	ResourceQuotaController ResourceQuotaControllerConfiguration `json:"-"`
	// SAControllerConfiguration holds configuration for ServiceAccountController
	// related features.
	SAController SAControllerConfiguration `json:"-"`
	// ServiceControllerConfiguration holds configuration for ServiceController
	// related features.
	ServiceController ServiceControllerConfiguration `json:"-"`
	// TTLAfterFinishedControllerConfiguration holds configuration for
	// TTLAfterFinishedController related features.
	TTLAfterFinishedController TTLAfterFinishedControllerConfiguration `json:"-"`
}

// GenericControllerManagerConfiguration holds configuration for a generic controller-manager.
type GenericControllerManagerConfiguration struct {
	// Port is the port that the controller-manager's http service runs on.
	Port int32 `json:"-"`
	// Address is the IP address to serve on (set to 0.0.0.0 for all interfaces).
	Address string `json:"-"`
	// MinResyncPeriod is the resync period in reflectors; will be random between
	// minResyncPeriod and 2*minResyncPeriod.
	MinResyncPeriod metav1.Duration `json:"-"`
	// ClientConnection specifies the kubeconfig file and client connection
	// settings for the proxy server to use when communicating with the apiserver.
	ClientConnection apimachineryconfigv1alpha1.ClientConnectionConfiguration `json:"-"`
	// ControllerStartInterval specifies how long to wait between starting
	// controller managers
	ControllerStartInterval metav1.Duration `json:"-"`
	// LeaderElection defines the configuration of leader election client.
	LeaderElection apiserverconfigv1alpha1.LeaderElectionConfiguration `json:"-"`
	// Controllers is the list of controllers to enable or disable
	// '*' means "all enabled by default controllers"
	// 'foo' means "enable 'foo'"
	// '-foo' means "disable 'foo'"
	// first item for a particular name wins
	Controllers []string `json:"-"`
	// DebuggingConfiguration holds configuration for Debugging related features.
	Debugging apiserverconfigv1alpha1.DebuggingConfiguration `json:"-"`
}

// KubeCloudSharedConfiguration contains elements shared by both kube-controller manager
// and cloud-controller manager, but not genericconfig.
type KubeCloudSharedConfiguration struct {
	// CloudProviderConfiguration holds configuration for CloudProvider related features.
	CloudProvider CloudProviderConfiguration `json:"-"`
	// ExternalCloudVolumePlugin specifies the plugin to use when cloudProvider is "external".
	// It is currently used by the in repo cloud providers to handle node and volume control in the KCM.
	ExternalCloudVolumePlugin string `json:"-"`
	// UseServiceAccountCredentials indicates whether controllers should be run with
	// individual service account credentials.
	UseServiceAccountCredentials bool `json:"-"`
	// AllowUntaggedCloud indicates whether to run with untagged cloud instances
	AllowUntaggedCloud bool `json:"-"`
	// RouteReconciliationPeriod is the period for reconciling routes created for Nodes by cloud provider..
	RouteReconciliationPeriod metav1.Duration `json:"-"`
	// NodeMonitorPeriod is the period for syncing NodeStatus in NodeController.
	NodeMonitorPeriod metav1.Duration `json:"-"`
	// ClusterName is the instance prefix for the cluster.
	ClusterName string `json:"-"`
	// ClusterCIDR is CIDR Range for Pods in cluster.
	ClusterCIDR string `json:"-"`
	// AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if
	// ConfigureCloudRoutes is true, to be set on the cloud provider.
	AllocateNodeCIDRs bool `json:"-"`
	// CIDRAllocatorType determines what kind of pod CIDR allocator will be used.
	CIDRAllocatorType string `json:"-"`
	// ConfigureCloudRoutes enables CIDRs allocated with allocateNodeCIDRs
	// to be configured on the cloud provider.
	ConfigureCloudRoutes *bool `json:"-"`
	// NodeSyncPeriod is the period for syncing nodes from cloudprovider. Longer
	// periods will result in fewer calls to cloud provider, but may delay addition
	// of new nodes to cluster.
	NodeSyncPeriod metav1.Duration `json:"-"`
}

// AttachDetachControllerConfiguration contains elements describing AttachDetachController.
type AttachDetachControllerConfiguration struct {
	// Reconciler runs a periodic loop to reconcile the desired state of the with
	// the actual state of the world by triggering attach detach operations.
	// This flag enables or disables reconcile.  Is false by default, and thus enabled.
	DisableAttachDetachReconcilerSync bool `json:"-"`
	// ReconcilerSyncLoopPeriod is the amount of time the reconciler sync states loop
	// wait between successive executions. Is set to 5 sec by default.
	ReconcilerSyncLoopPeriod metav1.Duration `json:"-"`
}

// CloudProviderConfiguration contains basically elements about cloud provider.
type CloudProviderConfiguration struct {
	// Name is the provider for cloud services.
	Name string `json:"-"`
	// CloudConfigFile is the path to the cloud provider configuration file.
	CloudConfigFile string `json:"-"`
}

// CSRSigningControllerConfiguration contains elements describing CSRSigningController.
type CSRSigningControllerConfiguration struct {
	// ClusterSigningCertFile is the filename containing a PEM-encoded
	// X509 CA certificate used to issue cluster-scoped certificates
	ClusterSigningCertFile string `json:"-"`
	// ClusterSigningCertFile is the filename containing a PEM-encoded
	// RSA or ECDSA private key used to issue cluster-scoped certificates
	ClusterSigningKeyFile string `json:"-"`
	// ClusterSigningDuration is the length of duration signed certificates
	// will be given.
	ClusterSigningDuration metav1.Duration `json:"-"`
}

// DaemonSetControllerConfiguration contains elements describing DaemonSetController.
type DaemonSetControllerConfiguration struct {
	// ConcurrentDaemonSetSyncs is the number of daemonset objects that are
	// allowed to sync concurrently. Larger number = more responsive daemonset,
	// but more CPU (and network) load.
	ConcurrentDaemonSetSyncs int32 `json:"-"`
}

// DeploymentControllerConfiguration contains elements describing DeploymentController.
type DeploymentControllerConfiguration struct {
	// ConcurrentDeploymentSyncs is the number of deployment objects that are
	// allowed to sync concurrently. Larger number = more responsive deployments,
	// but more CPU (and network) load.
	ConcurrentDeploymentSyncs int32 `json:"-"`
	// DeploymentControllerSyncPeriod is the period for syncing the deployments.
	DeploymentControllerSyncPeriod metav1.Duration `json:"-"`
}

// DeprecatedControllerConfiguration contains elements be deprecated.
type DeprecatedControllerConfiguration struct {
	// DEPRECATED: DeletingPodsQps is the number of nodes per second on which pods are deleted in
	// case of node failure.
	DeletingPodsQPS float32 `json:"-"`
	// DEPRECATED: DeletingPodsBurst is the number of nodes on which pods are bursty deleted in
	// case of node failure. For more details look into RateLimiter.
	DeletingPodsBurst int32 `json:"-"`
	// RegisterRetryCount is the number of retries for initial node registration.
	// Retry interval equals node-sync-period.
	RegisterRetryCount int32 `json:"-"`
}

// EndpointControllerConfiguration contains elements describing EndpointController.
type EndpointControllerConfiguration struct {
	// ConcurrentEndpointSyncs is the number of endpoint syncing operations
	// that will be done concurrently. Larger number = faster endpoint updating,
	// but more CPU (and network) load.
	ConcurrentEndpointSyncs int32 `json:"-"`
}

// GarbageCollectorControllerConfiguration contains elements describing GarbageCollectorController.
type GarbageCollectorControllerConfiguration struct {
	// EnableGarbageCollector enables the generic garbage collector. MUST be
	// synced with the corresponding flag of the kube-apiserver. WARNING: the
	// generic garbage collector is an alpha feature.
	EnableGarbageCollector *bool `json:"-"`
	// ConcurrentGCSyncs is the number of garbage collector workers that are
	// allowed to sync concurrently.
	ConcurrentGCSyncs int32 `json:"-"`
	// GCIgnoredResources is the list of GroupResources that garbage collection should ignore.
	GCIgnoredResources []GroupResource `json:"-"`
}

// HPAControllerConfiguration contains elements describing HPAController.
type HPAControllerConfiguration struct {
	// HorizontalPodAutoscalerSyncPeriod is the period for syncing the number of
	// pods in horizontal pod autoscaler.
	HorizontalPodAutoscalerSyncPeriod metav1.Duration `json:"-"`
	// HorizontalPodAutoscalerUpscaleForbiddenWindow is a period after which next upscale allowed.
	HorizontalPodAutoscalerUpscaleForbiddenWindow metav1.Duration `json:"-"`
	// HorizontalPodAutoscalerDowncaleStabilizationWindow is a period for which autoscaler will look
	// backwards and not scale down below any recommendation it made during that period.
	HorizontalPodAutoscalerDownscaleStabilizationWindow metav1.Duration `json:"-"`
	// HorizontalPodAutoscalerDownscaleForbiddenWindow is a period after which next downscale allowed.
	HorizontalPodAutoscalerDownscaleForbiddenWindow metav1.Duration `json:"-"`
	// HorizontalPodAutoscalerTolerance is the tolerance for when
	// resource usage suggests upscaling/downscaling
	HorizontalPodAutoscalerTolerance float64 `json:"-"`
	// HorizontalPodAutoscalerUseRESTClients causes the HPA controller to use REST clients
	// through the kube-aggregator when enabled, instead of using the legacy metrics client
	// through the API server proxy.
	HorizontalPodAutoscalerUseRESTClients *bool `json:"-"`
	// HorizontalPodAutoscalerCPUInitializationPeriod is the period after pod start when CPU samples
	// might be skipped.
	HorizontalPodAutoscalerCPUInitializationPeriod metav1.Duration `json:"-"`
	// HorizontalPodAutoscalerInitialReadinessDelay is period after pod start during which readiness
	// changes are treated as readiness being set for the first time. The only effect of this is that
	// HPA will disregard CPU samples from unready pods that had last readiness change during that
	// period.
	HorizontalPodAutoscalerInitialReadinessDelay metav1.Duration `json:"-"`
}

// JobControllerConfiguration contains elements describing JobController.
type JobControllerConfiguration struct {
	// ConcurrentJobSyncs is the number of job objects that are
	// allowed to sync concurrently. Larger number = more responsive jobs,
	// but more CPU (and network) load.
	ConcurrentJobSyncs int32 `json:"-"`
}

// NamespaceControllerConfiguration contains elements describing NamespaceController.
type NamespaceControllerConfiguration struct {
	// NamespaceSyncPeriod is the period for syncing namespace life-cycle
	// updates.
	NamespaceSyncPeriod metav1.Duration `json:"-"`
	// ConcurrentNamespaceSyncs is the number of namespace objects that are
	// allowed to sync concurrently.
	ConcurrentNamespaceSyncs int32 `json:"-"`
}

// NodeIPAMControllerConfiguration contains elements describing NodeIpamController.
type NodeIPAMControllerConfiguration struct {
	// ServiceCIDR is CIDR Range for Services in cluster.
	ServiceCIDR string `json:"-"`
	// NodeCIDRMaskSize is the mask size for node cidr in cluster.
	NodeCIDRMaskSize int32 `json:"-"`
}

// NodeLifecycleControllerConfiguration contains elements describing NodeLifecycleController.
type NodeLifecycleControllerConfiguration struct {
	// EnableTaintManager enables NoExecute Taints and will evict all not-tolerating
	// Pod running on Nodes tainted with this kind of Taints.
	EnableTaintManager *bool `json:"-"`
	// NodeEvictionRate is the number of nodes per second on which pods are deleted in case of node failure when a zone is healthy
	NodeEvictionRate float32 `json:"-"`
	// SecondaryNodeEvictionRate is the number of nodes per second on which pods are deleted in case of node failure when a zone is unhealthy
	SecondaryNodeEvictionRate float32 `json:"-"`
	// NodeStartupGracePeriod is the amount of time which we allow starting a node to
	// be unresponsive before marking it unhealthy.
	NodeStartupGracePeriod metav1.Duration `json:"-"`
	// NodeMonitorGracePeriod is the amount of time which we allow a running node to be
	// unresponsive before marking it unhealthy. Must be N times more than kubelet's
	// nodeStatusUpdateFrequency, where N means number of retries allowed for kubelet
	// to post node status.
	NodeMonitorGracePeriod metav1.Duration `json:"-"`
	// PodEvictionTimeout is the grace period for deleting pods on failed nodes.
	PodEvictionTimeout metav1.Duration `json:"-"`
	// LargeClusterSizeThreshold is the number below which the secondaryNodeEvictionRate is implicitly overridden to 0
	LargeClusterSizeThreshold int32 `json:"-"`
	// Zone is treated as unhealthy in nodeEvictionRate and secondaryNodeEvictionRate when at least
	// unhealthyZoneThreshold (no less than 3) of Nodes in the zone are NotReady
	UnhealthyZoneThreshold float32 `json:"-"`
}

// PersistentVolumeBinderControllerConfiguration contains elements describing
// PersistentVolumeBinderController.
type PersistentVolumeBinderControllerConfiguration struct {
	// PVClaimBinderSyncPeriod is the period for syncing persistent volumes
	// and persistent volume claims.
	PVClaimBinderSyncPeriod metav1.Duration `json:"-"`
	// VolumeConfiguration holds configuration for volume related features.
	VolumeConfiguration VolumeConfiguration `json:"-"`
}

// PodGCControllerConfiguration contains elements describing PodGCController.
type PodGCControllerConfiguration struct {
	// TerminatedPodGCThreshold is the number of terminated pods that can exist
	// before the terminated pod garbage collector starts deleting terminated pods.
	// If <= 0, the terminated pod garbage collector is disabled.
	TerminatedPodGCThreshold int32 `json:"-"`
}

// ReplicaSetControllerConfiguration contains elements describing ReplicaSetController.
type ReplicaSetControllerConfiguration struct {
	// ConcurrentRSSyncs is the number of replica sets that are  allowed to sync
	// concurrently. Larger number = more responsive replica  management, but more
	// CPU (and network) load.
	ConcurrentRSSyncs int32 `json:"-"`
}

// ReplicationControllerConfiguration contains elements describing ReplicationController.
type ReplicationControllerConfiguration struct {
	// ConcurrentRCSyncs is the number of replication controllers that are
	// allowed to sync concurrently. Larger number = more responsive replica
	// management, but more CPU (and network) load.
	ConcurrentRCSyncs int32 `json:"-"`
}

// ResourceQuotaControllerConfiguration contains elements describing ResourceQuotaController.
type ResourceQuotaControllerConfiguration struct {
	// ResourceQuotaSyncPeriod is the period for syncing quota usage status
	// in the system.
	ResourceQuotaSyncPeriod metav1.Duration `json:"-"`
	// ConcurrentResourceQuotaSyncs is the number of resource quotas that are
	// allowed to sync concurrently. Larger number = more responsive quota
	// management, but more CPU (and network) load.
	ConcurrentResourceQuotaSyncs int32 `json:"-"`
}

// SAControllerConfiguration contains elements describing ServiceAccountController.
type SAControllerConfiguration struct {
	// ServiceAccountKeyFile is the filename containing a PEM-encoded private RSA key
	// used to sign service account tokens.
	ServiceAccountKeyFile string `json:"-"`
	// ConcurrentSATokenSyncs is the number of service account token syncing operations
	// that will be done concurrently.
	ConcurrentSATokenSyncs int32 `json:"-"`
	// RootCAFile is the root certificate authority will be included in service
	// account's token secret. This must be a valid PEM-encoded CA bundle.
	RootCAFile string `json:"-"`
}

// ServiceControllerConfiguration contains elements describing ServiceController.
type ServiceControllerConfiguration struct {
	// ConcurrentServiceSyncs is the number of services that are
	// allowed to sync concurrently. Larger number = more responsive service
	// management, but more CPU (and network) load.
	ConcurrentServiceSyncs int32 `json:"-"`
}

// TTLAfterFinishedControllerConfiguration contains elements describing TTLAfterFinishedController.
type TTLAfterFinishedControllerConfiguration struct {
	// ConcurrentTTLSyncs is the number of TTL-after-finished collector workers that are
	// allowed to sync concurrently.
	ConcurrentTTLSyncs int32 `json:"-"`
}
