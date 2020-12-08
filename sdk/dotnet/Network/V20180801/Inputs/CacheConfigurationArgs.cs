// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20180801.Inputs
{

    /// <summary>
    /// Caching settings for a caching-type route. To disable caching, do not provide a cacheConfiguration object.
    /// </summary>
    public sealed class CacheConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to use dynamic compression for cached content
        /// </summary>
        [Input("dynamicCompression")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20180801.DynamicCompressionEnabled>? DynamicCompression { get; set; }

        /// <summary>
        /// Treatment of URL query terms when forming the cache key.
        /// </summary>
        [Input("queryParameterStripDirective")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20180801.FrontDoorQuery>? QueryParameterStripDirective { get; set; }

        public CacheConfigurationArgs()
        {
        }
    }
}
