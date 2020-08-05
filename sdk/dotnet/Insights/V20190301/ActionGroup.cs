// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20190301
{
    /// <summary>
    /// An action group resource.
    /// </summary>
    public partial class ActionGroup : Pulumi.CustomResource
    {
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
        /// The action groups properties of the resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ActionGroupResponseResult> Properties { get; private set; } = null!;

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
        /// Create a ActionGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionGroup(string name, ActionGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:insights/v20190301:ActionGroup", name, args ?? new ActionGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:insights/v20190301:ActionGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ActionGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ActionGroup(name, id, options);
        }
    }

    public sealed class ActionGroupArgs : Pulumi.ResourceArgs
    {
        [Input("armRoleReceivers")]
        private InputList<Inputs.ArmRoleReceiverArgs>? _armRoleReceivers;

        /// <summary>
        /// The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
        /// </summary>
        public InputList<Inputs.ArmRoleReceiverArgs> ArmRoleReceivers
        {
            get => _armRoleReceivers ?? (_armRoleReceivers = new InputList<Inputs.ArmRoleReceiverArgs>());
            set => _armRoleReceivers = value;
        }

        [Input("automationRunbookReceivers")]
        private InputList<Inputs.AutomationRunbookReceiverArgs>? _automationRunbookReceivers;

        /// <summary>
        /// The list of AutomationRunbook receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.AutomationRunbookReceiverArgs> AutomationRunbookReceivers
        {
            get => _automationRunbookReceivers ?? (_automationRunbookReceivers = new InputList<Inputs.AutomationRunbookReceiverArgs>());
            set => _automationRunbookReceivers = value;
        }

        [Input("azureAppPushReceivers")]
        private InputList<Inputs.AzureAppPushReceiverArgs>? _azureAppPushReceivers;

        /// <summary>
        /// The list of AzureAppPush receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.AzureAppPushReceiverArgs> AzureAppPushReceivers
        {
            get => _azureAppPushReceivers ?? (_azureAppPushReceivers = new InputList<Inputs.AzureAppPushReceiverArgs>());
            set => _azureAppPushReceivers = value;
        }

        [Input("azureFunctionReceivers")]
        private InputList<Inputs.AzureFunctionReceiverArgs>? _azureFunctionReceivers;

        /// <summary>
        /// The list of azure function receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.AzureFunctionReceiverArgs> AzureFunctionReceivers
        {
            get => _azureFunctionReceivers ?? (_azureFunctionReceivers = new InputList<Inputs.AzureFunctionReceiverArgs>());
            set => _azureFunctionReceivers = value;
        }

        [Input("emailReceivers")]
        private InputList<Inputs.EmailReceiverArgs>? _emailReceivers;

        /// <summary>
        /// The list of email receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.EmailReceiverArgs> EmailReceivers
        {
            get => _emailReceivers ?? (_emailReceivers = new InputList<Inputs.EmailReceiverArgs>());
            set => _emailReceivers = value;
        }

        /// <summary>
        /// Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The short name of the action group. This will be used in SMS messages.
        /// </summary>
        [Input("groupShortName", required: true)]
        public Input<string> GroupShortName { get; set; } = null!;

        [Input("itsmReceivers")]
        private InputList<Inputs.ItsmReceiverArgs>? _itsmReceivers;

        /// <summary>
        /// The list of ITSM receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.ItsmReceiverArgs> ItsmReceivers
        {
            get => _itsmReceivers ?? (_itsmReceivers = new InputList<Inputs.ItsmReceiverArgs>());
            set => _itsmReceivers = value;
        }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("logicAppReceivers")]
        private InputList<Inputs.LogicAppReceiverArgs>? _logicAppReceivers;

        /// <summary>
        /// The list of logic app receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.LogicAppReceiverArgs> LogicAppReceivers
        {
            get => _logicAppReceivers ?? (_logicAppReceivers = new InputList<Inputs.LogicAppReceiverArgs>());
            set => _logicAppReceivers = value;
        }

        /// <summary>
        /// The name of the action group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("smsReceivers")]
        private InputList<Inputs.SmsReceiverArgs>? _smsReceivers;

        /// <summary>
        /// The list of SMS receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.SmsReceiverArgs> SmsReceivers
        {
            get => _smsReceivers ?? (_smsReceivers = new InputList<Inputs.SmsReceiverArgs>());
            set => _smsReceivers = value;
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

        [Input("voiceReceivers")]
        private InputList<Inputs.VoiceReceiverArgs>? _voiceReceivers;

        /// <summary>
        /// The list of voice receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.VoiceReceiverArgs> VoiceReceivers
        {
            get => _voiceReceivers ?? (_voiceReceivers = new InputList<Inputs.VoiceReceiverArgs>());
            set => _voiceReceivers = value;
        }

        [Input("webhookReceivers")]
        private InputList<Inputs.WebhookReceiverArgs>? _webhookReceivers;

        /// <summary>
        /// The list of webhook receivers that are part of this action group.
        /// </summary>
        public InputList<Inputs.WebhookReceiverArgs> WebhookReceivers
        {
            get => _webhookReceivers ?? (_webhookReceivers = new InputList<Inputs.WebhookReceiverArgs>());
            set => _webhookReceivers = value;
        }

        public ActionGroupArgs()
        {
        }
    }
}
