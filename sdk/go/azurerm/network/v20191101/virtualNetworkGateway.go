// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A common class for general resource information.
type VirtualNetworkGateway struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the virtual network gateway.
	Properties VirtualNetworkGatewayPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualNetworkGateway registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkGateway(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGateway, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VirtualNetworkGatewayArgs{}
	}
	var resource VirtualNetworkGateway
	err := ctx.RegisterResource("azurerm:network/v20191101:VirtualNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkGateway gets an existing VirtualNetworkGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayState, opts ...pulumi.ResourceOption) (*VirtualNetworkGateway, error) {
	var resource VirtualNetworkGateway
	err := ctx.ReadResource("azurerm:network/v20191101:VirtualNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkGateway resources.
type virtualNetworkGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the virtual network gateway.
	Properties *VirtualNetworkGatewayPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type VirtualNetworkGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the virtual network gateway.
	Properties VirtualNetworkGatewayPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (VirtualNetworkGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayState)(nil)).Elem()
}

type virtualNetworkGatewayArgs struct {
	// ActiveActive flag.
	ActiveActive *bool `pulumi:"activeActive"`
	// Virtual network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `pulumi:"bgpSettings"`
	// The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
	CustomRoutes *AddressSpace `pulumi:"customRoutes"`
	// Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Whether dns forwarding is enabled or not.
	EnableDnsForwarding *bool `pulumi:"enableDnsForwarding"`
	// The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite *SubResource `pulumi:"gatewayDefaultSite"`
	// The type of this virtual network gateway.
	GatewayType *string `pulumi:"gatewayType"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IP configurations for virtual network gateway.
	IpConfigurations []VirtualNetworkGatewayIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the virtual network gateway.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
	Sku *VirtualNetworkGatewaySku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
	VpnClientConfiguration *VpnClientConfiguration `pulumi:"vpnClientConfiguration"`
	// The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
	VpnGatewayGeneration *string `pulumi:"vpnGatewayGeneration"`
	// The type of this virtual network gateway.
	VpnType *string `pulumi:"vpnType"`
}

// The set of arguments for constructing a VirtualNetworkGateway resource.
type VirtualNetworkGatewayArgs struct {
	// ActiveActive flag.
	ActiveActive pulumi.BoolPtrInput
	// Virtual network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
	CustomRoutes AddressSpacePtrInput
	// Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp pulumi.BoolPtrInput
	// Whether dns forwarding is enabled or not.
	EnableDnsForwarding pulumi.BoolPtrInput
	// The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite SubResourcePtrInput
	// The type of this virtual network gateway.
	GatewayType pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// IP configurations for virtual network gateway.
	IpConfigurations VirtualNetworkGatewayIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the virtual network gateway.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
	Sku VirtualNetworkGatewaySkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
	VpnClientConfiguration VpnClientConfigurationPtrInput
	// The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
	VpnGatewayGeneration pulumi.StringPtrInput
	// The type of this virtual network gateway.
	VpnType pulumi.StringPtrInput
}

func (VirtualNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayArgs)(nil)).Elem()
}
