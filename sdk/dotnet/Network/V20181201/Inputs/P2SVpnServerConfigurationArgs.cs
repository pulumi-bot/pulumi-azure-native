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
    /// P2SVpnServerConfiguration Resource.
    /// </summary>
    public sealed class P2SVpnServerConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("p2SVpnServerConfigRadiusClientRootCertificates")]
        private InputList<Inputs.P2SVpnServerConfigRadiusClientRootCertificateArgs>? _p2SVpnServerConfigRadiusClientRootCertificates;

        /// <summary>
        /// Radius client root certificate of P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.P2SVpnServerConfigRadiusClientRootCertificateArgs> P2SVpnServerConfigRadiusClientRootCertificates
        {
            get => _p2SVpnServerConfigRadiusClientRootCertificates ?? (_p2SVpnServerConfigRadiusClientRootCertificates = new InputList<Inputs.P2SVpnServerConfigRadiusClientRootCertificateArgs>());
            set => _p2SVpnServerConfigRadiusClientRootCertificates = value;
        }

        [Input("p2SVpnServerConfigRadiusServerRootCertificates")]
        private InputList<Inputs.P2SVpnServerConfigRadiusServerRootCertificateArgs>? _p2SVpnServerConfigRadiusServerRootCertificates;

        /// <summary>
        /// Radius Server root certificate of P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.P2SVpnServerConfigRadiusServerRootCertificateArgs> P2SVpnServerConfigRadiusServerRootCertificates
        {
            get => _p2SVpnServerConfigRadiusServerRootCertificates ?? (_p2SVpnServerConfigRadiusServerRootCertificates = new InputList<Inputs.P2SVpnServerConfigRadiusServerRootCertificateArgs>());
            set => _p2SVpnServerConfigRadiusServerRootCertificates = value;
        }

        [Input("p2SVpnServerConfigVpnClientRevokedCertificates")]
        private InputList<Inputs.P2SVpnServerConfigVpnClientRevokedCertificateArgs>? _p2SVpnServerConfigVpnClientRevokedCertificates;

        /// <summary>
        /// VPN client revoked certificate of P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.P2SVpnServerConfigVpnClientRevokedCertificateArgs> P2SVpnServerConfigVpnClientRevokedCertificates
        {
            get => _p2SVpnServerConfigVpnClientRevokedCertificates ?? (_p2SVpnServerConfigVpnClientRevokedCertificates = new InputList<Inputs.P2SVpnServerConfigVpnClientRevokedCertificateArgs>());
            set => _p2SVpnServerConfigVpnClientRevokedCertificates = value;
        }

        [Input("p2SVpnServerConfigVpnClientRootCertificates")]
        private InputList<Inputs.P2SVpnServerConfigVpnClientRootCertificateArgs>? _p2SVpnServerConfigVpnClientRootCertificates;

        /// <summary>
        /// VPN client root certificate of P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.P2SVpnServerConfigVpnClientRootCertificateArgs> P2SVpnServerConfigVpnClientRootCertificates
        {
            get => _p2SVpnServerConfigVpnClientRootCertificates ?? (_p2SVpnServerConfigVpnClientRootCertificates = new InputList<Inputs.P2SVpnServerConfigVpnClientRootCertificateArgs>());
            set => _p2SVpnServerConfigVpnClientRootCertificates = value;
        }

        /// <summary>
        /// The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Input("radiusServerAddress")]
        public Input<string>? RadiusServerAddress { get; set; }

        /// <summary>
        /// The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Input("radiusServerSecret")]
        public Input<string>? RadiusServerSecret { get; set; }

        [Input("vpnClientIpsecPolicies")]
        private InputList<Inputs.IpsecPolicyArgs>? _vpnClientIpsecPolicies;

        /// <summary>
        /// VpnClientIpsecPolicies for P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.IpsecPolicyArgs> VpnClientIpsecPolicies
        {
            get => _vpnClientIpsecPolicies ?? (_vpnClientIpsecPolicies = new InputList<Inputs.IpsecPolicyArgs>());
            set => _vpnClientIpsecPolicies = value;
        }

        [Input("vpnProtocols")]
        private InputList<string>? _vpnProtocols;

        /// <summary>
        /// VPN protocols for the P2SVpnServerConfiguration.
        /// </summary>
        public InputList<string> VpnProtocols
        {
            get => _vpnProtocols ?? (_vpnProtocols = new InputList<string>());
            set => _vpnProtocols = value;
        }

        public P2SVpnServerConfigurationArgs()
        {
        }
    }
}
