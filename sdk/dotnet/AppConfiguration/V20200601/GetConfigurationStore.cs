// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppConfiguration.V20200601
{
    public static class GetConfigurationStore
    {
        public static Task<GetConfigurationStoreResult> InvokeAsync(GetConfigurationStoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationStoreResult>("azurerm:appconfiguration/v20200601:getConfigurationStore", args ?? new GetConfigurationStoreArgs(), options.WithVersion());
    }


    public sealed class GetConfigurationStoreArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the configuration store.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConfigurationStoreArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigurationStoreResult
    {
        /// <summary>
        /// The managed identity information, if configured.
        /// </summary>
        public readonly Outputs.ResourceIdentityResponseResult? Identity;
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of a configuration store.
        /// </summary>
        public readonly Outputs.ConfigurationStorePropertiesResponseResult Properties;
        /// <summary>
        /// The sku of the configuration store.
        /// </summary>
        public readonly Outputs.SkuResponseResult Sku;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConfigurationStoreResult(
            Outputs.ResourceIdentityResponseResult? identity,

            string location,

            string name,

            Outputs.ConfigurationStorePropertiesResponseResult properties,

            Outputs.SkuResponseResult sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}