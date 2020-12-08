// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CostManagement.Latest.Inputs
{

    /// <summary>
    /// The definition of data present in the report.
    /// </summary>
    public sealed class ReportConfigDatasetArgs : Pulumi.ResourceArgs
    {
        [Input("aggregation")]
        private InputMap<Inputs.ReportConfigAggregationArgs>? _aggregation;

        /// <summary>
        /// Dictionary of aggregation expression to use in the report. The key of each item in the dictionary is the alias for the aggregated column. Report can have up to 2 aggregation clauses.
        /// </summary>
        public InputMap<Inputs.ReportConfigAggregationArgs> Aggregation
        {
            get => _aggregation ?? (_aggregation = new InputMap<Inputs.ReportConfigAggregationArgs>());
            set => _aggregation = value;
        }

        /// <summary>
        /// Has configuration information for the data in the report. The configuration will be ignored if aggregation and grouping are provided.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.ReportConfigDatasetConfigurationArgs>? Configuration { get; set; }

        /// <summary>
        /// Has filter expression to use in the report.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.ReportConfigFilterArgs>? Filter { get; set; }

        /// <summary>
        /// The granularity of rows in the report.
        /// </summary>
        [Input("granularity")]
        public InputUnion<string, Pulumi.AzureNextGen.CostManagement.Latest.ReportGranularityType>? Granularity { get; set; }

        [Input("grouping")]
        private InputList<Inputs.ReportConfigGroupingArgs>? _grouping;

        /// <summary>
        /// Array of group by expression to use in the report. Report can have up to 2 group by clauses.
        /// </summary>
        public InputList<Inputs.ReportConfigGroupingArgs> Grouping
        {
            get => _grouping ?? (_grouping = new InputList<Inputs.ReportConfigGroupingArgs>());
            set => _grouping = value;
        }

        [Input("sorting")]
        private InputList<Inputs.ReportConfigSortingArgs>? _sorting;

        /// <summary>
        /// Array of order by expression to use in the report.
        /// </summary>
        public InputList<Inputs.ReportConfigSortingArgs> Sorting
        {
            get => _sorting ?? (_sorting = new InputList<Inputs.ReportConfigSortingArgs>());
            set => _sorting = value;
        }

        public ReportConfigDatasetArgs()
        {
        }
    }
}
