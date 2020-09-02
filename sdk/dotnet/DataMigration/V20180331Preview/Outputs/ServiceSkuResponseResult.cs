// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180331Preview.Outputs
{

    [OutputType]
    public sealed class ServiceSkuResponseResult
    {
        /// <summary>
        /// The capacity of the SKU, if it supports scaling
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// The SKU family, used when the service has multiple performance classes within a tier, such as 'A', 'D', etc. for virtual machines
        /// </summary>
        public readonly string? Family;
        /// <summary>
        /// The unique name of the SKU, such as 'P3'
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The size of the SKU, used when the name alone does not denote a service size or when a SKU has multiple performance classes within a family, e.g. 'A1' for virtual machines
        /// </summary>
        public readonly string? Size;
        /// <summary>
        /// The tier of the SKU, such as 'Free', 'Basic', 'Standard', or 'Premium'
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private ServiceSkuResponseResult(
            int? capacity,

            string? family,

            string? name,

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
