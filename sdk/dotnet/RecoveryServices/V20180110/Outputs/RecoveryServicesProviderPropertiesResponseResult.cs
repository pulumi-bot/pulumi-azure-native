// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20180110.Outputs
{

    [OutputType]
    public sealed class RecoveryServicesProviderPropertiesResponseResult
    {
        /// <summary>
        /// The scenarios allowed on this provider.
        /// </summary>
        public readonly ImmutableArray<string> AllowedScenarios;
        /// <summary>
        /// The authentication identity details.
        /// </summary>
        public readonly Outputs.IdentityProviderDetailsResponseResult? AuthenticationIdentityDetails;
        /// <summary>
        /// A value indicating whether DRA is responsive.
        /// </summary>
        public readonly string? ConnectionStatus;
        /// <summary>
        /// The DRA Id.
        /// </summary>
        public readonly string? DraIdentifier;
        /// <summary>
        /// The fabric friendly name.
        /// </summary>
        public readonly string? FabricFriendlyName;
        /// <summary>
        /// Type of the site.
        /// </summary>
        public readonly string? FabricType;
        /// <summary>
        /// Friendly name of the DRA.
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// The recovery services provider health error details.
        /// </summary>
        public readonly ImmutableArray<Outputs.HealthErrorResponseResult> HealthErrorDetails;
        /// <summary>
        /// Time when last heartbeat was sent by the DRA.
        /// </summary>
        public readonly string? LastHeartBeat;
        /// <summary>
        /// Number of protected VMs currently managed by the DRA.
        /// </summary>
        public readonly int? ProtectedItemCount;
        /// <summary>
        /// The provider version.
        /// </summary>
        public readonly string? ProviderVersion;
        /// <summary>
        /// The provider version details.
        /// </summary>
        public readonly Outputs.VersionDetailsResponseResult? ProviderVersionDetails;
        /// <summary>
        /// Expiry date of the version.
        /// </summary>
        public readonly string? ProviderVersionExpiryDate;
        /// <summary>
        /// DRA version status.
        /// </summary>
        public readonly string? ProviderVersionState;
        /// <summary>
        /// The resource access identity details.
        /// </summary>
        public readonly Outputs.IdentityProviderDetailsResponseResult? ResourceAccessIdentityDetails;
        /// <summary>
        /// The fabric provider.
        /// </summary>
        public readonly string? ServerVersion;

        [OutputConstructor]
        private RecoveryServicesProviderPropertiesResponseResult(
            ImmutableArray<string> allowedScenarios,

            Outputs.IdentityProviderDetailsResponseResult? authenticationIdentityDetails,

            string? connectionStatus,

            string? draIdentifier,

            string? fabricFriendlyName,

            string? fabricType,

            string? friendlyName,

            ImmutableArray<Outputs.HealthErrorResponseResult> healthErrorDetails,

            string? lastHeartBeat,

            int? protectedItemCount,

            string? providerVersion,

            Outputs.VersionDetailsResponseResult? providerVersionDetails,

            string? providerVersionExpiryDate,

            string? providerVersionState,

            Outputs.IdentityProviderDetailsResponseResult? resourceAccessIdentityDetails,

            string? serverVersion)
        {
            AllowedScenarios = allowedScenarios;
            AuthenticationIdentityDetails = authenticationIdentityDetails;
            ConnectionStatus = connectionStatus;
            DraIdentifier = draIdentifier;
            FabricFriendlyName = fabricFriendlyName;
            FabricType = fabricType;
            FriendlyName = friendlyName;
            HealthErrorDetails = healthErrorDetails;
            LastHeartBeat = lastHeartBeat;
            ProtectedItemCount = protectedItemCount;
            ProviderVersion = providerVersion;
            ProviderVersionDetails = providerVersionDetails;
            ProviderVersionExpiryDate = providerVersionExpiryDate;
            ProviderVersionState = providerVersionState;
            ResourceAccessIdentityDetails = resourceAccessIdentityDetails;
            ServerVersion = serverVersion;
        }
    }
}