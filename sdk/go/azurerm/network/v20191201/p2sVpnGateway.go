// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// P2SVpnGateway Resource.
//
// ## Example Usage
// ### P2SVpnGatewayPut
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20191201"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewP2sVpnGateway(ctx, "p2sVpnGateway", &network.P2sVpnGatewayArgs{
// 			GatewayName: pulumi.String("p2sVpnGateway1"),
// 			Location:    pulumi.String("West US"),
// 			P2SConnectionConfigurations: network.P2SConnectionConfigurationArray{
// 				&network.P2SConnectionConfigurationArgs{
// 					Id:   pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1"),
// 					Name: pulumi.String("P2SConnectionConfig1"),
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
// 			},
// 			VirtualHub: &network.SubResourceArgs{
// 				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
// 			},
// 			VpnGatewayScaleUnit: pulumi.Int(1),
// 			VpnServerConfiguration: &network.SubResourceArgs{
// 				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type P2sVpnGateway struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations P2SConnectionConfigurationResponseArrayOutput `pulumi:"p2SConnectionConfigurations"`
	// The provisioning state of the P2S VPN gateway resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
	// All P2S VPN clients' connection health status.
	VpnClientConnectionHealth VpnClientConnectionHealthResponseOutput `pulumi:"vpnClientConnectionHealth"`
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit pulumi.IntPtrOutput `pulumi:"vpnGatewayScaleUnit"`
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration SubResourceResponsePtrOutput `pulumi:"vpnServerConfiguration"`
}

// NewP2sVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewP2sVpnGateway(ctx *pulumi.Context,
	name string, args *P2sVpnGatewayArgs, opts ...pulumi.ResourceOption) (*P2sVpnGateway, error) {
	if args == nil || args.GatewayName == nil {
		return nil, errors.New("missing required argument 'GatewayName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &P2sVpnGatewayArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:P2sVpnGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource P2sVpnGateway
	err := ctx.RegisterResource("azurerm:network/v20191201:P2sVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetP2sVpnGateway gets an existing P2sVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetP2sVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *P2sVpnGatewayState, opts ...pulumi.ResourceOption) (*P2sVpnGateway, error) {
	var resource P2sVpnGateway
	err := ctx.ReadResource("azurerm:network/v20191201:P2sVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering P2sVpnGateway resources.
type p2sVpnGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations []P2SConnectionConfigurationResponse `pulumi:"p2SConnectionConfigurations"`
	// The provisioning state of the P2S VPN gateway resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
	// All P2S VPN clients' connection health status.
	VpnClientConnectionHealth *VpnClientConnectionHealthResponse `pulumi:"vpnClientConnectionHealth"`
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration *SubResourceResponse `pulumi:"vpnServerConfiguration"`
}

type P2sVpnGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations P2SConnectionConfigurationResponseArrayInput
	// The provisioning state of the P2S VPN gateway resource.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The VirtualHub to which the gateway belongs.
	VirtualHub SubResourceResponsePtrInput
	// All P2S VPN clients' connection health status.
	VpnClientConnectionHealth VpnClientConnectionHealthResponsePtrInput
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit pulumi.IntPtrInput
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration SubResourceResponsePtrInput
}

func (P2sVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*p2sVpnGatewayState)(nil)).Elem()
}

type p2sVpnGatewayArgs struct {
	// The name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations []P2SConnectionConfiguration `pulumi:"p2SConnectionConfigurations"`
	// The resource group name of the P2SVpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub *SubResource `pulumi:"virtualHub"`
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration *SubResource `pulumi:"vpnServerConfiguration"`
}

// The set of arguments for constructing a P2sVpnGateway resource.
type P2sVpnGatewayArgs struct {
	// The name of the gateway.
	GatewayName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations P2SConnectionConfigurationArrayInput
	// The resource group name of the P2SVpnGateway.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The VirtualHub to which the gateway belongs.
	VirtualHub SubResourcePtrInput
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit pulumi.IntPtrInput
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration SubResourcePtrInput
}

func (P2sVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*p2sVpnGatewayArgs)(nil)).Elem()
}
