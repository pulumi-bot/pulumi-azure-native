// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.OperationalInsights.V20201001
{
    public static class GetWorkspace
    {
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("azure-nextgen:operationalinsights/v20201001:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithVersion());
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
        /// Indicates whether customer managed storage is mandatory for query management.
        /// </summary>
        public readonly bool? ForceCmkForQuery;
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
        public readonly ImmutableArray<Outputs.PrivateLinkScopedResourceResponse> PrivateLinkScopedResources;
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
        /// The workspace data retention in days, between 30 and 730.
        /// </summary>
        public readonly int? RetentionInDays;
        /// <summary>
        /// The SKU of the workspace.
        /// </summary>
        public readonly Outputs.WorkspaceSkuResponse? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The daily volume cap for ingestion.
        /// </summary>
        public readonly Outputs.WorkspaceCappingResponse? WorkspaceCapping;

        [OutputConstructor]
        private GetWorkspaceResult(
            string customerId,

            string? eTag,

            bool? forceCmkForQuery,

            string location,

            string name,

            ImmutableArray<Outputs.PrivateLinkScopedResourceResponse> privateLinkScopedResources,

            string? provisioningState,

            string? publicNetworkAccessForIngestion,

            string? publicNetworkAccessForQuery,

            int? retentionInDays,

            Outputs.WorkspaceSkuResponse? sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.WorkspaceCappingResponse? workspaceCapping)
        {
            CustomerId = customerId;
            ETag = eTag;
            ForceCmkForQuery = forceCmkForQuery;
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
