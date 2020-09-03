// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201
{
    /// <summary>
    /// VirtualWAN Resource.
    /// 
    /// ## Example Usage
    /// ### VirtualWANCreate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var virtualWan = new AzureRM.Network.V20191201.VirtualWan("virtualWan", new AzureRM.Network.V20191201.VirtualWanArgs
    ///         {
    ///             DisableVpnEncryption = false,
    ///             Location = "West US",
    ///             ResourceGroupName = "rg1",
    ///             Tags = 
    ///             {
    ///                 { "key1", "value1" },
    ///             },
    ///             Type = "Basic",
    ///             VirtualWANName = "wan1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class VirtualWan : Pulumi.CustomResource
    {
        /// <summary>
        /// True if branch to branch traffic is allowed.
        /// </summary>
        [Output("allowBranchToBranchTraffic")]
        public Output<bool?> AllowBranchToBranchTraffic { get; private set; } = null!;

        /// <summary>
        /// True if Vnet to Vnet traffic is allowed.
        /// </summary>
        [Output("allowVnetToVnetTraffic")]
        public Output<bool?> AllowVnetToVnetTraffic { get; private set; } = null!;

        /// <summary>
        /// Vpn encryption to be disabled or not.
        /// </summary>
        [Output("disableVpnEncryption")]
        public Output<bool?> DisableVpnEncryption { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
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
        /// The office local breakout category.
        /// </summary>
        [Output("office365LocalBreakoutCategory")]
        public Output<string?> Office365LocalBreakoutCategory { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the virtual WAN resource.
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
        /// List of VirtualHubs in the VirtualWAN.
        /// </summary>
        [Output("virtualHubs")]
        public Output<ImmutableArray<Outputs.SubResourceResponseResult>> VirtualHubs { get; private set; } = null!;

        /// <summary>
        /// List of VpnSites in the VirtualWAN.
        /// </summary>
        [Output("vpnSites")]
        public Output<ImmutableArray<Outputs.SubResourceResponseResult>> VpnSites { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualWan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualWan(string name, VirtualWanArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191201:VirtualWan", name, args ?? new VirtualWanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualWan(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191201:VirtualWan", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180401:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180601:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180701:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:VirtualWan"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:VirtualWan"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualWan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualWan Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualWan(name, id, options);
        }
    }

    public sealed class VirtualWanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// True if branch to branch traffic is allowed.
        /// </summary>
        [Input("allowBranchToBranchTraffic")]
        public Input<bool>? AllowBranchToBranchTraffic { get; set; }

        /// <summary>
        /// True if Vnet to Vnet traffic is allowed.
        /// </summary>
        [Input("allowVnetToVnetTraffic")]
        public Input<bool>? AllowVnetToVnetTraffic { get; set; }

        /// <summary>
        /// Vpn encryption to be disabled or not.
        /// </summary>
        [Input("disableVpnEncryption")]
        public Input<bool>? DisableVpnEncryption { get; set; }

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
        /// The office local breakout category.
        /// </summary>
        [Input("office365LocalBreakoutCategory")]
        public Input<string>? Office365LocalBreakoutCategory { get; set; }

        /// <summary>
        /// The resource group name of the VirtualWan.
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
        /// The type of the VirtualWAN.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The name of the VirtualWAN being created or updated.
        /// </summary>
        [Input("virtualWANName", required: true)]
        public Input<string> VirtualWANName { get; set; } = null!;

        public VirtualWanArgs()
        {
        }
    }
}
