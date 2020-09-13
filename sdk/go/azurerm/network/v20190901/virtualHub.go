// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VirtualHub Resource.
//
// ## Example Usage
// ### VirtualHubPut
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
// 		_, err := network.NewVirtualHub(ctx, "virtualHub", &network.VirtualHubArgs{
// 			AddressPrefix:     pulumi.String("10.168.0.0/24"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			Sku:               pulumi.String("Basic"),
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
// 			},
// 			VirtualHubName: pulumi.String("virtualHub2"),
// 			VirtualWan: &network.SubResourceArgs{
// 				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
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
type VirtualHub struct {
	pulumi.CustomResourceState

	// Address-prefix for this VirtualHub.
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall SubResourceResponsePtrOutput `pulumi:"azureFirewall"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway SubResourceResponsePtrOutput `pulumi:"expressRouteGateway"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway SubResourceResponsePtrOutput `pulumi:"p2SVpnGateway"`
	// The provisioning state of the virtual hub resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The routeTable associated with this virtual hub.
	RouteTable VirtualHubRouteTableResponsePtrOutput `pulumi:"routeTable"`
	// The Security Provider name.
	SecurityProviderName pulumi.StringPtrOutput `pulumi:"securityProviderName"`
	// The sku of this VirtualHub.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of all virtual hub route table v2s associated with this VirtualHub.
	VirtualHubRouteTableV2s VirtualHubRouteTableV2ResponseArrayOutput `pulumi:"virtualHubRouteTableV2s"`
	// List of all vnet connections with this VirtualHub.
	VirtualNetworkConnections HubVirtualNetworkConnectionResponseArrayOutput `pulumi:"virtualNetworkConnections"`
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan SubResourceResponsePtrOutput `pulumi:"virtualWan"`
	// The VpnGateway associated with this VirtualHub.
	VpnGateway SubResourceResponsePtrOutput `pulumi:"vpnGateway"`
}

// NewVirtualHub registers a new resource with the given unique name, arguments, and options.
func NewVirtualHub(ctx *pulumi.Context,
	name string, args *VirtualHubArgs, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualHubName == nil {
		return nil, errors.New("missing required argument 'VirtualHubName'")
	}
	if args == nil {
		args = &VirtualHubArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:VirtualHub"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:VirtualHub"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualHub
	err := ctx.RegisterResource("azurerm:network/v20190901:VirtualHub", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20190901:VirtualHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHub resources.
type virtualHubState struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall *SubResourceResponse `pulumi:"azureFirewall"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway *SubResourceResponse `pulumi:"expressRouteGateway"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway *SubResourceResponse `pulumi:"p2SVpnGateway"`
	// The provisioning state of the virtual hub resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The routeTable associated with this virtual hub.
	RouteTable *VirtualHubRouteTableResponse `pulumi:"routeTable"`
	// The Security Provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// The sku of this VirtualHub.
	Sku *string `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// List of all virtual hub route table v2s associated with this VirtualHub.
	VirtualHubRouteTableV2s []VirtualHubRouteTableV2Response `pulumi:"virtualHubRouteTableV2s"`
	// List of all vnet connections with this VirtualHub.
	VirtualNetworkConnections []HubVirtualNetworkConnectionResponse `pulumi:"virtualNetworkConnections"`
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan *SubResourceResponse `pulumi:"virtualWan"`
	// The VpnGateway associated with this VirtualHub.
	VpnGateway *SubResourceResponse `pulumi:"vpnGateway"`
}

type VirtualHubState struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix pulumi.StringPtrInput
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall SubResourceResponsePtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway SubResourceResponsePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway SubResourceResponsePtrInput
	// The provisioning state of the virtual hub resource.
	ProvisioningState pulumi.StringPtrInput
	// The routeTable associated with this virtual hub.
	RouteTable VirtualHubRouteTableResponsePtrInput
	// The Security Provider name.
	SecurityProviderName pulumi.StringPtrInput
	// The sku of this VirtualHub.
	Sku pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// List of all virtual hub route table v2s associated with this VirtualHub.
	VirtualHubRouteTableV2s VirtualHubRouteTableV2ResponseArrayInput
	// List of all vnet connections with this VirtualHub.
	VirtualNetworkConnections HubVirtualNetworkConnectionResponseArrayInput
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan SubResourceResponsePtrInput
	// The VpnGateway associated with this VirtualHub.
	VpnGateway SubResourceResponsePtrInput
}

func (VirtualHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubState)(nil)).Elem()
}

type virtualHubArgs struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall *SubResource `pulumi:"azureFirewall"`
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway *SubResource `pulumi:"expressRouteGateway"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway *SubResource `pulumi:"p2SVpnGateway"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The routeTable associated with this virtual hub.
	RouteTable *VirtualHubRouteTable `pulumi:"routeTable"`
	// The Security Provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// The sku of this VirtualHub.
	Sku *string `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
	// List of all virtual hub route table v2s associated with this VirtualHub.
	VirtualHubRouteTableV2s []VirtualHubRouteTableV2Type `pulumi:"virtualHubRouteTableV2s"`
	// List of all vnet connections with this VirtualHub.
	VirtualNetworkConnections []HubVirtualNetworkConnection `pulumi:"virtualNetworkConnections"`
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan *SubResource `pulumi:"virtualWan"`
	// The VpnGateway associated with this VirtualHub.
	VpnGateway *SubResource `pulumi:"vpnGateway"`
}

// The set of arguments for constructing a VirtualHub resource.
type VirtualHubArgs struct {
	// Address-prefix for this VirtualHub.
	AddressPrefix pulumi.StringPtrInput
	// The azureFirewall associated with this VirtualHub.
	AzureFirewall SubResourcePtrInput
	// The expressRouteGateway associated with this VirtualHub.
	ExpressRouteGateway SubResourcePtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The P2SVpnGateway associated with this VirtualHub.
	P2SVpnGateway SubResourcePtrInput
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput
	// The routeTable associated with this virtual hub.
	RouteTable VirtualHubRouteTablePtrInput
	// The Security Provider name.
	SecurityProviderName pulumi.StringPtrInput
	// The sku of this VirtualHub.
	Sku pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
	// List of all virtual hub route table v2s associated with this VirtualHub.
	VirtualHubRouteTableV2s VirtualHubRouteTableV2TypeArrayInput
	// List of all vnet connections with this VirtualHub.
	VirtualNetworkConnections HubVirtualNetworkConnectionArrayInput
	// The VirtualWAN to which the VirtualHub belongs.
	VirtualWan SubResourcePtrInput
	// The VpnGateway associated with this VirtualHub.
	VpnGateway SubResourcePtrInput
}

func (VirtualHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubArgs)(nil)).Elem()
}
