// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Private Endpoint Connection resource.
// Latest API Version: 2020-09-30.
type DiskAccessAPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDiskAccessAPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *DiskAccessAPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*DiskAccessAPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskAccessName == nil {
		return nil, errors.New("invalid value for required argument 'DiskAccessName'")
	}
	if args.PrivateEndpointConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateEndpointConnectionName'")
	}
	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:DiskAccessAPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource DiskAccessAPrivateEndpointConnection
	err := ctx.RegisterResource("azure-nextgen:compute/latest:DiskAccessAPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskAccessAPrivateEndpointConnection gets an existing DiskAccessAPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAccessAPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*DiskAccessAPrivateEndpointConnection, error) {
	var resource DiskAccessAPrivateEndpointConnection
	err := ctx.ReadResource("azure-nextgen:compute/latest:DiskAccessAPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskAccessAPrivateEndpointConnection resources.
type diskAccessAPrivateEndpointConnectionState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource type
	Type *string `pulumi:"type"`
}

type DiskAccessAPrivateEndpointConnectionState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrInput
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrInput
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (DiskAccessAPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessAPrivateEndpointConnectionState)(nil)).Elem()
}

type diskAccessAPrivateEndpointConnectionArgs struct {
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	DiskAccessName string `pulumi:"diskAccessName"`
	// The name of the private endpoint connection
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DiskAccessAPrivateEndpointConnection resource.
type DiskAccessAPrivateEndpointConnectionArgs struct {
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	DiskAccessName pulumi.StringInput
	// The name of the private endpoint connection
	PrivateEndpointConnectionName pulumi.StringInput
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (DiskAccessAPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessAPrivateEndpointConnectionArgs)(nil)).Elem()
}

type DiskAccessAPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput
	ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput
}

func (*DiskAccessAPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskAccessAPrivateEndpointConnection)(nil))
}

func (i *DiskAccessAPrivateEndpointConnection) ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput {
	return i.ToDiskAccessAPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *DiskAccessAPrivateEndpointConnection) ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAccessAPrivateEndpointConnectionOutput)
}

type DiskAccessAPrivateEndpointConnectionOutput struct {
	*pulumi.OutputState
}

func (DiskAccessAPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskAccessAPrivateEndpointConnection)(nil))
}

func (o DiskAccessAPrivateEndpointConnectionOutput) ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput {
	return o
}

func (o DiskAccessAPrivateEndpointConnectionOutput) ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DiskAccessAPrivateEndpointConnectionOutput{})
}
