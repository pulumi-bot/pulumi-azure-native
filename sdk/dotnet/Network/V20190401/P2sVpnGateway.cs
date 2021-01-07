// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190401
{
    /// <summary>
    /// P2SVpnGateway Resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:network/v20190401:P2sVpnGateway")]
    public partial class P2sVpnGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// The reference of the address space resource which represents the custom routes specified by the customer for P2SVpnGateway and P2S VpnClient.
        /// </summary>
        [Output("customRoutes")]
        public Output<Outputs.AddressSpaceResponse?> CustomRoutes { get; private set; } = null!;

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The P2SVpnServerConfiguration to which the p2sVpnGateway is attached to.
        /// </summary>
        [Output("p2SVpnServerConfiguration")]
        public Output<Outputs.SubResourceResponse?> P2SVpnServerConfiguration { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// The VirtualHub to which the gateway belongs.
        /// </summary>
        [Output("virtualHub")]
        public Output<Outputs.SubResourceResponse?> VirtualHub { get; private set; } = null!;

        /// <summary>
        /// The reference of the address space resource which represents Address space for P2S VpnClient.
        /// </summary>
        [Output("vpnClientAddressPool")]
        public Output<Outputs.AddressSpaceResponse?> VpnClientAddressPool { get; private set; } = null!;

        /// <summary>
        /// All P2S VPN clients' connection health status.
        /// </summary>
        [Output("vpnClientConnectionHealth")]
        public Output<Outputs.VpnClientConnectionHealthResponse> VpnClientConnectionHealth { get; private set; } = null!;

        /// <summary>
        /// The scale unit for this p2s vpn gateway.
        /// </summary>
        [Output("vpnGatewayScaleUnit")]
        public Output<int?> VpnGatewayScaleUnit { get; private set; } = null!;


        /// <summary>
        /// Create a P2sVpnGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public P2sVpnGateway(string name, P2sVpnGatewayArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190401:P2sVpnGateway", name, args ?? new P2sVpnGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private P2sVpnGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190401:P2sVpnGateway", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181001:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:P2sVpnGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200801:P2sVpnGateway"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing P2sVpnGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static P2sVpnGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new P2sVpnGateway(name, id, options);
        }
    }

    public sealed class P2sVpnGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reference of the address space resource which represents the custom routes specified by the customer for P2SVpnGateway and P2S VpnClient.
        /// </summary>
        [Input("customRoutes")]
        public Input<Inputs.AddressSpaceArgs>? CustomRoutes { get; set; }

        /// <summary>
        /// The name of the gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The P2SVpnServerConfiguration to which the p2sVpnGateway is attached to.
        /// </summary>
        [Input("p2SVpnServerConfiguration")]
        public Input<Inputs.SubResourceArgs>? P2SVpnServerConfiguration { get; set; }

        /// <summary>
        /// The resource group name of the P2SVpnGateway.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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
        /// The VirtualHub to which the gateway belongs.
        /// </summary>
        [Input("virtualHub")]
        public Input<Inputs.SubResourceArgs>? VirtualHub { get; set; }

        /// <summary>
        /// The reference of the address space resource which represents Address space for P2S VpnClient.
        /// </summary>
        [Input("vpnClientAddressPool")]
        public Input<Inputs.AddressSpaceArgs>? VpnClientAddressPool { get; set; }

        /// <summary>
        /// The scale unit for this p2s vpn gateway.
        /// </summary>
        [Input("vpnGatewayScaleUnit")]
        public Input<int>? VpnGatewayScaleUnit { get; set; }

        public P2sVpnGatewayArgs()
        {
        }
    }
}
