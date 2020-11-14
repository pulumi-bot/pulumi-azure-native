// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVirtualNetworkPeering(ctx *pulumi.Context, args *LookupVirtualNetworkPeeringArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkPeeringResult, error) {
	var rv LookupVirtualNetworkPeeringResult
	err := ctx.Invoke("azure-nextgen:network/latest:getVirtualNetworkPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkPeeringArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
	// The name of the virtual network peering.
	VirtualNetworkPeeringName string `pulumi:"virtualNetworkPeeringName"`
}

// Peerings in a virtual network resource.
type LookupVirtualNetworkPeeringResult struct {
	// Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
	AllowForwardedTraffic *bool `pulumi:"allowForwardedTraffic"`
	// If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool `pulumi:"allowGatewayTransit"`
	// Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
	AllowVirtualNetworkAccess *bool `pulumi:"allowVirtualNetworkAccess"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The status of the virtual network peering.
	PeeringState *string `pulumi:"peeringState"`
	// The peering sync status of the virtual network peering.
	PeeringSyncLevel *string `pulumi:"peeringSyncLevel"`
	// The provisioning state of the virtual network peering resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The reference to the address space peered with the remote virtual network.
	RemoteAddressSpace *AddressSpaceResponse `pulumi:"remoteAddressSpace"`
	// The reference to the remote virtual network's Bgp Communities.
	RemoteBgpCommunities *VirtualNetworkBgpCommunitiesResponse `pulumi:"remoteBgpCommunities"`
	// The reference to the remote virtual network. The remote virtual network can be in the same or different region (preview). See here to register for the preview and learn more (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
	RemoteVirtualNetwork *SubResourceResponse `pulumi:"remoteVirtualNetwork"`
	// The reference to the current address space of the remote virtual network.
	RemoteVirtualNetworkAddressSpace *AddressSpaceResponse `pulumi:"remoteVirtualNetworkAddressSpace"`
	// Provided when user wants to sync the peering with address space on the remote virtual network after the address space is updated.
	SyncRemoteAddressSpace *bool `pulumi:"syncRemoteAddressSpace"`
	// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
	UseRemoteGateways *bool `pulumi:"useRemoteGateways"`
}
