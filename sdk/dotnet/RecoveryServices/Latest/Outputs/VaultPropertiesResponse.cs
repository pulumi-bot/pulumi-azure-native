// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.Latest.Outputs
{

    [OutputType]
    public sealed class VaultPropertiesResponse
    {
        /// <summary>
        /// Customer Managed Key details of the resource.
        /// </summary>
        public readonly Outputs.VaultPropertiesResponseEncryption? Encryption;
        /// <summary>
        /// List of private endpoint connection.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionVaultPropertiesResponse> PrivateEndpointConnections;
        /// <summary>
        /// Private endpoint state for backup.
        /// </summary>
        public readonly string PrivateEndpointStateForBackup;
        /// <summary>
        /// Private endpoint state for site recovery.
        /// </summary>
        public readonly string PrivateEndpointStateForSiteRecovery;
        /// <summary>
        /// Provisioning State.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Details for upgrading vault.
        /// </summary>
        public readonly Outputs.UpgradeDetailsResponse? UpgradeDetails;

        [OutputConstructor]
        private VaultPropertiesResponse(
            Outputs.VaultPropertiesResponseEncryption? encryption,

            ImmutableArray<Outputs.PrivateEndpointConnectionVaultPropertiesResponse> privateEndpointConnections,

            string privateEndpointStateForBackup,

            string privateEndpointStateForSiteRecovery,

            string provisioningState,

            Outputs.UpgradeDetailsResponse? upgradeDetails)
        {
            Encryption = encryption;
            PrivateEndpointConnections = privateEndpointConnections;
            PrivateEndpointStateForBackup = privateEndpointStateForBackup;
            PrivateEndpointStateForSiteRecovery = privateEndpointStateForSiteRecovery;
            ProvisioningState = provisioningState;
            UpgradeDetails = upgradeDetails;
        }
    }
}
