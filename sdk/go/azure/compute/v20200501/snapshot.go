// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Snapshot resource.
type Snapshot struct {
	pulumi.CustomResourceState

	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponseOutput `pulumi:"creationData"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId pulumi.StringPtrOutput `pulumi:"diskAccessId"`
	// The size of the disk in bytes. This field is read only.
	DiskSizeBytes pulumi.Float64Output `pulumi:"diskSizeBytes"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrOutput `pulumi:"diskSizeGB"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption EncryptionResponsePtrOutput `pulumi:"encryption"`
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection EncryptionSettingsCollectionResponsePtrOutput `pulumi:"encryptionSettingsCollection"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental pulumi.BoolPtrOutput `pulumi:"incremental"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Unused. Always Null.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy pulumi.StringPtrOutput `pulumi:"networkAccessPolicy"`
	// The Operating System type.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The disk provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
	Sku SnapshotSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The time when the snapshot was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique Guid identifying the resource.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreationData == nil {
		return nil, errors.New("invalid value for required argument 'CreationData'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SnapshotName == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/latest:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200630:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-nextgen:compute/v20200501:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-nextgen:compute/v20200501:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationDataResponse `pulumi:"creationData"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `pulumi:"diskAccessId"`
	// The size of the disk in bytes. This field is read only.
	DiskSizeBytes *float64 `pulumi:"diskSizeBytes"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *EncryptionResponse `pulumi:"encryption"`
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental *bool `pulumi:"incremental"`
	// Resource location
	Location *string `pulumi:"location"`
	// Unused. Always Null.
	ManagedBy *string `pulumi:"managedBy"`
	// Resource name
	Name *string `pulumi:"name"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy *string `pulumi:"networkAccessPolicy"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// The disk provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
	Sku *SnapshotSkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The time when the snapshot was created.
	TimeCreated *string `pulumi:"timeCreated"`
	// Resource type
	Type *string `pulumi:"type"`
	// Unique Guid identifying the resource.
	UniqueId *string `pulumi:"uniqueId"`
}

type SnapshotState struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponsePtrInput
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId pulumi.StringPtrInput
	// The size of the disk in bytes. This field is read only.
	DiskSizeBytes pulumi.Float64PtrInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption EncryptionResponsePtrInput
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection EncryptionSettingsCollectionResponsePtrInput
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrInput
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental pulumi.BoolPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Unused. Always Null.
	ManagedBy pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Policy for accessing the disk via network.
	NetworkAccessPolicy pulumi.StringPtrInput
	// The Operating System type.
	OsType pulumi.StringPtrInput
	// The disk provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
	Sku SnapshotSkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The time when the snapshot was created.
	TimeCreated pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// Unique Guid identifying the resource.
	UniqueId pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationData `pulumi:"creationData"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `pulumi:"diskAccessId"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption `pulumi:"encryption"`
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection `pulumi:"encryptionSettingsCollection"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental *bool `pulumi:"incremental"`
	// Resource location
	Location *string `pulumi:"location"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy *string `pulumi:"networkAccessPolicy"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
	Sku *SnapshotSku `pulumi:"sku"`
	// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
	SnapshotName string `pulumi:"snapshotName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataInput
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId pulumi.StringPtrInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption EncryptionPtrInput
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection EncryptionSettingsCollectionPtrInput
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrInput
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental pulumi.BoolPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Policy for accessing the disk via network.
	NetworkAccessPolicy pulumi.StringPtrInput
	// The Operating System type.
	OsType *OperatingSystemTypes
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
	Sku SnapshotSkuPtrInput
	// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
	SnapshotName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct {
	*pulumi.OutputState
}

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
}
