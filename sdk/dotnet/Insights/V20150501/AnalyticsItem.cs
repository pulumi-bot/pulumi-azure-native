// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20150501
{
    /// <summary>
    /// Properties that define an Analytics item that is associated to an Application Insights component.
    /// 
    /// ## Example Usage
    /// ### AnalyticsItemPut
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var analyticsItem = new AzureRM.Insights.V20150501.AnalyticsItem("analyticsItem", new AzureRM.Insights.V20150501.AnalyticsItemArgs
    ///         {
    ///             Content = @"let newExceptionsTimeRange = 1d;
    /// let timeRangeToCheckBefore = 7d;
    /// exceptions
    /// | where timestamp &lt; ago(timeRangeToCheckBefore)
    /// | summarize count() by problemId
    /// | join kind= rightanti (
    /// exceptions
    /// | where timestamp &gt;= ago(newExceptionsTimeRange)
    /// | extend stack = tostring(details[0].rawStack)
    /// | summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  
    /// ) on problemId 
    /// | order by  count_ desc
    /// ",
    ///             Name = "Exceptions - New in the last 24 hours",
    ///             ResourceGroupName = "my-resource-group",
    ///             ResourceName = "my-component",
    ///             Scope = "shared",
    ///             ScopePath = "analyticsItems",
    ///             Type = "query",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class AnalyticsItem : Pulumi.CustomResource
    {
        /// <summary>
        /// The content of this item
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// The user-defined name of the item.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ApplicationInsightsComponentAnalyticsItemPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// Date and time in UTC when this item was created.
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// Date and time in UTC of the last modification that was made to this item.
        /// </summary>
        [Output("timeModified")]
        public Output<string> TimeModified { get; private set; } = null!;

        /// <summary>
        /// Enum indicating the type of the Analytics item.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// This instance's version of the data model. This can change as new features are added.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AnalyticsItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AnalyticsItem(string name, AnalyticsItemArgs args, CustomResourceOptions? options = null)
            : base("azurerm:insights/v20150501:AnalyticsItem", name, args ?? new AnalyticsItemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AnalyticsItem(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:insights/v20150501:AnalyticsItem", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:insights/latest:AnalyticsItem"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AnalyticsItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AnalyticsItem Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AnalyticsItem(name, id, options);
        }
    }

    public sealed class AnalyticsItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content of this item
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Internally assigned unique id of the item definition.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The user-defined name of the item.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Flag indicating whether or not to force save an item. This allows overriding an item if it already exists.
        /// </summary>
        [Input("overrideItem")]
        public Input<bool>? OverrideItem { get; set; }

        /// <summary>
        /// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ApplicationInsightsComponentAnalyticsItemPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Application Insights component resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application Insights component.
        /// </summary>
        [Input("scopePath", required: true)]
        public Input<string> ScopePath { get; set; } = null!;

        /// <summary>
        /// Enum indicating the type of the Analytics item.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AnalyticsItemArgs()
        {
        }
    }
}
