// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Peerings in a virtual network resource.
type VirtualNetworkPeering struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       pulumi.StringPtrOutput                              `pulumi:"name"`
	Properties VirtualNetworkPeeringPropertiesFormatResponseOutput `pulumi:"properties"`
}

// NewVirtualNetworkPeering registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkPeering(ctx *pulumi.Context,
	name string, args *VirtualNetworkPeeringArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkPeering, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualNetworkName == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkName'")
	}
	if args == nil {
		args = &VirtualNetworkPeeringArgs{}
	}
	var resource VirtualNetworkPeering
	err := ctx.RegisterResource("azurerm:network/v20161201:VirtualNetworkPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkPeering gets an existing VirtualNetworkPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkPeeringState, opts ...pulumi.ResourceOption) (*VirtualNetworkPeering, error) {
	var resource VirtualNetworkPeering
	err := ctx.ReadResource("azurerm:network/v20161201:VirtualNetworkPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkPeering resources.
type virtualNetworkPeeringState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       *string                                        `pulumi:"name"`
	Properties *VirtualNetworkPeeringPropertiesFormatResponse `pulumi:"properties"`
}

type VirtualNetworkPeeringState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name       pulumi.StringPtrInput
	Properties VirtualNetworkPeeringPropertiesFormatResponsePtrInput
}

func (VirtualNetworkPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkPeeringState)(nil)).Elem()
}

type virtualNetworkPeeringArgs struct {
	// Whether the forwarded traffic from the VMs in the remote virtual network will be allowed/disallowed.
	AllowForwardedTraffic *bool `pulumi:"allowForwardedTraffic"`
	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool `pulumi:"allowGatewayTransit"`
	// Whether the VMs in the linked virtual network space would be able to access all the VMs in local Virtual network space.
	AllowVirtualNetworkAccess *bool `pulumi:"allowVirtualNetworkAccess"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the peering.
	Name string `pulumi:"name"`
	// The status of the virtual network peering. Possible values are 'Initiated', 'Connected', and 'Disconnected'.
	PeeringState *string `pulumi:"peeringState"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The reference of the remote virtual network.
	RemoteVirtualNetwork *SubResource `pulumi:"remoteVirtualNetwork"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
	UseRemoteGateways *bool `pulumi:"useRemoteGateways"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// The set of arguments for constructing a VirtualNetworkPeering resource.
type VirtualNetworkPeeringArgs struct {
	// Whether the forwarded traffic from the VMs in the remote virtual network will be allowed/disallowed.
	AllowForwardedTraffic pulumi.BoolPtrInput
	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit pulumi.BoolPtrInput
	// Whether the VMs in the linked virtual network space would be able to access all the VMs in local Virtual network space.
	AllowVirtualNetworkAccess pulumi.BoolPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the peering.
	Name pulumi.StringInput
	// The status of the virtual network peering. Possible values are 'Initiated', 'Connected', and 'Disconnected'.
	PeeringState pulumi.StringPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The reference of the remote virtual network.
	RemoteVirtualNetwork SubResourcePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
	UseRemoteGateways pulumi.BoolPtrInput
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringInput
}

func (VirtualNetworkPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkPeeringArgs)(nil)).Elem()
}
