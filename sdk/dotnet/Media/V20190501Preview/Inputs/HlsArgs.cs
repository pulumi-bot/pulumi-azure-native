// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20190501Preview.Inputs
{

    /// <summary>
    /// The HLS configuration.
    /// </summary>
    public sealed class HlsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of fragments per HTTP Live Streaming (HLS) segment.
        /// </summary>
        [Input("fragmentsPerTsSegment")]
        public Input<int>? FragmentsPerTsSegment { get; set; }

        public HlsArgs()
        {
        }
    }
}
