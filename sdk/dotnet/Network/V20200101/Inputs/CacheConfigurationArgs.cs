// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200101.Inputs
{

    /// <summary>
    /// Caching settings for a caching-type route. To disable caching, do not provide a cacheConfiguration object.
    /// </summary>
    public sealed class CacheConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The duration for which the content needs to be cached. Allowed format is in ISO 8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations). HTTP requires the value to be no more than a year
        /// </summary>
        [Input("cacheDuration")]
        public Input<string>? CacheDuration { get; set; }

        /// <summary>
        /// Whether to use dynamic compression for cached content
        /// </summary>
        [Input("dynamicCompression")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200101.DynamicCompressionEnabled>? DynamicCompression { get; set; }

        /// <summary>
        /// Treatment of URL query terms when forming the cache key.
        /// </summary>
        [Input("queryParameterStripDirective")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200101.FrontDoorQuery>? QueryParameterStripDirective { get; set; }

        /// <summary>
        /// query parameters to include or exclude (comma separated).
        /// </summary>
        [Input("queryParameters")]
        public Input<string>? QueryParameters { get; set; }

        public CacheConfigurationArgs()
        {
        }
    }
}
