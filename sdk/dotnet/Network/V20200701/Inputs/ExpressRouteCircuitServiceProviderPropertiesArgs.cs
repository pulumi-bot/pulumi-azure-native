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
    /// Contains ServiceProviderProperties in an ExpressRouteCircuit.
    /// </summary>
    public sealed class ExpressRouteCircuitServiceProviderPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The BandwidthInMbps.
        /// </summary>
        [Input("bandwidthInMbps")]
        public Input<int>? BandwidthInMbps { get; set; }

        /// <summary>
        /// The peering location.
        /// </summary>
        [Input("peeringLocation")]
        public Input<string>? PeeringLocation { get; set; }

        /// <summary>
        /// The serviceProviderName.
        /// </summary>
        [Input("serviceProviderName")]
        public Input<string>? ServiceProviderName { get; set; }

        public ExpressRouteCircuitServiceProviderPropertiesArgs()
        {
        }
    }
}
