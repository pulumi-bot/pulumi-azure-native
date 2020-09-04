// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180801.Inputs
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
        public Input<string>? DynamicCompression { get; set; }

        /// <summary>
        /// Treatment of URL query terms when forming the cache key.
        /// </summary>
        [Input("queryParameterStripDirective")]
        public Input<string>? QueryParameterStripDirective { get; set; }

        public CacheConfigurationArgs()
        {
        }
    }
}
