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
    /// Parameters for VpnServerConfiguration.
    /// </summary>
    public sealed class VpnServerConfigurationPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The set of aad vpn authentication parameters.
        /// </summary>
        [Input("aadAuthenticationParameters")]
        public Input<Inputs.AadAuthenticationParametersArgs>? AadAuthenticationParameters { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name of the VpnServerConfiguration that is unique within a resource group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("radiusClientRootCertificates")]
        private InputList<Inputs.VpnServerConfigRadiusClientRootCertificateArgs>? _radiusClientRootCertificates;

        /// <summary>
        /// Radius client root certificate of VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.VpnServerConfigRadiusClientRootCertificateArgs> RadiusClientRootCertificates
        {
            get => _radiusClientRootCertificates ?? (_radiusClientRootCertificates = new InputList<Inputs.VpnServerConfigRadiusClientRootCertificateArgs>());
            set => _radiusClientRootCertificates = value;
        }

        /// <summary>
        /// The radius server address property of the VpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Input("radiusServerAddress")]
        public Input<string>? RadiusServerAddress { get; set; }

        [Input("radiusServerRootCertificates")]
        private InputList<Inputs.VpnServerConfigRadiusServerRootCertificateArgs>? _radiusServerRootCertificates;

        /// <summary>
        /// Radius Server root certificate of VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.VpnServerConfigRadiusServerRootCertificateArgs> RadiusServerRootCertificates
        {
            get => _radiusServerRootCertificates ?? (_radiusServerRootCertificates = new InputList<Inputs.VpnServerConfigRadiusServerRootCertificateArgs>());
            set => _radiusServerRootCertificates = value;
        }

        /// <summary>
        /// The radius secret property of the VpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Input("radiusServerSecret")]
        public Input<string>? RadiusServerSecret { get; set; }

        [Input("vpnAuthenticationTypes")]
        private InputList<string>? _vpnAuthenticationTypes;

        /// <summary>
        /// VPN authentication types for the VpnServerConfiguration.
        /// </summary>
        public InputList<string> VpnAuthenticationTypes
        {
            get => _vpnAuthenticationTypes ?? (_vpnAuthenticationTypes = new InputList<string>());
            set => _vpnAuthenticationTypes = value;
        }

        [Input("vpnClientIpsecPolicies")]
        private InputList<Inputs.IpsecPolicyArgs>? _vpnClientIpsecPolicies;

        /// <summary>
        /// VpnClientIpsecPolicies for VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.IpsecPolicyArgs> VpnClientIpsecPolicies
        {
            get => _vpnClientIpsecPolicies ?? (_vpnClientIpsecPolicies = new InputList<Inputs.IpsecPolicyArgs>());
            set => _vpnClientIpsecPolicies = value;
        }

        [Input("vpnClientRevokedCertificates")]
        private InputList<Inputs.VpnServerConfigVpnClientRevokedCertificateArgs>? _vpnClientRevokedCertificates;

        /// <summary>
        /// VPN client revoked certificate of VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.VpnServerConfigVpnClientRevokedCertificateArgs> VpnClientRevokedCertificates
        {
            get => _vpnClientRevokedCertificates ?? (_vpnClientRevokedCertificates = new InputList<Inputs.VpnServerConfigVpnClientRevokedCertificateArgs>());
            set => _vpnClientRevokedCertificates = value;
        }

        [Input("vpnClientRootCertificates")]
        private InputList<Inputs.VpnServerConfigVpnClientRootCertificateArgs>? _vpnClientRootCertificates;

        /// <summary>
        /// VPN client root certificate of VpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.VpnServerConfigVpnClientRootCertificateArgs> VpnClientRootCertificates
        {
            get => _vpnClientRootCertificates ?? (_vpnClientRootCertificates = new InputList<Inputs.VpnServerConfigVpnClientRootCertificateArgs>());
            set => _vpnClientRootCertificates = value;
        }

        [Input("vpnProtocols")]
        private InputList<string>? _vpnProtocols;

        /// <summary>
        /// VPN protocols for the VpnServerConfiguration.
        /// </summary>
        public InputList<string> VpnProtocols
        {
            get => _vpnProtocols ?? (_vpnProtocols = new InputList<string>());
            set => _vpnProtocols = value;
        }

        public VpnServerConfigurationPropertiesArgs()
        {
        }
    }
}