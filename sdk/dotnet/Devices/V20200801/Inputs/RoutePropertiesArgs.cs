// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200801.Inputs
{

    /// <summary>
    /// The properties of a routing rule that your IoT hub uses to route messages to endpoints.
    /// </summary>
    public sealed class RoutePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The condition that is evaluated to apply the routing rule. If no condition is provided, it evaluates to true by default. For grammar, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        [Input("endpointNames", required: true)]
        private InputList<string>? _endpointNames;

        /// <summary>
        /// The list of endpoints to which messages that satisfy the condition are routed. Currently only one endpoint is allowed.
        /// </summary>
        public InputList<string> EndpointNames
        {
            get => _endpointNames ?? (_endpointNames = new InputList<string>());
            set => _endpointNames = value;
        }

        /// <summary>
        /// Used to specify whether a route is enabled.
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        /// <summary>
        /// The name of the route. The name can only include alphanumeric characters, periods, underscores, hyphens, has a maximum length of 64 characters, and must be unique.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The source that the routing rule is to be applied to, such as DeviceMessages.
        /// </summary>
        [Input("source", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Devices.V20200801.RoutingSource> Source { get; set; } = null!;

        public RoutePropertiesArgs()
        {
        }
    }
}
