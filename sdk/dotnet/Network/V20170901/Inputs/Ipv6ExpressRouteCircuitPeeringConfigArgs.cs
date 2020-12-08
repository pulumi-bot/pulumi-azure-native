// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20170901.Inputs
{

    /// <summary>
    /// Contains IPv6 peering config.
    /// </summary>
    public sealed class Ipv6ExpressRouteCircuitPeeringConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        [Input("microsoftPeeringConfig")]
        public Input<Inputs.ExpressRouteCircuitPeeringConfigArgs>? MicrosoftPeeringConfig { get; set; }

        /// <summary>
        /// The primary address prefix.
        /// </summary>
        [Input("primaryPeerAddressPrefix")]
        public Input<string>? PrimaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The reference of the RouteFilter resource.
        /// </summary>
        [Input("routeFilter")]
        public Input<Inputs.RouteFilterArgs>? RouteFilter { get; set; }

        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        [Input("secondaryPeerAddressPrefix")]
        public Input<string>? SecondaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The state of peering. Possible values are: 'Disabled' and 'Enabled'
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20170901.ExpressRouteCircuitPeeringState>? State { get; set; }

        public Ipv6ExpressRouteCircuitPeeringConfigArgs()
        {
        }
    }
}
