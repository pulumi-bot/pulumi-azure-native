// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180901Preview.Outputs
{

    [OutputType]
    public sealed class DiagnosticsDescriptionResponseResult
    {
        /// <summary>
        /// The sinks to be used if diagnostics is enabled. Sink choices can be overridden at the service and code package level.
        /// </summary>
        public readonly ImmutableArray<string> DefaultSinkRefs;
        /// <summary>
        /// Status of whether or not sinks are enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// List of supported sinks that can be referenced.
        /// </summary>
        public readonly ImmutableArray<Outputs.DiagnosticsSinkPropertiesResponseResult> Sinks;

        [OutputConstructor]
        private DiagnosticsDescriptionResponseResult(
            ImmutableArray<string> defaultSinkRefs,

            bool? enabled,

            ImmutableArray<Outputs.DiagnosticsSinkPropertiesResponseResult> sinks)
        {
            DefaultSinkRefs = defaultSinkRefs;
            Enabled = enabled;
            Sinks = sinks;
        }
    }
}
