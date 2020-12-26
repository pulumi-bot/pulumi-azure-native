// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DevOps.V20200713Preview.Inputs
{

    /// <summary>
    /// Configuration used to bootstrap a Pipeline.
    /// </summary>
    public sealed class BootstrapConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Repository containing the source code for the pipeline. Currently only 'azurePipeline' pipeline type supports this.
        /// </summary>
        [Input("sourceRepository")]
        public Input<Inputs.CodeRepositoryArgs>? SourceRepository { get; set; }

        /// <summary>
        /// Template used to bootstrap the pipeline.
        /// </summary>
        [Input("template", required: true)]
        public Input<Inputs.PipelineTemplateArgs> Template { get; set; } = null!;

        public BootstrapConfigurationArgs()
        {
        }
    }
}
