// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20190101Preview
{
    /// <summary>
    /// Describes the suppression rule
    /// </summary>
    public partial class AlertsSuppressionRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Type of the alert to automatically suppress. For all alert types, use '*'
        /// </summary>
        [Output("alertType")]
        public Output<string> AlertType { get; private set; } = null!;

        /// <summary>
        /// Any comment regarding the rule
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
        /// </summary>
        [Output("expirationDateUtc")]
        public Output<string?> ExpirationDateUtc { get; private set; } = null!;

        /// <summary>
        /// The last time this rule was modified
        /// </summary>
        [Output("lastModifiedUtc")]
        public Output<string> LastModifiedUtc { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The reason for dismissing the alert
        /// </summary>
        [Output("reason")]
        public Output<string> Reason { get; private set; } = null!;

        /// <summary>
        /// Possible states of the rule
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The suppression conditions
        /// </summary>
        [Output("suppressionAlertsScope")]
        public Output<Outputs.SuppressionAlertsScopeResponseResult?> SuppressionAlertsScope { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AlertsSuppressionRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertsSuppressionRule(string name, AlertsSuppressionRuleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:security/v20190101preview:AlertsSuppressionRule", name, args ?? new AlertsSuppressionRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertsSuppressionRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:security/v20190101preview:AlertsSuppressionRule", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AlertsSuppressionRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertsSuppressionRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AlertsSuppressionRule(name, id, options);
        }
    }

    public sealed class AlertsSuppressionRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the alert to automatically suppress. For all alert types, use '*'
        /// </summary>
        [Input("alertType", required: true)]
        public Input<string> AlertType { get; set; } = null!;

        /// <summary>
        /// The unique name of the suppression alert rule
        /// </summary>
        [Input("alertsSuppressionRuleName", required: true)]
        public Input<string> AlertsSuppressionRuleName { get; set; } = null!;

        /// <summary>
        /// Any comment regarding the rule
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
        /// </summary>
        [Input("expirationDateUtc")]
        public Input<string>? ExpirationDateUtc { get; set; }

        /// <summary>
        /// The reason for dismissing the alert
        /// </summary>
        [Input("reason", required: true)]
        public Input<string> Reason { get; set; } = null!;

        /// <summary>
        /// Possible states of the rule
        /// </summary>
        [Input("state", required: true)]
        public Input<string> State { get; set; } = null!;

        /// <summary>
        /// The suppression conditions
        /// </summary>
        [Input("suppressionAlertsScope")]
        public Input<Inputs.SuppressionAlertsScopeArgs>? SuppressionAlertsScope { get; set; }

        public AlertsSuppressionRuleArgs()
        {
        }
    }
}
