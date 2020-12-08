// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cdn.V20200901.Inputs
{

    /// <summary>
    /// Defines the parameters for RequestHeader match conditions
    /// </summary>
    public sealed class RequestHeaderMatchConditionParametersArgs : Pulumi.ResourceArgs
    {
        [Input("matchValues")]
        private InputList<string>? _matchValues;

        /// <summary>
        /// The match value for the condition of the delivery rule
        /// </summary>
        public InputList<string> MatchValues
        {
            get => _matchValues ?? (_matchValues = new InputList<string>());
            set => _matchValues = value;
        }

        /// <summary>
        /// Describes if this is negate condition or not
        /// </summary>
        [Input("negateCondition")]
        public Input<bool>? NegateCondition { get; set; }

        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        /// <summary>
        /// Describes operator to be matched
        /// </summary>
        [Input("operator", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Cdn.V20200901.RequestHeaderOperator> Operator { get; set; } = null!;

        /// <summary>
        /// Name of Header to be matched
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        [Input("transforms")]
        private InputList<Union<string, Pulumi.AzureNextGen.Cdn.V20200901.Transform>>? _transforms;

        /// <summary>
        /// List of transforms
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNextGen.Cdn.V20200901.Transform>> Transforms
        {
            get => _transforms ?? (_transforms = new InputList<Union<string, Pulumi.AzureNextGen.Cdn.V20200901.Transform>>());
            set => _transforms = value;
        }

        public RequestHeaderMatchConditionParametersArgs()
        {
        }
    }
}
