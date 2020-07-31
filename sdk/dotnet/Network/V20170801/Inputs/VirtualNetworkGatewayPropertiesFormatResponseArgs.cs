// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Inputs
{

    /// <summary>
    /// VirtualNetworkGateway properties
    /// </summary>
    public sealed class VirtualNetworkGatewayPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ActiveActive flag
        /// </summary>
        [Input("activeActive")]
        public bool? ActiveActive { get; set; }

        /// <summary>
        /// Virtual network gateway's BGP speaker settings.
        /// </summary>
        [Input("bgpSettings")]
        public Inputs.BgpSettingsResponseArgs? BgpSettings { get; set; }

        /// <summary>
        /// Whether BGP is enabled for this virtual network gateway or not.
        /// </summary>
        [Input("enableBgp")]
        public bool? EnableBgp { get; set; }

        /// <summary>
        /// The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
        /// </summary>
        [Input("gatewayDefaultSite")]
        public Inputs.SubResourceResponseArgs? GatewayDefaultSite { get; set; }

        /// <summary>
        /// The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
        /// </summary>
        [Input("gatewayType")]
        public string? GatewayType { get; set; }

        [Input("ipConfigurations")]
        private List<Inputs.VirtualNetworkGatewayIPConfigurationResponseArgs>? _ipConfigurations;

        /// <summary>
        /// IP configurations for virtual network gateway.
        /// </summary>
        public List<Inputs.VirtualNetworkGatewayIPConfigurationResponseArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new List<Inputs.VirtualNetworkGatewayIPConfigurationResponseArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The provisioning state of the VirtualNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState", required: true)]
        public string ProvisioningState { get; set; } = null!;

        /// <summary>
        /// The resource GUID property of the VirtualNetworkGateway resource.
        /// </summary>
        [Input("resourceGuid")]
        public string? ResourceGuid { get; set; }

        /// <summary>
        /// The reference of the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
        /// </summary>
        [Input("sku")]
        public Inputs.VirtualNetworkGatewaySkuResponseArgs? Sku { get; set; }

        /// <summary>
        /// The reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
        /// </summary>
        [Input("vpnClientConfiguration")]
        public Inputs.VpnClientConfigurationResponseArgs? VpnClientConfiguration { get; set; }

        /// <summary>
        /// The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'.
        /// </summary>
        [Input("vpnType")]
        public string? VpnType { get; set; }

        public VirtualNetworkGatewayPropertiesFormatResponseArgs()
        {
        }
    }
}