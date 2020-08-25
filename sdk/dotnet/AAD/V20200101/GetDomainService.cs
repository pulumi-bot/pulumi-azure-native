// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Aad.V20200101
{
    public static class GetDomainService
    {
        public static Task<GetDomainServiceResult> InvokeAsync(GetDomainServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainServiceResult>("azurerm:aad/v20200101:getDomainService", args ?? new GetDomainServiceArgs(), options.WithVersion());
    }


    public sealed class GetDomainServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDomainServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainServiceResult
    {
        /// <summary>
        /// Deployment Id
        /// </summary>
        public readonly string DeploymentId;
        /// <summary>
        /// The name of the Azure domain that the user would like to deploy Domain Services to.
        /// </summary>
        public readonly string? DomainName;
        /// <summary>
        /// DomainSecurity Settings
        /// </summary>
        public readonly Outputs.DomainSecuritySettingsResponseResult? DomainSecuritySettings;
        /// <summary>
        /// Resource etag
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Enabled or Disabled flag to turn on Group-based filtered sync
        /// </summary>
        public readonly string? FilteredSync;
        /// <summary>
        /// Secure LDAP Settings
        /// </summary>
        public readonly Outputs.LdapsSettingsResponseResult? LdapsSettings;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Notification Settings
        /// </summary>
        public readonly Outputs.NotificationSettingsResponseResult? NotificationSettings;
        /// <summary>
        /// the current deployment or provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// List of ReplicaSets
        /// </summary>
        public readonly ImmutableArray<Outputs.ReplicaSetResponseResult> ReplicaSets;
        /// <summary>
        /// SyncOwner ReplicaSet Id
        /// </summary>
        public readonly string SyncOwner;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure Active Directory Tenant Id
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Data Model Version
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GetDomainServiceResult(
            string deploymentId,

            string? domainName,

            Outputs.DomainSecuritySettingsResponseResult? domainSecuritySettings,

            string? etag,

            string? filteredSync,

            Outputs.LdapsSettingsResponseResult? ldapsSettings,

            string? location,

            string name,

            Outputs.NotificationSettingsResponseResult? notificationSettings,

            string provisioningState,

            ImmutableArray<Outputs.ReplicaSetResponseResult> replicaSets,

            string syncOwner,

            ImmutableDictionary<string, string>? tags,

            string tenantId,

            string type,

            int version)
        {
            DeploymentId = deploymentId;
            DomainName = domainName;
            DomainSecuritySettings = domainSecuritySettings;
            Etag = etag;
            FilteredSync = filteredSync;
            LdapsSettings = ldapsSettings;
            Location = location;
            Name = name;
            NotificationSettings = notificationSettings;
            ProvisioningState = provisioningState;
            ReplicaSets = replicaSets;
            SyncOwner = syncOwner;
            Tags = tags;
            TenantId = tenantId;
            Type = type;
            Version = version;
        }
    }
}
