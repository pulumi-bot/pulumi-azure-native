// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights.V20191101Preview.Inputs
{

    /// <summary>
    /// Definition of which performance counters will be collected and how they will be collected by this data collection rule.
    /// Collected from both Windows and Linux machines where the counter is present.
    /// </summary>
    public sealed class PerfCounterDataSourceArgs : Pulumi.ResourceArgs
    {
        [Input("counterSpecifiers", required: true)]
        private InputList<string>? _counterSpecifiers;

        /// <summary>
        /// A list of specifier names of the performance counters you want to collect.
        /// Use a wildcard (*) to collect a counter for all instances.
        /// To get a list of performance counters on Windows, run the command 'typeperf'.
        /// </summary>
        public InputList<string> CounterSpecifiers
        {
            get => _counterSpecifiers ?? (_counterSpecifiers = new InputList<string>());
            set => _counterSpecifiers = value;
        }

        /// <summary>
        /// A friendly name for the data source. 
        /// This name should be unique across all data sources (regardless of type) within the data collection rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The number of seconds between consecutive counter measurements (samples).
        /// </summary>
        [Input("samplingFrequencyInSeconds", required: true)]
        public Input<int> SamplingFrequencyInSeconds { get; set; } = null!;

        /// <summary>
        /// The interval between data uploads (scheduled transfers), rounded up to the nearest minute.
        /// </summary>
        [Input("scheduledTransferPeriod", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Insights.V20191101Preview.KnownPerfCounterDataSourceScheduledTransferPeriod> ScheduledTransferPeriod { get; set; } = null!;

        [Input("streams", required: true)]
        private InputList<Union<string, Pulumi.AzureNative.Insights.V20191101Preview.KnownPerfCounterDataSourceStreams>>? _streams;

        /// <summary>
        /// List of streams that this data source will be sent to.
        /// A stream indicates what schema will be used for this data and usually what table in Log Analytics the data will be sent to.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Insights.V20191101Preview.KnownPerfCounterDataSourceStreams>> Streams
        {
            get => _streams ?? (_streams = new InputList<Union<string, Pulumi.AzureNative.Insights.V20191101Preview.KnownPerfCounterDataSourceStreams>>());
            set => _streams = value;
        }

        public PerfCounterDataSourceArgs()
        {
        }
    }
}
