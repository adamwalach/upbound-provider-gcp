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

type ProjectBucketConfigCmekSettingsInitParameters struct {

	// The resource name for the configured Cloud KMS key.
	// KMS key name format:
	// 'projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]'
	// To enable CMEK for the bucket, set this field to a valid kmsKeyName for which the associated service account has the required cloudkms.cryptoKeyEncrypterDecrypter roles assigned for the key.
	// The Cloud KMS key used by the bucket can be updated by changing the kmsKeyName to a new valid key name. Encryption operations that are in progress will be completed with the key that was in use when they started. Decryption operations will be completed using the key that was used at the time of encryption unless access to that key has been revoked.
	// See Enabling CMEK for Logging Buckets for more information.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`
}

type ProjectBucketConfigCmekSettingsObservation struct {

	// The resource name for the configured Cloud KMS key.
	// KMS key name format:
	// 'projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]'
	// To enable CMEK for the bucket, set this field to a valid kmsKeyName for which the associated service account has the required cloudkms.cryptoKeyEncrypterDecrypter roles assigned for the key.
	// The Cloud KMS key used by the bucket can be updated by changing the kmsKeyName to a new valid key name. Encryption operations that are in progress will be completed with the key that was in use when they started. Decryption operations will be completed using the key that was used at the time of encryption unless access to that key has been revoked.
	// See Enabling CMEK for Logging Buckets for more information.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The CryptoKeyVersion resource name for the configured Cloud KMS key.
	// KMS key name format:
	// 'projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]/cryptoKeyVersions/[VERSION]'
	// For example:
	// "projects/my-project/locations/us-central1/keyRings/my-ring/cryptoKeys/my-key/cryptoKeyVersions/1"
	// This is a read-only field used to convey the specific configured CryptoKeyVersion of kms_key that has been configured. It will be populated in cases where the CMEK settings are bound to a single key version.
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty" tf:"kms_key_version_name,omitempty"`

	// The resource name of the CMEK settings.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The service account associated with a project for which CMEK will apply.
	// Before enabling CMEK for a logging bucket, you must first assign the cloudkms.cryptoKeyEncrypterDecrypter role to the service account associated with the project for which CMEK will apply. Use v2.getCmekSettings to obtain the service account ID.
	// See Enabling CMEK for Logging Buckets for more information.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

type ProjectBucketConfigCmekSettingsParameters struct {

	// The resource name for the configured Cloud KMS key.
	// KMS key name format:
	// 'projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]'
	// To enable CMEK for the bucket, set this field to a valid kmsKeyName for which the associated service account has the required cloudkms.cryptoKeyEncrypterDecrypter roles assigned for the key.
	// The Cloud KMS key used by the bucket can be updated by changing the kmsKeyName to a new valid key name. Encryption operations that are in progress will be completed with the key that was in use when they started. Decryption operations will be completed using the key that was used at the time of encryption unless access to that key has been revoked.
	// See Enabling CMEK for Logging Buckets for more information.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameRef *v1.Reference `json:"kmsKeyNameRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKeyName.
	// +kubebuilder:validation:Optional
	KMSKeyNameSelector *v1.Selector `json:"kmsKeyNameSelector,omitempty" tf:"-"`
}

type ProjectBucketConfigIndexConfigsInitParameters struct {

	// The LogEntry field path to index.
	// Note that some paths are automatically indexed, and other paths are not eligible for indexing. See indexing documentation for details.
	FieldPath *string `json:"fieldPath,omitempty" tf:"field_path,omitempty"`

	// The type of data in this index. Allowed types include INDEX_TYPE_UNSPECIFIED, INDEX_TYPE_STRING and INDEX_TYPE_INTEGER.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProjectBucketConfigIndexConfigsObservation struct {

	// The LogEntry field path to index.
	// Note that some paths are automatically indexed, and other paths are not eligible for indexing. See indexing documentation for details.
	FieldPath *string `json:"fieldPath,omitempty" tf:"field_path,omitempty"`

	// The type of data in this index. Allowed types include INDEX_TYPE_UNSPECIFIED, INDEX_TYPE_STRING and INDEX_TYPE_INTEGER.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProjectBucketConfigIndexConfigsParameters struct {

	// The LogEntry field path to index.
	// Note that some paths are automatically indexed, and other paths are not eligible for indexing. See indexing documentation for details.
	// +kubebuilder:validation:Optional
	FieldPath *string `json:"fieldPath" tf:"field_path,omitempty"`

	// The type of data in this index. Allowed types include INDEX_TYPE_UNSPECIFIED, INDEX_TYPE_STRING and INDEX_TYPE_INTEGER.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ProjectBucketConfigInitParameters struct {

	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	CmekSettings []ProjectBucketConfigCmekSettingsInitParameters `json:"cmekSettings,omitempty" tf:"cmek_settings,omitempty"`

	// Describes this bucket.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the Log Analytics page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics *bool `json:"enableAnalytics,omitempty" tf:"enable_analytics,omitempty"`

	// A list of indexed fields and related configuration data. Structure is documented below.
	IndexConfigs []ProjectBucketConfigIndexConfigsInitParameters `json:"indexConfigs,omitempty" tf:"index_configs,omitempty"`

	// Whether the bucket is locked. The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

type ProjectBucketConfigObservation struct {

	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	BucketID *string `json:"bucketId,omitempty" tf:"bucket_id,omitempty"`

	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	CmekSettings []ProjectBucketConfigCmekSettingsObservation `json:"cmekSettings,omitempty" tf:"cmek_settings,omitempty"`

	// Describes this bucket.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the Log Analytics page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics *bool `json:"enableAnalytics,omitempty" tf:"enable_analytics,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of indexed fields and related configuration data. Structure is documented below.
	IndexConfigs []ProjectBucketConfigIndexConfigsObservation `json:"indexConfigs,omitempty" tf:"index_configs,omitempty"`

	// The bucket's lifecycle such as active or deleted. See LifecycleState.
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state,omitempty"`

	// The location of the bucket.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether the bucket is locked. The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parent resource that contains the logging bucket.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

type ProjectBucketConfigParameters struct {

	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	// +kubebuilder:validation:Required
	BucketID *string `json:"bucketId" tf:"bucket_id,omitempty"`

	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	// +kubebuilder:validation:Optional
	CmekSettings []ProjectBucketConfigCmekSettingsParameters `json:"cmekSettings,omitempty" tf:"cmek_settings,omitempty"`

	// Describes this bucket.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the Log Analytics page using SQL queries. Cannot be disabled once enabled.
	// +kubebuilder:validation:Optional
	EnableAnalytics *bool `json:"enableAnalytics,omitempty" tf:"enable_analytics,omitempty"`

	// A list of indexed fields and related configuration data. Structure is documented below.
	// +kubebuilder:validation:Optional
	IndexConfigs []ProjectBucketConfigIndexConfigsParameters `json:"indexConfigs,omitempty" tf:"index_configs,omitempty"`

	// The location of the bucket.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Whether the bucket is locked. The retention period on a locked bucket cannot be changed. Locked buckets may only be deleted if they are empty.
	// +kubebuilder:validation:Optional
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`

	// The parent resource that contains the logging bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("project_id",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a Project in cloudplatform to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	// +kubebuilder:validation:Optional
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

// ProjectBucketConfigSpec defines the desired state of ProjectBucketConfig
type ProjectBucketConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectBucketConfigParameters `json:"forProvider"`
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
	InitProvider ProjectBucketConfigInitParameters `json:"initProvider,omitempty"`
}

// ProjectBucketConfigStatus defines the observed state of ProjectBucketConfig.
type ProjectBucketConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectBucketConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProjectBucketConfig is the Schema for the ProjectBucketConfigs API. Manages a project-level logging bucket config.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectBucketConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectBucketConfigSpec   `json:"spec"`
	Status            ProjectBucketConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectBucketConfigList contains a list of ProjectBucketConfigs
type ProjectBucketConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectBucketConfig `json:"items"`
}

// Repository type metadata.
var (
	ProjectBucketConfig_Kind             = "ProjectBucketConfig"
	ProjectBucketConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectBucketConfig_Kind}.String()
	ProjectBucketConfig_KindAPIVersion   = ProjectBucketConfig_Kind + "." + CRDGroupVersion.String()
	ProjectBucketConfig_GroupVersionKind = CRDGroupVersion.WithKind(ProjectBucketConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectBucketConfig{}, &ProjectBucketConfigList{})
}
