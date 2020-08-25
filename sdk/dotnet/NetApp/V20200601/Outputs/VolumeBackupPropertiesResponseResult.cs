// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NetApp.V20200601.Outputs
{

    [OutputType]
    public sealed class VolumeBackupPropertiesResponseResult
    {
        /// <summary>
        /// Backup Enabled
        /// </summary>
        public readonly bool? BackupEnabled;
        /// <summary>
        /// Backup Policy Resource ID
        /// </summary>
        public readonly string? BackupPolicyId;
        /// <summary>
        /// Policy Enforced
        /// </summary>
        public readonly bool? PolicyEnforced;
        /// <summary>
        /// Vault Resource ID
        /// </summary>
        public readonly string? VaultId;

        [OutputConstructor]
        private VolumeBackupPropertiesResponseResult(
            bool? backupEnabled,

            string? backupPolicyId,

            bool? policyEnforced,

            string? vaultId)
        {
            BackupEnabled = backupEnabled;
            BackupPolicyId = backupPolicyId;
            PolicyEnforced = policyEnforced;
            VaultId = vaultId;
        }
    }
}
