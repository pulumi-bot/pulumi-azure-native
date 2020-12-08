// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20191201.Inputs
{

    /// <summary>
    /// Define match conditions.
    /// </summary>
    public sealed class MatchConditionArgs : Pulumi.ResourceArgs
    {
        [Input("matchValues", required: true)]
        private InputList<string>? _matchValues;

        /// <summary>
        /// Match value.
        /// </summary>
        public InputList<string> MatchValues
        {
            get => _matchValues ?? (_matchValues = new InputList<string>());
            set => _matchValues = value;
        }

        [Input("matchVariables", required: true)]
        private InputList<Inputs.MatchVariableArgs>? _matchVariables;

        /// <summary>
        /// List of match variables.
        /// </summary>
        public InputList<Inputs.MatchVariableArgs> MatchVariables
        {
            get => _matchVariables ?? (_matchVariables = new InputList<Inputs.MatchVariableArgs>());
            set => _matchVariables = value;
        }

        /// <summary>
        /// Whether this is negate condition or not.
        /// </summary>
        [Input("negationConditon")]
        public Input<bool>? NegationConditon { get; set; }

        /// <summary>
        /// The operator to be matched.
        /// </summary>
        [Input("operator", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20191201.WebApplicationFirewallOperator> Operator { get; set; } = null!;

        [Input("transforms")]
        private InputList<Union<string, Pulumi.AzureNextGen.Network.V20191201.WebApplicationFirewallTransform>>? _transforms;

        /// <summary>
        /// List of transforms.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNextGen.Network.V20191201.WebApplicationFirewallTransform>> Transforms
        {
            get => _transforms ?? (_transforms = new InputList<Union<string, Pulumi.AzureNextGen.Network.V20191201.WebApplicationFirewallTransform>>());
            set => _transforms = value;
        }

        public MatchConditionArgs()
        {
        }
    }
}
