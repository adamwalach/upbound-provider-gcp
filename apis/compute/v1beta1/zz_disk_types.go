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

type AsyncPrimaryDiskInitParameters struct {

	// Primary disk for asynchronous disk replication.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Disk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// Reference to a Disk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskRef *v1.Reference `json:"diskRef,omitempty" tf:"-"`

	// Selector for a Disk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskSelector *v1.Selector `json:"diskSelector,omitempty" tf:"-"`
}

type AsyncPrimaryDiskObservation struct {

	// Primary disk for asynchronous disk replication.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`
}

type AsyncPrimaryDiskParameters struct {

	// Primary disk for asynchronous disk replication.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Disk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// Reference to a Disk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskRef *v1.Reference `json:"diskRef,omitempty" tf:"-"`

	// Selector for a Disk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskSelector *v1.Selector `json:"diskSelector,omitempty" tf:"-"`
}

type DiskEncryptionKeyInitParameters struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`
}

type DiskEncryptionKeyObservation struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// (Output)
	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	// encryption key that protects this resource.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type DiskEncryptionKeyParameters struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	// +kubebuilder:validation:Optional
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	RawKeySecretRef *v1.SecretKeySelector `json:"rawKeySecretRef,omitempty" tf:"-"`

	// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit
	// customer-supplied encryption key to either encrypt or decrypt
	// this resource. You can provide either the rawKey or the rsaEncryptedKey.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Optional
	RsaEncryptedKeySecretRef *v1.SecretKeySelector `json:"rsaEncryptedKeySecretRef,omitempty" tf:"-"`
}

type DiskInitParameters struct {

	// A nested object resource
	// Structure is documented below.
	AsyncPrimaryDisk []AsyncPrimaryDiskInitParameters `json:"asyncPrimaryDisk,omitempty" tf:"async_primary_disk,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey []DiskEncryptionKeyInitParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// Whether this disk is using confidential compute mode.
	// Note: Only supported on hyperdisk skus, disk_encryption_key is required when setting to true
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty" tf:"enable_confidential_compute,omitempty"`

	// A list of features to enable on the guest operating system.
	// Applicable only for bootable disks.
	// Structure is documented below.
	GuestOsFeatures []GuestOsFeaturesInitParameters `json:"guestOsFeatures,omitempty" tf:"guest_os_features,omitempty"`

	// The image from which to initialize this disk. This can be
	// one of: the image's self_link, projects/{project}/global/images/{image},
	// projects/{project}/global/images/family/{family}, global/images/{image},
	// global/images/family/{family}, family/{family}, {project}/{family},
	// {project}/{image}, {family}, or {image}. If referred by family, the
	// images names must include the family name. If they don't, use the
	// google_compute_image data source.
	// For instance, the image centos-6-v20180104 includes its family name centos-6.
	// These images can be referred by family name here.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Labels to apply to this disk.  A list of key->value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Any applicable license URI.
	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes *float64 `json:"physicalBlockSizeBytes,omitempty" tf:"physical_block_size_bytes,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Indicates how many IOPS must be provisioned for the disk.
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of IOPS every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	ProvisionedIops *float64 `json:"provisionedIops,omitempty" tf:"provisioned_iops,omitempty"`

	// Indicates how much Throughput must be provisioned for the disk.
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of Throughput every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	ProvisionedThroughput *float64 `json:"provisionedThroughput,omitempty" tf:"provisioned_throughput,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the image or
	// snapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with image or snapshot,
	// the value must not be less than the size of the image
	// or the size of the snapshot.
	// You can add lifecycle.prevent_destroy in the config to prevent destroying
	// and recreating.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. If the snapshot is in another
	// project than this disk, you must supply a full URL. For example, the
	// following are valid values:
	Snapshot *string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`

	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource.
	// For example, the following are valid values:
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The customer-supplied encryption key of the source image. Required if
	// the source image is protected by a customer-supplied encryption key.
	// Structure is documented below.
	SourceImageEncryptionKey []SourceImageEncryptionKeyInitParameters `json:"sourceImageEncryptionKey,omitempty" tf:"source_image_encryption_key,omitempty"`

	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey []SourceSnapshotEncryptionKeyInitParameters `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`

	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DiskObservation struct {

	// A nested object resource
	// Structure is documented below.
	AsyncPrimaryDisk []AsyncPrimaryDiskObservation `json:"asyncPrimaryDisk,omitempty" tf:"async_primary_disk,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	DiskEncryptionKey []DiskEncryptionKeyObservation `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// The unique identifier for the resource. This identifier is defined by the server.
	DiskID *string `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// for all of the labels present on the resource.
	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// Whether this disk is using confidential compute mode.
	// Note: Only supported on hyperdisk skus, disk_encryption_key is required when setting to true
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty" tf:"enable_confidential_compute,omitempty"`

	// A list of features to enable on the guest operating system.
	// Applicable only for bootable disks.
	// Structure is documented below.
	GuestOsFeatures []GuestOsFeaturesObservation `json:"guestOsFeatures,omitempty" tf:"guest_os_features,omitempty"`

	// an identifier for the resource with format projects/{{project}}/zones/{{zone}}/disks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image from which to initialize this disk. This can be
	// one of: the image's self_link, projects/{project}/global/images/{image},
	// projects/{project}/global/images/family/{family}, global/images/{image},
	// global/images/family/{family}, family/{family}, {project}/{family},
	// {project}/{image}, {family}, or {image}. If referred by family, the
	// images names must include the family name. If they don't, use the
	// google_compute_image data source.
	// For instance, the image centos-6-v20180104 includes its family name centos-6.
	// These images can be referred by family name here.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The fingerprint used for optimistic locking of this resource.  Used
	// internally during updates.
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	// Labels to apply to this disk.  A list of key->value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp *string `json:"lastAttachTimestamp,omitempty" tf:"last_attach_timestamp,omitempty"`

	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp *string `json:"lastDetachTimestamp,omitempty" tf:"last_detach_timestamp,omitempty"`

	// Any applicable license URI.
	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	PhysicalBlockSizeBytes *float64 `json:"physicalBlockSizeBytes,omitempty" tf:"physical_block_size_bytes,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Indicates how many IOPS must be provisioned for the disk.
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of IOPS every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	ProvisionedIops *float64 `json:"provisionedIops,omitempty" tf:"provisioned_iops,omitempty"`

	// Indicates how much Throughput must be provisioned for the disk.
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of Throughput every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	ProvisionedThroughput *float64 `json:"provisionedThroughput,omitempty" tf:"provisioned_throughput,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the image or
	// snapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with image or snapshot,
	// the value must not be less than the size of the image
	// or the size of the snapshot.
	// You can add lifecycle.prevent_destroy in the config to prevent destroying
	// and recreating.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. If the snapshot is in another
	// project than this disk, you must supply a full URL. For example, the
	// following are valid values:
	Snapshot *string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`

	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource.
	// For example, the following are valid values:
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The ID value of the disk used to create this image. This value may
	// be used to determine whether the image was taken from the current
	// or a previous instance of a given disk name.
	SourceDiskID *string `json:"sourceDiskId,omitempty" tf:"source_disk_id,omitempty"`

	// The customer-supplied encryption key of the source image. Required if
	// the source image is protected by a customer-supplied encryption key.
	// Structure is documented below.
	SourceImageEncryptionKey []SourceImageEncryptionKeyObservation `json:"sourceImageEncryptionKey,omitempty" tf:"source_image_encryption_key,omitempty"`

	// The ID value of the image used to create this disk. This value
	// identifies the exact image that was used to create this persistent
	// disk. For example, if you created the persistent disk from an image
	// that was later deleted and recreated under the same name, the source
	// image ID would identify the exact version of the image that was used.
	SourceImageID *string `json:"sourceImageId,omitempty" tf:"source_image_id,omitempty"`

	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	SourceSnapshotEncryptionKey []SourceSnapshotEncryptionKeyObservation `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`

	// The unique ID of the snapshot used to create this disk. This value
	// identifies the exact snapshot that was used to create this persistent
	// disk. For example, if you created the persistent disk from a snapshot
	// that was later deleted and recreated under the same name, the source
	// snapshot ID would identify the exact version of the snapshot that was
	// used.
	SourceSnapshotID *string `json:"sourceSnapshotId,omitempty" tf:"source_snapshot_id,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Links to the users of the disk (attached instances) in form:
	// project/zones/zone/instances/instance
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`

	// A reference to the zone where the disk resides.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskParameters struct {

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AsyncPrimaryDisk []AsyncPrimaryDiskParameters `json:"asyncPrimaryDisk,omitempty" tf:"async_primary_disk,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Encrypts the disk using a customer-supplied encryption key.
	// After you encrypt a disk with a customer-supplied key, you must
	// provide the same key if you use the disk later (e.g. to create a disk
	// snapshot or an image, or to attach the disk to a virtual machine).
	// Customer-supplied encryption keys do not protect access to metadata of
	// the disk.
	// If you do not provide an encryption key when creating the disk, then
	// the disk will be encrypted using an automatically generated key and
	// you do not need to provide a key to use the disk later.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DiskEncryptionKey []DiskEncryptionKeyParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// Whether this disk is using confidential compute mode.
	// Note: Only supported on hyperdisk skus, disk_encryption_key is required when setting to true
	// +kubebuilder:validation:Optional
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute,omitempty" tf:"enable_confidential_compute,omitempty"`

	// A list of features to enable on the guest operating system.
	// Applicable only for bootable disks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GuestOsFeatures []GuestOsFeaturesParameters `json:"guestOsFeatures,omitempty" tf:"guest_os_features,omitempty"`

	// The image from which to initialize this disk. This can be
	// one of: the image's self_link, projects/{project}/global/images/{image},
	// projects/{project}/global/images/family/{family}, global/images/{image},
	// global/images/family/{family}, family/{family}, {project}/{family},
	// {project}/{image}, {family}, or {image}. If referred by family, the
	// images names must include the family name. If they don't, use the
	// google_compute_image data source.
	// For instance, the image centos-6-v20180104 includes its family name centos-6.
	// These images can be referred by family name here.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Labels to apply to this disk.  A list of key->value pairs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Any applicable license URI.
	// +kubebuilder:validation:Optional
	Licenses []*string `json:"licenses,omitempty" tf:"licenses,omitempty"`

	// Physical block size of the persistent disk, in bytes. If not present
	// in a request, a default value is used. Currently supported sizes
	// are 4096 and 16384, other sizes may be added in the future.
	// If an unsupported value is requested, the error message will list
	// the supported values for the caller's project.
	// +kubebuilder:validation:Optional
	PhysicalBlockSizeBytes *float64 `json:"physicalBlockSizeBytes,omitempty" tf:"physical_block_size_bytes,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Indicates how many IOPS must be provisioned for the disk.
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of IOPS every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	// +kubebuilder:validation:Optional
	ProvisionedIops *float64 `json:"provisionedIops,omitempty" tf:"provisioned_iops,omitempty"`

	// Indicates how much Throughput must be provisioned for the disk.
	// Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
	// allows for an update of Throughput every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it
	// +kubebuilder:validation:Optional
	ProvisionedThroughput *float64 `json:"provisionedThroughput,omitempty" tf:"provisioned_throughput,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the image or
	// snapshot parameter, or specify it alone to create an empty
	// persistent disk.
	// If you specify this field along with image or snapshot,
	// the value must not be less than the size of the image
	// or the size of the snapshot.
	// You can add lifecycle.prevent_destroy in the config to prevent destroying
	// and recreating.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot used to create this disk. You can provide this as
	// a partial or full URL to the resource. If the snapshot is in another
	// project than this disk, you must supply a full URL. For example, the
	// following are valid values:
	// +kubebuilder:validation:Optional
	Snapshot *string `json:"snapshot,omitempty" tf:"snapshot,omitempty"`

	// The source disk used to create this disk. You can provide this as a partial or full URL to the resource.
	// For example, the following are valid values:
	// +kubebuilder:validation:Optional
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The customer-supplied encryption key of the source image. Required if
	// the source image is protected by a customer-supplied encryption key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SourceImageEncryptionKey []SourceImageEncryptionKeyParameters `json:"sourceImageEncryptionKey,omitempty" tf:"source_image_encryption_key,omitempty"`

	// The customer-supplied encryption key of the source snapshot. Required
	// if the source snapshot is protected by a customer-supplied encryption
	// key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SourceSnapshotEncryptionKey []SourceSnapshotEncryptionKeyParameters `json:"sourceSnapshotEncryptionKey,omitempty" tf:"source_snapshot_encryption_key,omitempty"`

	// URL of the disk type resource describing which disk type to use to
	// create the disk. Provide this when creating the disk.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A reference to the zone where the disk resides.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type GuestOsFeaturesInitParameters struct {

	// The type of supported feature. Read Enabling guest operating system features to see a list of available options.
	// Possible values are: MULTI_IP_SUBNET, SECURE_BOOT, SEV_CAPABLE, UEFI_COMPATIBLE, VIRTIO_SCSI_MULTIQUEUE, WINDOWS, GVNIC, SEV_LIVE_MIGRATABLE, SEV_SNP_CAPABLE, SUSPEND_RESUME_COMPATIBLE, TDX_CAPABLE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GuestOsFeaturesObservation struct {

	// The type of supported feature. Read Enabling guest operating system features to see a list of available options.
	// Possible values are: MULTI_IP_SUBNET, SECURE_BOOT, SEV_CAPABLE, UEFI_COMPATIBLE, VIRTIO_SCSI_MULTIQUEUE, WINDOWS, GVNIC, SEV_LIVE_MIGRATABLE, SEV_SNP_CAPABLE, SUSPEND_RESUME_COMPATIBLE, TDX_CAPABLE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GuestOsFeaturesParameters struct {

	// The type of supported feature. Read Enabling guest operating system features to see a list of available options.
	// Possible values are: MULTI_IP_SUBNET, SECURE_BOOT, SEV_CAPABLE, UEFI_COMPATIBLE, VIRTIO_SCSI_MULTIQUEUE, WINDOWS, GVNIC, SEV_LIVE_MIGRATABLE, SEV_SNP_CAPABLE, SUSPEND_RESUME_COMPATIBLE, TDX_CAPABLE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type SourceImageEncryptionKeyInitParameters struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type SourceImageEncryptionKeyObservation struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`

	// (Output)
	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	// encryption key that protects this resource.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type SourceImageEncryptionKeyParameters struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	// +kubebuilder:validation:Optional
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// +kubebuilder:validation:Optional
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type SourceSnapshotEncryptionKeyInitParameters struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

