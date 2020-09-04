// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningCompute.V20170801Preview
{
    public static class ListOperationalizationClusterKeys
    {
        public static Task<ListOperationalizationClusterKeysResult> InvokeAsync(ListOperationalizationClusterKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListOperationalizationClusterKeysResult>("azurerm:machinelearningcompute/v20170801preview:listOperationalizationClusterKeys", args ?? new ListOperationalizationClusterKeysArgs(), options.WithVersion());
    }


    public sealed class ListOperationalizationClusterKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group in which the cluster is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListOperationalizationClusterKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListOperationalizationClusterKeysResult
    {
        /// <summary>
        /// Credentials for Azure AppInsights.
        /// </summary>
        public readonly Outputs.AppInsightsCredentialsResponseResult? AppInsights;
        /// <summary>
        /// Credentials for Azure Container Registry.
        /// </summary>
        public readonly Outputs.ContainerRegistryCredentialsResponseResult? ContainerRegistry;
        /// <summary>
        /// Credentials for Azure Container Service.
        /// </summary>
        public readonly Outputs.ContainerServiceCredentialsResponseResult? ContainerService;
        /// <summary>
        /// Global authorization keys for all user services deployed in cluster. These are used if the service does not have auth keys.
        /// </summary>
        public readonly Outputs.ServiceAuthConfigurationResponseResult? ServiceAuthConfiguration;
        /// <summary>
        /// The SSL configuration for the services.
        /// </summary>
        public readonly Outputs.SslConfigurationResponseResult? SslConfiguration;
        /// <summary>
        /// Credentials for the Storage Account.
        /// </summary>
        public readonly Outputs.StorageAccountCredentialsResponseResult? StorageAccount;

        [OutputConstructor]
        private ListOperationalizationClusterKeysResult(
            Outputs.AppInsightsCredentialsResponseResult? appInsights,

            Outputs.ContainerRegistryCredentialsResponseResult? containerRegistry,

            Outputs.ContainerServiceCredentialsResponseResult? containerService,

            Outputs.ServiceAuthConfigurationResponseResult? serviceAuthConfiguration,

            Outputs.SslConfigurationResponseResult? sslConfiguration,

            Outputs.StorageAccountCredentialsResponseResult? storageAccount)
        {
            AppInsights = appInsights;
            ContainerRegistry = containerRegistry;
            ContainerService = containerService;
            ServiceAuthConfiguration = serviceAuthConfiguration;
            SslConfiguration = sslConfiguration;
            StorageAccount = storageAccount;
        }
    }
}
