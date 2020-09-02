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
    /// Definition of which Windows Event Log events will be collected and how they will be collected.
    /// Only collected from Windows machines.
    /// </summary>
    public sealed class WindowsEventLogDataSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A friendly name for the data source. 
        /// This name should be unique across all data sources (regardless of type) within the data collection rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The interval between data uploads (scheduled transfers), rounded up to the nearest minute.
        /// </summary>
        [Input("scheduledTransferPeriod", required: true)]
        public Input<string> ScheduledTransferPeriod { get; set; } = null!;

        [Input("streams", required: true)]
        private InputList<string>? _streams;

        /// <summary>
        /// List of streams that this data source will be sent to.
        /// A stream indicates what schema will be used for this data and usually what table in Log Analytics the data will be sent to.
        /// </summary>
        public InputList<string> Streams
        {
            get => _streams ?? (_streams = new InputList<string>());
            set => _streams = value;
        }

        [Input("xPathQueries", required: true)]
        private InputList<string>? _xPathQueries;

        /// <summary>
        /// A list of Windows Event Log queries in XPATH format.
        /// </summary>
        public InputList<string> XPathQueries
        {
            get => _xPathQueries ?? (_xPathQueries = new InputList<string>());
            set => _xPathQueries = value;
        }

        public WindowsEventLogDataSourceArgs()
        {
        }
    }
}
