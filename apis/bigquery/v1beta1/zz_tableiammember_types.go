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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TableIAMMemberConditionObservation struct {
}

type TableIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type TableIAMMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TableIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TableIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	TableID *string `json:"tableId" tf:"table_id,omitempty"`
}

// TableIAMMemberSpec defines the desired state of TableIAMMember
type TableIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableIAMMemberParameters `json:"forProvider"`
}

// TableIAMMemberStatus defines the observed state of TableIAMMember.
type TableIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMMember is the Schema for the TableIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TableIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableIAMMemberSpec   `json:"spec"`
	Status            TableIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMMemberList contains a list of TableIAMMembers
type TableIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableIAMMember `json:"items"`
}

// Repository type metadata.
var (
	TableIAMMember_Kind             = "TableIAMMember"
	TableIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableIAMMember_Kind}.String()
	TableIAMMember_KindAPIVersion   = TableIAMMember_Kind + "." + CRDGroupVersion.String()
	TableIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(TableIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&TableIAMMember{}, &TableIAMMemberList{})
}
