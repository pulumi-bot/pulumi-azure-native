// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Inputs
{

    /// <summary>
    /// Contains IPv6 peering config.
    /// </summary>
    public sealed class Ipv6ExpressRouteCircuitPeeringConfigResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        [Input("microsoftPeeringConfig")]
        public Inputs.ExpressRouteCircuitPeeringConfigResponseArgs? MicrosoftPeeringConfig { get; set; }

        /// <summary>
        /// The primary address prefix.
        /// </summary>
        [Input("primaryPeerAddressPrefix")]
        public string? PrimaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The reference of the RouteFilter resource.
        /// </summary>
        [Input("routeFilter")]
        public Inputs.RouteFilterResponseArgs? RouteFilter { get; set; }

        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        [Input("secondaryPeerAddressPrefix")]
        public string? SecondaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The state of peering. Possible values are: 'Disabled' and 'Enabled'
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public Ipv6ExpressRouteCircuitPeeringConfigResponseArgs()
        {
        }
    }
}
