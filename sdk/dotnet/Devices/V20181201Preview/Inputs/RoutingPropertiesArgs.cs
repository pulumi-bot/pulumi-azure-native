// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20181201Preview.Inputs
{

    /// <summary>
    /// The routing related properties of the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging
    /// </summary>
    public sealed class RoutingPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The properties related to the custom endpoints to which your IoT hub routes messages based on the routing rules. A maximum of 10 custom endpoints are allowed across all endpoint types for paid hubs and only 1 custom endpoint is allowed across all endpoint types for free hubs.
        /// </summary>
        [Input("endpoints")]
        public Input<Inputs.RoutingEndpointsArgs>? Endpoints { get; set; }

        /// <summary>
        /// The properties of the route that is used as a fall-back route when none of the conditions specified in the 'routes' section are met. This is an optional parameter. When this property is not set, the messages which do not meet any of the conditions specified in the 'routes' section get routed to the built-in eventhub endpoint.
        /// </summary>
        [Input("fallbackRoute")]
        public Input<Inputs.FallbackRoutePropertiesArgs>? FallbackRoute { get; set; }

        [Input("routes")]
        private InputList<Inputs.RoutePropertiesArgs>? _routes;

        /// <summary>
        /// The list of user-provided routing rules that the IoT hub uses to route messages to built-in and custom endpoints. A maximum of 100 routing rules are allowed for paid hubs and a maximum of 5 routing rules are allowed for free hubs.
        /// </summary>
        public InputList<Inputs.RoutePropertiesArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.RoutePropertiesArgs>());
            set => _routes = value;
        }

        public RoutingPropertiesArgs()
        {
        }
    }
}
