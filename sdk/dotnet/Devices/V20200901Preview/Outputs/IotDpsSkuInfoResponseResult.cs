// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200901Preview.Outputs
{

    [OutputType]
    public sealed class IotDpsSkuInfoResponseResult
    {
        /// <summary>
        /// The number of units to provision
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// Sku name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Pricing tier name of the provisioning service.
        /// </summary>
        public readonly string Tier;

        [OutputConstructor]
        private IotDpsSkuInfoResponseResult(
            int? capacity,

            string? name,

            string tier)
        {
            Capacity = capacity;
            Name = name;
            Tier = tier;
        }
    }
}
