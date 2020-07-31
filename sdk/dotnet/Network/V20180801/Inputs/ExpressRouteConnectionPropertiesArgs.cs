// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180801.Inputs
{

    /// <summary>
    /// Properties of the ExpressRouteConnection subresource.
    /// </summary>
    public sealed class ExpressRouteConnectionPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authorization key to establish the connection.
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// The ExpressRoute circuit peering.
        /// </summary>
        [Input("expressRouteCircuitPeering", required: true)]
        public Input<Inputs.ExpressRouteCircuitPeeringIdArgs> ExpressRouteCircuitPeering { get; set; } = null!;

        /// <summary>
        /// The routing weight associated to the connection.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        public ExpressRouteConnectionPropertiesArgs()
        {
        }
    }
}