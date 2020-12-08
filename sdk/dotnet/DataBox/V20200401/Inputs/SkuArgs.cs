// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBox.V20200401.Inputs
{

    /// <summary>
    /// The Sku.
    /// </summary>
    public sealed class SkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the sku.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The sku family.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// The sku name.
        /// </summary>
        [Input("name", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.DataBox.V20200401.SkuName> Name { get; set; } = null!;

        public SkuArgs()
        {
        }
    }
}
