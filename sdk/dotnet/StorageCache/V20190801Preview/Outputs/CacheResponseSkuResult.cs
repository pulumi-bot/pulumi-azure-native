// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageCache.V20190801Preview.Outputs
{

    [OutputType]
    public sealed class CacheResponseSkuResult
    {
        /// <summary>
        /// Sku name for this cache.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private CacheResponseSkuResult(string? name)
        {
            Name = name;
        }
    }
}
