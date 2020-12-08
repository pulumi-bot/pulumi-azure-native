// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Consumption.Latest.Inputs
{

    /// <summary>
    /// The comparison expression to be used in the budgets.
    /// </summary>
    public sealed class BudgetComparisonExpressionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the column to use in comparison.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The operator to use for comparison.
        /// </summary>
        [Input("operator", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Consumption.Latest.BudgetOperatorType> Operator { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// Array of values to use for comparison
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public BudgetComparisonExpressionArgs()
        {
        }
    }
}
