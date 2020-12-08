// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Relay.Latest.Inputs
{

    /// <summary>
    /// SKU of the namespace.
    /// </summary>
    public sealed class SkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of this SKU.
        /// </summary>
        [Input("name", required: true)]
        public Input<Pulumi.AzureNextGen.Relay.Latest.SkuName> Name { get; set; } = null!;

        /// <summary>
        /// The tier of this SKU.
        /// </summary>
        [Input("tier")]
        public Input<Pulumi.AzureNextGen.Relay.Latest.SkuTier>? Tier { get; set; }

        public SkuArgs()
        {
        }
    }
}
