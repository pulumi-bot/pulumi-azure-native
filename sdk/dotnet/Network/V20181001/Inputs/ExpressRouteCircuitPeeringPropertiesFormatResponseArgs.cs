// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181001.Inputs
{

    public sealed class ExpressRouteCircuitPeeringPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Azure ASN.
        /// </summary>
        [Input("azureASN")]
        public int? AzureASN { get; set; }

        [Input("connections")]
        private List<Inputs.ExpressRouteCircuitConnectionResponseArgs>? _connections;

        /// <summary>
        /// The list of circuit connections associated with Azure Private Peering for this circuit.
        /// </summary>
        public List<Inputs.ExpressRouteCircuitConnectionResponseArgs> Connections
        {
            get => _connections ?? (_connections = new List<Inputs.ExpressRouteCircuitConnectionResponseArgs>());
            set => _connections = value;
        }

        /// <summary>
        /// The ExpressRoute connection.
        /// </summary>
        [Input("expressRouteConnection")]
        public Inputs.ExpressRouteConnectionIdResponseArgs? ExpressRouteConnection { get; set; }

        /// <summary>
        /// The GatewayManager Etag.
        /// </summary>
        [Input("gatewayManagerEtag")]
        public string? GatewayManagerEtag { get; set; }

        /// <summary>
        /// The IPv6 peering configuration.
        /// </summary>
        [Input("ipv6PeeringConfig")]
        public Inputs.Ipv6ExpressRouteCircuitPeeringConfigResponseArgs? Ipv6PeeringConfig { get; set; }

        /// <summary>
        /// Gets whether the provider or the customer last modified the peering.
        /// </summary>
        [Input("lastModifiedBy")]
        public string? LastModifiedBy { get; set; }

        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        [Input("microsoftPeeringConfig")]
        public Inputs.ExpressRouteCircuitPeeringConfigResponseArgs? MicrosoftPeeringConfig { get; set; }

        /// <summary>
        /// The peer ASN.
        /// </summary>
        [Input("peerASN")]
        public int? PeerASN { get; set; }

        /// <summary>
        /// The peering type.
        /// </summary>
        [Input("peeringType")]
        public string? PeeringType { get; set; }

        /// <summary>
        /// The primary port.
        /// </summary>
        [Input("primaryAzurePort")]
        public string? PrimaryAzurePort { get; set; }

        /// <summary>
        /// The primary address prefix.
        /// </summary>
        [Input("primaryPeerAddressPrefix")]
        public string? PrimaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        /// <summary>
        /// The reference of the RouteFilter resource.
        /// </summary>
        [Input("routeFilter")]
        public Inputs.RouteFilterResponseArgs? RouteFilter { get; set; }

        /// <summary>
        /// The secondary port.
        /// </summary>
        [Input("secondaryAzurePort")]
        public string? SecondaryAzurePort { get; set; }

        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        [Input("secondaryPeerAddressPrefix")]
        public string? SecondaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The shared key.
        /// </summary>
        [Input("sharedKey")]
        public string? SharedKey { get; set; }

        /// <summary>
        /// The peering state.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        /// <summary>
        /// Gets peering stats.
        /// </summary>
        [Input("stats")]
        public Inputs.ExpressRouteCircuitStatsResponseArgs? Stats { get; set; }

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Input("vlanId")]
        public int? VlanId { get; set; }

        public ExpressRouteCircuitPeeringPropertiesFormatResponseArgs()
        {
        }
    }
}
