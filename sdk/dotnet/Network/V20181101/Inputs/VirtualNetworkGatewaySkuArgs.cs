// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20181101.Inputs
{

    /// <summary>
    /// VirtualNetworkGatewaySku details
    /// </summary>
    public sealed class VirtualNetworkGatewaySkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The capacity.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Gateway SKU name.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20181101.VirtualNetworkGatewaySkuName>? Name { get; set; }

        /// <summary>
        /// Gateway SKU tier.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20181101.VirtualNetworkGatewaySkuTier>? Tier { get; set; }

        public VirtualNetworkGatewaySkuArgs()
        {
        }
    }
}
