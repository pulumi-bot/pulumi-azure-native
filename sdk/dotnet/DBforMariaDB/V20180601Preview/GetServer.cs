// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforMariaDB.V20180601Preview
{
    public static class GetServer
    {
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("azurerm:dbformariadb/v20180601preview:getServer", args ?? new GetServerArgs(), options.WithVersion());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetServerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        /// </summary>
        public readonly string? AdministratorLogin;
        /// <summary>
        /// Earliest restore point creation time (ISO8601 format)
        /// </summary>
        public readonly string? EarliestRestoreDate;
        /// <summary>
        /// The fully qualified domain name of a server.
        /// </summary>
        public readonly string? FullyQualifiedDomainName;
        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        public readonly Outputs.ResourceIdentityResponseResult? Identity;
        /// <summary>
        /// The location the resource resides in.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The master server id of a replica server.
        /// </summary>
        public readonly string? MasterServerId;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The maximum number of replicas that a master server can have.
        /// </summary>
        public readonly int? ReplicaCapacity;
        /// <summary>
        /// The replication role of the server.
        /// </summary>
        public readonly string? ReplicationRole;
        /// <summary>
        /// The SKU (pricing tier) of the server.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// Enable ssl enforcement or not when connect to server.
        /// </summary>
        public readonly string? SslEnforcement;
        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        public readonly Outputs.StorageProfileResponseResult? StorageProfile;
        /// <summary>
        /// Application-specific metadata in the form of key-value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A state of a server that is visible to user.
        /// </summary>
        public readonly string? UserVisibleState;
        /// <summary>
        /// Server version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetServerResult(
            string? administratorLogin,

            string? earliestRestoreDate,

            string? fullyQualifiedDomainName,

            Outputs.ResourceIdentityResponseResult? identity,

            string location,

            string? masterServerId,

            string name,

            int? replicaCapacity,

            string? replicationRole,

            Outputs.SkuResponseResult? sku,

            string? sslEnforcement,

            Outputs.StorageProfileResponseResult? storageProfile,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? userVisibleState,

            string? version)
        {
            AdministratorLogin = administratorLogin;
            EarliestRestoreDate = earliestRestoreDate;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            Identity = identity;
            Location = location;
            MasterServerId = masterServerId;
            Name = name;
            ReplicaCapacity = replicaCapacity;
            ReplicationRole = replicationRole;
            Sku = sku;
            SslEnforcement = sslEnforcement;
            StorageProfile = storageProfile;
            Tags = tags;
            Type = type;
            UserVisibleState = userVisibleState;
            Version = version;
        }
    }
}
