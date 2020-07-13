// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Outputs
{

    [OutputType]
    public sealed class RouteFilterPropertiesFormatResponse
    {
        /// <summary>
        /// A collection of references to express route circuit ipv6 peerings.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExpressRouteCircuitPeeringResponse> Ipv6Peerings;
        /// <summary>
        /// A collection of references to express route circuit peerings.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExpressRouteCircuitPeeringResponse> Peerings;
        /// <summary>
        /// The provisioning state of the route filter resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Collection of RouteFilterRules contained within a route filter.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouteFilterRuleResponse> Rules;

        [OutputConstructor]
        private RouteFilterPropertiesFormatResponse(
            ImmutableArray<Outputs.ExpressRouteCircuitPeeringResponse> ipv6Peerings,

            ImmutableArray<Outputs.ExpressRouteCircuitPeeringResponse> peerings,

            string provisioningState,

            ImmutableArray<Outputs.RouteFilterRuleResponse> rules)
        {
            Ipv6Peerings = ipv6Peerings;
            Peerings = peerings;
            ProvisioningState = provisioningState;
            Rules = rules;
        }
    }
}