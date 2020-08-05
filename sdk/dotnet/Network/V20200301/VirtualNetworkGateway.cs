// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200301
{
    /// <summary>
    /// A common class for general resource information.
    /// </summary>
    public partial class VirtualNetworkGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the virtual network gateway.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualNetworkGatewayPropertiesFormatResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetworkGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetworkGateway(string name, VirtualNetworkGatewayArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200301:VirtualNetworkGateway", name, args ?? new VirtualNetworkGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetworkGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200301:VirtualNetworkGateway", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualNetworkGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetworkGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualNetworkGateway(name, id, options);
        }
    }

    public sealed class VirtualNetworkGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ActiveActive flag.
        /// </summary>
        [Input("activeActive")]
        public Input<bool>? ActiveActive { get; set; }

        /// <summary>
        /// Virtual network gateway's BGP speaker settings.
        /// </summary>
        [Input("bgpSettings")]
        public Input<Inputs.BgpSettingsArgs>? BgpSettings { get; set; }

        /// <summary>
        /// The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
        /// </summary>
        [Input("customRoutes")]
        public Input<Inputs.AddressSpaceArgs>? CustomRoutes { get; set; }

        /// <summary>
        /// Whether BGP is enabled for this virtual network gateway or not.
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// Whether dns forwarding is enabled or not.
        /// </summary>
        [Input("enableDnsForwarding")]
        public Input<bool>? EnableDnsForwarding { get; set; }

        /// <summary>
        /// Whether private IP needs to be enabled on this gateway for connections or not.
        /// </summary>
        [Input("enablePrivateIpAddress")]
        public Input<bool>? EnablePrivateIpAddress { get; set; }

        /// <summary>
        /// The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
        /// </summary>
        [Input("gatewayDefaultSite")]
        public Input<Inputs.SubResourceArgs>? GatewayDefaultSite { get; set; }

        /// <summary>
        /// The type of this virtual network gateway.
        /// </summary>
        [Input("gatewayType")]
        public Input<string>? GatewayType { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.VirtualNetworkGatewayIPConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// IP configurations for virtual network gateway.
        /// </summary>
        public InputList<Inputs.VirtualNetworkGatewayIPConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.VirtualNetworkGatewayIPConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the virtual network gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.VirtualNetworkGatewaySkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
        /// </summary>
        [Input("vpnClientConfiguration")]
        public Input<Inputs.VpnClientConfigurationArgs>? VpnClientConfiguration { get; set; }

        /// <summary>
        /// The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
        /// </summary>
        [Input("vpnGatewayGeneration")]
        public Input<string>? VpnGatewayGeneration { get; set; }

        /// <summary>
        /// The type of this virtual network gateway.
        /// </summary>
        [Input("vpnType")]
        public Input<string>? VpnType { get; set; }

        public VirtualNetworkGatewayArgs()
        {
        }
    }
}
