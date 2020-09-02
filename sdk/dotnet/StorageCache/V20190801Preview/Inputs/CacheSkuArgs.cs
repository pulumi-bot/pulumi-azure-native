// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageCache.V20190801Preview.Inputs
{

    /// <summary>
    /// Sku for the cache.
    /// </summary>
    public sealed class CacheSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sku name for this cache.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CacheSkuArgs()
        {
        }
    }
}
