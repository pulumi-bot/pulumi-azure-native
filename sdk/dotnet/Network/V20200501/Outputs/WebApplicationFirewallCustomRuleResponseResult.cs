// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Outputs
{

    [OutputType]
    public sealed class WebApplicationFirewallCustomRuleResponseResult
    {
        /// <summary>
        /// Type of Actions.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// List of match conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.MatchConditionResponseResult> MatchConditions;
        /// <summary>
        /// The name of the resource that is unique within a policy. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// The rule type.
        /// </summary>
        public readonly string RuleType;

        [OutputConstructor]
        private WebApplicationFirewallCustomRuleResponseResult(
            string action,

            string etag,

            ImmutableArray<Outputs.MatchConditionResponseResult> matchConditions,

            string? name,

            int priority,

            string ruleType)
        {
            Action = action;
            Etag = etag;
            MatchConditions = matchConditions;
            Name = name;
            Priority = priority;
            RuleType = ruleType;
        }
    }
}
