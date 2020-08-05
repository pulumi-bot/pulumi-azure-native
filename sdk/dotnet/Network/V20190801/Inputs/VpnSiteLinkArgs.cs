// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190801.Inputs
{

    /// <summary>
    /// VpnSiteLink Resource.
    /// </summary>
    public sealed class VpnSiteLinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The set of bgp properties.
        /// </summary>
        [Input("bgpProperties")]
        public Input<Inputs.VpnLinkBgpSettingsArgs>? BgpProperties { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

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
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The provisioning state of the VPN site link resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public VpnSiteLinkArgs()
        {
        }
    }
}
