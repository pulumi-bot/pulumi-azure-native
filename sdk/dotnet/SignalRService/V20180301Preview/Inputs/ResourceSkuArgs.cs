// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SignalRService.V20180301Preview.Inputs
{

    /// <summary>
    /// The billing information of the resource.(e.g. basic vs. standard)
    /// </summary>
    public sealed class ResourceSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional, integer. If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not 
        /// possible for the resource this may be omitted.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Optional, string. If the service has different generations of hardware, for the same SKU, then that can be captured here.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// The name of the SKU. This is typically a letter + number code, such as A0 or P3.  Required (if sku is specified)
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Optional, string. When the name field is the combination of tier and some other value, this would be the standalone code.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// Optional tier of this particular SKU. `Basic` is deprecated, use `Standard` instead for Basic tier
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public ResourceSkuArgs()
        {
        }
    }
}
