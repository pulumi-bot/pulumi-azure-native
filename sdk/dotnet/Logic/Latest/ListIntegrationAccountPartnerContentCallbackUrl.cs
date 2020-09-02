// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.Latest
{
    public static class ListIntegrationAccountPartnerContentCallbackUrl
    {
        public static Task<ListIntegrationAccountPartnerContentCallbackUrlResult> InvokeAsync(ListIntegrationAccountPartnerContentCallbackUrlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListIntegrationAccountPartnerContentCallbackUrlResult>("azurerm:logic/latest:listIntegrationAccountPartnerContentCallbackUrl", args ?? new ListIntegrationAccountPartnerContentCallbackUrlArgs(), options.WithVersion());
    }


    public sealed class ListIntegrationAccountPartnerContentCallbackUrlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public string IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The key type.
        /// </summary>
        [Input("keyType")]
        public string? KeyType { get; set; }

        /// <summary>
        /// The expiry time.
        /// </summary>
        [Input("notAfter")]
        public string? NotAfter { get; set; }

        /// <summary>
        /// The integration account partner name.
        /// </summary>
        [Input("partnerName", required: true)]
        public string PartnerName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListIntegrationAccountPartnerContentCallbackUrlArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListIntegrationAccountPartnerContentCallbackUrlResult
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
        private ListIntegrationAccountPartnerContentCallbackUrlResult(
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