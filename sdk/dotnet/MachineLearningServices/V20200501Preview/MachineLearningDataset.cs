// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.MachineLearningServices.V20200501Preview
{
    /// <summary>
    /// Machine Learning dataset object wrapped into ARM resource envelope.
    /// </summary>
    public partial class MachineLearningDataset : Pulumi.CustomResource
    {
        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Dataset properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.DatasetResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The sku of the workspace.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Contains resource tags defined as key/value pairs.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a MachineLearningDataset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MachineLearningDataset(string name, MachineLearningDatasetArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:machinelearningservices/v20200501preview:MachineLearningDataset", name, args ?? new MachineLearningDatasetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MachineLearningDataset(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:machinelearningservices/v20200501preview:MachineLearningDataset", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MachineLearningDataset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MachineLearningDataset Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MachineLearningDataset(name, id, options);
        }
    }

    public sealed class MachineLearningDatasetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Dataset name.
        /// </summary>
        [Input("datasetName", required: true)]
        public Input<string> DatasetName { get; set; } = null!;

        /// <summary>
        /// Specifies dataset type.
        /// </summary>
        [Input("datasetType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.MachineLearningServices.V20200501Preview.DatasetType> DatasetType { get; set; } = null!;

        [Input("parameters", required: true)]
        public Input<Inputs.DatasetCreateRequestParametersArgs> Parameters { get; set; } = null!;

        [Input("registration", required: true)]
        public Input<Inputs.DatasetCreateRequestRegistrationArgs> Registration { get; set; } = null!;

        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Skip validation that ensures data can be loaded from the dataset before registration.
        /// </summary>
        [Input("skipValidation")]
        public Input<bool>? SkipValidation { get; set; }

        [Input("timeSeries")]
        public Input<Inputs.DatasetCreateRequestTimeSeriesArgs>? TimeSeries { get; set; }

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public MachineLearningDatasetArgs()
        {
        }
    }
}
