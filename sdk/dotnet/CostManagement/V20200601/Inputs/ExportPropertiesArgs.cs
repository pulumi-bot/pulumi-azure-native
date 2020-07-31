// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CostManagement.V20200601.Inputs
{

    /// <summary>
    /// The properties of the export.
    /// </summary>
    public sealed class ExportPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Has the definition for the export.
        /// </summary>
        [Input("definition", required: true)]
        public Input<Inputs.ExportDefinitionArgs> Definition { get; set; } = null!;

        /// <summary>
        /// Has delivery information for the export.
        /// </summary>
        [Input("deliveryInfo", required: true)]
        public Input<Inputs.ExportDeliveryInfoArgs> DeliveryInfo { get; set; } = null!;

        /// <summary>
        /// The format of the export being delivered. Currently only 'Csv' is supported.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// If requested, has the most recent execution history for the export.
        /// </summary>
        [Input("runHistory")]
        public Input<Inputs.ExportExecutionListResultArgs>? RunHistory { get; set; }

        /// <summary>
        /// Has schedule information for the export.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ExportScheduleArgs>? Schedule { get; set; }

        public ExportPropertiesArgs()
        {
        }
    }
}