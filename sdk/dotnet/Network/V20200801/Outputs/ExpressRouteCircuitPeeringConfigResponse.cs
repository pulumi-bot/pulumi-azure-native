// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801.Outputs
{

    [OutputType]
    public sealed class ExpressRouteCircuitPeeringConfigResponse
    {
        /// <summary>
        /// The communities of bgp peering. Specified for microsoft peering.
        /// </summary>
        public readonly ImmutableArray<string> AdvertisedCommunities;
        /// <summary>
        /// The reference to AdvertisedPublicPrefixes.
        /// </summary>
        public readonly ImmutableArray<string> AdvertisedPublicPrefixes;
        /// <summary>
        /// The advertised public prefix state of the Peering resource.
        /// </summary>
        public readonly string AdvertisedPublicPrefixesState;
        /// <summary>
        /// The CustomerASN of the peering.
        /// </summary>
        public readonly int? CustomerASN;
        /// <summary>
        /// The legacy mode of the peering.
        /// </summary>
        public readonly int? LegacyMode;
        /// <summary>
        /// The RoutingRegistryName of the configuration.
        /// </summary>
        public readonly string? RoutingRegistryName;

        [OutputConstructor]
        private ExpressRouteCircuitPeeringConfigResponse(
            ImmutableArray<string> advertisedCommunities,

            ImmutableArray<string> advertisedPublicPrefixes,

            string advertisedPublicPrefixesState,

            int? customerASN,

            int? legacyMode,

            string? routingRegistryName)
        {
            AdvertisedCommunities = advertisedCommunities;
            AdvertisedPublicPrefixes = advertisedPublicPrefixes;
            AdvertisedPublicPrefixesState = advertisedPublicPrefixesState;
            CustomerASN = customerASN;
            LegacyMode = legacyMode;
            RoutingRegistryName = routingRegistryName;
        }
    }
}
