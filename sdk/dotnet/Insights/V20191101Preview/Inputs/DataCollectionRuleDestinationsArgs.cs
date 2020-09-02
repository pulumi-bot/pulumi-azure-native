// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20191101Preview.Inputs
{

    /// <summary>
    /// The specification of destinations.
    /// </summary>
    public sealed class DataCollectionRuleDestinationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure Monitor Metrics destination.
        /// </summary>
        [Input("azureMonitorMetrics")]
        public Input<Inputs.DestinationsSpecAzureMonitorMetricsArgs>? AzureMonitorMetrics { get; set; }

        [Input("logAnalytics")]
        private InputList<Inputs.LogAnalyticsDestinationArgs>? _logAnalytics;

        /// <summary>
        /// List of Log Analytics destinations.
        /// </summary>
        public InputList<Inputs.LogAnalyticsDestinationArgs> LogAnalytics
        {
            get => _logAnalytics ?? (_logAnalytics = new InputList<Inputs.LogAnalyticsDestinationArgs>());
            set => _logAnalytics = value;
        }

        public DataCollectionRuleDestinationsArgs()
        {
        }
    }
}
