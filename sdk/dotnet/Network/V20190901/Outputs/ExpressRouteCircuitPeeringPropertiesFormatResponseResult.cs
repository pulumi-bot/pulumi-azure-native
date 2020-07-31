// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190901.Outputs
{

    [OutputType]
    public sealed class ExpressRouteCircuitPeeringPropertiesFormatResponseResult
    {
        /// <summary>
        /// The Azure ASN.
        /// </summary>
        public readonly int? AzureASN;
        /// <summary>
        /// The list of circuit connections associated with Azure Private Peering for this circuit.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExpressRouteCircuitConnectionResponseResult> Connections;
        /// <summary>
        /// The ExpressRoute connection.
        /// </summary>
        public readonly Outputs.ExpressRouteConnectionIdResponseResult? ExpressRouteConnection;
        /// <summary>
        /// The GatewayManager Etag.
        /// </summary>
        public readonly string? GatewayManagerEtag;
        /// <summary>
        /// The IPv6 peering configuration.
        /// </summary>
        public readonly Outputs.Ipv6ExpressRouteCircuitPeeringConfigResponseResult? Ipv6PeeringConfig;
        /// <summary>
        /// Who was the last to modify the peering.
        /// </summary>
        public readonly string LastModifiedBy;
        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        public readonly Outputs.ExpressRouteCircuitPeeringConfigResponseResult? MicrosoftPeeringConfig;
        /// <summary>
        /// The peer ASN.
        /// </summary>
        public readonly int? PeerASN;
        /// <summary>
        /// The list of peered circuit connections associated with Azure Private Peering for this circuit.
        /// </summary>
        public readonly ImmutableArray<Outputs.PeerExpressRouteCircuitConnectionResponseResult> PeeredConnections;
        /// <summary>
        /// The peering type.
        /// </summary>
        public readonly string? PeeringType;
        /// <summary>
        /// The primary port.
        /// </summary>
        public readonly string? PrimaryAzurePort;
        /// <summary>
        /// The primary address prefix.
        /// </summary>
        public readonly string? PrimaryPeerAddressPrefix;
        /// <summary>
        /// The provisioning state of the express route circuit peering resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The reference of the RouteFilter resource.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? RouteFilter;
        /// <summary>
        /// The secondary port.
        /// </summary>
        public readonly string? SecondaryAzurePort;
        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        public readonly string? SecondaryPeerAddressPrefix;
        /// <summary>
        /// The shared key.
        /// </summary>
        public readonly string? SharedKey;
        /// <summary>
        /// The peering state.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The peering stats of express route circuit.
        /// </summary>
        public readonly Outputs.ExpressRouteCircuitStatsResponseResult? Stats;
        /// <summary>
        /// The VLAN ID.
        /// </summary>
        public readonly int? VlanId;

        [OutputConstructor]
        private ExpressRouteCircuitPeeringPropertiesFormatResponseResult(
            int? azureASN,

            ImmutableArray<Outputs.ExpressRouteCircuitConnectionResponseResult> connections,

            Outputs.ExpressRouteConnectionIdResponseResult? expressRouteConnection,

            string? gatewayManagerEtag,

            Outputs.Ipv6ExpressRouteCircuitPeeringConfigResponseResult? ipv6PeeringConfig,

            string lastModifiedBy,

            Outputs.ExpressRouteCircuitPeeringConfigResponseResult? microsoftPeeringConfig,

            int? peerASN,

            ImmutableArray<Outputs.PeerExpressRouteCircuitConnectionResponseResult> peeredConnections,

            string? peeringType,

            string? primaryAzurePort,

            string? primaryPeerAddressPrefix,

            string provisioningState,

            Outputs.SubResourceResponseResult? routeFilter,

            string? secondaryAzurePort,

            string? secondaryPeerAddressPrefix,

            string? sharedKey,

            string? state,

            Outputs.ExpressRouteCircuitStatsResponseResult? stats,

            int? vlanId)
        {
            AzureASN = azureASN;
            Connections = connections;
            ExpressRouteConnection = expressRouteConnection;
            GatewayManagerEtag = gatewayManagerEtag;
            Ipv6PeeringConfig = ipv6PeeringConfig;
            LastModifiedBy = lastModifiedBy;
            MicrosoftPeeringConfig = microsoftPeeringConfig;
            PeerASN = peerASN;
            PeeredConnections = peeredConnections;
            PeeringType = peeringType;
            PrimaryAzurePort = primaryAzurePort;
            PrimaryPeerAddressPrefix = primaryPeerAddressPrefix;
            ProvisioningState = provisioningState;
            RouteFilter = routeFilter;
            SecondaryAzurePort = secondaryAzurePort;
            SecondaryPeerAddressPrefix = secondaryPeerAddressPrefix;
            SharedKey = sharedKey;
            State = state;
            Stats = stats;
            VlanId = vlanId;
        }
    }
}