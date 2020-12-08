// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerInstance.V20191201.Inputs
{

    /// <summary>
    /// The GPU resource.
    /// </summary>
    public sealed class GpuResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The count of the GPU resource.
        /// </summary>
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        /// <summary>
        /// The SKU of the GPU resource.
        /// </summary>
        [Input("sku", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerInstance.V20191201.GpuSku> Sku { get; set; } = null!;

        public GpuResourceArgs()
        {
        }
    }
}
