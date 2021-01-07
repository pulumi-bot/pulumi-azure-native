// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801.Inputs
{

    /// <summary>
    /// Firewall Policy Filter Rule Collection.
    /// </summary>
    public sealed class FirewallPolicyFilterRuleCollectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action type of a Filter rule collection.
        /// </summary>
        [Input("action")]
        public Input<Inputs.FirewallPolicyFilterRuleCollectionActionArgs>? Action { get; set; }

        /// <summary>
        /// The name of the rule collection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority of the Firewall Policy Rule Collection resource.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The type of the rule collection.
        /// Expected value is 'FirewallPolicyFilterRuleCollection'.
        /// </summary>
        [Input("ruleCollectionType", required: true)]
        public Input<string> RuleCollectionType { get; set; } = null!;

        [Input("rules")]
        private InputList<object>? _rules;

        /// <summary>
        /// List of rules included in a rule collection.
        /// </summary>
        public InputList<object> Rules
        {
            get => _rules ?? (_rules = new InputList<object>());
            set => _rules = value;
        }

        public FirewallPolicyFilterRuleCollectionArgs()
        {
        }
    }
}
