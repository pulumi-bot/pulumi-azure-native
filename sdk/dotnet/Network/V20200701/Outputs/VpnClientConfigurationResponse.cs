// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Outputs
{

    [OutputType]
    public sealed class VpnClientConfigurationResponse
    {
        /// <summary>
        /// The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.
        /// </summary>
        public readonly string? AadAudience;
        /// <summary>
        /// The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.
        /// </summary>
        public readonly string? AadIssuer;
        /// <summary>
        /// The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.
        /// </summary>
        public readonly string? AadTenant;
        /// <summary>
        /// The radius server address property of the VirtualNetworkGateway resource for vpn client connection.
        /// </summary>
        public readonly string? RadiusServerAddress;
        /// <summary>
        /// The radius secret property of the VirtualNetworkGateway resource for vpn client connection.
        /// </summary>
        public readonly string? RadiusServerSecret;
        /// <summary>
        /// The radiusServers property for multiple radius server configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.RadiusServerResponse> RadiusServers;
        /// <summary>
        /// The reference to the address space resource which represents Address space for P2S VpnClient.
        /// </summary>
        public readonly Outputs.AddressSpaceResponse? VpnClientAddressPool;
        /// <summary>
        /// VpnClientIpsecPolicies for virtual network gateway P2S client.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpsecPolicyResponse> VpnClientIpsecPolicies;
        /// <summary>
        /// VpnClientProtocols for Virtual network gateway.
        /// </summary>
        public readonly ImmutableArray<string> VpnClientProtocols;
        /// <summary>
        /// VpnClientRevokedCertificate for Virtual network gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpnClientRevokedCertificateResponse> VpnClientRevokedCertificates;
        /// <summary>
        /// VpnClientRootCertificate for virtual network gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpnClientRootCertificateResponse> VpnClientRootCertificates;

        [OutputConstructor]
        private VpnClientConfigurationResponse(
            string? aadAudience,

            string? aadIssuer,

            string? aadTenant,

            string? radiusServerAddress,

            string? radiusServerSecret,

            ImmutableArray<Outputs.RadiusServerResponse> radiusServers,

            Outputs.AddressSpaceResponse? vpnClientAddressPool,

            ImmutableArray<Outputs.IpsecPolicyResponse> vpnClientIpsecPolicies,

            ImmutableArray<string> vpnClientProtocols,

            ImmutableArray<Outputs.VpnClientRevokedCertificateResponse> vpnClientRevokedCertificates,

            ImmutableArray<Outputs.VpnClientRootCertificateResponse> vpnClientRootCertificates)
        {
            AadAudience = aadAudience;
            AadIssuer = aadIssuer;
            AadTenant = aadTenant;
            RadiusServerAddress = radiusServerAddress;
            RadiusServerSecret = radiusServerSecret;
            RadiusServers = radiusServers;
            VpnClientAddressPool = vpnClientAddressPool;
            VpnClientIpsecPolicies = vpnClientIpsecPolicies;
            VpnClientProtocols = vpnClientProtocols;
            VpnClientRevokedCertificates = vpnClientRevokedCertificates;
            VpnClientRootCertificates = vpnClientRootCertificates;
        }
    }
}
