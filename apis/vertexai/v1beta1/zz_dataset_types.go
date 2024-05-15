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

type DatasetInitParameters struct {

	// The user-defined name of the Dataset. The name can be up to 128 characters long and can be consist of any UTF-8 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	// Structure is documented below.
	EncryptionSpec []EncryptionSpecInitParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// A set of key/value label pairs to assign to this Workflow.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	MetadataSchemaURI *string `json:"metadataSchemaUri,omitempty" tf:"metadata_schema_uri,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the dataset. eg us-central1
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DatasetObservation struct {

	// The timestamp of when the dataset was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The user-defined name of the Dataset. The name can be up to 128 characters long and can be consist of any UTF-8 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// for all of the labels present on the resource.
	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	// Structure is documented below.
	EncryptionSpec []EncryptionSpecObservation `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to this Workflow.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	MetadataSchemaURI *string `json:"metadataSchemaUri,omitempty" tf:"metadata_schema_uri,omitempty"`

	// The resource name of the Dataset. This value is set by Google.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the dataset. eg us-central1
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// The timestamp of when the dataset was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type DatasetParameters struct {

	// The user-defined name of the Dataset. The name can be up to 128 characters long and can be consist of any UTF-8 characters.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionSpec []EncryptionSpecParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// A set of key/value label pairs to assign to this Workflow.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	// +kubebuilder:validation:Optional
	MetadataSchemaURI *string `json:"metadataSchemaUri,omitempty" tf:"metadata_schema_uri,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the dataset. eg us-central1
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EncryptionSpecInitParameters struct {

	// Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type EncryptionSpecObservation struct {

	// Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type EncryptionSpecParameters struct {

	// Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetParameters `json:"forProvider"`
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
	InitProvider DatasetInitParameters `json:"initProvider,omitempty"`
}

// DatasetStatus defines the observed state of Dataset.
type DatasetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Dataset is the Schema for the Datasets API. A collection of DataItems and Annotations on them.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metadataSchemaUri) || (has(self.initProvider) && has(self.initProvider.metadataSchemaUri))",message="spec.forProvider.metadataSchemaUri is a required parameter"
	Spec   DatasetSpec   `json:"spec"`
	Status DatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetList contains a list of Datasets
type DatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataset `json:"items"`
}

// Repository type metadata.
var (
	Dataset_Kind             = "Dataset"
	Dataset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dataset_Kind}.String()
	Dataset_KindAPIVersion   = Dataset_Kind + "." + CRDGroupVersion.String()
	Dataset_GroupVersionKind = CRDGroupVersion.WithKind(Dataset_Kind)
)

func init() {
	SchemeBuilder.Register(&Dataset{}, &DatasetList{})
}
