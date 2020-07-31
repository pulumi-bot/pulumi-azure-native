// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170601.Inputs
{

    public sealed class ExpressRouteCircuitPeeringPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure ASN.
        /// </summary>
        [Input("azureASN")]
        public Input<int>? AzureASN { get; set; }

        /// <summary>
        /// The GatewayManager Etag.
        /// </summary>
        [Input("gatewayManagerEtag")]
        public Input<string>? GatewayManagerEtag { get; set; }

        /// <summary>
        /// The IPv6 peering configuration.
        /// </summary>
        [Input("ipv6PeeringConfig")]
        public Input<Inputs.Ipv6ExpressRouteCircuitPeeringConfigArgs>? Ipv6PeeringConfig { get; set; }

        /// <summary>
        /// Gets whether the provider or the customer last modified the peering.
        /// </summary>
        [Input("lastModifiedBy")]
        public Input<string>? LastModifiedBy { get; set; }

        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        [Input("microsoftPeeringConfig")]
        public Input<Inputs.ExpressRouteCircuitPeeringConfigArgs>? MicrosoftPeeringConfig { get; set; }

        /// <summary>
        /// The peer ASN.
        /// </summary>
        [Input("peerASN")]
        public Input<int>? PeerASN { get; set; }

        /// <summary>
        /// The PeeringType. Possible values are: 'AzurePublicPeering', 'AzurePrivatePeering', and 'MicrosoftPeering'.
        /// </summary>
        [Input("peeringType")]
        public Input<string>? PeeringType { get; set; }

        /// <summary>
        /// The primary port.
        /// </summary>
        [Input("primaryAzurePort")]
        public Input<string>? PrimaryAzurePort { get; set; }

        /// <summary>
        /// The primary address prefix.
        /// </summary>
        [Input("primaryPeerAddressPrefix")]
        public Input<string>? PrimaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The reference of the RouteFilter resource.
        /// </summary>
        [Input("routeFilter")]
        public Input<Inputs.RouteFilterDefinitionArgs>? RouteFilter { get; set; }

        /// <summary>
        /// The secondary port.
        /// </summary>
        [Input("secondaryAzurePort")]
        public Input<string>? SecondaryAzurePort { get; set; }

        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        [Input("secondaryPeerAddressPrefix")]
        public Input<string>? SecondaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The shared key.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        /// <summary>
        /// The state of peering. Possible values are: 'Disabled' and 'Enabled'
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Gets peering stats.
        /// </summary>
        [Input("stats")]
        public Input<Inputs.ExpressRouteCircuitStatsArgs>? Stats { get; set; }

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public ExpressRouteCircuitPeeringPropertiesFormatArgs()
        {
        }
    }
}