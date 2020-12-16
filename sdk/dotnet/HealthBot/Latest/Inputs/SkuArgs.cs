// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HealthBot.Latest.Inputs
{

    /// <summary>
    /// The resource model definition representing SKU
    /// </summary>
    public sealed class SkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// If the service has different generations of hardware, for the same SKU, then that can be captured here.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// The name of the HealthBot SKU
        /// </summary>
        [Input("name", required: true)]
        public Input<Pulumi.AzureNextGen.HealthBot.Latest.SkuName> Name { get; set; } = null!;

        /// <summary>
        /// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code. 
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
        /// </summary>
        [Input("tier")]
        public Input<Pulumi.AzureNextGen.HealthBot.Latest.SkuTier>? Tier { get; set; }

        public SkuArgs()
        {
        }
    }
}
