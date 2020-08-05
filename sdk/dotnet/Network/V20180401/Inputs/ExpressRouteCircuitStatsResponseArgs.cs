// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180401.Inputs
{

    /// <summary>
    /// Contains stats associated with the peering.
    /// </summary>
    public sealed class ExpressRouteCircuitStatsResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Gets BytesIn of the peering.
        /// </summary>
        [Input("primarybytesIn")]
        public int? PrimarybytesIn { get; set; }

        /// <summary>
        /// Gets BytesOut of the peering.
        /// </summary>
        [Input("primarybytesOut")]
        public int? PrimarybytesOut { get; set; }

        /// <summary>
        /// Gets BytesIn of the peering.
        /// </summary>
        [Input("secondarybytesIn")]
        public int? SecondarybytesIn { get; set; }

        /// <summary>
        /// Gets BytesOut of the peering.
        /// </summary>
        [Input("secondarybytesOut")]
        public int? SecondarybytesOut { get; set; }

        public ExpressRouteCircuitStatsResponseArgs()
        {
        }
    }
}
