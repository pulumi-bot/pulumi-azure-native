// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureData
{
    public static class GetPostgresInstance
    {
        /// <summary>
        /// A Postgres Instance.
        /// API Version: 2020-09-08-preview.
        /// </summary>
        public static Task<GetPostgresInstanceResult> InvokeAsync(GetPostgresInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPostgresInstanceResult>("azure-native:azuredata:getPostgresInstance", args ?? new GetPostgresInstanceArgs(), options.WithVersion());
    }


    public sealed class GetPostgresInstanceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Postgres Instance
        /// </summary>
        [Input("postgresInstanceName", required: true)]
        public string PostgresInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPostgresInstanceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPostgresInstanceResult
    {
        /// <summary>
        /// The instance admin
        /// </summary>
        public readonly string? Admin;
        /// <summary>
        /// The data controller id
        /// </summary>
        public readonly string? DataControllerId;
        /// <summary>
        /// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The raw kubernetes information
        /// </summary>
        public readonly object? K8sRaw;
        /// <summary>
        /// Last uploaded date from on premise cluster. Defaults to current date time
        /// </summary>
        public readonly string? LastUploadedDate;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Read only system data
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPostgresInstanceResult(
            string? admin,

            string? dataControllerId,

            string id,

            object? k8sRaw,

            string? lastUploadedDate,

            string location,

            string name,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Admin = admin;
            DataControllerId = dataControllerId;
            Id = id;
            K8sRaw = k8sRaw;
            LastUploadedDate = lastUploadedDate;
            Location = location;
            Name = name;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
