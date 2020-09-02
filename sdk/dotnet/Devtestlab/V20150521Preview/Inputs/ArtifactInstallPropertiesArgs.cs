// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20150521Preview.Inputs
{

    /// <summary>
    /// Properties of an artifact.
    /// </summary>
    public sealed class ArtifactInstallPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The artifact's identifier.
        /// </summary>
        [Input("artifactId")]
        public Input<string>? ArtifactId { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ArtifactParameterPropertiesArgs>? _parameters;

        /// <summary>
        /// The parameters of the artifact.
        /// </summary>
        public InputList<Inputs.ArtifactParameterPropertiesArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ArtifactParameterPropertiesArgs>());
            set => _parameters = value;
        }

        public ArtifactInstallPropertiesArgs()
        {
        }
    }
}
