// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181101.Outputs
{

    [OutputType]
    public sealed class PublicIPAddressPropertiesFormatResponseResult
    {
        /// <summary>
        /// The DDoS protection custom policy associated with the public IP address.
        /// </summary>
        public readonly Outputs.DdosSettingsResponseResult? DdosSettings;
        /// <summary>
        /// The FQDN of the DNS record associated with the public IP address.
        /// </summary>
        public readonly Outputs.PublicIPAddressDnsSettingsResponseResult? DnsSettings;
        /// <summary>
        /// The idle timeout of the public IP address.
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        /// <summary>
        /// The IP address associated with the public IP address resource.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// The IP configuration associated with the public IP address.
        /// </summary>
        public readonly Outputs.IPConfigurationResponseResult IpConfiguration;
        /// <summary>
        /// The list of tags associated with the public IP address.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpTagResponseResult> IpTags;
        /// <summary>
        /// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
        /// </summary>
        public readonly string? PublicIPAddressVersion;
        /// <summary>
        /// The public IP allocation method. Possible values are: 'Static' and 'Dynamic'.
        /// </summary>
        public readonly string? PublicIPAllocationMethod;
        /// <summary>
        /// The Public IP Prefix this Public IP Address should be allocated from.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? PublicIPPrefix;
        /// <summary>
        /// The resource GUID property of the public IP resource.
        /// </summary>
        public readonly string? ResourceGuid;

        [OutputConstructor]
        private PublicIPAddressPropertiesFormatResponseResult(
            Outputs.DdosSettingsResponseResult? ddosSettings,

            Outputs.PublicIPAddressDnsSettingsResponseResult? dnsSettings,

            int? idleTimeoutInMinutes,

            string? ipAddress,

            Outputs.IPConfigurationResponseResult ipConfiguration,

            ImmutableArray<Outputs.IpTagResponseResult> ipTags,

            string? provisioningState,

            string? publicIPAddressVersion,

            string? publicIPAllocationMethod,

            Outputs.SubResourceResponseResult? publicIPPrefix,

            string? resourceGuid)
        {
            DdosSettings = ddosSettings;
            DnsSettings = dnsSettings;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            IpAddress = ipAddress;
            IpConfiguration = ipConfiguration;
            IpTags = ipTags;
            ProvisioningState = provisioningState;
            PublicIPAddressVersion = publicIPAddressVersion;
            PublicIPAllocationMethod = publicIPAllocationMethod;
            PublicIPPrefix = publicIPPrefix;
            ResourceGuid = resourceGuid;
        }
    }
}