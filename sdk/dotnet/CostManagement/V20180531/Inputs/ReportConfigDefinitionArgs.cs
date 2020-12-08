// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CostManagement.V20180531.Inputs
{

    /// <summary>
    /// The definition of a report config.
    /// </summary>
    public sealed class ReportConfigDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Has definition for data in this report config.
        /// </summary>
        [Input("dataset")]
        public Input<Inputs.ReportConfigDatasetArgs>? Dataset { get; set; }

        /// <summary>
        /// Has time period for pulling data for the report.
        /// </summary>
        [Input("timePeriod")]
        public Input<Inputs.ReportConfigTimePeriodArgs>? TimePeriod { get; set; }

        /// <summary>
        /// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
        /// </summary>
        [Input("timeframe", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.CostManagement.V20180531.TimeframeType> Timeframe { get; set; } = null!;

        /// <summary>
        /// The type of the report.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.CostManagement.V20180531.ReportType> Type { get; set; } = null!;

        public ReportConfigDefinitionArgs()
        {
        }
    }
}
