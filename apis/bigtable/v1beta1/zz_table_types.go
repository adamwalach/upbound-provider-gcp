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

type ColumnFamilyInitParameters struct {

	// The name of the column family.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`
}

type ColumnFamilyObservation struct {

	// The name of the column family.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`
}

type ColumnFamilyParameters struct {

	// The name of the column family.
	// +kubebuilder:validation:Optional
	Family *string `json:"family" tf:"family,omitempty"`
}

type TableInitParameters struct {

	// Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.
	ChangeStreamRetention *string `json:"changeStreamRetention,omitempty" tf:"change_stream_retention,omitempty"`

	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamily []ColumnFamilyInitParameters `json:"columnFamily,omitempty" tf:"column_family,omitempty"`

	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A list of predefined keys to split the table on.
	SplitKeys []*string `json:"splitKeys,omitempty" tf:"split_keys,omitempty"`
}

type TableObservation struct {

	// Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.
	ChangeStreamRetention *string `json:"changeStreamRetention,omitempty" tf:"change_stream_retention,omitempty"`

	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamily []ColumnFamilyObservation `json:"columnFamily,omitempty" tf:"column_family,omitempty"`

	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// an identifier for the resource with format projects/{{project}}/instances/{{instance_name}}/tables/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Bigtable instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A list of predefined keys to split the table on.
	SplitKeys []*string `json:"splitKeys,omitempty" tf:"split_keys,omitempty"`
}

type TableParameters struct {

	// Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.
	// +kubebuilder:validation:Optional
	ChangeStreamRetention *string `json:"changeStreamRetention,omitempty" tf:"change_stream_retention,omitempty"`

	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	// +kubebuilder:validation:Optional
	ColumnFamily []ColumnFamilyParameters `json:"columnFamily,omitempty" tf:"column_family,omitempty"`

	// A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
	// +kubebuilder:validation:Optional
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The name of the Bigtable instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigtable/v1beta1.Instance
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in bigtable to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in bigtable to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A list of predefined keys to split the table on.
	// +kubebuilder:validation:Optional
	SplitKeys []*string `json:"splitKeys,omitempty" tf:"split_keys,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
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
	InitProvider TableInitParameters `json:"initProvider,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Table is the Schema for the Tables API. Creates a Google Cloud Bigtable table inside an instance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec"`
	Status            TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
