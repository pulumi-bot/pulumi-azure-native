// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180901Preview.Inputs
{

    /// <summary>
    /// Describes a container and its runtime properties.
    /// </summary>
    public sealed class ContainerCodePackagePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// Command array to execute within the container in exec form.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        /// <summary>
        /// Reference to sinks in DiagnosticsDescription.
        /// </summary>
        [Input("diagnostics")]
        public Input<Inputs.DiagnosticsRefArgs>? Diagnostics { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.EndpointPropertiesArgs>? _endpoints;

        /// <summary>
        /// The endpoints exposed by this container.
        /// </summary>
        public InputList<Inputs.EndpointPropertiesArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.EndpointPropertiesArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// Override for the default entry point in the container.
        /// </summary>
        [Input("entrypoint")]
        public Input<string>? Entrypoint { get; set; }

        [Input("environmentVariables")]
        private InputList<Inputs.EnvironmentVariableArgs>? _environmentVariables;

        /// <summary>
        /// The environment variables to set in this container
        /// </summary>
        public InputList<Inputs.EnvironmentVariableArgs> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputList<Inputs.EnvironmentVariableArgs>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The Container image to use.
        /// </summary>
        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        /// <summary>
        /// Image registry credential.
        /// </summary>
        [Input("imageRegistryCredential")]
        public Input<Inputs.ImageRegistryCredentialArgs>? ImageRegistryCredential { get; set; }

        [Input("labels")]
        private InputList<Inputs.ContainerLabelArgs>? _labels;

        /// <summary>
        /// The labels to set in this container.
        /// </summary>
        public InputList<Inputs.ContainerLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.ContainerLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the code package.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("reliableCollectionsRefs")]
        private InputList<Inputs.ReliableCollectionsRefArgs>? _reliableCollectionsRefs;

        /// <summary>
        /// A list of ReliableCollection resources used by this particular code package. Please refer to ReliableCollectionsRef for more details.
        /// </summary>
        public InputList<Inputs.ReliableCollectionsRefArgs> ReliableCollectionsRefs
        {
            get => _reliableCollectionsRefs ?? (_reliableCollectionsRefs = new InputList<Inputs.ReliableCollectionsRefArgs>());
            set => _reliableCollectionsRefs = value;
        }

        /// <summary>
        /// The resources required by this container.
        /// </summary>
        [Input("resources", required: true)]
        public Input<Inputs.ResourceRequirementsArgs> Resources { get; set; } = null!;

        [Input("settings")]
        private InputList<Inputs.SettingArgs>? _settings;

        /// <summary>
        /// The settings to set in this container. The setting file path can be fetched from environment variable "Fabric_SettingPath". The path for Windows container is "C:\\secrets". The path for Linux container is "/var/secrets".
        /// </summary>
        public InputList<Inputs.SettingArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.SettingArgs>());
            set => _settings = value;
        }

        [Input("volumeRefs")]
        private InputList<Inputs.VolumeReferenceArgs>? _volumeRefs;

        /// <summary>
        /// Volumes to be attached to the container. The lifetime of these volumes is independent of the application's lifetime.
        /// </summary>
        public InputList<Inputs.VolumeReferenceArgs> VolumeRefs
        {
            get => _volumeRefs ?? (_volumeRefs = new InputList<Inputs.VolumeReferenceArgs>());
            set => _volumeRefs = value;
        }

        [Input("volumes")]
        private InputList<Inputs.ApplicationScopedVolumeArgs>? _volumes;

        /// <summary>
        /// Volumes to be attached to the container. The lifetime of these volumes is scoped to the application's lifetime.
        /// </summary>
        public InputList<Inputs.ApplicationScopedVolumeArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.ApplicationScopedVolumeArgs>());
            set => _volumes = value;
        }

        public ContainerCodePackagePropertiesArgs()
        {
        }
    }
}
