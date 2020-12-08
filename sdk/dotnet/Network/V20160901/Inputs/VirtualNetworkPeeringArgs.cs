// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20160901.Inputs
{

    /// <summary>
    /// Peerings in a virtual network resource.
    /// </summary>
    public sealed class VirtualNetworkPeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the forwarded traffic from the VMs in the remote virtual network will be allowed/disallowed.
        /// </summary>
        [Input("allowForwardedTraffic")]
        public Input<bool>? AllowForwardedTraffic { get; set; }

        /// <summary>
        /// If gateway links can be used in remote virtual networking to link to this virtual network.
        /// </summary>
        [Input("allowGatewayTransit")]
        public Input<bool>? AllowGatewayTransit { get; set; }

        /// <summary>
        /// Whether the VMs in the linked virtual network space would be able to access all the VMs in local Virtual network space.
        /// </summary>
        [Input("allowVirtualNetworkAccess")]
        public Input<bool>? AllowVirtualNetworkAccess { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The status of the virtual network peering. Possible values are 'Initiated', 'Connected', and 'Disconnected'.
        /// </summary>
        [Input("peeringState")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20160901.VirtualNetworkPeeringState>? PeeringState { get; set; }

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The reference of the remote virtual network.
        /// </summary>
        [Input("remoteVirtualNetwork")]
        public Input<Inputs.SubResourceArgs>? RemoteVirtualNetwork { get; set; }

        /// <summary>
        /// If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
        /// </summary>
        [Input("useRemoteGateways")]
        public Input<bool>? UseRemoteGateways { get; set; }

        public VirtualNetworkPeeringArgs()
        {
        }
    }
}
