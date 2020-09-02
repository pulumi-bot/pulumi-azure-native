// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Storage Sync Service object.
type StorageSyncService struct {
	pulumi.CustomResourceState

	// Incoming Traffic Policy
	IncomingTrafficPolicy pulumi.StringPtrOutput `pulumi:"incomingTrafficPolicy"`
	// Resource Last Operation Name
	LastOperationName pulumi.StringOutput `pulumi:"lastOperationName"`
	// StorageSyncService lastWorkflowId
	LastWorkflowId pulumi.StringOutput `pulumi:"lastWorkflowId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connection associated with the specified storage sync service
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// StorageSyncService Provisioning State
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Storage Sync service status.
	StorageSyncServiceStatus pulumi.IntOutput `pulumi:"storageSyncServiceStatus"`
	// Storage Sync service Uid
	StorageSyncServiceUid pulumi.StringOutput `pulumi:"storageSyncServiceUid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageSyncService registers a new resource with the given unique name, arguments, and options.
func NewStorageSyncService(ctx *pulumi.Context,
	name string, args *StorageSyncServiceArgs, opts ...pulumi.ResourceOption) (*StorageSyncService, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageSyncServiceName == nil {
		return nil, errors.New("missing required argument 'StorageSyncServiceName'")
	}
	if args == nil {
		args = &StorageSyncServiceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storagesync/v20170605preview:StorageSyncService"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20180402:StorageSyncService"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20180701:StorageSyncService"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20181001:StorageSyncService"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190201:StorageSyncService"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190301:StorageSyncService"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190601:StorageSyncService"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20191001:StorageSyncService"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20200301:StorageSyncService"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageSyncService
	err := ctx.RegisterResource("azurerm:storagesync/latest:StorageSyncService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageSyncService gets an existing StorageSyncService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageSyncService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageSyncServiceState, opts ...pulumi.ResourceOption) (*StorageSyncService, error) {
	var resource StorageSyncService
	err := ctx.ReadResource("azurerm:storagesync/latest:StorageSyncService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageSyncService resources.
type storageSyncServiceState struct {
	// Incoming Traffic Policy
	IncomingTrafficPolicy *string `pulumi:"incomingTrafficPolicy"`
	// Resource Last Operation Name
	LastOperationName *string `pulumi:"lastOperationName"`
	// StorageSyncService lastWorkflowId
	LastWorkflowId *string `pulumi:"lastWorkflowId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// List of private endpoint connection associated with the specified storage sync service
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// StorageSyncService Provisioning State
	ProvisioningState *string `pulumi:"provisioningState"`
	// Storage Sync service status.
	StorageSyncServiceStatus *int `pulumi:"storageSyncServiceStatus"`
	// Storage Sync service Uid
	StorageSyncServiceUid *string `pulumi:"storageSyncServiceUid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type StorageSyncServiceState struct {
	// Incoming Traffic Policy
	IncomingTrafficPolicy pulumi.StringPtrInput
	// Resource Last Operation Name
	LastOperationName pulumi.StringPtrInput
	// StorageSyncService lastWorkflowId
	LastWorkflowId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// List of private endpoint connection associated with the specified storage sync service
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput
	// StorageSyncService Provisioning State
	ProvisioningState pulumi.StringPtrInput
	// Storage Sync service status.
	StorageSyncServiceStatus pulumi.IntPtrInput
	// Storage Sync service Uid
	StorageSyncServiceUid pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (StorageSyncServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSyncServiceState)(nil)).Elem()
}

type storageSyncServiceArgs struct {
	// Incoming Traffic Policy
	IncomingTrafficPolicy *string `pulumi:"incomingTrafficPolicy"`
	// Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageSyncService resource.
type StorageSyncServiceArgs struct {
	// Incoming Traffic Policy
	IncomingTrafficPolicy pulumi.StringPtrInput
	// Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
	Location pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (StorageSyncServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSyncServiceArgs)(nil)).Elem()
}
