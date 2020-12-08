// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cache.Latest.Inputs
{

    /// <summary>
    /// SKU parameters supplied to the create Redis operation.
    /// </summary>
    public sealed class SkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4).
        /// </summary>
        [Input("capacity", required: true)]
        public Input<int> Capacity { get; set; } = null!;

        /// <summary>
        /// The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
        /// </summary>
        [Input("family", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Cache.Latest.SkuFamily> Family { get; set; } = null!;

        /// <summary>
        /// The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
        /// </summary>
        [Input("name", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Cache.Latest.SkuName> Name { get; set; } = null!;

        public SkuArgs()
        {
        }
    }
}
