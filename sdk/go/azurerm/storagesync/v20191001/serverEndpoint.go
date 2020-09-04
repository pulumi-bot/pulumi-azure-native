// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Server Endpoint object.
type ServerEndpoint struct {
	pulumi.CustomResourceState

	// Cloud Tiering.
	CloudTiering pulumi.StringPtrOutput `pulumi:"cloudTiering"`
	// Cloud tiering status. Only populated if cloud tiering is enabled.
	CloudTieringStatus ServerEndpointCloudTieringStatusResponseOutput `pulumi:"cloudTieringStatus"`
	// Friendly Name
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// Resource Last Operation Name
	LastOperationName pulumi.StringOutput `pulumi:"lastOperationName"`
	// ServerEndpoint lastWorkflowId
	LastWorkflowId pulumi.StringOutput `pulumi:"lastWorkflowId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Offline data transfer
	OfflineDataTransfer pulumi.StringPtrOutput `pulumi:"offlineDataTransfer"`
	// Offline data transfer share name
	OfflineDataTransferShareName pulumi.StringPtrOutput `pulumi:"offlineDataTransferShareName"`
	// Offline data transfer storage account resource ID
	OfflineDataTransferStorageAccountResourceId pulumi.StringOutput `pulumi:"offlineDataTransferStorageAccountResourceId"`
	// Offline data transfer storage account tenant ID
	OfflineDataTransferStorageAccountTenantId pulumi.StringOutput `pulumi:"offlineDataTransferStorageAccountTenantId"`
	// ServerEndpoint Provisioning State
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Recall status. Only populated if cloud tiering is enabled.
	RecallStatus ServerEndpointRecallStatusResponseOutput `pulumi:"recallStatus"`
	// Server Local path.
	ServerLocalPath pulumi.StringPtrOutput `pulumi:"serverLocalPath"`
	// Server Resource Id.
	ServerResourceId pulumi.StringPtrOutput `pulumi:"serverResourceId"`
	// Server Endpoint sync status
	SyncStatus ServerEndpointSyncStatusResponseOutput `pulumi:"syncStatus"`
	// Tier files older than days.
	TierFilesOlderThanDays pulumi.IntPtrOutput `pulumi:"tierFilesOlderThanDays"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent pulumi.IntPtrOutput `pulumi:"volumeFreeSpacePercent"`
}

