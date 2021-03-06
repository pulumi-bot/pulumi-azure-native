// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Latest
{
    [Obsolete(@"The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:devtestlab:listScheduleApplicable'.")]
    public static class ListScheduleApplicable
    {
        /// <summary>
        /// The response of a list operation.
        /// Latest API Version: 2018-09-15.
        /// </summary>
        public static Task<ListScheduleApplicableResult> InvokeAsync(ListScheduleApplicableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListScheduleApplicableResult>("azure-native:devtestlab/latest:listScheduleApplicable", args ?? new ListScheduleApplicableArgs(), options.WithVersion());
    }


    public sealed class ListScheduleApplicableArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public string LabName { get; set; } = null!;

        /// <summary>
        /// The name of the schedule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListScheduleApplicableArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListScheduleApplicableResult
    {
        /// <summary>
        /// Link for next set of results.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// Results of the list operation.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScheduleResponse> Value;

        [OutputConstructor]
        private ListScheduleApplicableResult(
            string? nextLink,

            ImmutableArray<Outputs.ScheduleResponse> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
