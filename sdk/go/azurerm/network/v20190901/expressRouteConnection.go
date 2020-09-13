// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ExpressRouteConnection resource.
//
// ## Example Usage
// ### ExpressRouteConnectionCreate
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20190901"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewExpressRouteConnection(ctx, "expressRouteConnection", &network.ExpressRouteConnectionArgs{
// 			AuthorizationKey: pulumi.String("authorizationKey"),
// 			ConnectionName:   pulumi.String("connectionName"),
// 			ExpressRouteCircuitPeering: &network.ExpressRouteCircuitPeeringIdArgs{
// 				Id: pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering"),
// 			},
// 			ExpressRouteGatewayName: pulumi.String("gateway-2"),
// 			Id:                      pulumi.String("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName"),
// 			Name:                    pulumi.String("connectionName"),
// 			ResourceGroupName:       pulumi.String("resourceGroupName"),
// 			RoutingWeight:           pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ExpressRouteConnection struct {
	pulumi.CustomResourceState

	// Authorization key to establish the connection.
	AuthorizationKey pulumi.StringPtrOutput `pulumi:"authorizationKey"`
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrOutput `pulumi:"enableInternetSecurity"`
	// The ExpressRoute circuit peering.
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringIdResponseOutput `pulumi:"expressRouteCircuitPeering"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the express route connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The routing weight associated to the connection.
	RoutingWeight pulumi.IntPtrOutput `pulumi:"routingWeight"`
}

// NewExpressRouteConnection registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteConnection(ctx *pulumi.Context,
	name string, args *ExpressRouteConnectionArgs, opts ...pulumi.ResourceOption) (*ExpressRouteConnection, error) {
	if args == nil || args.ConnectionName == nil {
		return nil, errors.New("missing required argument 'ConnectionName'")
	}
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:ExpressRouteConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteConnection
	err := ctx.RegisterResource("azurerm:network/v20190901:ExpressRouteConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20190901:ExpressRouteConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteConnection resources.
type expressRouteConnectionState struct {
	// Authorization key to establish the connection.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// The ExpressRoute circuit peering.
	ExpressRouteCircuitPeering *ExpressRouteCircuitPeeringIdResponse `pulumi:"expressRouteCircuitPeering"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the express route connection resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The routing weight associated to the connection.
	RoutingWeight *int `pulumi:"routingWeight"`
}

type ExpressRouteConnectionState struct {
	// Authorization key to establish the connection.
	AuthorizationKey pulumi.StringPtrInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// The ExpressRoute circuit peering.
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringIdResponsePtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning state of the express route connection resource.
	ProvisioningState pulumi.StringPtrInput
	// The routing weight associated to the connection.
	RoutingWeight pulumi.IntPtrInput
}

func (ExpressRouteConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteConnectionState)(nil)).Elem()
}

type expressRouteConnectionArgs struct {
	// Authorization key to establish the connection.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The name of the connection subresource.
	ConnectionName string `pulumi:"connectionName"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// The ExpressRoute circuit peering.
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringId `pulumi:"expressRouteCircuitPeering"`
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName string `pulumi:"expressRouteGatewayName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource.
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
	// The name of the connection subresource.
	ConnectionName pulumi.StringInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// The ExpressRoute circuit peering.
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringIdInput
	// The name of the ExpressRoute gateway.
	ExpressRouteGatewayName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The routing weight associated to the connection.
	RoutingWeight pulumi.IntPtrInput
}

func (ExpressRouteConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteConnectionArgs)(nil)).Elem()
}
