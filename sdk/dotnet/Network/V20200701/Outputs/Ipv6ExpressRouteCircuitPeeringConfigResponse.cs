// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Outputs
{

    [OutputType]
    public sealed class Ipv6ExpressRouteCircuitPeeringConfigResponse
    {
        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        public readonly Outputs.ExpressRouteCircuitPeeringConfigResponse? MicrosoftPeeringConfig;
        /// <summary>
        /// The primary address prefix.
        /// </summary>
        public readonly string? PrimaryPeerAddressPrefix;
        /// <summary>
        /// The reference to the RouteFilter resource.
        /// </summary>
        public readonly Outputs.SubResourceResponse? RouteFilter;
        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        public readonly string? SecondaryPeerAddressPrefix;
        /// <summary>
        /// The state of peering.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private Ipv6ExpressRouteCircuitPeeringConfigResponse(
            Outputs.ExpressRouteCircuitPeeringConfigResponse? microsoftPeeringConfig,

            string? primaryPeerAddressPrefix,

            Outputs.SubResourceResponse? routeFilter,

            string? secondaryPeerAddressPrefix,

            string? state)
        {
            MicrosoftPeeringConfig = microsoftPeeringConfig;
            PrimaryPeerAddressPrefix = primaryPeerAddressPrefix;
            RouteFilter = routeFilter;
            SecondaryPeerAddressPrefix = secondaryPeerAddressPrefix;
            State = state;
        }
    }
}
