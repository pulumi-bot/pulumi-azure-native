// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventHub.V20170401
{
    /// <summary>
    /// Single item in a List or Get AuthorizationRule operation
    /// 
    /// ## Example Usage
    /// ### EventHubAuthorizationRuleCreate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var eventHubAuthorizationRule = new AzureRM.EventHub.V20170401.EventHubAuthorizationRule("eventHubAuthorizationRule", new AzureRM.EventHub.V20170401.EventHubAuthorizationRuleArgs
    ///         {
    ///             AuthorizationRuleName = "sdk-Authrules-2513",
    ///             EventHubName = "sdk-EventHub-532",
    ///             NamespaceName = "sdk-Namespace-960",
    ///             ResourceGroupName = "ArunMonocle",
    ///             Rights = 
    ///             {
    ///                 "Listen",
    ///                 "Send",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class EventHubAuthorizationRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        [Output("rights")]
        public Output<ImmutableArray<string>> Rights { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EventHubAuthorizationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventHubAuthorizationRule(string name, EventHubAuthorizationRuleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:eventhub/v20170401:EventHubAuthorizationRule", name, args ?? new EventHubAuthorizationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventHubAuthorizationRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:eventhub/v20170401:EventHubAuthorizationRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:eventhub/latest:EventHubAuthorizationRule"},
                    new Pulumi.Alias { Type = "azurerm:eventhub/v20140901:EventHubAuthorizationRule"},
                    new Pulumi.Alias { Type = "azurerm:eventhub/v20150801:EventHubAuthorizationRule"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventHubAuthorizationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventHubAuthorizationRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventHubAuthorizationRule(name, id, options);
        }
    }

    public sealed class EventHubAuthorizationRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization rule name.
        /// </summary>
        [Input("authorizationRuleName", required: true)]
        public Input<string> AuthorizationRuleName { get; set; } = null!;

        /// <summary>
        /// The Event Hub name
        /// </summary>
        [Input("eventHubName", required: true)]
        public Input<string> EventHubName { get; set; } = null!;

        /// <summary>
        /// The Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("rights", required: true)]
        private InputList<string>? _rights;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        public InputList<string> Rights
        {
            get => _rights ?? (_rights = new InputList<string>());
            set => _rights = value;
        }

        public EventHubAuthorizationRuleArgs()
        {
        }
    }
}
