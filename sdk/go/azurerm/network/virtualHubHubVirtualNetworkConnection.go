// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// HubVirtualNetworkConnection Resource.
type VirtualHubHubVirtualNetworkConnection struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the hub virtual network connection.
	Properties HubVirtualNetworkConnectionPropertiesResponseOutput `pulumi:"properties"`
}

// NewVirtualHubHubVirtualNetworkConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualHubHubVirtualNetworkConnection(ctx *pulumi.Context,
	name string, args *VirtualHubHubVirtualNetworkConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualHubHubVirtualNetworkConnection, error) {
	if args == nil || args.ConnectionName == nil {
		return nil, errors.New("missing required argument 'ConnectionName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualHubName == nil {
		return nil, errors.New("missing required argument 'VirtualHubName'")
	}
	if args == nil {
		args = &VirtualHubHubVirtualNetworkConnectionArgs{}
	}
	var resource VirtualHubHubVirtualNetworkConnection
	err := ctx.RegisterResource("azurerm:network:VirtualHubHubVirtualNetworkConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHubHubVirtualNetworkConnection gets an existing VirtualHubHubVirtualNetworkConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHubHubVirtualNetworkConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubHubVirtualNetworkConnectionState, opts ...pulumi.ResourceOption) (*VirtualHubHubVirtualNetworkConnection, error) {
	var resource VirtualHubHubVirtualNetworkConnection
	err := ctx.ReadResource("azurerm:network:VirtualHubHubVirtualNetworkConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHubHubVirtualNetworkConnection resources.
type virtualHubHubVirtualNetworkConnectionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the hub virtual network connection.
	Properties *HubVirtualNetworkConnectionPropertiesResponse `pulumi:"properties"`
}

type VirtualHubHubVirtualNetworkConnectionState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the hub virtual network connection.
	Properties HubVirtualNetworkConnectionPropertiesResponsePtrInput
}

func (VirtualHubHubVirtualNetworkConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubHubVirtualNetworkConnectionState)(nil)).Elem()
}

type virtualHubHubVirtualNetworkConnectionArgs struct {
	// The name of the HubVirtualNetworkConnection.
	ConnectionName string `pulumi:"connectionName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the hub virtual network connection.
	Properties *HubVirtualNetworkConnectionProperties `pulumi:"properties"`
	// The resource group name of the HubVirtualNetworkConnection.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The set of arguments for constructing a VirtualHubHubVirtualNetworkConnection resource.
type VirtualHubHubVirtualNetworkConnectionArgs struct {
	// The name of the HubVirtualNetworkConnection.
	ConnectionName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the hub virtual network connection.
	Properties HubVirtualNetworkConnectionPropertiesPtrInput
	// The resource group name of the HubVirtualNetworkConnection.
	ResourceGroupName pulumi.StringInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
}

func (VirtualHubHubVirtualNetworkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubHubVirtualNetworkConnectionArgs)(nil)).Elem()
}