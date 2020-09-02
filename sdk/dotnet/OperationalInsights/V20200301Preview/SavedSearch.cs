// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.V20200301Preview
{
    /// <summary>
    /// Value object for saved search results.
    /// </summary>
    public partial class SavedSearch : Pulumi.CustomResource
    {
        /// <summary>
        /// The category of the saved search. This helps the user to find a saved search faster. 
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Saved search display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The ETag of the saved search.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The function alias if query serves as a function.
        /// </summary>
        [Output("functionAlias")]
        public Output<string?> FunctionAlias { get; private set; } = null!;

        /// <summary>
        /// The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
        /// </summary>
        [Output("functionParameters")]
        public Output<string?> FunctionParameters { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The query expression for the saved search.
        /// </summary>
        [Output("query")]
        public Output<string> Query { get; private set; } = null!;

        /// <summary>
        /// The tags attached to the saved search.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.TagResponseResult>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The version number of the query language. The current version is 2 and is the default.
        /// </summary>
        [Output("version")]
        public Output<int?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a SavedSearch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SavedSearch(string name, SavedSearchArgs args, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/v20200301preview:SavedSearch", name, args ?? new SavedSearchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SavedSearch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/v20200301preview:SavedSearch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:operationalinsights/latest:SavedSearch"},
                    new Pulumi.Alias { Type = "azurerm:operationalinsights/v20150320:SavedSearch"},
                    new Pulumi.Alias { Type = "azurerm:operationalinsights/v20200801:SavedSearch"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SavedSearch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SavedSearch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SavedSearch(name, id, options);
        }
    }

    public sealed class SavedSearchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The category of the saved search. This helps the user to find a saved search faster. 
        /// </summary>
        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        /// <summary>
        /// Saved search display name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The ETag of the saved search.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The function alias if query serves as a function.
        /// </summary>
        [Input("functionAlias")]
        public Input<string>? FunctionAlias { get; set; }

        /// <summary>
        /// The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
        /// </summary>
        [Input("functionParameters")]
        public Input<string>? FunctionParameters { get; set; }

        /// <summary>
        /// The query expression for the saved search.
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The id of the saved search.
        /// </summary>
        [Input("savedSearchId", required: true)]
        public Input<string> SavedSearchId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.TagArgs>? _tags;

        /// <summary>
        /// The tags attached to the saved search.
        /// </summary>
        public InputList<Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The version number of the query language. The current version is 2 and is the default.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public SavedSearchArgs()
        {
        }
    }
}
