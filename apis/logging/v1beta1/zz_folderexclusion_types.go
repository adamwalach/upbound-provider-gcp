// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FolderExclusionInitParameters struct {

	// A human-readable description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`
}

type FolderExclusionObservation struct {

	// A human-readable description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// an identifier for the resource with format folders/{{folder}}/exclusions/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FolderExclusionParameters struct {

	// A human-readable description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See Advanced Log Filters for information on how to
	// write a filter.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
	// accepted.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Folder
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// Reference to a Folder in cloudplatform to populate folder.
	// +kubebuilder:validation:Optional
	FolderRef *v1.Reference `json:"folderRef,omitempty" tf:"-"`

	// Selector for a Folder in cloudplatform to populate folder.
	// +kubebuilder:validation:Optional
	FolderSelector *v1.Selector `json:"folderSelector,omitempty" tf:"-"`
}

// FolderExclusionSpec defines the desired state of FolderExclusion
type FolderExclusionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderExclusionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FolderExclusionInitParameters `json:"initProvider,omitempty"`
}

// FolderExclusionStatus defines the observed state of FolderExclusion.
type FolderExclusionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderExclusionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FolderExclusion is the Schema for the FolderExclusions API. Manages a folder-level logging exclusion.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type FolderExclusion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filter) || (has(self.initProvider) && has(self.initProvider.filter))",message="spec.forProvider.filter is a required parameter"
	Spec   FolderExclusionSpec   `json:"spec"`
	Status FolderExclusionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderExclusionList contains a list of FolderExclusions
type FolderExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderExclusion `json:"items"`
}

// Repository type metadata.
var (
	FolderExclusion_Kind             = "FolderExclusion"
	FolderExclusion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderExclusion_Kind}.String()
	FolderExclusion_KindAPIVersion   = FolderExclusion_Kind + "." + CRDGroupVersion.String()
	FolderExclusion_GroupVersionKind = CRDGroupVersion.WithKind(FolderExclusion_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderExclusion{}, &FolderExclusionList{})
}
