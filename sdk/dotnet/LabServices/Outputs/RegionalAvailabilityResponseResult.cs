// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LabServices.Outputs
{

    [OutputType]
    public sealed class RegionalAvailabilityResponseResult
    {
        /// <summary>
        /// Corresponding region
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// List of all the size information for the region
        /// </summary>
        public readonly ImmutableArray<Outputs.SizeAvailabilityResponseResult> SizeAvailabilities;

        [OutputConstructor]
        private RegionalAvailabilityResponseResult(
            string? region,

            ImmutableArray<Outputs.SizeAvailabilityResponseResult> sizeAvailabilities)
        {
            Region = region;
            SizeAvailabilities = sizeAvailabilities;
        }
    }
}
