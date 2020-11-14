// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701
{
    public static class GetVpnGateway
    {
        public static Task<GetVpnGatewayResult> InvokeAsync(GetVpnGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpnGatewayResult>("azure-nextgen:network/v20200701:getVpnGateway", args ?? new GetVpnGatewayArgs(), options.WithVersion());
    }


    public sealed class GetVpnGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public string GatewayName { get; set; } = null!;

        /// <summary>
        /// The resource group name of the VpnGateway.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVpnGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpnGatewayResult
    {
        /// <summary>
        /// Local network gateway's BGP speaker settings.
        /// </summary>
        public readonly Outputs.BgpSettingsResponse? BgpSettings;
        /// <summary>
        /// List of all vpn connections to the gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpnConnectionResponse> Connections;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// List of all IPs configured on the gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpnGatewayIpConfigurationResponse> IpConfigurations;
        /// <summary>
        /// Enable Routing Preference property for the Public IP Interface of the VpnGateway.
        /// </summary>
        public readonly bool? IsRoutingPreferenceInternet;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the VPN gateway resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The VirtualHub to which the gateway belongs.
        /// </summary>
        public readonly Outputs.SubResourceResponse? VirtualHub;
        /// <summary>
        /// The scale unit for this vpn gateway.
        /// </summary>
        public readonly int? VpnGatewayScaleUnit;

        [OutputConstructor]
        private GetVpnGatewayResult(
            Outputs.BgpSettingsResponse? bgpSettings,

            ImmutableArray<Outputs.VpnConnectionResponse> connections,

            string etag,

            ImmutableArray<Outputs.VpnGatewayIpConfigurationResponse> ipConfigurations,

            bool? isRoutingPreferenceInternet,

            string location,

            string name,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.SubResourceResponse? virtualHub,

            int? vpnGatewayScaleUnit)
        {
            BgpSettings = bgpSettings;
            Connections = connections;
            Etag = etag;
            IpConfigurations = ipConfigurations;
            IsRoutingPreferenceInternet = isRoutingPreferenceInternet;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
            VirtualHub = virtualHub;
            VpnGatewayScaleUnit = vpnGatewayScaleUnit;
        }
    }
}
