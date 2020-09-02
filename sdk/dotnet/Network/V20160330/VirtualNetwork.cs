// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330
{
    /// <summary>
    /// Virtual Network resource
    /// </summary>
    public partial class VirtualNetwork : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets AddressSpace that contains an array of IP address ranges that can be used by subnets
        /// </summary>
        [Output("addressSpace")]
        public Output<Outputs.AddressSpaceResponseResult?> AddressSpace { get; private set; } = null!;

        /// <summary>
        /// Gets or sets DHCPOptions that contains an array of DNS servers available to VMs deployed in the virtual network
        /// </summary>
        [Output("dhcpOptions")]
        public Output<Outputs.DhcpOptionsResponseResult?> DhcpOptions { get; private set; } = null!;

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets resource GUID property of the VirtualNetwork resource
        /// </summary>
        [Output("resourceGuid")]
        public Output<string?> ResourceGuid { get; private set; } = null!;

        /// <summary>
        /// Gets or sets List of subnets in a VirtualNetwork
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<Outputs.SubnetResponseResult>> Subnets { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetwork(string name, VirtualNetworkArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160330:VirtualNetwork", name, args ?? new VirtualNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetwork(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160330:VirtualNetwork", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20150501preview:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20150615:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160601:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160901:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20161201:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170301:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170601:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170801:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170901:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171001:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171101:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180101:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180201:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180401:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180601:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180701:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:VirtualNetwork"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:VirtualNetwork"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetwork Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualNetwork(name, id, options);
        }
    }

    public sealed class VirtualNetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets AddressSpace that contains an array of IP address ranges that can be used by subnets
        /// </summary>
        [Input("addressSpace")]
        public Input<Inputs.AddressSpaceArgs>? AddressSpace { get; set; }

        /// <summary>
        /// Gets or sets DHCPOptions that contains an array of DNS servers available to VMs deployed in the virtual network
        /// </summary>
        [Input("dhcpOptions")]
        public Input<Inputs.DhcpOptionsArgs>? DhcpOptions { get; set; }

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets resource GUID property of the VirtualNetwork resource
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        [Input("subnets")]
        private InputList<Inputs.SubnetArgs>? _subnets;

        /// <summary>
        /// Gets or sets List of subnets in a VirtualNetwork
        /// </summary>
        public InputList<Inputs.SubnetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.SubnetArgs>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the virtual network.
        /// </summary>
        [Input("virtualNetworkName", required: true)]
        public Input<string> VirtualNetworkName { get; set; } = null!;

        public VirtualNetworkArgs()
        {
        }
    }
}
