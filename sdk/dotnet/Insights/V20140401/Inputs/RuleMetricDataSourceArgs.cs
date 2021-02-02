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
    /// A rule metric data source. The discriminator value is always RuleMetricDataSource in this case.
    /// </summary>
    public sealed class RuleMetricDataSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the name of the metric that defines what the rule monitors.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// specifies the type of data source. There are two types of rule data sources: RuleMetricDataSource and RuleManagementEventDataSource
        /// Expected value is 'Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource'.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        /// <summary>
        /// the resource identifier of the resource the rule monitors. **NOTE**: this property cannot be updated for an existing rule.
        /// </summary>
        [Input("resourceUri")]
        public Input<string>? ResourceUri { get; set; }

        public RuleMetricDataSourceArgs()
        {
        }
    }
}