// NewServerEndpoint registers a new resource with the given unique name, arguments, and options.
func NewServerEndpoint(ctx *pulumi.Context,
	name string, args *ServerEndpointArgs, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerEndpointName == nil {
		return nil, errors.New("missing required argument 'ServerEndpointName'")
	}
	if args == nil || args.StorageSyncServiceName == nil {
		return nil, errors.New("missing required argument 'StorageSyncServiceName'")
	}
	if args == nil || args.SyncGroupName == nil {
		return nil, errors.New("missing required argument 'SyncGroupName'")
	}
	if args == nil {
		args = &ServerEndpointArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storagesync/latest:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20170605preview:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20180402:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20180701:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20181001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190201:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190601:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20200301:ServerEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerEndpoint
	err := ctx.RegisterResource("azurerm:storagesync/v20191001:ServerEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerEndpoint gets an existing ServerEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerEndpointState, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
	var resource ServerEndpoint
	err := ctx.ReadResource("azurerm:storagesync/v20191001:ServerEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerEndpoint resources.
type serverEndpointState struct {
	// Cloud Tiering.
	CloudTiering *string `pulumi:"cloudTiering"`
	// Cloud tiering status. Only populated if cloud tiering is enabled.
	CloudTieringStatus *ServerEndpointCloudTieringStatusResponse `pulumi:"cloudTieringStatus"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// Resource Last Operation Name
	LastOperationName *string `pulumi:"lastOperationName"`
	// ServerEndpoint lastWorkflowId
	LastWorkflowId *string `pulumi:"lastWorkflowId"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Offline data transfer
	OfflineDataTransfer *string `pulumi:"offlineDataTransfer"`
	// Offline data transfer share name
	OfflineDataTransferShareName *string `pulumi:"offlineDataTransferShareName"`
	// Offline data transfer storage account resource ID
	OfflineDataTransferStorageAccountResourceId *string `pulumi:"offlineDataTransferStorageAccountResourceId"`
	// Offline data transfer storage account tenant ID
	OfflineDataTransferStorageAccountTenantId *string `pulumi:"offlineDataTransferStorageAccountTenantId"`
	// ServerEndpoint Provisioning State
	ProvisioningState *string `pulumi:"provisioningState"`
	// Recall status. Only populated if cloud tiering is enabled.
	RecallStatus *ServerEndpointRecallStatusResponse `pulumi:"recallStatus"`
	// Server Local path.
	ServerLocalPath *string `pulumi:"serverLocalPath"`
	// Server Resource Id.
	ServerResourceId *string `pulumi:"serverResourceId"`
	// Server Endpoint sync status
	SyncStatus *ServerEndpointSyncStatusResponse `pulumi:"syncStatus"`
	// Tier files older than days.
	TierFilesOlderThanDays *int `pulumi:"tierFilesOlderThanDays"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent *int `pulumi:"volumeFreeSpacePercent"`
}

type ServerEndpointState struct {
	// Cloud Tiering.
	CloudTiering pulumi.StringPtrInput
	// Cloud tiering status. Only populated if cloud tiering is enabled.
	CloudTieringStatus ServerEndpointCloudTieringStatusResponsePtrInput
	// Friendly Name
	FriendlyName pulumi.StringPtrInput
	// Resource Last Operation Name
	LastOperationName pulumi.StringPtrInput
	// ServerEndpoint lastWorkflowId
	LastWorkflowId pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Offline data transfer
	OfflineDataTransfer pulumi.StringPtrInput
	// Offline data transfer share name
	OfflineDataTransferShareName pulumi.StringPtrInput
	// Offline data transfer storage account resource ID
	OfflineDataTransferStorageAccountResourceId pulumi.StringPtrInput
	// Offline data transfer storage account tenant ID
	OfflineDataTransferStorageAccountTenantId pulumi.StringPtrInput
	// ServerEndpoint Provisioning State
	ProvisioningState pulumi.StringPtrInput
	// Recall status. Only populated if cloud tiering is enabled.
	RecallStatus ServerEndpointRecallStatusResponsePtrInput
	// Server Local path.
	ServerLocalPath pulumi.StringPtrInput
	// Server Resource Id.
	ServerResourceId pulumi.StringPtrInput
	// Server Endpoint sync status
	SyncStatus ServerEndpointSyncStatusResponsePtrInput
	// Tier files older than days.
	TierFilesOlderThanDays pulumi.IntPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent pulumi.IntPtrInput
}

func (ServerEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverEndpointState)(nil)).Elem()
}

type serverEndpointArgs struct {
	// Cloud Tiering.
	CloudTiering *string `pulumi:"cloudTiering"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// Offline data transfer
	OfflineDataTransfer *string `pulumi:"offlineDataTransfer"`
	// Offline data transfer share name
	OfflineDataTransferShareName *string `pulumi:"offlineDataTransferShareName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Server Endpoint object.
	ServerEndpointName string `pulumi:"serverEndpointName"`
	// Server Local path.
	ServerLocalPath *string `pulumi:"serverLocalPath"`
	// Server Resource Id.
	ServerResourceId *string `pulumi:"serverResourceId"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Name of Sync Group resource.
	SyncGroupName string `pulumi:"syncGroupName"`
	// Tier files older than days.
	TierFilesOlderThanDays *int `pulumi:"tierFilesOlderThanDays"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent *int `pulumi:"volumeFreeSpacePercent"`
}

// The set of arguments for constructing a ServerEndpoint resource.
type ServerEndpointArgs struct {
	// Cloud Tiering.
	CloudTiering pulumi.StringPtrInput
	// Friendly Name
	FriendlyName pulumi.StringPtrInput
	// Offline data transfer
	OfflineDataTransfer pulumi.StringPtrInput
	// Offline data transfer share name
	OfflineDataTransferShareName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Server Endpoint object.
	ServerEndpointName pulumi.StringInput
	// Server Local path.
	ServerLocalPath pulumi.StringPtrInput
	// Server Resource Id.
	ServerResourceId pulumi.StringPtrInput
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput
	// Name of Sync Group resource.
	SyncGroupName pulumi.StringInput
	// Tier files older than days.
	TierFilesOlderThanDays pulumi.IntPtrInput
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent pulumi.IntPtrInput
}

func (ServerEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverEndpointArgs)(nil)).Elem()
}
