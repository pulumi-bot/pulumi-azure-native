// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview.Outputs
{

    [OutputType]
    public sealed class LiveEventEncodingResponseResult
    {
        /// <summary>
        /// The encoding type for Live Event.
        /// </summary>
        public readonly string? EncodingType;
        /// <summary>
        /// The encoding preset name.
        /// </summary>
        public readonly string? PresetName;

        [OutputConstructor]
        private LiveEventEncodingResponseResult(
            string? encodingType,

            string? presetName)
        {
            EncodingType = encodingType;
            PresetName = presetName;
        }
    }
}
