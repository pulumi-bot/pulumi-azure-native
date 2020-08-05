// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A common class for general resource information
type VirtualNetworkGateway struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VirtualNetworkGateway properties
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
	err := ctx.RegisterResource("azurerm:network/v20170301:VirtualNetworkGateway", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20170301:VirtualNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkGateway resources.
type virtualNetworkGatewayState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// VirtualNetworkGateway properties
	Properties *VirtualNetworkGatewayPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type VirtualNetworkGatewayState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// VirtualNetworkGateway properties
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
	// ActiveActive flag
	ActiveActive *bool `pulumi:"activeActive"`
	// Virtual network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `pulumi:"bgpSettings"`
	// Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite *SubResource `pulumi:"gatewayDefaultSite"`
	// The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
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
	// The resource GUID property of the VirtualNetworkGateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// The reference of the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
	Sku *VirtualNetworkGatewaySku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
	VpnClientConfiguration *VpnClientConfiguration `pulumi:"vpnClientConfiguration"`
	// The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'.
	VpnType *string `pulumi:"vpnType"`
}

// The set of arguments for constructing a VirtualNetworkGateway resource.
type VirtualNetworkGatewayArgs struct {
	// ActiveActive flag
	ActiveActive pulumi.BoolPtrInput
	// Virtual network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp pulumi.BoolPtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite SubResourcePtrInput
	// The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
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
	// The resource GUID property of the VirtualNetworkGateway resource.
	ResourceGuid pulumi.StringPtrInput
	// The reference of the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
	Sku VirtualNetworkGatewaySkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
	VpnClientConfiguration VpnClientConfigurationPtrInput
	// The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'.
	VpnType pulumi.StringPtrInput
}

func (VirtualNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayArgs)(nil)).Elem()
}
