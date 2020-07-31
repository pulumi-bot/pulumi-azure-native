// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Inputs
{

    /// <summary>
    /// Parameters for VpnSite.
    /// </summary>
    public sealed class VpnSiteLinkPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The set of bgp properties.
        /// </summary>
        [Input("bgpProperties")]
        public Input<Inputs.VpnLinkBgpSettingsArgs>? BgpProperties { get; set; }

        /// <summary>
        /// The ip-address for the vpn-site-link.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The link provider properties.
        /// </summary>
        [Input("linkProperties")]
        public Input<Inputs.VpnLinkProviderPropertiesArgs>? LinkProperties { get; set; }

        /// <summary>
        /// The provisioning state of the VPN site link resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public VpnSiteLinkPropertiesArgs()
        {
        }
    }
}