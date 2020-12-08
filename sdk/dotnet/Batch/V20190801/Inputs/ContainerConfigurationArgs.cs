// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20190801.Inputs
{

    public sealed class ContainerConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("containerImageNames")]
        private InputList<string>? _containerImageNames;

        /// <summary>
        /// This is the full image reference, as would be specified to "docker pull". An image will be sourced from the default Docker registry unless the image is fully qualified with an alternative registry.
        /// </summary>
        public InputList<string> ContainerImageNames
        {
            get => _containerImageNames ?? (_containerImageNames = new InputList<string>());
            set => _containerImageNames = value;
        }

        [Input("containerRegistries")]
        private InputList<Inputs.ContainerRegistryArgs>? _containerRegistries;

        /// <summary>
        /// If any images must be downloaded from a private registry which requires credentials, then those credentials must be provided here.
        /// </summary>
        public InputList<Inputs.ContainerRegistryArgs> ContainerRegistries
        {
            get => _containerRegistries ?? (_containerRegistries = new InputList<Inputs.ContainerRegistryArgs>());
            set => _containerRegistries = value;
        }

        [Input("type", required: true)]
        public Input<Pulumi.AzureNextGen.Batch.V20190801.ContainerType> Type { get; set; } = null!;

        public ContainerConfigurationArgs()
        {
        }
    }
}
