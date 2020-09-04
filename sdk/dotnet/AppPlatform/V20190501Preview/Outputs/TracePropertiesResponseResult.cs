// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppPlatform.V20190501Preview.Outputs
{

    [OutputType]
    public sealed class TracePropertiesResponseResult
    {
        /// <summary>
        /// Target application insight instrumentation key
        /// </summary>
        public readonly string? AppInsightInstrumentationKey;
        /// <summary>
        /// Indicates whether enable the tracing functionality
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Error when apply trace proxy changes.
        /// </summary>
        public readonly Outputs.ErrorResponseResult? Error;
        /// <summary>
        /// State of the trace proxy.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private TracePropertiesResponseResult(
            string? appInsightInstrumentationKey,

            bool? enabled,

            Outputs.ErrorResponseResult? error,

            string state)
        {
            AppInsightInstrumentationKey = appInsightInstrumentationKey;
            Enabled = enabled;
            Error = error;
            State = state;
        }
    }
}
