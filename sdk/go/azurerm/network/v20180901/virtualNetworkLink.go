// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a link to virtual network for a Private DNS zone.
type VirtualNetworkLink struct {
	pulumi.CustomResourceState

	// The ETag of the virtual network link.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the virtual network link to the Private DNS zone.
	Properties VirtualNetworkLinkPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualNetworkLink registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkLink(ctx *pulumi.Context,
	name string, args *VirtualNetworkLinkArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkLink, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PrivateZoneName == nil {
		return nil, errors.New("missing required argument 'PrivateZoneName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VirtualNetworkLinkArgs{}
	}
	var resource VirtualNetworkLink
	err := ctx.RegisterResource("azurerm:network/v20180901:VirtualNetworkLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkLink gets an existing VirtualNetworkLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkLinkState, opts ...pulumi.ResourceOption) (*VirtualNetworkLink, error) {
	var resource VirtualNetworkLink
	err := ctx.ReadResource("azurerm:network/v20180901:VirtualNetworkLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkLink resources.
type virtualNetworkLinkState struct {
	// The ETag of the virtual network link.
	Etag *string `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Properties of the virtual network link to the Private DNS zone.
	Properties *VirtualNetworkLinkPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `pulumi:"type"`
}

type VirtualNetworkLinkState struct {
	// The ETag of the virtual network link.
	Etag pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Properties of the virtual network link to the Private DNS zone.
	Properties VirtualNetworkLinkPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type pulumi.StringPtrInput
}

func (VirtualNetworkLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkLinkState)(nil)).Elem()
}

type virtualNetworkLinkArgs struct {
	// The ETag of the virtual network link.
	Etag *string `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the virtual network link.
	Name string `pulumi:"name"`
	// The name of the Private DNS zone (without a terminating dot).
	PrivateZoneName string `pulumi:"privateZoneName"`
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled?
	RegistrationEnabled *bool `pulumi:"registrationEnabled"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The reference of the virtual network.
	VirtualNetwork *SubResource `pulumi:"virtualNetwork"`
}

// The set of arguments for constructing a VirtualNetworkLink resource.
type VirtualNetworkLinkArgs struct {
	// The ETag of the virtual network link.
	Etag pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the virtual network link.
	Name pulumi.StringInput
	// The name of the Private DNS zone (without a terminating dot).
	PrivateZoneName pulumi.StringInput
	// Is auto-registration of virtual machine records in the virtual network in the Private DNS zone enabled?
	RegistrationEnabled pulumi.BoolPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The reference of the virtual network.
	VirtualNetwork SubResourcePtrInput
}

func (VirtualNetworkLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkLinkArgs)(nil)).Elem()
}
