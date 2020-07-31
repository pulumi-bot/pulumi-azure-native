// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20160601.Inputs
{

    /// <summary>
    /// The integration account map.
    /// </summary>
    public sealed class IntegrationAccountMapPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The content type.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The map type.
        /// </summary>
        [Input("mapType", required: true)]
        public Input<string> MapType { get; set; } = null!;

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// The metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The parameters schema of integration account map.
        /// </summary>
        [Input("parametersSchema")]
        public Input<Inputs.IntegrationAccountMapPropertiesPropertiesArgs>? ParametersSchema { get; set; }

        public IntegrationAccountMapPropertiesArgs()
        {
        }
    }
}