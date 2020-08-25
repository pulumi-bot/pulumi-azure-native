// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NetApp.V20200601.Inputs
{

    /// <summary>
    /// Volume Backup Properties
    /// </summary>
    public sealed class VolumeBackupPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backup Enabled
        /// </summary>
        [Input("backupEnabled")]
        public Input<bool>? BackupEnabled { get; set; }

        /// <summary>
        /// Backup Policy Resource ID
        /// </summary>
        [Input("backupPolicyId")]
        public Input<string>? BackupPolicyId { get; set; }

        /// <summary>
        /// Policy Enforced
        /// </summary>
        [Input("policyEnforced")]
        public Input<bool>? PolicyEnforced { get; set; }

        /// <summary>
        /// Vault Resource ID
        /// </summary>
        [Input("vaultId")]
        public Input<string>? VaultId { get; set; }

        public VolumeBackupPropertiesArgs()
        {
        }
    }
}
