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

type ConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type MembershipIAMMemberInitParameters struct {
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/gkehub/v1beta1.Membership
	MembershipID *string `json:"membershipId,omitempty" tf:"membership_id,omitempty"`

	// Reference to a Membership in gkehub to populate membershipId.
	// +kubebuilder:validation:Optional
	MembershipIDRef *v1.Reference `json:"membershipIdRef,omitempty" tf:"-"`

	// Selector for a Membership in gkehub to populate membershipId.
	// +kubebuilder:validation:Optional
	MembershipIDSelector *v1.Selector `json:"membershipIdSelector,omitempty" tf:"-"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type MembershipIAMMemberObservation struct {
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	MembershipID *string `json:"membershipId,omitempty" tf:"membership_id,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type MembershipIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/gkehub/v1beta1.Membership
	// +kubebuilder:validation:Optional
	MembershipID *string `json:"membershipId,omitempty" tf:"membership_id,omitempty"`

	// Reference to a Membership in gkehub to populate membershipId.
	// +kubebuilder:validation:Optional
	MembershipIDRef *v1.Reference `json:"membershipIdRef,omitempty" tf:"-"`

	// Selector for a Membership in gkehub to populate membershipId.
	// +kubebuilder:validation:Optional
	MembershipIDSelector *v1.Selector `json:"membershipIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// MembershipIAMMemberSpec defines the desired state of MembershipIAMMember
type MembershipIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MembershipIAMMemberParameters `json:"forProvider"`
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
	InitProvider MembershipIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// MembershipIAMMemberStatus defines the observed state of MembershipIAMMember.
type MembershipIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MembershipIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MembershipIAMMember is the Schema for the MembershipIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type MembershipIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || (has(self.initProvider) && has(self.initProvider.member))",message="spec.forProvider.member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   MembershipIAMMemberSpec   `json:"spec"`
	Status MembershipIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MembershipIAMMemberList contains a list of MembershipIAMMembers
type MembershipIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MembershipIAMMember `json:"items"`
}

// Repository type metadata.
var (
	MembershipIAMMember_Kind             = "MembershipIAMMember"
	MembershipIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MembershipIAMMember_Kind}.String()
	MembershipIAMMember_KindAPIVersion   = MembershipIAMMember_Kind + "." + CRDGroupVersion.String()
	MembershipIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(MembershipIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&MembershipIAMMember{}, &MembershipIAMMemberList{})
}
