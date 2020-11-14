// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Inputs
{

    /// <summary>
    /// Routing Configuration indicating the associated and propagated route tables for this connection.
    /// </summary>
    public sealed class RoutingConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource id RouteTable associated with this RoutingConfiguration.
        /// </summary>
        [Input("associatedRouteTable")]
        public Input<Inputs.SubResourceArgs>? AssociatedRouteTable { get; set; }

        /// <summary>
        /// The list of RouteTables to advertise the routes to.
        /// </summary>
        [Input("propagatedRouteTables")]
        public Input<Inputs.PropagatedRouteTableArgs>? PropagatedRouteTables { get; set; }

        /// <summary>
        /// List of routes that control routing from VirtualHub into a virtual network connection.
        /// </summary>
        [Input("vnetRoutes")]
        public Input<Inputs.VnetRouteArgs>? VnetRoutes { get; set; }

        public RoutingConfigurationArgs()
        {
        }
    }
}