type SourceSnapshotEncryptionKeyObservation struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`

	// (Output)
	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	// encryption key that protects this resource.
	Sha256 *string `json:"sha256,omitempty" tf:"sha256,omitempty"`
}

type SourceSnapshotEncryptionKeyParameters struct {

	// The self link of the encryption key used to encrypt the disk. Also called KmsKeyName
	// in the cloud console. Your project's Compute Engine System service account
	// (service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com) must have
	// roles/cloudkms.cryptoKeyEncrypterDecrypter to use this feature.
	// See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// The service account used for the encryption request for the given KMS key.
	// If absent, the Compute Engine Service Agent service account is used.
	// +kubebuilder:validation:Optional
	KMSKeyServiceAccount *string `json:"kmsKeyServiceAccount,omitempty" tf:"kms_key_service_account,omitempty"`

	// Specifies a 256-bit customer-supplied encryption key, encoded in
	// RFC 4648 base64 to either encrypt or decrypt this resource.
	// +kubebuilder:validation:Optional
	RawKey *string `json:"rawKey,omitempty" tf:"raw_key,omitempty"`
}

// DiskSpec defines the desired state of Disk
type DiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskParameters `json:"forProvider"`
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
	InitProvider DiskInitParameters `json:"initProvider,omitempty"`
}

// DiskStatus defines the observed state of Disk.
type DiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Disk is the Schema for the Disks API. Persistent disks are durable storage devices that function similarly to the physical disks in a desktop or a server.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskSpec   `json:"spec"`
	Status            DiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskList contains a list of Disks
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

// Repository type metadata.
var (
	Disk_Kind             = "Disk"
	Disk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Disk_Kind}.String()
	Disk_KindAPIVersion   = Disk_Kind + "." + CRDGroupVersion.String()
	Disk_GroupVersionKind = CRDGroupVersion.WithKind(Disk_Kind)
)

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
