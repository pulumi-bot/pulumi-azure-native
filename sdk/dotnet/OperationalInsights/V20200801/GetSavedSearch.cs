// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.OperationalInsights.V20200801
{
    public static class GetSavedSearch
    {
        public static Task<GetSavedSearchResult> InvokeAsync(GetSavedSearchArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSavedSearchResult>("azure-nextgen:operationalinsights/v20200801:getSavedSearch", args ?? new GetSavedSearchArgs(), options.WithVersion());
    }


    public sealed class GetSavedSearchArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The id of the saved search.
        /// </summary>
        [Input("savedSearchId", required: true)]
        public string SavedSearchId { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetSavedSearchArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSavedSearchResult
    {
        /// <summary>
        /// The category of the saved search. This helps the user to find a saved search faster. 
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Saved search display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The ETag of the saved search.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The function alias if query serves as a function.
        /// </summary>
        public readonly string? FunctionAlias;
        /// <summary>
        /// The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
        /// </summary>
        public readonly string? FunctionParameters;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The query expression for the saved search.
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// The tags attached to the saved search.
        /// </summary>
        public readonly ImmutableArray<Outputs.TagResponse> Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The version number of the query language. The current version is 2 and is the default.
        /// </summary>
        public readonly int? Version;

        [OutputConstructor]
        private GetSavedSearchResult(
            string category,

            string displayName,

            string? etag,

            string? functionAlias,

            string? functionParameters,

            string name,

            string query,

            ImmutableArray<Outputs.TagResponse> tags,

            string type,

            int? version)
        {
            Category = category;
            DisplayName = displayName;
            Etag = etag;
            FunctionAlias = functionAlias;
            FunctionParameters = functionParameters;
            Name = name;
            Query = query;
            Tags = tags;
            Type = type;
            Version = version;
        }
    }
}
