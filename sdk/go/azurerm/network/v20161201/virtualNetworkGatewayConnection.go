// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A common class for general resource information
type VirtualNetworkGatewayConnection struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VirtualNetworkGatewayConnection properties
	Properties VirtualNetworkGatewayConnectionPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualNetworkGatewayConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	if args == nil || args.ConnectionType == nil {
		return nil, errors.New("missing required argument 'ConnectionType'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualNetworkGateway1 == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkGateway1'")
	}
	if args == nil {
		args = &VirtualNetworkGatewayConnectionArgs{}
	}
	var resource VirtualNetworkGatewayConnection
	err := ctx.RegisterResource("azurerm:network/v20161201:VirtualNetworkGatewayConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkGatewayConnection gets an existing VirtualNetworkGatewayConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayConnectionState, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	var resource VirtualNetworkGatewayConnection
	err := ctx.ReadResource("azurerm:network/v20161201:VirtualNetworkGatewayConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkGatewayConnection resources.
type virtualNetworkGatewayConnectionState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// VirtualNetworkGatewayConnection properties
	Properties *VirtualNetworkGatewayConnectionPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type VirtualNetworkGatewayConnectionState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// VirtualNetworkGatewayConnection properties
	Properties VirtualNetworkGatewayConnectionPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (VirtualNetworkGatewayConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayConnectionState)(nil)).Elem()
}

type virtualNetworkGatewayConnectionArgs struct {
	// The authorizationKey.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// Gateway connection type. Possible values are: 'IPsec','Vnet2Vnet','ExpressRoute', and 'VPNClient.
	ConnectionType string `pulumi:"connectionType"`
	// EnableBgp flag
	EnableBgp *bool `pulumi:"enableBgp"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A common class for general resource information
	LocalNetworkGateway2 *LocalNetworkGatewayType `pulumi:"localNetworkGateway2"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the virtual network gateway connection.
	Name string `pulumi:"name"`
	// The reference to peerings resource.
	Peer *SubResource `pulumi:"peer"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource GUID property of the VirtualNetworkGatewayConnection resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// The routing weight.
	RoutingWeight *int `pulumi:"routingWeight"`
	// The IPSec shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A common class for general resource information
	VirtualNetworkGateway1 VirtualNetworkGatewayType `pulumi:"virtualNetworkGateway1"`
	// A common class for general resource information
	VirtualNetworkGateway2 *VirtualNetworkGatewayType `pulumi:"virtualNetworkGateway2"`
}

// The set of arguments for constructing a VirtualNetworkGatewayConnection resource.
type VirtualNetworkGatewayConnectionArgs struct {
	// The authorizationKey.
	AuthorizationKey pulumi.StringPtrInput
	// Gateway connection type. Possible values are: 'IPsec','Vnet2Vnet','ExpressRoute', and 'VPNClient.
	ConnectionType pulumi.StringInput
	// EnableBgp flag
	EnableBgp pulumi.BoolPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// A common class for general resource information
	LocalNetworkGateway2 LocalNetworkGatewayTypePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the virtual network gateway connection.
	Name pulumi.StringInput
	// The reference to peerings resource.
	Peer SubResourcePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource GUID property of the VirtualNetworkGatewayConnection resource.
	ResourceGuid pulumi.StringPtrInput
	// The routing weight.
	RoutingWeight pulumi.IntPtrInput
	// The IPSec shared key.
	SharedKey pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A common class for general resource information
	VirtualNetworkGateway1 VirtualNetworkGatewayTypeInput
	// A common class for general resource information
	VirtualNetworkGateway2 VirtualNetworkGatewayTypePtrInput
}

func (VirtualNetworkGatewayConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayConnectionArgs)(nil)).Elem()
}
