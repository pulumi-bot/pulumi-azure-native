// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Insights.V20140401.Inputs
{

    /// <summary>
    /// A management event rule condition.
    /// </summary>
    public sealed class ManagementEventRuleConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// How the data that is collected should be combined over time and when the alert is activated. Note that for management event alerts aggregation is optional – if it is not provided then any event will cause the alert to activate.
        /// </summary>
        [Input("aggregation")]
        public Input<Inputs.ManagementEventAggregationConditionArgs>? Aggregation { get; set; }

        /// <summary>
        /// the resource from which the rule collects its data. For this type dataSource will always be of type RuleMetricDataSource.
        /// </summary>
        [Input("dataSource")]
        public InputUnion<Inputs.RuleManagementEventDataSourceArgs, Inputs.RuleMetricDataSourceArgs>? DataSource { get; set; }

        /// <summary>
        /// specifies the type of condition. This can be one of three types: ManagementEventRuleCondition (occurrences of management events), LocationThresholdRuleCondition (based on the number of failures of a web test), and ThresholdRuleCondition (based on the threshold of a metric).
        /// Expected value is 'Microsoft.Azure.Management.Insights.Models.ManagementEventRuleCondition'.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        public ManagementEventRuleConditionArgs()
        {
        }
    }
}
