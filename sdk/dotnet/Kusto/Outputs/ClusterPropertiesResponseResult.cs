// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.Outputs
{

    [OutputType]
    public sealed class ClusterPropertiesResponseResult
    {
        /// <summary>
        /// The cluster data ingestion URI.
        /// </summary>
        public readonly string DataIngestionUri;
        /// <summary>
        /// A boolean value that indicates if the cluster's disks are encrypted.
        /// </summary>
        public readonly bool? EnableDiskEncryption;
        /// <summary>
        /// A boolean value that indicates if double encryption is enabled.
        /// </summary>
        public readonly bool? EnableDoubleEncryption;
        /// <summary>
        /// A boolean value that indicates if the purge operations are enabled.
        /// </summary>
        public readonly bool? EnablePurge;
        /// <summary>
        /// A boolean value that indicates if the streaming ingest is enabled.
        /// </summary>
        public readonly bool? EnableStreamingIngest;
        /// <summary>
        /// KeyVault properties for the cluster encryption.
        /// </summary>
        public readonly Outputs.KeyVaultPropertiesResponseResult? KeyVaultProperties;
        /// <summary>
        /// List of the cluster's language extensions.
        /// </summary>
        public readonly Outputs.LanguageExtensionsListResponseResult LanguageExtensions;
        /// <summary>
        /// Optimized auto scale definition.
        /// </summary>
        public readonly Outputs.OptimizedAutoscaleResponseResult? OptimizedAutoscale;
        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The state of the resource.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The reason for the cluster's current state.
        /// </summary>
        public readonly string StateReason;
        /// <summary>
        /// The cluster's external tenants.
        /// </summary>
        public readonly ImmutableArray<Outputs.TrustedExternalTenantResponseResult> TrustedExternalTenants;
        /// <summary>
        /// The cluster URI.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// Virtual network definition.
        /// </summary>
        public readonly Outputs.VirtualNetworkConfigurationResponseResult? VirtualNetworkConfiguration;

        [OutputConstructor]
        private ClusterPropertiesResponseResult(
            string dataIngestionUri,

            bool? enableDiskEncryption,

            bool? enableDoubleEncryption,

            bool? enablePurge,

            bool? enableStreamingIngest,

            Outputs.KeyVaultPropertiesResponseResult? keyVaultProperties,

            Outputs.LanguageExtensionsListResponseResult languageExtensions,

            Outputs.OptimizedAutoscaleResponseResult? optimizedAutoscale,

            string provisioningState,

            string state,

            string stateReason,

            ImmutableArray<Outputs.TrustedExternalTenantResponseResult> trustedExternalTenants,

            string uri,

            Outputs.VirtualNetworkConfigurationResponseResult? virtualNetworkConfiguration)
        {
            DataIngestionUri = dataIngestionUri;
            EnableDiskEncryption = enableDiskEncryption;
            EnableDoubleEncryption = enableDoubleEncryption;
            EnablePurge = enablePurge;
            EnableStreamingIngest = enableStreamingIngest;
            KeyVaultProperties = keyVaultProperties;
            LanguageExtensions = languageExtensions;
            OptimizedAutoscale = optimizedAutoscale;
            ProvisioningState = provisioningState;
            State = state;
            StateReason = stateReason;
            TrustedExternalTenants = trustedExternalTenants;
            Uri = uri;
            VirtualNetworkConfiguration = virtualNetworkConfiguration;
        }
    }
}