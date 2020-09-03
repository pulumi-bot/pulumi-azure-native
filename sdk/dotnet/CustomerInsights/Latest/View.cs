// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.Latest
{
    /// <summary>
    /// The view resource format.
    /// 
    /// ## Example Usage
    /// ### Views_CreateOrUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var view = new AzureRM.CustomerInsights.Latest.View("view", new AzureRM.CustomerInsights.Latest.ViewArgs
    ///         {
    ///             Definition = "{\\\"isProfileType\\\":false,\\\"profileTypes\\\":[],\\\"widgets\\\":[],\\\"style\\\":[]}",
    ///             DisplayName = 
    ///             {
    ///                 { "en", "some name" },
    ///             },
    ///             HubName = "sdkTestHub",
    ///             ResourceGroupName = "TestHubRG",
    ///             UserId = "testUser",
    ///             ViewName = "testView",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class View : Pulumi.CustomResource
    {
        /// <summary>
        /// Date time when view was last modified.
        /// </summary>
        [Output("changed")]
        public Output<string> Changed { get; private set; } = null!;

        /// <summary>
        /// Date time when view was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// View definition.
        /// </summary>
        [Output("definition")]
        public Output<string> Definition { get; private set; } = null!;

        /// <summary>
        /// Localized display name for the view.
        /// </summary>
        [Output("displayName")]
        public Output<ImmutableDictionary<string, string>?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// the hub name.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// the user ID.
        /// </summary>
        [Output("userId")]
        public Output<string?> UserId { get; private set; } = null!;

        /// <summary>
        /// Name of the view.
        /// </summary>
        [Output("viewName")]
        public Output<string> ViewName { get; private set; } = null!;


        /// <summary>
        /// Create a View resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public View(string name, ViewArgs args, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/latest:View", name, args ?? new ViewArgs(), MakeResourceOptions(options, ""))
        {
        }

        private View(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:customerinsights/latest:View", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:customerinsights/v20170101:View"},
                    new Pulumi.Alias { Type = "azurerm:customerinsights/v20170426:View"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing View resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static View Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new View(name, id, options);
        }
    }

    public sealed class ViewArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// View definition.
        /// </summary>
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        [Input("displayName")]
        private InputMap<string>? _displayName;

        /// <summary>
        /// Localized display name for the view.
        /// </summary>
        public InputMap<string> DisplayName
        {
            get => _displayName ?? (_displayName = new InputMap<string>());
            set => _displayName = value;
        }

        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public Input<string> HubName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// the user ID.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// The name of the view.
        /// </summary>
        [Input("viewName", required: true)]
        public Input<string> ViewName { get; set; } = null!;

        public ViewArgs()
        {
        }
    }
}
