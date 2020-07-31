// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Outputs
{

    [OutputType]
    public sealed class VpnSitePropertiesResponseResult
    {
        /// <summary>
        /// The AddressSpace that contains an array of IP address ranges.
        /// </summary>
        public readonly Outputs.AddressSpaceResponseResult? AddressSpace;
        /// <summary>
        /// The set of bgp properties.
        /// </summary>
        public readonly Outputs.BgpSettingsResponseResult? BgpProperties;
        /// <summary>
        /// The device properties.
        /// </summary>
        public readonly Outputs.DevicePropertiesResponseResult? DeviceProperties;
        /// <summary>
        /// The ip-address for the vpn-site.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// IsSecuritySite flag.
        /// </summary>
        public readonly bool? IsSecuritySite;
        /// <summary>
        /// The provisioning state of the VPN site resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The key for vpn-site that can be used for connections.
        /// </summary>
        public readonly string? SiteKey;
        /// <summary>
        /// The VirtualWAN to which the vpnSite belongs.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? VirtualWan;
        /// <summary>
        /// List of all vpn site links.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpnSiteLinkResponseResult> VpnSiteLinks;

        [OutputConstructor]
        private VpnSitePropertiesResponseResult(
            Outputs.AddressSpaceResponseResult? addressSpace,

            Outputs.BgpSettingsResponseResult? bgpProperties,

            Outputs.DevicePropertiesResponseResult? deviceProperties,

            string? ipAddress,

            bool? isSecuritySite,

            string? provisioningState,

            string? siteKey,

            Outputs.SubResourceResponseResult? virtualWan,

            ImmutableArray<Outputs.VpnSiteLinkResponseResult> vpnSiteLinks)
        {
            AddressSpace = addressSpace;
            BgpProperties = bgpProperties;
            DeviceProperties = deviceProperties;
            IpAddress = ipAddress;
            IsSecuritySite = isSecuritySite;
            ProvisioningState = provisioningState;
            SiteKey = siteKey;
            VirtualWan = virtualWan;
            VpnSiteLinks = vpnSiteLinks;
        }
    }
}