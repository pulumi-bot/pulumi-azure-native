// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.Latest
{
    /// <summary>
    /// An activity log alert resource.
    /// </summary>
    public partial class ActivityLogAlert : Pulumi.CustomResource
    {
        /// <summary>
        /// The actions that will activate when the condition is met.
        /// </summary>
        [Output("actions")]
        public Output<Outputs.ActivityLogAlertActionListResponseResult> Actions { get; private set; } = null!;

        /// <summary>
        /// The condition that will cause this alert to activate.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.ActivityLogAlertAllOfConditionResponseResult> Condition { get; private set; } = null!;

        /// <summary>
        /// A description of this activity log alert.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this activity log alert is enabled. If an activity log alert is not enabled, then none of its actions will be activated.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Azure resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of resourceIds that will be used as prefixes. The alert will only apply to activityLogs with resourceIds that fall under one of these prefixes. This list must include at least one item.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Azure resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ActivityLogAlert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActivityLogAlert(string name, ActivityLogAlertArgs args, CustomResourceOptions? options = null)
            : base("azurerm:insights/latest:ActivityLogAlert", name, args ?? new ActivityLogAlertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActivityLogAlert(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:insights/latest:ActivityLogAlert", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:insights/v20170301preview:ActivityLogAlert"},
                    new Pulumi.Alias { Type = "azurerm:insights/v20170401:ActivityLogAlert"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ActivityLogAlert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActivityLogAlert Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ActivityLogAlert(name, id, options);
        }
    }

    public sealed class ActivityLogAlertArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The actions that will activate when the condition is met.
        /// </summary>
        [Input("actions", required: true)]
        public Input<Inputs.ActivityLogAlertActionListArgs> Actions { get; set; } = null!;

        /// <summary>
        /// The name of the activity log alert.
        /// </summary>
        [Input("activityLogAlertName", required: true)]
        public Input<string> ActivityLogAlertName { get; set; } = null!;

        /// <summary>
        /// The condition that will cause this alert to activate.
        /// </summary>
        [Input("condition", required: true)]
        public Input<Inputs.ActivityLogAlertAllOfConditionArgs> Condition { get; set; } = null!;

        /// <summary>
        /// A description of this activity log alert.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether this activity log alert is enabled. If an activity log alert is not enabled, then none of its actions will be activated.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;

        /// <summary>
        /// A list of resourceIds that will be used as prefixes. The alert will only apply to activityLogs with resourceIds that fall under one of these prefixes. This list must include at least one item.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ActivityLogAlertArgs()
        {
        }
    }
}
