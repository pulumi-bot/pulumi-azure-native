// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190801.Outputs
{

    [OutputType]
    public sealed class VirtualNetworkPeeringPropertiesFormatResponseResult
    {
        /// <summary>
        /// Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
        /// </summary>
        public readonly bool? AllowForwardedTraffic;
        /// <summary>
        /// If gateway links can be used in remote virtual networking to link to this virtual network.
        /// </summary>
        public readonly bool? AllowGatewayTransit;
        /// <summary>
        /// Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
        /// </summary>
        public readonly bool? AllowVirtualNetworkAccess;
        /// <summary>
        /// The status of the virtual network peering.
        /// </summary>
        public readonly string? PeeringState;
        /// <summary>
        /// The provisioning state of the virtual network peering resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The reference of the remote virtual network address space.
        /// </summary>
        public readonly Outputs.AddressSpaceResponseResult? RemoteAddressSpace;
        /// <summary>
        /// The reference of the remote virtual network. The remote virtual network can be in the same or different region (preview). See here to register for the preview and learn more (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? RemoteVirtualNetwork;
        /// <summary>
        /// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
        /// </summary>
        public readonly bool? UseRemoteGateways;

        [OutputConstructor]
        private VirtualNetworkPeeringPropertiesFormatResponseResult(
            bool? allowForwardedTraffic,

            bool? allowGatewayTransit,

            bool? allowVirtualNetworkAccess,

            string? peeringState,

            string? provisioningState,

            Outputs.AddressSpaceResponseResult? remoteAddressSpace,

            Outputs.SubResourceResponseResult? remoteVirtualNetwork,

            bool? useRemoteGateways)
        {
            AllowForwardedTraffic = allowForwardedTraffic;
            AllowGatewayTransit = allowGatewayTransit;
            AllowVirtualNetworkAccess = allowVirtualNetworkAccess;
            PeeringState = peeringState;
            ProvisioningState = provisioningState;
            RemoteAddressSpace = remoteAddressSpace;
            RemoteVirtualNetwork = remoteVirtualNetwork;
            UseRemoteGateways = useRemoteGateways;
        }
    }
}