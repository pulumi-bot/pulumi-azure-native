// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20170801Preview.Inputs
{

    /// <summary>
    /// A custom alert rule that checks if a value (depends on the custom alert type) is allowed
    /// </summary>
    public sealed class AllowlistCustomAlertRuleArgs : Pulumi.ResourceArgs
    {
        [Input("allowlistValues", required: true)]
        private InputList<string>? _allowlistValues;

        /// <summary>
        /// The values to allow. The format of the values depends on the rule type.
        /// </summary>
        public InputList<string> AllowlistValues
        {
            get => _allowlistValues ?? (_allowlistValues = new InputList<string>());
            set => _allowlistValues = value;
        }

        /// <summary>
        /// Whether the custom alert is enabled.
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        /// <summary>
        /// The type of the custom alert rule.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<string> RuleType { get; set; } = null!;

        public AllowlistCustomAlertRuleArgs()
        {
        }
    }
}
