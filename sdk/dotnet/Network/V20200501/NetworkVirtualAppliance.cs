// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501
{
    /// <summary>
    /// NetworkVirtualAppliance Resource.
    /// 
    /// ## Example Usage
    /// ### Create NetworkVirtualAppliance
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var networkVirtualAppliance = new AzureRM.Network.V20200501.NetworkVirtualAppliance("networkVirtualAppliance", new AzureRM.Network.V20200501.NetworkVirtualApplianceArgs
    ///         {
    ///             BootStrapConfigurationBlobs = 
    ///             {
    ///                 "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig",
    ///             },
    ///             CloudInitConfigurationBlobs = 
    ///             {
    ///                 "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig",
    ///             },
    ///             Identity = new AzureRM.Network.V20200501.Inputs.ManagedServiceIdentityArgs
    ///             {
    ///                 Type = "UserAssigned",
    ///             },
    ///             Location = "West US",
    ///             NetworkVirtualApplianceName = "nva",
    ///             NvaSku = new AzureRM.Network.V20200501.Inputs.VirtualApplianceSkuPropertiesArgs
    ///             {
    ///                 BundledScaleUnit = "1",
    ///                 MarketPlaceVersion = "12.1",
    ///                 Vendor = "Cisco SDWAN",
    ///             },
    ///             ResourceGroupName = "rg1",
    ///             Tags = 
    ///             {
    ///                 { "key1", "value1" },
    ///             },
    ///             VirtualApplianceAsn = 10000,
    ///             VirtualHub = new AzureRM.Network.V20200501.Inputs.SubResourceArgs
    ///             {
    ///                 Id = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class NetworkVirtualAppliance : Pulumi.CustomResource
    {
        /// <summary>
        /// BootStrapConfigurationBlobs storage URLs.
        /// </summary>
        [Output("bootStrapConfigurationBlobs")]
        public Output<ImmutableArray<string>> BootStrapConfigurationBlobs { get; private set; } = null!;

        /// <summary>
        /// CloudInitConfiguration string in plain text.
        /// </summary>
        [Output("cloudInitConfiguration")]
        public Output<string?> CloudInitConfiguration { get; private set; } = null!;

        /// <summary>
        /// CloudInitConfigurationBlob storage URLs.
        /// </summary>
        [Output("cloudInitConfigurationBlobs")]
        public Output<ImmutableArray<string>> CloudInitConfigurationBlobs { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The service principal that has read access to cloud-init and config blob.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponseResult?> Identity { get; private set; } = null!;

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
        /// Network Virtual Appliance SKU.
        /// </summary>
        [Output("nvaSku")]
        public Output<Outputs.VirtualApplianceSkuPropertiesResponseResult?> NvaSku { get; private set; } = null!;

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
        /// VirtualAppliance ASN.
        /// </summary>
        [Output("virtualApplianceAsn")]
        public Output<int?> VirtualApplianceAsn { get; private set; } = null!;

        /// <summary>
        /// List of Virtual Appliance Network Interfaces.
        /// </summary>
        [Output("virtualApplianceNics")]
        public Output<ImmutableArray<Outputs.VirtualApplianceNicPropertiesResponseResult>> VirtualApplianceNics { get; private set; } = null!;

        /// <summary>
        /// List of references to VirtualApplianceSite.
        /// </summary>
        [Output("virtualApplianceSites")]
        public Output<ImmutableArray<Outputs.SubResourceResponseResult>> VirtualApplianceSites { get; private set; } = null!;

        /// <summary>
        /// The Virtual Hub where Network Virtual Appliance is being deployed.
        /// </summary>
        [Output("virtualHub")]
        public Output<Outputs.SubResourceResponseResult?> VirtualHub { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkVirtualAppliance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkVirtualAppliance(string name, NetworkVirtualApplianceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200501:NetworkVirtualAppliance", name, args ?? new NetworkVirtualApplianceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkVirtualAppliance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200501:NetworkVirtualAppliance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:NetworkVirtualAppliance"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:NetworkVirtualAppliance"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:NetworkVirtualAppliance"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:NetworkVirtualAppliance"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:NetworkVirtualAppliance"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkVirtualAppliance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkVirtualAppliance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkVirtualAppliance(name, id, options);
        }
    }

    public sealed class NetworkVirtualApplianceArgs : Pulumi.ResourceArgs
    {
        [Input("bootStrapConfigurationBlobs")]
        private InputList<string>? _bootStrapConfigurationBlobs;

        /// <summary>
        /// BootStrapConfigurationBlobs storage URLs.
        /// </summary>
        public InputList<string> BootStrapConfigurationBlobs
        {
            get => _bootStrapConfigurationBlobs ?? (_bootStrapConfigurationBlobs = new InputList<string>());
            set => _bootStrapConfigurationBlobs = value;
        }

        /// <summary>
        /// CloudInitConfiguration string in plain text.
        /// </summary>
        [Input("cloudInitConfiguration")]
        public Input<string>? CloudInitConfiguration { get; set; }

        [Input("cloudInitConfigurationBlobs")]
        private InputList<string>? _cloudInitConfigurationBlobs;

        /// <summary>
        /// CloudInitConfigurationBlob storage URLs.
        /// </summary>
        public InputList<string> CloudInitConfigurationBlobs
        {
            get => _cloudInitConfigurationBlobs ?? (_cloudInitConfigurationBlobs = new InputList<string>());
            set => _cloudInitConfigurationBlobs = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The service principal that has read access to cloud-init and config blob.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of Network Virtual Appliance.
        /// </summary>
        [Input("networkVirtualApplianceName", required: true)]
        public Input<string> NetworkVirtualApplianceName { get; set; } = null!;

        /// <summary>
        /// Network Virtual Appliance SKU.
        /// </summary>
        [Input("nvaSku")]
        public Input<Inputs.VirtualApplianceSkuPropertiesArgs>? NvaSku { get; set; }

        /// <summary>
        /// The name of the resource group.
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
        /// VirtualAppliance ASN.
        /// </summary>
        [Input("virtualApplianceAsn")]
        public Input<int>? VirtualApplianceAsn { get; set; }

        /// <summary>
        /// The Virtual Hub where Network Virtual Appliance is being deployed.
        /// </summary>
        [Input("virtualHub")]
        public Input<Inputs.SubResourceArgs>? VirtualHub { get; set; }

        public NetworkVirtualApplianceArgs()
        {
        }
    }
}
