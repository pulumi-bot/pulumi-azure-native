// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LabServices.Latest.Outputs
{

    [OutputType]
    public sealed class SizeAvailabilityResponseResult
    {
        /// <summary>
        /// Whether or not this size category is available
        /// </summary>
        public readonly bool? IsAvailable;
        /// <summary>
        /// The category of the size (Basic, Standard, Performance).
        /// </summary>
        public readonly string? SizeCategory;

        [OutputConstructor]
        private SizeAvailabilityResponseResult(
            bool? isAvailable,

            string? sizeCategory)
        {
            IsAvailable = isAvailable;
            SizeCategory = sizeCategory;
        }
    }
}
