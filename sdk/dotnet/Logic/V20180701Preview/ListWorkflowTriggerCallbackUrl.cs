// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview
{
    public static class ListWorkflowTriggerCallbackUrl
    {
        public static Task<ListWorkflowTriggerCallbackUrlResult> InvokeAsync(ListWorkflowTriggerCallbackUrlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListWorkflowTriggerCallbackUrlResult>("azurerm:logic/v20180701preview:listWorkflowTriggerCallbackUrl", args ?? new ListWorkflowTriggerCallbackUrlArgs(), options.WithVersion());
    }


    public sealed class ListWorkflowTriggerCallbackUrlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The workflow trigger name.
        /// </summary>
        [Input("triggerName", required: true)]
        public string TriggerName { get; set; } = null!;

        /// <summary>
        /// The workflow name.
        /// </summary>
        [Input("workflowName", required: true)]
        public string WorkflowName { get; set; } = null!;

        public ListWorkflowTriggerCallbackUrlArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListWorkflowTriggerCallbackUrlResult
    {
        /// <summary>
        /// Gets the workflow trigger callback URL base path.
        /// </summary>
        public readonly string BasePath;
        /// <summary>
        /// Gets the workflow trigger callback URL HTTP method.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// Gets the workflow trigger callback URL query parameters.
        /// </summary>
        public readonly Outputs.WorkflowTriggerListCallbackUrlQueriesResponseResult? Queries;
        /// <summary>
        /// Gets the workflow trigger callback URL relative path.
        /// </summary>
        public readonly string RelativePath;
        /// <summary>
        /// Gets the workflow trigger callback URL relative path parameters.
        /// </summary>
        public readonly ImmutableArray<string> RelativePathParameters;
        /// <summary>
        /// Gets the workflow trigger callback URL.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ListWorkflowTriggerCallbackUrlResult(
            string basePath,

            string method,

            Outputs.WorkflowTriggerListCallbackUrlQueriesResponseResult? queries,

            string relativePath,

            ImmutableArray<string> relativePathParameters,

            string value)
        {
            BasePath = basePath;
            Method = method;
            Queries = queries;
            RelativePath = relativePath;
            RelativePathParameters = relativePathParameters;
            Value = value;
        }
    }
}
