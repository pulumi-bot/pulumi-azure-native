// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SqlVirtualMachine.V20170301Preview
{
    public static class GetSqlVirtualMachine
    {
        public static Task<GetSqlVirtualMachineResult> InvokeAsync(GetSqlVirtualMachineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSqlVirtualMachineResult>("azurerm:sqlvirtualmachine/v20170301preview:getSqlVirtualMachine", args ?? new GetSqlVirtualMachineArgs(), options.WithVersion());
    }


    public sealed class GetSqlVirtualMachineArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The child resources to include in the response.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the SQL virtual machine.
        /// </summary>
        [Input("sqlVirtualMachineName", required: true)]
        public string SqlVirtualMachineName { get; set; } = null!;

        public GetSqlVirtualMachineArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSqlVirtualMachineResult
    {
        /// <summary>
        /// Auto backup settings for SQL Server.
        /// </summary>
        public readonly Outputs.AutoBackupSettingsResponseResult? AutoBackupSettings;
        /// <summary>
        /// Auto patching settings for applying critical security updates to SQL virtual machine.
        /// </summary>
        public readonly Outputs.AutoPatchingSettingsResponseResult? AutoPatchingSettings;
        /// <summary>
        /// Azure Active Directory identity of the server.
        /// </summary>
        public readonly Outputs.ResourceIdentityResponseResult? Identity;
        /// <summary>
        /// Key vault credential settings.
        /// </summary>
        public readonly Outputs.KeyVaultCredentialSettingsResponseResult? KeyVaultCredentialSettings;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state to track the async operation status.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// SQL Server configuration management settings.
        /// </summary>
        public readonly Outputs.ServerConfigurationsManagementSettingsResponseResult? ServerConfigurationsManagementSettings;
        /// <summary>
        /// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
        /// </summary>
        public readonly string? SqlImageOffer;
        /// <summary>
        /// SQL Server edition type.
        /// </summary>
        public readonly string? SqlImageSku;
        /// <summary>
        /// SQL Server Management type.
        /// </summary>
        public readonly string? SqlManagement;
        /// <summary>
        /// SQL Server license type.
        /// </summary>
        public readonly string? SqlServerLicenseType;
        /// <summary>
        /// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
        /// </summary>
        public readonly string? SqlVirtualMachineGroupResourceId;
        /// <summary>
        /// Storage Configuration Settings.
        /// </summary>
        public readonly Outputs.StorageConfigurationSettingsResponseResult? StorageConfigurationSettings;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// ARM Resource id of underlying virtual machine created from SQL marketplace image.
        /// </summary>
        public readonly string? VirtualMachineResourceId;
        /// <summary>
        /// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
        /// </summary>
        public readonly Outputs.WsfcDomainCredentialsResponseResult? WsfcDomainCredentials;

        [OutputConstructor]
        private GetSqlVirtualMachineResult(
            Outputs.AutoBackupSettingsResponseResult? autoBackupSettings,

            Outputs.AutoPatchingSettingsResponseResult? autoPatchingSettings,

            Outputs.ResourceIdentityResponseResult? identity,

            Outputs.KeyVaultCredentialSettingsResponseResult? keyVaultCredentialSettings,

            string location,

            string name,

            string provisioningState,

            Outputs.ServerConfigurationsManagementSettingsResponseResult? serverConfigurationsManagementSettings,

            string? sqlImageOffer,

            string? sqlImageSku,

            string? sqlManagement,

            string? sqlServerLicenseType,

            string? sqlVirtualMachineGroupResourceId,

            Outputs.StorageConfigurationSettingsResponseResult? storageConfigurationSettings,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? virtualMachineResourceId,

            Outputs.WsfcDomainCredentialsResponseResult? wsfcDomainCredentials)
        {
            AutoBackupSettings = autoBackupSettings;
            AutoPatchingSettings = autoPatchingSettings;
            Identity = identity;
            KeyVaultCredentialSettings = keyVaultCredentialSettings;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ServerConfigurationsManagementSettings = serverConfigurationsManagementSettings;
            SqlImageOffer = sqlImageOffer;
            SqlImageSku = sqlImageSku;
            SqlManagement = sqlManagement;
            SqlServerLicenseType = sqlServerLicenseType;
            SqlVirtualMachineGroupResourceId = sqlVirtualMachineGroupResourceId;
            StorageConfigurationSettings = storageConfigurationSettings;
            Tags = tags;
            Type = type;
            VirtualMachineResourceId = virtualMachineResourceId;
            WsfcDomainCredentials = wsfcDomainCredentials;
        }
    }
}
