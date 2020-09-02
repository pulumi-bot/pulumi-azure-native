// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20170907privatePreview.Inputs
{

    public sealed class AzureSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// SKU capacity.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// SKU name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// SKU tier.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public AzureSkuArgs()
        {
        }
    }
}
