// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation
{
    /// <summary>
    /// Definition of the webhook type.
    /// </summary>
    public partial class AutomationAccountWebhook : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the webhook properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.WebhookPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AutomationAccountWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutomationAccountWebhook(string name, AutomationAccountWebhookArgs args, CustomResourceOptions? options = null)
            : base("azurerm:automation:AutomationAccountWebhook", name, args ?? new AutomationAccountWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutomationAccountWebhook(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:automation:AutomationAccountWebhook", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AutomationAccountWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutomationAccountWebhook Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AutomationAccountWebhook(name, id, options);
        }
    }

    public sealed class AutomationAccountWebhookArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The webhook name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Gets or sets the properties of the webhook.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.WebhookCreateOrUpdatePropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AutomationAccountWebhookArgs()
        {
        }
    }
}
