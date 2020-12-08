// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20180101.Inputs
{

    /// <summary>
    /// Contains SKU in an ExpressRouteCircuit.
    /// </summary>
    public sealed class ExpressRouteCircuitSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The family of the SKU. Possible values are: 'UnlimitedData' and 'MeteredData'.
        /// </summary>
        [Input("family")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20180101.ExpressRouteCircuitSkuFamily>? Family { get; set; }

        /// <summary>
        /// The name of the SKU.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The tier of the SKU. Possible values are 'Standard' and 'Premium'.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20180101.ExpressRouteCircuitSkuTier>? Tier { get; set; }

        public ExpressRouteCircuitSkuArgs()
        {
        }
    }
}
