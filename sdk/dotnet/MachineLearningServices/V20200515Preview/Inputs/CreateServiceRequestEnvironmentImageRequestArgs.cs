// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200515Preview.Inputs
{

    /// <summary>
    /// The Environment, models and assets needed for inferencing.
    /// </summary>
    public sealed class CreateServiceRequestEnvironmentImageRequestArgs : Pulumi.ResourceArgs
    {
        [Input("assets")]
        private InputList<Inputs.ImageAssetArgs>? _assets;

        /// <summary>
        /// The list of assets.
        /// </summary>
        public InputList<Inputs.ImageAssetArgs> Assets
        {
            get => _assets ?? (_assets = new InputList<Inputs.ImageAssetArgs>());
            set => _assets = value;
        }

        /// <summary>
        /// The name of the driver file.
        /// </summary>
        [Input("driverProgram")]
        public Input<string>? DriverProgram { get; set; }

        /// <summary>
        /// The details of the AZURE ML environment.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.EnvironmentImageRequestEnvironmentArgs>? Environment { get; set; }

        /// <summary>
        /// The unique identifying details of the AZURE ML environment.
        /// </summary>
        [Input("environmentReference")]
        public Input<Inputs.EnvironmentImageRequestEnvironmentReferenceArgs>? EnvironmentReference { get; set; }

        [Input("modelIds")]
        private InputList<string>? _modelIds;

        /// <summary>
        /// The list of model Ids.
        /// </summary>
        public InputList<string> ModelIds
        {
            get => _modelIds ?? (_modelIds = new InputList<string>());
            set => _modelIds = value;
        }

        [Input("models")]
        private InputList<Inputs.ModelArgs>? _models;

        /// <summary>
        /// The list of models.
        /// </summary>
        public InputList<Inputs.ModelArgs> Models
        {
            get => _models ?? (_models = new InputList<Inputs.ModelArgs>());
            set => _models = value;
        }

        public CreateServiceRequestEnvironmentImageRequestArgs()
        {
        }
    }
}
