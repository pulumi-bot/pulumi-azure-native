// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20200801Preview.Outputs
{

    [OutputType]
    public sealed class LastAccessTimeTrackingPolicyResponse
    {
        /// <summary>
        /// An array of predefined supported blob types. Only blockBlob is the supported value. This field is currently read only
        /// </summary>
        public readonly ImmutableArray<string> BlobType;
        /// <summary>
        /// When set to true last access time based tracking is enabled.
        /// </summary>
        public readonly bool Enable;
        /// <summary>
        /// Name of the policy. The valid value is AccessTimeTracking. This field is currently read only
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The field specifies blob object tracking granularity in days, typically how often the blob object should be tracked.This field is currently read only with value as 1
        /// </summary>
        public readonly int? TrackingGranularityInDays;

        [OutputConstructor]
        private LastAccessTimeTrackingPolicyResponse(
            ImmutableArray<string> blobType,

            bool enable,

            string? name,

            int? trackingGranularityInDays)
        {
            BlobType = blobType;
            Enable = enable;
            Name = name;
            TrackingGranularityInDays = trackingGranularityInDays;
        }
    }
}
