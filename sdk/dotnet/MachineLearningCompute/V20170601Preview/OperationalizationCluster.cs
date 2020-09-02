// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningCompute.V20170601Preview
{
    /// <summary>
    /// Instance of an Azure ML Operationalization Cluster resource.
    /// </summary>
    public partial class OperationalizationCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// AppInsights configuration
        /// </summary>
        [Output("appInsights")]
        public Output<Outputs.AppInsightsCredentialsResponseResult?> AppInsights { get; private set; } = null!;

        /// <summary>
        /// The cluster type.
        /// </summary>
        [Output("clusterType")]
        public Output<string> ClusterType { get; private set; } = null!;

        /// <summary>
        /// Container Registry properties.
        /// </summary>
        [Output("containerRegistry")]
        public Output<Outputs.ContainerRegistryPropertiesResponseResult?> ContainerRegistry { get; private set; } = null!;

        /// <summary>
        /// Parameters for the Azure Container Service cluster.
        /// </summary>
        [Output("containerService")]
        public Output<Outputs.AcsClusterPropertiesResponseResult> ContainerService { get; private set; } = null!;

        /// <summary>
        /// The date and time when the cluster was created.
        /// </summary>
        [Output("createdOn")]
        public Output<string> CreatedOn { get; private set; } = null!;

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Contains global configuration for the web services in the cluster.
        /// </summary>
        [Output("globalServiceConfiguration")]
        public Output<Outputs.GlobalServiceConfigurationResponseResult?> GlobalServiceConfiguration { get; private set; } = null!;

        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The date and time when the cluster was last modified.
        /// </summary>
        [Output("modifiedOn")]
        public Output<string> ModifiedOn { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provision state of the cluster. Valid values are Unknown, Updating, Provisioning, Succeeded, and Failed.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Storage Account properties.
        /// </summary>
        [Output("storageAccount")]
        public Output<Outputs.StorageAccountPropertiesResponseResult?> StorageAccount { get; private set; } = null!;

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
        /// Create a OperationalizationCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OperationalizationCluster(string name, OperationalizationClusterArgs args, CustomResourceOptions? options = null)
            : base("azurerm:machinelearningcompute/v20170601preview:OperationalizationCluster", name, args ?? new OperationalizationClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OperationalizationCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:machinelearningcompute/v20170601preview:OperationalizationCluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:machinelearningcompute/v20170801preview:OperationalizationCluster"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OperationalizationCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OperationalizationCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OperationalizationCluster(name, id, options);
        }
    }

    public sealed class OperationalizationClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// AppInsights configuration
        /// </summary>
        [Input("appInsights")]
        public Input<Inputs.AppInsightsCredentialsArgs>? AppInsights { get; set; }

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The cluster type.
        /// </summary>
        [Input("clusterType", required: true)]
        public Input<string> ClusterType { get; set; } = null!;

        /// <summary>
        /// Container Registry properties.
        /// </summary>
        [Input("containerRegistry")]
        public Input<Inputs.ContainerRegistryPropertiesArgs>? ContainerRegistry { get; set; }

        /// <summary>
        /// Parameters for the Azure Container Service cluster.
        /// </summary>
        [Input("containerService", required: true)]
        public Input<Inputs.AcsClusterPropertiesArgs> ContainerService { get; set; } = null!;

        /// <summary>
        /// The description of the cluster.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Contains global configuration for the web services in the cluster.
        /// </summary>
        [Input("globalServiceConfiguration")]
        public Input<Inputs.GlobalServiceConfigurationArgs>? GlobalServiceConfiguration { get; set; }

        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Name of the resource group in which the cluster is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Storage Account properties.
        /// </summary>
        [Input("storageAccount")]
        public Input<Inputs.StorageAccountPropertiesArgs>? StorageAccount { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Contains resource tags defined as key/value pairs.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public OperationalizationClusterArgs()
        {
        }
    }
}
