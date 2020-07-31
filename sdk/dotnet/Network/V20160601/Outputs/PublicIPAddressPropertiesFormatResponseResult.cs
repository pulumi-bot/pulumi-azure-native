// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601.Outputs
{

    [OutputType]
    public sealed class PublicIPAddressPropertiesFormatResponseResult
    {
        /// <summary>
        /// Gets or sets FQDN of the DNS record associated with the public IP address
        /// </summary>
        public readonly Outputs.PublicIPAddressDnsSettingsResponseResult? DnsSettings;
        /// <summary>
        /// Gets or sets the idle timeout of the public IP address
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        public readonly string? IpAddress;
        /// <summary>
        /// IPConfiguration
        /// </summary>
        public readonly Outputs.IPConfigurationResponseResult IpConfiguration;
        /// <summary>
        /// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets PublicIP address version (IPv4/IPv6)
        /// </summary>
        public readonly string? PublicIPAddressVersion;
        /// <summary>
        /// Gets or sets PublicIP allocation method (Static/Dynamic)
        /// </summary>
        public readonly string? PublicIPAllocationMethod;
        /// <summary>
        /// Gets or sets resource guid property of the PublicIP resource
        /// </summary>
        public readonly string? ResourceGuid;

        [OutputConstructor]
        private PublicIPAddressPropertiesFormatResponseResult(
            Outputs.PublicIPAddressDnsSettingsResponseResult? dnsSettings,

            int? idleTimeoutInMinutes,

            string? ipAddress,

            Outputs.IPConfigurationResponseResult ipConfiguration,

            string? provisioningState,

            string? publicIPAddressVersion,

            string? publicIPAllocationMethod,

            string? resourceGuid)
        {
            DnsSettings = dnsSettings;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            IpAddress = ipAddress;
            IpConfiguration = ipConfiguration;
            ProvisioningState = provisioningState;
            PublicIPAddressVersion = publicIPAddressVersion;
            PublicIPAllocationMethod = publicIPAllocationMethod;
            ResourceGuid = resourceGuid;
        }
    }
}