// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20170501Preview.Inputs
{

    /// <summary>
    /// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
    /// </summary>
    public sealed class MetricSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// a value indicating whether this category is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// the retention policy for this category.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.RetentionPolicyArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// the timegrain of the metric in ISO8601 format.
        /// </summary>
        [Input("timeGrain")]
        public Input<string>? TimeGrain { get; set; }

        public MetricSettingsArgs()
        {
        }
    }
}
