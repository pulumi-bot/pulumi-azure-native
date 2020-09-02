// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Synapse.V20190601Preview.Inputs
{

    /// <summary>
    /// SQL pool SKU
    /// </summary>
    public sealed class SkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SKU name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The service tier
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public SkuArgs()
        {
        }
    }
}
