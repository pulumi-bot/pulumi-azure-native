// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageCache.Latest.Inputs
{

    /// <summary>
    /// Cache identity properties.
    /// </summary>
    public sealed class CacheIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of identity used for the cache
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNextGen.StorageCache.Latest.CacheIdentityType>? Type { get; set; }

        public CacheIdentityArgs()
        {
        }
    }
}
