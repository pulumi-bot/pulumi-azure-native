// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20190615Preview.Inputs
{

    /// <summary>
    /// Defines a managed rule set.
    /// </summary>
    public sealed class ManagedRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Verizon only : If the rule set supports anomaly detection mode, this describes the threshold for blocking requests.
        /// </summary>
        [Input("anomalyScore")]
        public Input<int>? AnomalyScore { get; set; }

        [Input("ruleGroupOverrides")]
        private InputList<Inputs.ManagedRuleGroupOverrideArgs>? _ruleGroupOverrides;

        /// <summary>
        /// Defines the rule overrides to apply to the rule set.
        /// </summary>
        public InputList<Inputs.ManagedRuleGroupOverrideArgs> RuleGroupOverrides
        {
            get => _ruleGroupOverrides ?? (_ruleGroupOverrides = new InputList<Inputs.ManagedRuleGroupOverrideArgs>());
            set => _ruleGroupOverrides = value;
        }

        /// <summary>
        /// Defines the rule set type to use.
        /// </summary>
        [Input("ruleSetType", required: true)]
        public Input<string> RuleSetType { get; set; } = null!;

        /// <summary>
        /// Defines the version of the rule set to use.
        /// </summary>
        [Input("ruleSetVersion", required: true)]
        public Input<string> RuleSetVersion { get; set; } = null!;

        public ManagedRuleSetArgs()
        {
        }
    }
}
