// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBForMySql.V20200701privatePreview
{
    public static class GetServer
    {
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("azurerm:dbformysql/v20200701privatepreview:getServer", args ?? new GetServerArgs(), options.WithVersion());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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
        /// The password of the administrator login (required for server creation).
        /// </summary>
        public readonly string? AdministratorLoginPassword;
        /// <summary>
        /// The mode to create a new MySQL server.
        /// </summary>
        public readonly string? CreateMode;
        /// <summary>
        /// Earliest restore point creation time (ISO8601 format)
        /// </summary>
        public readonly string? EarliestRestoreDate;
        /// <summary>
        /// The fully qualified domain name of a server.
        /// </summary>
        public readonly string FullyQualifiedDomainName;
        /// <summary>
        /// The state of a HA server.
        /// </summary>
        public readonly string HaState;
        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        public readonly Outputs.IdentityResponseResult? Identity;
        /// <summary>
        /// Status showing whether the server enabled infrastructure encryption.
        /// </summary>
        public readonly string? InfrastructureEncryption;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The primary server id of a replica server.
        /// </summary>
        public readonly string PrimaryServerId;
        /// <summary>
        /// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The maximum number of replicas that a primary server can have.
        /// </summary>
        public readonly int? ReplicaCapacity;
        /// <summary>
        /// The replication role.
        /// </summary>
        public readonly string? ReplicationRole;
        /// <summary>
        /// Restore point creation time (ISO8601 format), specifying the time to restore from.
        /// </summary>
        public readonly string? RestorePointInTime;
        /// <summary>
        /// The SKU (pricing tier) of the server.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// The source MySQL server name to restore from.
        /// </summary>
        public readonly string? SourceServerId;
        /// <summary>
        /// Enable ssl enforcement or not when connect to server.
        /// </summary>
        public readonly string? SslEnforcement;
        /// <summary>
        /// stand by count value can be either 0 or 1
        /// </summary>
        public readonly int? StandbyCount;
        /// <summary>
        /// The state of a server.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        public readonly Outputs.StorageProfileResponseResult? StorageProfile;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Server version.
        /// </summary>
        public readonly string? Version;
        /// <summary>
        /// Vnet arguments.
        /// </summary>
        public readonly Outputs.VnetInjArgsResponseResult? VnetInjArgs;

        [OutputConstructor]
        private GetServerResult(
            string? administratorLogin,

            string? administratorLoginPassword,

            string? createMode,

            string? earliestRestoreDate,

            string fullyQualifiedDomainName,

            string haState,

            Outputs.IdentityResponseResult? identity,

            string? infrastructureEncryption,

            string location,

            string name,

            string primaryServerId,

            string? publicNetworkAccess,

            int? replicaCapacity,

            string? replicationRole,

            string? restorePointInTime,

            Outputs.SkuResponseResult? sku,

            string? sourceServerId,

            string? sslEnforcement,

            int? standbyCount,

            string state,

            Outputs.StorageProfileResponseResult? storageProfile,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? version,

            Outputs.VnetInjArgsResponseResult? vnetInjArgs)
        {
            AdministratorLogin = administratorLogin;
            AdministratorLoginPassword = administratorLoginPassword;
            CreateMode = createMode;
            EarliestRestoreDate = earliestRestoreDate;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            HaState = haState;
            Identity = identity;
            InfrastructureEncryption = infrastructureEncryption;
            Location = location;
            Name = name;
            PrimaryServerId = primaryServerId;
            PublicNetworkAccess = publicNetworkAccess;
            ReplicaCapacity = replicaCapacity;
            ReplicationRole = replicationRole;
            RestorePointInTime = restorePointInTime;
            Sku = sku;
            SourceServerId = sourceServerId;
            SslEnforcement = sslEnforcement;
            StandbyCount = standbyCount;
            State = state;
            StorageProfile = storageProfile;
            Tags = tags;
            Type = type;
            Version = version;
            VnetInjArgs = vnetInjArgs;
        }
    }
}
