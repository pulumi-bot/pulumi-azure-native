// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20200930.Outputs
{

    [OutputType]
    public sealed class DiskSkuResponse
    {
        /// <summary>
        /// The sku name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The sku tier.
        /// </summary>
        public readonly string Tier;

        [OutputConstructor]
        private DiskSkuResponse(
            string? name,

            string tier)
        {
            Name = name;
            Tier = tier;
        }
    }
}
