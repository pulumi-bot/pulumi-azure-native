// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.V20180701Preview
{
    public static class ListWorkflowCallbackUrl
    {
        public static Task<ListWorkflowCallbackUrlResult> InvokeAsync(ListWorkflowCallbackUrlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListWorkflowCallbackUrlResult>("azure-nextgen:logic/v20180701preview:listWorkflowCallbackUrl", args ?? new ListWorkflowCallbackUrlArgs(), options.WithVersion());
    }


    public sealed class ListWorkflowCallbackUrlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The key type.
        /// </summary>
        [Input("keyType")]
        public Union<string, Pulumi.AzureNextGen.Logic.V20180701Preview.KeyType>? KeyType { get; set; }

        /// <summary>
        /// The expiry time.
        /// </summary>
        [Input("notAfter")]
        public string? NotAfter { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The workflow name.
        /// </summary>
        [Input("workflowName", required: true)]
        public string WorkflowName { get; set; } = null!;

        public ListWorkflowCallbackUrlArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListWorkflowCallbackUrlResult
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
        private ListWorkflowCallbackUrlResult(
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
