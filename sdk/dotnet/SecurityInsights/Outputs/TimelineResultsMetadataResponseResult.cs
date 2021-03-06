// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Outputs
{

    [OutputType]
    public sealed class TimelineResultsMetadataResponseResult
    {
        /// <summary>
        /// timeline aggregation per kind
        /// </summary>
        public readonly ImmutableArray<Outputs.TimelineAggregationResponseResult> Aggregations;
        /// <summary>
        /// information about the failure queries
        /// </summary>
        public readonly ImmutableArray<Outputs.TimelineErrorResponseResult> Errors;
        /// <summary>
        /// the total items found for the timeline request
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private TimelineResultsMetadataResponseResult(
            ImmutableArray<Outputs.TimelineAggregationResponseResult> aggregations,

            ImmutableArray<Outputs.TimelineErrorResponseResult> errors,

            int totalCount)
        {
            Aggregations = aggregations;
            Errors = errors;
            TotalCount = totalCount;
        }
    }
}
