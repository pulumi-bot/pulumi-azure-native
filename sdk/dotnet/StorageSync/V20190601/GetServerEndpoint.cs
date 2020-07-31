// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageSync.V20190601
{
    public static class GetServerEndpoint
    {
        public static Task<GetServerEndpointResult> InvokeAsync(GetServerEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerEndpointResult>("azurerm:storagesync/v20190601:getServerEndpoint", args ?? new GetServerEndpointArgs(), options.WithVersion());
    }


    public sealed class GetServerEndpointArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Server Endpoint object.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Storage Sync Service resource.
        /// </summary>
        [Input("storageSyncServiceName", required: true)]
        public string StorageSyncServiceName { get; set; } = null!;

        /// <summary>
        /// Name of Sync Group resource.
        /// </summary>
        [Input("syncGroupName", required: true)]
        public string SyncGroupName { get; set; } = null!;

        public GetServerEndpointArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerEndpointResult
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Server Endpoint properties.
        /// </summary>
        public readonly Outputs.ServerEndpointPropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServerEndpointResult(
            string name,

            Outputs.ServerEndpointPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}