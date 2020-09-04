// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180701Preview.Inputs
{

    /// <summary>
    /// Describes the diagnostics options available
    /// </summary>
    public sealed class DiagnosticsDescriptionArgs : Pulumi.ResourceArgs
    {
        [Input("defaultSinkRefs")]
        private InputList<string>? _defaultSinkRefs;

        /// <summary>
        /// The sinks to be used if diagnostics is enabled. Sink choices can be overridden at the service and code package level.
        /// </summary>
        public InputList<string> DefaultSinkRefs
        {
            get => _defaultSinkRefs ?? (_defaultSinkRefs = new InputList<string>());
            set => _defaultSinkRefs = value;
        }

        /// <summary>
        /// Status of whether or not sinks are enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("sinks")]
        private InputList<Inputs.DiagnosticsSinkPropertiesArgs>? _sinks;

        /// <summary>
        /// List of supported sinks that can be referenced.
        /// </summary>
        public InputList<Inputs.DiagnosticsSinkPropertiesArgs> Sinks
        {
            get => _sinks ?? (_sinks = new InputList<Inputs.DiagnosticsSinkPropertiesArgs>());
            set => _sinks = value;
        }

        public DiagnosticsDescriptionArgs()
        {
        }
    }
}
