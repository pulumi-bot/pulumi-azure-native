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
    /// SKU of nat gateway.
    /// </summary>
    public sealed class NatGatewaySkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of Nat Gateway SKU.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public NatGatewaySkuArgs()
        {
        }
    }
}
