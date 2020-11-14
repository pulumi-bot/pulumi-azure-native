// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Peerings in a VirtualNetwork resource
type VNetPeering struct {
	pulumi.CustomResourceState

	// Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
	AllowForwardedTraffic pulumi.BoolPtrOutput `pulumi:"allowForwardedTraffic"`
	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit pulumi.BoolPtrOutput `pulumi:"allowGatewayTransit"`
	// Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
	AllowVirtualNetworkAccess pulumi.BoolPtrOutput `pulumi:"allowVirtualNetworkAccess"`
	// The reference to the databricks virtual network address space.
	DatabricksAddressSpace AddressSpaceResponsePtrOutput `pulumi:"databricksAddressSpace"`
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	DatabricksVirtualNetwork VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput `pulumi:"databricksVirtualNetwork"`
	// Name of the virtual network peering resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the virtual network peering.
	PeeringState pulumi.StringOutput `pulumi:"peeringState"`
	// The provisioning state of the virtual network peering resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The reference to the remote virtual network address space.
	RemoteAddressSpace AddressSpaceResponsePtrOutput `pulumi:"remoteAddressSpace"`
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	RemoteVirtualNetwork VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput `pulumi:"remoteVirtualNetwork"`
	// type of the virtual network peering resource
	Type pulumi.StringOutput `pulumi:"type"`
	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
	UseRemoteGateways pulumi.BoolPtrOutput `pulumi:"useRemoteGateways"`
}

// NewVNetPeering registers a new resource with the given unique name, arguments, and options.
func NewVNetPeering(ctx *pulumi.Context,
	name string, args *VNetPeeringArgs, opts ...pulumi.ResourceOption) (*VNetPeering, error) {
	if args == nil || args.PeeringName == nil {
		return nil, errors.New("missing required argument 'PeeringName'")
	}
	if args == nil || args.RemoteVirtualNetwork == nil {
		return nil, errors.New("missing required argument 'RemoteVirtualNetwork'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &VNetPeeringArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:databricks/latest:vNetPeering"),
		},
	})
	opts = append(opts, aliases)
	var resource VNetPeering
	err := ctx.RegisterResource("azure-nextgen:databricks/v20180401:vNetPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVNetPeering gets an existing VNetPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVNetPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VNetPeeringState, opts ...pulumi.ResourceOption) (*VNetPeering, error) {
	var resource VNetPeering
	err := ctx.ReadResource("azure-nextgen:databricks/v20180401:vNetPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VNetPeering resources.
type vnetPeeringState struct {
	// Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
	AllowForwardedTraffic *bool `pulumi:"allowForwardedTraffic"`
	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool `pulumi:"allowGatewayTransit"`
	// Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
	AllowVirtualNetworkAccess *bool `pulumi:"allowVirtualNetworkAccess"`
	// The reference to the databricks virtual network address space.
	DatabricksAddressSpace *AddressSpaceResponse `pulumi:"databricksAddressSpace"`
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	DatabricksVirtualNetwork *VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork `pulumi:"databricksVirtualNetwork"`
	// Name of the virtual network peering resource
	Name *string `pulumi:"name"`
	// The status of the virtual network peering.
	PeeringState *string `pulumi:"peeringState"`
	// The provisioning state of the virtual network peering resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The reference to the remote virtual network address space.
	RemoteAddressSpace *AddressSpaceResponse `pulumi:"remoteAddressSpace"`
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	RemoteVirtualNetwork *VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork `pulumi:"remoteVirtualNetwork"`
	// type of the virtual network peering resource
	Type *string `pulumi:"type"`
	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
	UseRemoteGateways *bool `pulumi:"useRemoteGateways"`
}

type VNetPeeringState struct {
	// Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
	AllowForwardedTraffic pulumi.BoolPtrInput
	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit pulumi.BoolPtrInput
	// Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
	AllowVirtualNetworkAccess pulumi.BoolPtrInput
	// The reference to the databricks virtual network address space.
	DatabricksAddressSpace AddressSpaceResponsePtrInput
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	DatabricksVirtualNetwork VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrInput
	// Name of the virtual network peering resource
	Name pulumi.StringPtrInput
	// The status of the virtual network peering.
	PeeringState pulumi.StringPtrInput
	// The provisioning state of the virtual network peering resource.
	ProvisioningState pulumi.StringPtrInput
	// The reference to the remote virtual network address space.
	RemoteAddressSpace AddressSpaceResponsePtrInput
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	RemoteVirtualNetwork VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrInput
	// type of the virtual network peering resource
	Type pulumi.StringPtrInput
	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
	UseRemoteGateways pulumi.BoolPtrInput
}

func (VNetPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*vnetPeeringState)(nil)).Elem()
}

type vnetPeeringArgs struct {
	// Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
	AllowForwardedTraffic *bool `pulumi:"allowForwardedTraffic"`
	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool `pulumi:"allowGatewayTransit"`
	// Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
	AllowVirtualNetworkAccess *bool `pulumi:"allowVirtualNetworkAccess"`
	// The reference to the databricks virtual network address space.
	DatabricksAddressSpace *AddressSpace `pulumi:"databricksAddressSpace"`
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	DatabricksVirtualNetwork *VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork `pulumi:"databricksVirtualNetwork"`
	// The name of the workspace vNet peering.
	PeeringName string `pulumi:"peeringName"`
	// The reference to the remote virtual network address space.
	RemoteAddressSpace *AddressSpace `pulumi:"remoteAddressSpace"`
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	RemoteVirtualNetwork VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork `pulumi:"remoteVirtualNetwork"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
	UseRemoteGateways *bool `pulumi:"useRemoteGateways"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a VNetPeering resource.
type VNetPeeringArgs struct {
	// Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
	AllowForwardedTraffic pulumi.BoolPtrInput
	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit pulumi.BoolPtrInput
	// Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
	AllowVirtualNetworkAccess pulumi.BoolPtrInput
	// The reference to the databricks virtual network address space.
	DatabricksAddressSpace AddressSpacePtrInput
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	DatabricksVirtualNetwork VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrInput
	// The name of the workspace vNet peering.
	PeeringName pulumi.StringInput
	// The reference to the remote virtual network address space.
	RemoteAddressSpace AddressSpacePtrInput
	//  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
	RemoteVirtualNetwork VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
	UseRemoteGateways pulumi.BoolPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (VNetPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vnetPeeringArgs)(nil)).Elem()
}

type VNetPeeringInput interface {
	pulumi.Input

	ToVNetPeeringOutput() VNetPeeringOutput
	ToVNetPeeringOutputWithContext(ctx context.Context) VNetPeeringOutput
}

func (VNetPeering) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetPeering)(nil)).Elem()
}

func (i VNetPeering) ToVNetPeeringOutput() VNetPeeringOutput {
	return i.ToVNetPeeringOutputWithContext(context.Background())
}

func (i VNetPeering) ToVNetPeeringOutputWithContext(ctx context.Context) VNetPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetPeeringOutput)
}

type VNetPeeringOutput struct {
	*pulumi.OutputState
}

func (VNetPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetPeeringOutput)(nil)).Elem()
}

func (o VNetPeeringOutput) ToVNetPeeringOutput() VNetPeeringOutput {
	return o
}

func (o VNetPeeringOutput) ToVNetPeeringOutputWithContext(ctx context.Context) VNetPeeringOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VNetPeeringOutput{})
}
