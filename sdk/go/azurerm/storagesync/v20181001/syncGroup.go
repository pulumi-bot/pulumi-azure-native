// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sync Group object.
//
// ## Example Usage
// ### SyncGroups_Create
//
// ```go
// package main
//
// import (
// 	storagesync "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/storagesync/v20181001"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagesync.NewSyncGroup(ctx, "syncGroup", &storagesync.SyncGroupArgs{
// 			ResourceGroupName:      pulumi.String("SampleResourceGroup_1"),
// 			StorageSyncServiceName: pulumi.String("SampleStorageSyncService_1"),
// 			SyncGroupName:          pulumi.String("SampleSyncGroup_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type SyncGroup struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Sync group status
	SyncGroupStatus pulumi.StringOutput `pulumi:"syncGroupStatus"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique Id
	UniqueId pulumi.StringPtrOutput `pulumi:"uniqueId"`
}

// NewSyncGroup registers a new resource with the given unique name, arguments, and options.
func NewSyncGroup(ctx *pulumi.Context,
	name string, args *SyncGroupArgs, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageSyncServiceName == nil {
		return nil, errors.New("missing required argument 'StorageSyncServiceName'")
	}
	if args == nil || args.SyncGroupName == nil {
		return nil, errors.New("missing required argument 'SyncGroupName'")
	}
	if args == nil {
		args = &SyncGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storagesync/latest:SyncGroup"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20170605preview:SyncGroup"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20180402:SyncGroup"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20180701:SyncGroup"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190201:SyncGroup"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190301:SyncGroup"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190601:SyncGroup"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20191001:SyncGroup"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20200301:SyncGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SyncGroup
	err := ctx.RegisterResource("azurerm:storagesync/v20181001:SyncGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncGroup gets an existing SyncGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncGroupState, opts ...pulumi.ResourceOption) (*SyncGroup, error) {
	var resource SyncGroup
	err := ctx.ReadResource("azurerm:storagesync/v20181001:SyncGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncGroup resources.
type syncGroupState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// Sync group status
	SyncGroupStatus *string `pulumi:"syncGroupStatus"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// Unique Id
	UniqueId *string `pulumi:"uniqueId"`
}

type SyncGroupState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// Sync group status
	SyncGroupStatus pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// Unique Id
	UniqueId pulumi.StringPtrInput
}

func (SyncGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupState)(nil)).Elem()
}

type syncGroupArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Name of Sync Group resource.
	SyncGroupName string `pulumi:"syncGroupName"`
}

// The set of arguments for constructing a SyncGroup resource.
type SyncGroupArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput
	// Name of Sync Group resource.
	SyncGroupName pulumi.StringInput
}

func (SyncGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncGroupArgs)(nil)).Elem()
}
