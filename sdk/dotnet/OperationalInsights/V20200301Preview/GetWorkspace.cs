// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.V20200301Preview
{
    public static class GetWorkspace
    {
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("azurerm:operationalinsights/v20200301preview:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithVersion());
    }


    public sealed class GetWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// This is a read-only property. Represents the ID associated with the workspace.
        /// </summary>
        public readonly string CustomerId;
        /// <summary>
        /// The ETag of the workspace.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of linked private link scope resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateLinkScopedResourceResponseResult> PrivateLinkScopedResources;
        /// <summary>
        /// The provisioning state of the workspace.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The network access type for accessing Log Analytics ingestion.
        /// </summary>
        public readonly string? PublicNetworkAccessForIngestion;
        /// <summary>
        /// The network access type for accessing Log Analytics query.
        /// </summary>
        public readonly string? PublicNetworkAccessForQuery;
        /// <summary>
        /// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus. 
        /// </summary>
        public readonly int? RetentionInDays;
        /// <summary>
        /// The SKU of the workspace.
        /// </summary>
        public readonly Outputs.WorkspaceSkuResponseResult? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The daily volume cap for ingestion.
        /// </summary>
        public readonly Outputs.WorkspaceCappingResponseResult? WorkspaceCapping;

        [OutputConstructor]
        private GetWorkspaceResult(
            string customerId,

            string? eTag,

            string location,

            string name,

            ImmutableArray<Outputs.PrivateLinkScopedResourceResponseResult> privateLinkScopedResources,

            string? provisioningState,

            string? publicNetworkAccessForIngestion,

            string? publicNetworkAccessForQuery,

            int? retentionInDays,

            Outputs.WorkspaceSkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.WorkspaceCappingResponseResult? workspaceCapping)
        {
            CustomerId = customerId;
            ETag = eTag;
            Location = location;
            Name = name;
            PrivateLinkScopedResources = privateLinkScopedResources;
            ProvisioningState = provisioningState;
            PublicNetworkAccessForIngestion = publicNetworkAccessForIngestion;
            PublicNetworkAccessForQuery = publicNetworkAccessForQuery;
            RetentionInDays = retentionInDays;
            Sku = sku;
            Tags = tags;
            Type = type;
            WorkspaceCapping = workspaceCapping;
        }
    }
}
