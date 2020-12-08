// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cdn.V20190615.Inputs
{

    /// <summary>
    /// Defines a managed rule group override setting.
    /// </summary>
    public sealed class ManagedRuleOverrideArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the override action to be applied when rule matches.
        /// </summary>
        [Input("action")]
        public InputUnion<string, Pulumi.AzureNextGen.Cdn.V20190615.ActionType>? Action { get; set; }

        /// <summary>
        /// Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
        /// </summary>
        [Input("enabledState")]
        public InputUnion<string, Pulumi.AzureNextGen.Cdn.V20190615.ManagedRuleEnabledState>? EnabledState { get; set; }

        /// <summary>
        /// Identifier for the managed rule.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        public ManagedRuleOverrideArgs()
        {
        }
    }
}
