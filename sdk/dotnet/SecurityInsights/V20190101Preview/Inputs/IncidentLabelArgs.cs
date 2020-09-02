// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SecurityInsights.V20190101Preview.Inputs
{

    /// <summary>
    /// Represents an incident label
    /// </summary>
    public sealed class IncidentLabelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the label
        /// </summary>
        [Input("labelName", required: true)]
        public Input<string> LabelName { get; set; } = null!;

        public IncidentLabelArgs()
        {
        }
    }
}
