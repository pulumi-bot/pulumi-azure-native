// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190401
{
    /// <summary>
    /// VpnGateway Resource.
    /// 
    /// ## Example Usage
    /// ### VpnGatewayPut
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var vpnGateway = new AzureRM.Network.V20190401.VpnGateway("vpnGateway", new AzureRM.Network.V20190401.VpnGatewayArgs
    ///         {
    ///             BgpSettings = new AzureRM.Network.V20190401.Inputs.BgpSettingsArgs
    ///             {
    ///                 Asn = 65515,
    ///                 BgpPeeringAddress = "10.0.1.30",
    ///                 PeerWeight = 0,
    ///             },
    ///             Connections = 
    ///             {
    ///                 new AzureRM.Network.V20190401.Inputs.VpnConnectionArgs
    ///                 {
    ///                     Name = "vpnConnection1",
    ///                 },
    ///             },
    ///             GatewayName = "gateway1",
    ///             Location = "West US",
    ///             ResourceGroupName = "rg1",
    ///             Tags = 
    ///             {
    ///                 { "key1", "value1" },
    ///             },
    ///             VirtualHub = new AzureRM.Network.V20190401.Inputs.SubResourceArgs
    ///             {
    ///                 Id = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class VpnGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// Local network gateway's BGP speaker settings.
        /// </summary>
        [Output("bgpSettings")]
        public Output<Outputs.BgpSettingsResponseResult?> BgpSettings { get; private set; } = null!;

        /// <summary>
        /// List of all vpn connections to the gateway.
        /// </summary>
        [Output("connections")]
        public Output<ImmutableArray<Outputs.VpnConnectionResponseResult>> Connections { get; private set; } = null!;

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
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

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
        public Output<Outputs.SubResourceResponseResult?> VirtualHub { get; private set; } = null!;

        /// <summary>
        /// The scale unit for this vpn gateway.
        /// </summary>
        [Output("vpnGatewayScaleUnit")]
        public Output<int?> VpnGatewayScaleUnit { get; private set; } = null!;


        /// <summary>
        /// Create a VpnGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnGateway(string name, VpnGatewayArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:VpnGateway", name, args ?? new VpnGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:VpnGateway", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180401:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180601:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180701:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:VpnGateway"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:VpnGateway"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpnGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpnGateway(name, id, options);
        }
    }

    public sealed class VpnGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Local network gateway's BGP speaker settings.
        /// </summary>
        [Input("bgpSettings")]
        public Input<Inputs.BgpSettingsArgs>? BgpSettings { get; set; }

        [Input("connections")]
        private InputList<Inputs.VpnConnectionArgs>? _connections;

        /// <summary>
        /// List of all vpn connections to the gateway.
        /// </summary>
        public InputList<Inputs.VpnConnectionArgs> Connections
        {
            get => _connections ?? (_connections = new InputList<Inputs.VpnConnectionArgs>());
            set => _connections = value;
        }

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
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The resource group name of the VpnGateway.
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
        /// The scale unit for this vpn gateway.
        /// </summary>
        [Input("vpnGatewayScaleUnit")]
        public Input<int>? VpnGatewayScaleUnit { get; set; }

        public VpnGatewayArgs()
        {
        }
    }
}
