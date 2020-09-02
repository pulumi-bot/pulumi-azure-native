// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AnalysisServices.V20170801beta.Outputs
{

    [OutputType]
    public sealed class ResourceSkuResponseResult
    {
        /// <summary>
        /// The number of instances in the read only query pool.
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// Name of the SKU level.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the Azure pricing tier to which the SKU applies.
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private ResourceSkuResponseResult(
            int? capacity,

            string name,

            string? tier)
        {
            Capacity = capacity;
            Name = name;
            Tier = tier;
        }
    }
}
