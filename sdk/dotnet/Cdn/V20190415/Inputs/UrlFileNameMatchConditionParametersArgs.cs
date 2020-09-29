// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cdn.V20190415.Inputs
{

    /// <summary>
    /// Defines the parameters for UrlFilename match conditions
    /// </summary>
    public sealed class UrlFileNameMatchConditionParametersArgs : Pulumi.ResourceArgs
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
        public Input<string> Operator { get; set; } = null!;

        [Input("transforms")]
        private InputList<string>? _transforms;

        /// <summary>
        /// List of transforms
        /// </summary>
        public InputList<string> Transforms
        {
            get => _transforms ?? (_transforms = new InputList<string>());
            set => _transforms = value;
        }

        public UrlFileNameMatchConditionParametersArgs()
        {
        }
    }
}
