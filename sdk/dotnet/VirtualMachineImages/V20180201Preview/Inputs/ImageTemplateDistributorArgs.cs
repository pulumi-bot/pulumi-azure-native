// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.VirtualMachineImages.V20180201Preview.Inputs
{

    /// <summary>
    /// Generic distribution object
    /// </summary>
    public sealed class ImageTemplateDistributorArgs : Pulumi.ResourceArgs
    {
        [Input("artifactTags")]
        private InputMap<string>? _artifactTags;

        /// <summary>
        /// Tags that will be applied to the artifact once it has been created/updated by the distributor.
        /// </summary>
        public InputMap<string> ArtifactTags
        {
            get => _artifactTags ?? (_artifactTags = new InputMap<string>());
            set => _artifactTags = value;
        }

        /// <summary>
        /// The name to be used for the associated RunOutput.
        /// </summary>
        [Input("runOutputName", required: true)]
        public Input<string> RunOutputName { get; set; } = null!;

        /// <summary>
        /// Type of distribution.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ImageTemplateDistributorArgs()
        {
        }
    }
}
