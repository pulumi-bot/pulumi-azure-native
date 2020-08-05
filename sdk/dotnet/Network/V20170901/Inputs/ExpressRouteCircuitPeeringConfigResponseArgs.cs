// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170901.Inputs
{

    /// <summary>
    /// Specifies the peering configuration.
    /// </summary>
    public sealed class ExpressRouteCircuitPeeringConfigResponseArgs : Pulumi.InvokeArgs
    {
        [Input("advertisedCommunities")]
        private List<string>? _advertisedCommunities;

        /// <summary>
        /// The communities of bgp peering. Specified for microsoft peering
        /// </summary>
        public List<string> AdvertisedCommunities
        {
            get => _advertisedCommunities ?? (_advertisedCommunities = new List<string>());
            set => _advertisedCommunities = value;
        }

        [Input("advertisedPublicPrefixes")]
        private List<string>? _advertisedPublicPrefixes;

        /// <summary>
        /// The reference of AdvertisedPublicPrefixes.
        /// </summary>
        public List<string> AdvertisedPublicPrefixes
        {
            get => _advertisedPublicPrefixes ?? (_advertisedPublicPrefixes = new List<string>());
            set => _advertisedPublicPrefixes = value;
        }

        /// <summary>
        /// AdvertisedPublicPrefixState of the Peering resource. Possible values are 'NotConfigured', 'Configuring', 'Configured', and 'ValidationNeeded'.
        /// </summary>
        [Input("advertisedPublicPrefixesState")]
        public string? AdvertisedPublicPrefixesState { get; set; }

        /// <summary>
        /// The CustomerASN of the peering.
        /// </summary>
        [Input("customerASN")]
        public int? CustomerASN { get; set; }

        /// <summary>
        /// The legacy mode of the peering.
        /// </summary>
        [Input("legacyMode")]
        public int? LegacyMode { get; set; }

        /// <summary>
        /// The RoutingRegistryName of the configuration.
        /// </summary>
        [Input("routingRegistryName")]
        public string? RoutingRegistryName { get; set; }

        public ExpressRouteCircuitPeeringConfigResponseArgs()
        {
        }
    }
}
