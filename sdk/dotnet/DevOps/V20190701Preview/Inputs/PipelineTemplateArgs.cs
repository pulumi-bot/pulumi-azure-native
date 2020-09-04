// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevOps.V20190701Preview.Inputs
{

    /// <summary>
    /// Template used to bootstrap the pipeline.
    /// </summary>
    public sealed class PipelineTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier of the pipeline template.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Dictionary of input parameters used in the pipeline template.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        public PipelineTemplateArgs()
        {
        }
    }
}
