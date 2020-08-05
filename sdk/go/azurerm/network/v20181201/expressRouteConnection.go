// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ExpressRouteConnection resource.
type ExpressRouteConnection struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the ExpressRouteConnection subresource.
	Properties ExpressRouteConnectionPropertiesResponseOutput `pulumi:"properties"`
}

// NewExpressRouteConnection registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteConnection(ctx *pulumi.Context,
	name string, args *ExpressRouteConnectionArgs, opts ...pulumi.ResourceOption) (*ExpressRouteConnection, error) {
	if args == nil || args.ExpressRouteCircuitPeering == nil {
		return nil, errors.New("missing required argument 'ExpressRouteCircuitPeering'")
	}
	if args == nil || args.ExpressRouteGatewayName == nil {
		return nil, errors.New("missing required argument 'ExpressRouteGatewayName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ExpressRouteConnectionArgs{}
	}
	var resource ExpressRouteConnection
	err := ctx.RegisterResource("azurerm:network/v20181201:ExpressRouteConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteConnection gets an existing ExpressRouteConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteConnectionState, opts ...pulumi.ResourceOption) (*ExpressRouteConnection, error) {
	var resource ExpressRouteConnection
	err := ctx.ReadResource("azurerm:network/v20181201:ExpressRouteConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteConnection resources.
type expressRouteConnectionState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Properties of the ExpressRouteConnection subresource.
	Properties *ExpressRouteConnectionPropertiesResponse `pulumi:"properties"`
}

type ExpressRouteConnectionState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Properties of the ExpressRouteConnection subresource.
	Properties ExpressRouteConnectionPropertiesResponsePtrInput
}

func (ExpressRouteConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteConnectionState)(nil)).Elem()
}

type expressRouteConnectionArgs struct {
	// Authorization key to establish the connection.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The ExpressRoute circuit peering.
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringId `pulumi:"expressRouteCircuitPeering"`
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName string `pulumi:"expressRouteGatewayName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the connection subresource.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The routing weight associated to the connection.
	RoutingWeight *int `pulumi:"routingWeight"`
}

// The set of arguments for constructing a ExpressRouteConnection resource.
type ExpressRouteConnectionArgs struct {
	// Authorization key to establish the connection.
	AuthorizationKey pulumi.StringPtrInput
	// The ExpressRoute circuit peering.
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringIdInput
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the connection subresource.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The routing weight associated to the connection.
	RoutingWeight pulumi.IntPtrInput
}

func (ExpressRouteConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteConnectionArgs)(nil)).Elem()
}
