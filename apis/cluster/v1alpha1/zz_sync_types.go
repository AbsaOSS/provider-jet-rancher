/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NodesObservation struct {
	Capacity map[string]string `json:"capacity,omitempty" tf:"capacity,omitempty"`

	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ExternalIPAddress *string `json:"externalIpAddress,omitempty" tf:"external_ip_address,omitempty"`

	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NodePoolID *string `json:"nodePoolId,omitempty" tf:"node_pool_id,omitempty"`

	NodeTemplateID *string `json:"nodeTemplateId,omitempty" tf:"node_template_id,omitempty"`

	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	RequestedHostname *string `json:"requestedHostname,omitempty" tf:"requested_hostname,omitempty"`

	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	SystemInfo map[string]string `json:"systemInfo,omitempty" tf:"system_info,omitempty"`
}

type NodesParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type SyncObservation struct {
	DefaultProjectID *string `json:"defaultProjectId,omitempty" tf:"default_project_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	SystemProjectID *string `json:"systemProjectId,omitempty" tf:"system_project_id,omitempty"`
}

type SyncParameters struct {

	// Cluster id to sync
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Cluster node pool ids
	// +kubebuilder:validation:Optional
	NodePoolIds []*string `json:"nodePoolIds,omitempty" tf:"node_pool_ids,omitempty"`

	// Wait until active status is confirmed a number of times (wait interval of 5s)
	// +kubebuilder:validation:Optional
	StateConfirm *int64 `json:"stateConfirm,omitempty" tf:"state_confirm,omitempty"`

	// +kubebuilder:validation:Optional
	Synced *bool `json:"synced,omitempty" tf:"synced,omitempty"`

	// Wait until alerting is up and running
	// +kubebuilder:validation:Optional
	WaitAlerting *bool `json:"waitAlerting,omitempty" tf:"wait_alerting,omitempty"`

	// Wait until all catalogs are downloaded and active
	// +kubebuilder:validation:Optional
	WaitCatalogs *bool `json:"waitCatalogs,omitempty" tf:"wait_catalogs,omitempty"`

	// Wait until monitoring is up and running
	// +kubebuilder:validation:Optional
	WaitMonitoring *bool `json:"waitMonitoring,omitempty" tf:"wait_monitoring,omitempty"`
}

// SyncSpec defines the desired state of Sync
type SyncSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SyncParameters `json:"forProvider"`
}

// SyncStatus defines the observed state of Sync.
type SyncStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SyncObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Sync is the Schema for the Syncs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type Sync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SyncSpec   `json:"spec"`
	Status            SyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SyncList contains a list of Syncs
type SyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sync `json:"items"`
}

// Repository type metadata.
var (
	Sync_Kind             = "Sync"
	Sync_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Sync_Kind}.String()
	Sync_KindAPIVersion   = Sync_Kind + "." + CRDGroupVersion.String()
	Sync_GroupVersionKind = CRDGroupVersion.WithKind(Sync_Kind)
)

func init() {
	SchemeBuilder.Register(&Sync{}, &SyncList{})
}
