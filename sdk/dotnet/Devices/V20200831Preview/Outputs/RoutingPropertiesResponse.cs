// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200831Preview.Outputs
{

    [OutputType]
    public sealed class RoutingPropertiesResponse
    {
        /// <summary>
        /// The properties related to the custom endpoints to which your IoT hub routes messages based on the routing rules. A maximum of 10 custom endpoints are allowed across all endpoint types for paid hubs and only 1 custom endpoint is allowed across all endpoint types for free hubs.
        /// </summary>
        public readonly Outputs.RoutingEndpointsResponse? Endpoints;
        /// <summary>
        /// The list of user-provided enrichments that the IoT hub applies to messages to be delivered to built-in and custom endpoints. See: https://aka.ms/telemetryoneventgrid
        /// </summary>
        public readonly ImmutableArray<Outputs.EnrichmentPropertiesResponse> Enrichments;
        /// <summary>
        /// The properties of the route that is used as a fall-back route when none of the conditions specified in the 'routes' section are met. This is an optional parameter. When this property is not set, the messages which do not meet any of the conditions specified in the 'routes' section get routed to the built-in eventhub endpoint.
        /// </summary>
        public readonly Outputs.FallbackRoutePropertiesResponse? FallbackRoute;
        /// <summary>
        /// The list of user-provided routing rules that the IoT hub uses to route messages to built-in and custom endpoints. A maximum of 100 routing rules are allowed for paid hubs and a maximum of 5 routing rules are allowed for free hubs.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoutePropertiesResponse> Routes;

        [OutputConstructor]
        private RoutingPropertiesResponse(
            Outputs.RoutingEndpointsResponse? endpoints,

            ImmutableArray<Outputs.EnrichmentPropertiesResponse> enrichments,

            Outputs.FallbackRoutePropertiesResponse? fallbackRoute,

            ImmutableArray<Outputs.RoutePropertiesResponse> routes)
        {
            Endpoints = endpoints;
            Enrichments = enrichments;
            FallbackRoute = fallbackRoute;
            Routes = routes;
        }
    }
}
