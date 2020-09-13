// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// HubVirtualNetworkConnection Resource.
//
// ## Example Usage
// ### HubVirtualNetworkConnectionPut
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewHubVirtualNetworkConnection(ctx, "hubVirtualNetworkConnection", &network.HubVirtualNetworkConnectionArgs{
// 			ConnectionName:         pulumi.String("connection1"),
// 			EnableInternetSecurity: pulumi.Bool(false),
// 			RemoteVirtualNetwork: &network.SubResourceArgs{
// 				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/SpokeVnet1"),
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			RoutingConfiguration: &network.RoutingConfigurationArgs{
// 				AssociatedRouteTable: &network.SubResourceArgs{
// 					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
// 				},
// 				PropagatedRouteTables: &network.PropagatedRouteTableArgs{
// 					Ids: network.SubResourceArray{
// 						&network.SubResourceArgs{
// 							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
// 						},
// 					},
// 					Labels: pulumi.StringArray{
// 						pulumi.String("label1"),
// 						pulumi.String("label2"),
// 					},
// 				},
// 				VnetRoutes: &network.VnetRouteArgs{
// 					StaticRoutes: network.StaticRouteArray{
// 						&network.StaticRouteArgs{
// 							AddressPrefixes: pulumi.StringArray{
// 								pulumi.String("10.1.0.0/16"),
// 								pulumi.String("10.2.0.0/16"),
// 							},
// 							Name:             pulumi.String("route1"),
// 							NextHopIpAddress: pulumi.String("10.0.0.68"),
// 						},
// 						&network.StaticRouteArgs{
// 							AddressPrefixes: pulumi.StringArray{
// 								pulumi.String("10.3.0.0/16"),
// 								pulumi.String("10.4.0.0/16"),
// 							},
// 							Name:             pulumi.String("route2"),
// 							NextHopIpAddress: pulumi.String("10.0.0.65"),
// 						},
// 					},
// 				},
// 			},
// 			VirtualHubName: pulumi.String("virtualHub1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type HubVirtualNetworkConnection struct {
	pulumi.CustomResourceState

	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit pulumi.BoolPtrOutput `pulumi:"allowHubToRemoteVnetTransit"`
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways pulumi.BoolPtrOutput `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrOutput `pulumi:"enableInternetSecurity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the hub virtual network connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Reference to the remote virtual network.
	RemoteVirtualNetwork SubResourceResponsePtrOutput `pulumi:"remoteVirtualNetwork"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationResponsePtrOutput `pulumi:"routingConfiguration"`
}

// NewHubVirtualNetworkConnection registers a new resource with the given unique name, arguments, and options.
func NewHubVirtualNetworkConnection(ctx *pulumi.Context,
	name string, args *HubVirtualNetworkConnectionArgs, opts ...pulumi.ResourceOption) (*HubVirtualNetworkConnection, error) {
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
		args = &HubVirtualNetworkConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:HubVirtualNetworkConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource HubVirtualNetworkConnection
	err := ctx.RegisterResource("azurerm:network/v20200601:HubVirtualNetworkConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHubVirtualNetworkConnection gets an existing HubVirtualNetworkConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHubVirtualNetworkConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubVirtualNetworkConnectionState, opts ...pulumi.ResourceOption) (*HubVirtualNetworkConnection, error) {
	var resource HubVirtualNetworkConnection
	err := ctx.ReadResource("azurerm:network/v20200601:HubVirtualNetworkConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HubVirtualNetworkConnection resources.
type hubVirtualNetworkConnectionState struct {
	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit *bool `pulumi:"allowHubToRemoteVnetTransit"`
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways *bool `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the hub virtual network connection resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Reference to the remote virtual network.
	RemoteVirtualNetwork *SubResourceResponse `pulumi:"remoteVirtualNetwork"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfigurationResponse `pulumi:"routingConfiguration"`
}

type HubVirtualNetworkConnectionState struct {
	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit pulumi.BoolPtrInput
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways pulumi.BoolPtrInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The provisioning state of the hub virtual network connection resource.
	ProvisioningState pulumi.StringPtrInput
	// Reference to the remote virtual network.
	RemoteVirtualNetwork SubResourceResponsePtrInput
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationResponsePtrInput
}

func (HubVirtualNetworkConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubVirtualNetworkConnectionState)(nil)).Elem()
}

type hubVirtualNetworkConnectionArgs struct {
	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit *bool `pulumi:"allowHubToRemoteVnetTransit"`
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways *bool `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	// The name of the HubVirtualNetworkConnection.
	ConnectionName string `pulumi:"connectionName"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Reference to the remote virtual network.
	RemoteVirtualNetwork *SubResource `pulumi:"remoteVirtualNetwork"`
	// The resource group name of the HubVirtualNetworkConnection.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfiguration `pulumi:"routingConfiguration"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The set of arguments for constructing a HubVirtualNetworkConnection resource.
type HubVirtualNetworkConnectionArgs struct {
	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit pulumi.BoolPtrInput
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways pulumi.BoolPtrInput
	// The name of the HubVirtualNetworkConnection.
	ConnectionName pulumi.StringInput
	// Enable internet security.
	EnableInternetSecurity pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Reference to the remote virtual network.
	RemoteVirtualNetwork SubResourcePtrInput
	// The resource group name of the HubVirtualNetworkConnection.
	ResourceGroupName pulumi.StringInput
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration RoutingConfigurationPtrInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
}

func (HubVirtualNetworkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubVirtualNetworkConnectionArgs)(nil)).Elem()
}
