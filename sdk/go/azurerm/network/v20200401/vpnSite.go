// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VpnSite Resource.
type VpnSite struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the VPN site.
	Properties VpnSitePropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVpnSite registers a new resource with the given unique name, arguments, and options.
func NewVpnSite(ctx *pulumi.Context,
	name string, args *VpnSiteArgs, opts ...pulumi.ResourceOption) (*VpnSite, error) {
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
		args = &VpnSiteArgs{}
	}
	var resource VpnSite
	err := ctx.RegisterResource("azurerm:network/v20200401:VpnSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnSite gets an existing VpnSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnSiteState, opts ...pulumi.ResourceOption) (*VpnSite, error) {
	var resource VpnSite
	err := ctx.ReadResource("azurerm:network/v20200401:VpnSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnSite resources.
type vpnSiteState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the VPN site.
	Properties *VpnSitePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type VpnSiteState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the VPN site.
	Properties VpnSitePropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (VpnSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteState)(nil)).Elem()
}

type vpnSiteArgs struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace *AddressSpace `pulumi:"addressSpace"`
	// The set of bgp properties.
	BgpProperties *BgpSettings `pulumi:"bgpProperties"`
	// The device properties.
	DeviceProperties *DeviceProperties `pulumi:"deviceProperties"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The ip-address for the vpn-site.
	IpAddress *string `pulumi:"ipAddress"`
	// IsSecuritySite flag.
	IsSecuritySite *bool `pulumi:"isSecuritySite"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the VpnSite being created or updated.
	Name string `pulumi:"name"`
	// The resource group name of the VpnSite.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The key for vpn-site that can be used for connections.
	SiteKey *string `pulumi:"siteKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan *SubResource `pulumi:"virtualWan"`
	// List of all vpn site links.
	VpnSiteLinks []VpnSiteLink `pulumi:"vpnSiteLinks"`
}

// The set of arguments for constructing a VpnSite resource.
type VpnSiteArgs struct {
	// The AddressSpace that contains an array of IP address ranges.
	AddressSpace AddressSpacePtrInput
	// The set of bgp properties.
	BgpProperties BgpSettingsPtrInput
	// The device properties.
	DeviceProperties DevicePropertiesPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The ip-address for the vpn-site.
	IpAddress pulumi.StringPtrInput
	// IsSecuritySite flag.
	IsSecuritySite pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The name of the VpnSite being created or updated.
	Name pulumi.StringInput
	// The resource group name of the VpnSite.
	ResourceGroupName pulumi.StringInput
	// The key for vpn-site that can be used for connections.
	SiteKey pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The VirtualWAN to which the vpnSite belongs.
	VirtualWan SubResourcePtrInput
	// List of all vpn site links.
	VpnSiteLinks VpnSiteLinkArrayInput
}

func (VpnSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnSiteArgs)(nil)).Elem()
}
