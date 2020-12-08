// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.SecurityInsights.V20190101Preview
{
    public static class GetEntitiesGetTimeline
    {
        public static Task<GetEntitiesGetTimelineResult> InvokeAsync(GetEntitiesGetTimelineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEntitiesGetTimelineResult>("azure-nextgen:securityinsights/v20190101preview:getEntitiesGetTimeline", args ?? new GetEntitiesGetTimelineArgs(), options.WithVersion());
    }


    public sealed class GetEntitiesGetTimelineArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The end timeline date, so the results returned are before this date.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// entity ID
        /// </summary>
        [Input("entityId", required: true)]
        public string EntityId { get; set; } = null!;

        [Input("kinds")]
        private List<Union<string, Pulumi.AzureNextGen.SecurityInsights.V20190101Preview.EntityTimelineKind>>? _kinds;

        /// <summary>
        /// Array of timeline Item kinds.
        /// </summary>
        public List<Union<string, Pulumi.AzureNextGen.SecurityInsights.V20190101Preview.EntityTimelineKind>> Kinds
        {
            get => _kinds ?? (_kinds = new List<Union<string, Pulumi.AzureNextGen.SecurityInsights.V20190101Preview.EntityTimelineKind>>());
            set => _kinds = value;
        }

        /// <summary>
        /// The number of bucket for timeline queries aggregation.
        /// </summary>
        [Input("numberOfBucket")]
        public int? NumberOfBucket { get; set; }

        /// <summary>
        /// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
        /// </summary>
        [Input("operationalInsightsResourceProvider", required: true)]
        public string OperationalInsightsResourceProvider { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The start timeline date, so the results returned are after this date.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetEntitiesGetTimelineArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEntitiesGetTimelineResult
    {
        /// <summary>
        /// The metadata from the timeline operation results.
        /// </summary>
        public readonly Outputs.TimelineResultsMetadataResponseResult? MetaData;
        /// <summary>
        /// The timeline result values.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.ActivityTimelineItemResponseResult, Union<Outputs.BookmarkTimelineItemResponseResult, Outputs.SecurityAlertTimelineItemResponseResult>>> Value;

        [OutputConstructor]
        private GetEntitiesGetTimelineResult(
            Outputs.TimelineResultsMetadataResponseResult? metaData,

            ImmutableArray<Union<Outputs.ActivityTimelineItemResponseResult, Union<Outputs.BookmarkTimelineItemResponseResult, Outputs.SecurityAlertTimelineItemResponseResult>>> value)
        {
            MetaData = metaData;
            Value = value;
        }
    }
}
