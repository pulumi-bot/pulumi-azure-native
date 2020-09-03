// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.Latest
{
    /// <summary>
    /// Represents a bookmark in Azure Security Insights.
    /// 
    /// ## Example Usage
    /// ### Creates or updates a bookmark.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bookmark = new AzureRM.OperationalInsights.Latest.Bookmark("bookmark", new AzureRM.OperationalInsights.Latest.BookmarkArgs
    ///         {
    ///             BookmarkId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
    ///             Created = "2019-01-01T13:15:30Z",
    ///             CreatedBy = new AzureRM.OperationalInsights.Latest.Inputs.UserInfoArgs
    ///             {
    ///                 ObjectId = "2046feea-040d-4a46-9e2b-91c2941bfa70",
    ///             },
    ///             DisplayName = "My bookmark",
    ///             Etag = "\"0300bf09-0000-0000-0000-5c37296e0000\"",
    ///             Labels = 
    ///             {
    ///                 "Tag1",
    ///                 "Tag2",
    ///             },
    ///             Notes = "Found a suspicious activity",
    ///             Query = "SecurityEvent | where TimeGenerated &gt; ago(1d) and TimeGenerated &lt; ago(2d)",
    ///             QueryResult = "Security Event query result",
    ///             ResourceGroupName = "myRg",
    ///             Updated = "2019-01-01T13:15:30Z",
    ///             UpdatedBy = new AzureRM.OperationalInsights.Latest.Inputs.UserInfoArgs
    ///             {
    ///                 ObjectId = "2046feea-040d-4a46-9e2b-91c2941bfa70",
    ///             },
    ///             WorkspaceName = "myWorkspace",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Bookmark : Pulumi.CustomResource
    {
        /// <summary>
        /// The time the bookmark was created
        /// </summary>
        [Output("created")]
        public Output<string?> Created { get; private set; } = null!;

        /// <summary>
        /// Describes a user that created the bookmark
        /// </summary>
        [Output("createdBy")]
        public Output<Outputs.UserInfoResponseResult?> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// The display name of the bookmark
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Describes an incident that relates to bookmark
        /// </summary>
        [Output("incidentInfo")]
        public Output<Outputs.IncidentInfoResponseResult?> IncidentInfo { get; private set; } = null!;

        /// <summary>
        /// List of labels relevant to this bookmark
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Azure resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The notes of the bookmark
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// The query of the bookmark.
        /// </summary>
        [Output("query")]
        public Output<string> Query { get; private set; } = null!;

        /// <summary>
        /// The query result of the bookmark.
        /// </summary>
        [Output("queryResult")]
        public Output<string?> QueryResult { get; private set; } = null!;

        /// <summary>
        /// Azure resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The last time the bookmark was updated
        /// </summary>
        [Output("updated")]
        public Output<string?> Updated { get; private set; } = null!;

        /// <summary>
        /// Describes a user that updated the bookmark
        /// </summary>
        [Output("updatedBy")]
        public Output<Outputs.UserInfoResponseResult?> UpdatedBy { get; private set; } = null!;


        /// <summary>
        /// Create a Bookmark resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bookmark(string name, BookmarkArgs args, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/latest:Bookmark", name, args ?? new BookmarkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Bookmark(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/latest:Bookmark", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:operationalinsights/v20200101:Bookmark"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Bookmark resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bookmark Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Bookmark(name, id, options);
        }
    }

    public sealed class BookmarkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bookmark ID
        /// </summary>
        [Input("bookmarkId", required: true)]
        public Input<string> BookmarkId { get; set; } = null!;

        /// <summary>
        /// The time the bookmark was created
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// Describes a user that created the bookmark
        /// </summary>
        [Input("createdBy")]
        public Input<Inputs.UserInfoArgs>? CreatedBy { get; set; }

        /// <summary>
        /// The display name of the bookmark
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Describes an incident that relates to bookmark
        /// </summary>
        [Input("incidentInfo")]
        public Input<Inputs.IncidentInfoArgs>? IncidentInfo { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// List of labels relevant to this bookmark
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The notes of the bookmark
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// The query of the bookmark.
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        /// <summary>
        /// The query result of the bookmark.
        /// </summary>
        [Input("queryResult")]
        public Input<string>? QueryResult { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The last time the bookmark was updated
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        /// <summary>
        /// Describes a user that updated the bookmark
        /// </summary>
        [Input("updatedBy")]
        public Input<Inputs.UserInfoArgs>? UpdatedBy { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public BookmarkArgs()
        {
        }
    }
}
