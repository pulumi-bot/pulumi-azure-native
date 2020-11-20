// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20201001Preview.Outputs
{

    [OutputType]
    public sealed class CloudServicePropertiesResponse
    {
        /// <summary>
        /// Specifies the XML service configuration (.cscfg) for the cloud service.
        /// </summary>
        public readonly string? Configuration;
        /// <summary>
        /// Specifies a URL that refers to the location of the service configuration in the Blob service. The service package URL  can be Shared Access Signature (SAS) URI from any storage account.
        /// This is a write-only property and is not returned in GET calls.
        /// </summary>
        public readonly string? ConfigurationUrl;
        /// <summary>
        /// Describes a cloud service extension profile.
        /// </summary>
        public readonly Outputs.CloudServiceExtensionProfileResponse? ExtensionProfile;
        /// <summary>
        /// Network Profile for the cloud service.
        /// </summary>
        public readonly Outputs.CloudServiceNetworkProfileResponse? NetworkProfile;
        /// <summary>
        /// Describes the OS profile for the cloud service.
        /// </summary>
        public readonly Outputs.CloudServiceOsProfileResponse? OsProfile;
        /// <summary>
        /// Specifies a URL that refers to the location of the service package in the Blob service. The service package URL can be Shared Access Signature (SAS) URI from any storage account.
        /// This is a write-only property and is not returned in GET calls.
        /// </summary>
        public readonly string? PackageUrl;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Describes the role profile for the cloud service.
        /// </summary>
        public readonly Outputs.CloudServiceRoleProfileResponse? RoleProfile;
        /// <summary>
        /// (Optional) Indicates whether to start the cloud service immediately after it is created. The default value is `true`.
        /// If false, the service model is still deployed, but the code is not run immediately. Instead, the service is PoweredOff until you call Start, at which time the service will be started. A deployed service still incurs charges, even if it is poweredoff.
        /// </summary>
        public readonly bool? StartCloudService;
        /// <summary>
        /// The unique identifier for the cloud service.
        /// </summary>
        public readonly string UniqueId;
        /// <summary>
        /// Update mode for the cloud service. Role instances are allocated to update domains when the service is deployed. Updates can be initiated manually in each update domain or initiated automatically in all update domains.
        /// Possible Values are &lt;br /&gt;&lt;br /&gt;**Auto**&lt;br /&gt;&lt;br /&gt;**Manual** &lt;br /&gt;&lt;br /&gt;**Simultaneous**&lt;br /&gt;&lt;br /&gt;
        /// If not specified, the default value is Auto. If set to Manual, PUT UpdateDomain must be called to apply the update. If set to Auto, the update is automatically applied to each update domain in sequence.
        /// </summary>
        public readonly string? UpgradeMode;

        [OutputConstructor]
        private CloudServicePropertiesResponse(
            string? configuration,

            string? configurationUrl,

            Outputs.CloudServiceExtensionProfileResponse? extensionProfile,

            Outputs.CloudServiceNetworkProfileResponse? networkProfile,

            Outputs.CloudServiceOsProfileResponse? osProfile,

            string? packageUrl,

            string provisioningState,

            Outputs.CloudServiceRoleProfileResponse? roleProfile,

            bool? startCloudService,

            string uniqueId,

            string? upgradeMode)
        {
            Configuration = configuration;
            ConfigurationUrl = configurationUrl;
            ExtensionProfile = extensionProfile;
            NetworkProfile = networkProfile;
            OsProfile = osProfile;
            PackageUrl = packageUrl;
            ProvisioningState = provisioningState;
            RoleProfile = roleProfile;
            StartCloudService = startCloudService;
            UniqueId = uniqueId;
            UpgradeMode = upgradeMode;
        }
    }
}
