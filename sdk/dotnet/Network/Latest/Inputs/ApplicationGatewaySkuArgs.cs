// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest.Inputs
{

    /// <summary>
    /// SKU of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewaySkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Capacity (instance count) of an application gateway.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Name of an application gateway SKU.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.Latest.ApplicationGatewaySkuName>? Name { get; set; }

        /// <summary>
        /// Tier of an application gateway.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.Latest.ApplicationGatewayTier>? Tier { get; set; }

        public ApplicationGatewaySkuArgs()
        {
        }
    }
}
