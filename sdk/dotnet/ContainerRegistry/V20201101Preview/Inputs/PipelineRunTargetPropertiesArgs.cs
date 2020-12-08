// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20201101Preview.Inputs
{

    public sealed class PipelineRunTargetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the target.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the target.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerRegistry.V20201101Preview.PipelineRunTargetType>? Type { get; set; }

        public PipelineRunTargetPropertiesArgs()
        {
        }
    }
}
