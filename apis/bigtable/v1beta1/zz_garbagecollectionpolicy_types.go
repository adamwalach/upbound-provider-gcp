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

type GarbageCollectionPolicyInitParameters struct {

	// The name of the column family.
	ColumnFamily *string `json:"columnFamily,omitempty" tf:"column_family,omitempty"`

	// The deletion policy for the GC policy.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted in a replicated instance.
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// Serialized JSON object to represent a more complex GC policy. Conflicts with mode, max_age and max_version. Conflicts with mode, max_age and max_version.
	GcRules *string `json:"gcRules,omitempty" tf:"gc_rules,omitempty"`

	// The name of the Bigtable instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigtable/v1beta1.Instance
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Reference to a Instance in bigtable to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameRef *v1.Reference `json:"instanceNameRef,omitempty" tf:"-"`

	// Selector for a Instance in bigtable to populate instanceName.
	// +kubebuilder:validation:Optional
	InstanceNameSelector *v1.Selector `json:"instanceNameSelector,omitempty" tf:"-"`

	// GC policy that applies to all cells older than the given age.
	MaxAge []MaxAgeInitParameters `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersion []MaxVersionInitParameters `json:"maxVersion,omitempty" tf:"max_version,omitempty"`

	// If multiple policies are set, you should choose between UNION OR INTERSECTION.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigtable/v1beta1.Table
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Reference to a Table in bigtable to populate table.
	// +kubebuilder:validation:Optional
	TableRef *v1.Reference `json:"tableRef,omitempty" tf:"-"`

	// Selector for a Table in bigtable to populate table.
	// +kubebuilder:validation:Optional
	TableSelector *v1.Selector `json:"tableSelector,omitempty" tf:"-"`
}

type GarbageCollectionPolicyObservation struct {

	// The name of the column family.
	ColumnFamily *string `json:"columnFamily,omitempty" tf:"column_family,omitempty"`

	// The deletion policy for the GC policy.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted in a replicated instance.
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// Serialized JSON object to represent a more complex GC policy. Conflicts with mode, max_age and max_version. Conflicts with mode, max_age and max_version.
	GcRules *string `json:"gcRules,omitempty" tf:"gc_rules,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Bigtable instance.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// GC policy that applies to all cells older than the given age.
	MaxAge []MaxAgeObservation `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// GC policy that applies to all versions of a cell except for the most recent.
	MaxVersion []MaxVersionObservation `json:"maxVersion,omitempty" tf:"max_version,omitempty"`

	// If multiple policies are set, you should choose between UNION OR INTERSECTION.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the table.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`
}

type GarbageCollectionPolicyParameters struct {

	// The name of the column family.
	// +kubebuilder:validation:Optional
	ColumnFamily *string `json:"columnFamily,omitempty" tf:"column_family,omitempty"`

	// The deletion policy for the GC policy.
	// Setting ABANDON allows the resource to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted in a replicated instance.
	// +kubebuilder:validation:Optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty" tf:"deletion_policy,omitempty"`

	// Serialized JSON object to represent a more complex GC policy. Conflicts with mode, max_age and max_version. Conflicts with mode, max_age and max_version.
	// +kubebuilder:validation:Optional
	GcRules *string `json:"gcRules,omitempty" tf:"gc_rules,omitempty"`

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

	// GC policy that applies to all cells older than the given age.
	// +kubebuilder:validation:Optional
	MaxAge []MaxAgeParameters `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// GC policy that applies to all versions of a cell except for the most recent.
	// +kubebuilder:validation:Optional
	MaxVersion []MaxVersionParameters `json:"maxVersion,omitempty" tf:"max_version,omitempty"`

	// If multiple policies are set, you should choose between UNION OR INTERSECTION.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigtable/v1beta1.Table
	// +kubebuilder:validation:Optional
	Table *string `json:"table,omitempty" tf:"table,omitempty"`

	// Reference to a Table in bigtable to populate table.
	// +kubebuilder:validation:Optional
	TableRef *v1.Reference `json:"tableRef,omitempty" tf:"-"`

	// Selector for a Table in bigtable to populate table.
	// +kubebuilder:validation:Optional
	TableSelector *v1.Selector `json:"tableSelector,omitempty" tf:"-"`
}

type MaxAgeInitParameters struct {

	// Number of days before applying GC policy.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Duration before applying GC policy (ex. "8h"). This is required when days isn't set
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type MaxAgeObservation struct {

	// Number of days before applying GC policy.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Duration before applying GC policy (ex. "8h"). This is required when days isn't set
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type MaxAgeParameters struct {

	// Number of days before applying GC policy.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Duration before applying GC policy (ex. "8h"). This is required when days isn't set
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type MaxVersionInitParameters struct {

	// Number of version before applying the GC policy.
	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`
}

type MaxVersionObservation struct {

	// Number of version before applying the GC policy.
	Number *float64 `json:"number,omitempty" tf:"number,omitempty"`
}

type MaxVersionParameters struct {

	// Number of version before applying the GC policy.
	// +kubebuilder:validation:Optional
	Number *float64 `json:"number" tf:"number,omitempty"`
}

// GarbageCollectionPolicySpec defines the desired state of GarbageCollectionPolicy
type GarbageCollectionPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GarbageCollectionPolicyParameters `json:"forProvider"`
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
	InitProvider GarbageCollectionPolicyInitParameters `json:"initProvider,omitempty"`
}

// GarbageCollectionPolicyStatus defines the observed state of GarbageCollectionPolicy.
type GarbageCollectionPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GarbageCollectionPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GarbageCollectionPolicy is the Schema for the GarbageCollectionPolicys API. Creates a Google Cloud Bigtable GC Policy inside a family.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type GarbageCollectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.columnFamily) || (has(self.initProvider) && has(self.initProvider.columnFamily))",message="spec.forProvider.columnFamily is a required parameter"
	Spec   GarbageCollectionPolicySpec   `json:"spec"`
	Status GarbageCollectionPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GarbageCollectionPolicyList contains a list of GarbageCollectionPolicys
type GarbageCollectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GarbageCollectionPolicy `json:"items"`
}

// Repository type metadata.
var (
	GarbageCollectionPolicy_Kind             = "GarbageCollectionPolicy"
	GarbageCollectionPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GarbageCollectionPolicy_Kind}.String()
	GarbageCollectionPolicy_KindAPIVersion   = GarbageCollectionPolicy_Kind + "." + CRDGroupVersion.String()
	GarbageCollectionPolicy_GroupVersionKind = CRDGroupVersion.WithKind(GarbageCollectionPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&GarbageCollectionPolicy{}, &GarbageCollectionPolicyList{})
}
