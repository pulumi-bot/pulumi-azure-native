// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20190501Preview.Outputs
{

    [OutputType]
    public sealed class HlsResponseResult
    {
        /// <summary>
        /// The amount of fragments per HTTP Live Streaming (HLS) segment.
        /// </summary>
        public readonly int? FragmentsPerTsSegment;

        [OutputConstructor]
        private HlsResponseResult(int? fragmentsPerTsSegment)
        {
            FragmentsPerTsSegment = fragmentsPerTsSegment;
        }
    }
}
