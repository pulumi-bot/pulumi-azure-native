// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.MachineLearning.Latest.Inputs
{

    /// <summary>
    /// Diagnostics settings for an Azure ML web service.
    /// </summary>
    public sealed class DiagnosticsConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the date and time when the logging will cease. If null, diagnostic collection is not time limited.
        /// </summary>
        [Input("expiry")]
        public Input<string>? Expiry { get; set; }

        /// <summary>
        /// Specifies the verbosity of the diagnostic output. Valid values are: None - disables tracing; Error - collects only error (stderr) traces; All - collects all traces (stdout and stderr).
        /// </summary>
        [Input("level", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.MachineLearning.Latest.DiagnosticsLevel> Level { get; set; } = null!;

        public DiagnosticsConfigurationArgs()
        {
        }
    }
}
