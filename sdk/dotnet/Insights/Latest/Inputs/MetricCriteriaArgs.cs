// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Insights.Latest.Inputs
{

    /// <summary>
    /// Criterion to filter metrics.
    /// </summary>
    public sealed class MetricCriteriaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the type of threshold criteria
        /// Expected value is 'StaticThresholdCriterion'.
        /// </summary>
        [Input("criterionType", required: true)]
        public Input<string> CriterionType { get; set; } = null!;

        [Input("dimensions")]
        private InputList<Inputs.MetricDimensionArgs>? _dimensions;

        /// <summary>
        /// List of dimension conditions.
        /// </summary>
        public InputList<Inputs.MetricDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.MetricDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Name of the metric.
        /// </summary>
        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        /// <summary>
        /// Namespace of the metric.
        /// </summary>
        [Input("metricNamespace")]
        public Input<string>? MetricNamespace { get; set; }

        /// <summary>
        /// Name of the criteria.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// the criteria operator.
        /// </summary>
        [Input("operator", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Insights.Latest.Operator> Operator { get; set; } = null!;

        /// <summary>
        /// Allows creating an alert rule on a custom metric that isn't yet emitted, by causing the metric validation to be skipped.
        /// </summary>
        [Input("skipMetricValidation")]
        public Input<bool>? SkipMetricValidation { get; set; }

        /// <summary>
        /// the criteria threshold value that activates the alert.
        /// </summary>
        [Input("threshold", required: true)]
        public Input<double> Threshold { get; set; } = null!;

        /// <summary>
        /// the criteria time aggregation types.
        /// </summary>
        [Input("timeAggregation", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Insights.Latest.AggregationTypeEnum> TimeAggregation { get; set; } = null!;

        public MetricCriteriaArgs()
        {
        }
    }
}
