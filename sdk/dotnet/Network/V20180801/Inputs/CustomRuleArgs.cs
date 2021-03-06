// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20180801.Inputs
{

    /// <summary>
    /// Defines contents of a web application rule
    /// </summary>
    public sealed class CustomRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Actions
        /// </summary>
        [Input("action", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.V20180801.Action> Action { get; set; } = null!;

        [Input("matchConditions", required: true)]
        private InputList<Inputs.MatchConditionArgs>? _matchConditions;

        /// <summary>
        /// List of match conditions
        /// </summary>
        public InputList<Inputs.MatchConditionArgs> MatchConditions
        {
            get => _matchConditions ?? (_matchConditions = new InputList<Inputs.MatchConditionArgs>());
            set => _matchConditions = value;
        }

        /// <summary>
        /// Gets name of the resource that is unique within a policy. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// Defines rate limit duration. Default - 1 minute
        /// </summary>
        [Input("rateLimitDurationInMinutes")]
        public Input<int>? RateLimitDurationInMinutes { get; set; }

        /// <summary>
        /// Defines rate limit threshold
        /// </summary>
        [Input("rateLimitThreshold")]
        public Input<int>? RateLimitThreshold { get; set; }

        /// <summary>
        /// Describes type of rule
        /// </summary>
        [Input("ruleType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.V20180801.RuleType> RuleType { get; set; } = null!;

        [Input("transforms")]
        private InputList<Union<string, Pulumi.AzureNative.Network.V20180801.Transform>>? _transforms;

        /// <summary>
        /// List of transforms
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Network.V20180801.Transform>> Transforms
        {
            get => _transforms ?? (_transforms = new InputList<Union<string, Pulumi.AzureNative.Network.V20180801.Transform>>());
            set => _transforms = value;
        }

        public CustomRuleArgs()
        {
        }
    }
}
