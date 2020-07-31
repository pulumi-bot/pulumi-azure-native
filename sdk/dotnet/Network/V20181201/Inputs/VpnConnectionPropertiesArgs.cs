// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181201.Inputs
{

    /// <summary>
    /// Parameters for VpnConnection
    /// </summary>
    public sealed class VpnConnectionPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expected bandwidth in MBPS.
        /// </summary>
        [Input("connectionBandwidth")]
        public Input<int>? ConnectionBandwidth { get; set; }

        /// <summary>
        /// The connection status.
        /// </summary>
        [Input("connectionStatus")]
        public Input<string>? ConnectionStatus { get; set; }

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// Enable internet security
        /// </summary>
        [Input("enableInternetSecurity")]
        public Input<bool>? EnableInternetSecurity { get; set; }

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Input("enableRateLimiting")]
        public Input<bool>? EnableRateLimiting { get; set; }

        [Input("ipsecPolicies")]
        private InputList<Inputs.IpsecPolicyArgs>? _ipsecPolicies;

        /// <summary>
        /// The IPSec Policies to be considered by this connection.
        /// </summary>
        public InputList<Inputs.IpsecPolicyArgs> IpsecPolicies
        {
            get => _ipsecPolicies ?? (_ipsecPolicies = new InputList<Inputs.IpsecPolicyArgs>());
            set => _ipsecPolicies = value;
        }

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Id of the connected vpn site.
        /// </summary>
        [Input("remoteVpnSite")]
        public Input<Inputs.SubResourceArgs>? RemoteVpnSite { get; set; }

        /// <summary>
        /// Routing weight for vpn connection.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        /// <summary>
        /// SharedKey for the vpn connection.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        /// <summary>
        /// Connection protocol used for this connection
        /// </summary>
        [Input("vpnConnectionProtocolType")]
        public Input<string>? VpnConnectionProtocolType { get; set; }

        public VpnConnectionPropertiesArgs()
        {
        }
    }
}