// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20180907Preview.Outputs
{

    [OutputType]
    public sealed class AzureSkuResponseResult
    {
        /// <summary>
        /// SKU capacity.
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// SKU name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// SKU tier.
        /// </summary>
        public readonly string Tier;

        [OutputConstructor]
        private AzureSkuResponseResult(
            int? capacity,

            string name,

            string tier)
        {
            Capacity = capacity;
            Name = name;
            Tier = tier;
        }
    }
}
