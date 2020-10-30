// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cache.V20201001Preview.Outputs
{

    [OutputType]
    public sealed class SkuResponse
    {
        /// <summary>
        /// The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, ...) for Enterprise SKUs and (3, 9, 15, ...) for Flash SKUs.
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// The type of RedisEnterprise cluster to deploy. Possible values: (Enterprise_E10, EnterpriseFlash_F300 etc.)
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private SkuResponse(
            int? capacity,

            string name)
        {
            Capacity = capacity;
            Name = name;
        }
    }
}
