// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20210101.Inputs
{

    public sealed class TaskContainerSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// These additional options are supplied as arguments to the "docker create" command, in addition to those controlled by the Batch Service.
        /// </summary>
        [Input("containerRunOptions")]
        public Input<string>? ContainerRunOptions { get; set; }

        /// <summary>
        /// This is the full image reference, as would be specified to "docker pull". If no tag is provided as part of the image name, the tag ":latest" is used as a default.
        /// </summary>
        [Input("imageName", required: true)]
        public Input<string> ImageName { get; set; } = null!;

        /// <summary>
        /// This setting can be omitted if was already provided at pool creation.
        /// </summary>
        [Input("registry")]
        public Input<Inputs.ContainerRegistryArgs>? Registry { get; set; }

        [Input("workingDirectory")]
        public Input<Pulumi.AzureNextGen.Batch.V20210101.ContainerWorkingDirectory>? WorkingDirectory { get; set; }

        public TaskContainerSettingsArgs()
        {
        }
    }
}
