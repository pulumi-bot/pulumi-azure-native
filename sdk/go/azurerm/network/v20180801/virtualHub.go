// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VirtualHub Resource.
type VirtualHub struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters for VirtualHub
	Properties VirtualHubPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualHub registers a new resource with the given unique name, arguments, and options.
func NewVirtualHub(ctx *pulumi.Context,
	name string, args *VirtualHubArgs, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VirtualHubArgs{}
	}
	var resource VirtualHub
	err := ctx.RegisterResource("azurerm:network/v20180801:VirtualHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHub gets an existing VirtualHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubState, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	var resource VirtualHub
	err := ctx.ReadResource("azurerm:network/v20180801:VirtualHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHub resources.
type virtualHubState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Parameters for VirtualHub
	Properties *VirtualHubPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type VirtualHubState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Parameters for VirtualHub
	Properties VirtualHubPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (VirtualHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubState)(nil)).Elem()
}

type virtualHubArgs struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The expressRouteGateway associated with this VirtualHub
	ExpressRouteGateway *SubResource `pulumi:"expressRouteGateway"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the VirtualHub.
	Name string `pulumi:"name"`
	// The P2SVpnGateway associated with this VirtualHub
	P2SVpnGateway *SubResource `pulumi:"p2SVpnGateway"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The routeTable associated with this virtual hub.
	RouteTable *VirtualHubRouteTable `pulumi:"routeTable"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// list of all vnet connections with this VirtualHub.
	VirtualNetworkConnections []HubVirtualNetworkConnection `pulumi:"virtualNetworkConnections"`
	// The VirtualWAN to which the VirtualHub belongs
	VirtualWan *SubResource `pulumi:"virtualWan"`
	// The VpnGateway associated with this VirtualHub
	VpnGateway *SubResource `pulumi:"vpnGateway"`
}

// The set of arguments for constructing a VirtualHub resource.
type VirtualHubArgs struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix pulumi.StringPtrInput
	// The expressRouteGateway associated with this VirtualHub
	ExpressRouteGateway SubResourcePtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The name of the VirtualHub.
	Name pulumi.StringInput
	// The P2SVpnGateway associated with this VirtualHub
	P2SVpnGateway SubResourcePtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput
	// The routeTable associated with this virtual hub.
	RouteTable VirtualHubRouteTablePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// list of all vnet connections with this VirtualHub.
	VirtualNetworkConnections HubVirtualNetworkConnectionArrayInput
	// The VirtualWAN to which the VirtualHub belongs
	VirtualWan SubResourcePtrInput
	// The VpnGateway associated with this VirtualHub
	VpnGateway SubResourcePtrInput
}

func (VirtualHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubArgs)(nil)).Elem()
}
