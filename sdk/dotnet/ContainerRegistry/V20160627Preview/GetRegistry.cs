// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20160627Preview
{
    public static class GetRegistry
    {
        public static Task<GetRegistryResult> InvokeAsync(GetRegistryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryResult>("azurerm:containerregistry/v20160627preview:getRegistry", args ?? new GetRegistryArgs(), options.WithVersion());
    }


    public sealed class GetRegistryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRegistryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryResult
    {
        /// <summary>
        /// The value that indicates whether the admin user is enabled. This value is false by default.
        /// </summary>
        public readonly bool? AdminUserEnabled;
        /// <summary>
        /// The creation date of the container registry in ISO8601 format.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The URL that can be used to log into the container registry.
        /// </summary>
        public readonly string LoginServer;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of the storage account for the container registry. If specified, the storage account must be in the same physical location as the container registry.
        /// </summary>
        public readonly Outputs.StorageAccountPropertiesResponseResult StorageAccount;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRegistryResult(
            bool? adminUserEnabled,

            string creationDate,

            string location,

            string loginServer,

            string name,

            Outputs.StorageAccountPropertiesResponseResult storageAccount,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AdminUserEnabled = adminUserEnabled;
            CreationDate = creationDate;
            Location = location;
            LoginServer = loginServer;
            Name = name;
            StorageAccount = storageAccount;
            Tags = tags;
            Type = type;
        }
    }
}
