// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20150501Preview.Outputs
{

    [OutputType]
    public sealed class SkuResponseResult
    {
        /// <summary>
        /// Capacity of the particular SKU.
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// If the service has different generations of hardware, for the same SKU, then that can be captured here.
        /// </summary>
        public readonly string? Family;
        /// <summary>
        /// The name of the SKU, typically, a letter + Number code, e.g. P3.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Size of the particular SKU
        /// </summary>
        public readonly string? Size;
        /// <summary>
        /// The tier or edition of the particular SKU, e.g. Basic, Premium.
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private SkuResponseResult(
            int? capacity,

            string? family,

            string name,

            string? size,

            string? tier)
        {
            Capacity = capacity;
            Family = family;
            Name = name;
            Size = size;
            Tier = tier;
        }
    }
}
