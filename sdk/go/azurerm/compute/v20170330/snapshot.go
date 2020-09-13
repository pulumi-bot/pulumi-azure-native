// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170330

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Snapshot resource.
//
// ## Example Usage
// ### Create a snapshot by importing an unmanaged blob from a different subscription.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20170330"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewSnapshot(ctx, "snapshot", &compute.SnapshotArgs{
// 			CreationData: &compute.CreationDataArgs{
// 				CreateOption:     pulumi.String("Import"),
// 				SourceUri:        pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
// 				StorageAccountId: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
// 			},
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			SnapshotName:      pulumi.String("mySnapshot1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a snapshot by importing an unmanaged blob from the same subscription.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20170330"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewSnapshot(ctx, "snapshot", &compute.SnapshotArgs{
// 			CreationData: &compute.CreationDataArgs{
// 				CreateOption: pulumi.String("Import"),
// 				SourceUri:    pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
// 			},
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			SnapshotName:      pulumi.String("mySnapshot1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a snapshot from an existing snapshot in the same or a different subscription.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20170330"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewSnapshot(ctx, "snapshot", &compute.SnapshotArgs{
// 			CreationData: &compute.CreationDataArgs{
// 				CreateOption:     pulumi.String("Copy"),
// 				SourceResourceId: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
// 			},
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			SnapshotName:      pulumi.String("mySnapshot2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponseOutput `pulumi:"creationData"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrOutput `pulumi:"diskSizeGB"`
	// Encryption settings for disk or snapshot
	EncryptionSettings EncryptionSettingsResponsePtrOutput `pulumi:"encryptionSettings"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Unused. Always Null.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The Operating System type.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The disk provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The disks and snapshots sku name. Can be Standard_LRS or Premium_LRS.
	Sku DiskSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The time when the disk was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil || args.CreationData == nil {
		return nil, errors.New("missing required argument 'CreationData'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SnapshotName == nil {
		return nil, errors.New("missing required argument 'SnapshotName'")
	}
	if args == nil {
		args = &SnapshotArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:compute/latest:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20160430preview:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20180401:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20180601:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20180930:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190301:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20200630:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azurerm:compute/v20170330:Snapshot", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:compute/v20170330:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationDataResponse `pulumi:"creationData"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption settings for disk or snapshot
	EncryptionSettings *EncryptionSettingsResponse `pulumi:"encryptionSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// Unused. Always Null.
	ManagedBy *string `pulumi:"managedBy"`
	// Resource name
	Name *string `pulumi:"name"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// The disk provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The disks and snapshots sku name. Can be Standard_LRS or Premium_LRS.
	Sku *DiskSkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The time when the disk was created.
	TimeCreated *string `pulumi:"timeCreated"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SnapshotState struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponsePtrInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption settings for disk or snapshot
	EncryptionSettings EncryptionSettingsResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Unused. Always Null.
	ManagedBy pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The Operating System type.
	OsType pulumi.StringPtrInput
	// The disk provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The disks and snapshots sku name. Can be Standard_LRS or Premium_LRS.
	Sku DiskSkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The time when the disk was created.
	TimeCreated pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationData `pulumi:"creationData"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption settings for disk or snapshot
	EncryptionSettings *EncryptionSettings `pulumi:"encryptionSettings"`
	// Resource location
	Location string `pulumi:"location"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The disks and snapshots sku name. Can be Standard_LRS or Premium_LRS.
	Sku *DiskSku `pulumi:"sku"`
	// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
	SnapshotName string `pulumi:"snapshotName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption settings for disk or snapshot
	EncryptionSettings EncryptionSettingsPtrInput
	// Resource location
	Location pulumi.StringInput
	// The Operating System type.
	OsType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The disks and snapshots sku name. Can be Standard_LRS or Premium_LRS.
	Sku DiskSkuPtrInput
	// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
	SnapshotName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}
