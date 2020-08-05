// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180601
{
    /// <summary>
    /// Virtual Network resource.
    /// </summary>
    public partial class VirtualNetwork : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

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
        /// Properties of the virtual network.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualNetworkPropertiesFormatResponseResult> Properties { get; private set; } = null!;

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
        /// Create a VirtualNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetwork(string name, VirtualNetworkArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180601:VirtualNetwork", name, args ?? new VirtualNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetwork(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180601:VirtualNetwork", name, null, MakeResourceOptions(options, id))
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
        /// The AddressSpace that contains an array of IP address ranges that can be used by subnets.
        /// </summary>
        [Input("addressSpace")]
        public Input<Inputs.AddressSpaceArgs>? AddressSpace { get; set; }

        /// <summary>
        /// The DDoS protection plan associated with the virtual network.
        /// </summary>
        [Input("ddosProtectionPlan")]
        public Input<Inputs.SubResourceArgs>? DdosProtectionPlan { get; set; }

        /// <summary>
        /// The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
        /// </summary>
        [Input("dhcpOptions")]
        public Input<Inputs.DhcpOptionsArgs>? DhcpOptions { get; set; }

        /// <summary>
        /// Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
        /// </summary>
        [Input("enableDdosProtection")]
        public Input<bool>? EnableDdosProtection { get; set; }

        /// <summary>
        /// Indicates if VM protection is enabled for all the subnets in the virtual network.
        /// </summary>
        [Input("enableVmProtection")]
        public Input<bool>? EnableVmProtection { get; set; }

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the virtual network.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resourceGuid property of the Virtual Network resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        [Input("subnets")]
        private InputList<Inputs.SubnetArgs>? _subnets;

        /// <summary>
        /// A list of subnets in a Virtual Network.
        /// </summary>
        public InputList<Inputs.SubnetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.SubnetArgs>());
            set => _subnets = value;
        }

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

        [Input("virtualNetworkPeerings")]
        private InputList<Inputs.VirtualNetworkPeeringArgs>? _virtualNetworkPeerings;

        /// <summary>
        /// A list of peerings in a Virtual Network.
        /// </summary>
        public InputList<Inputs.VirtualNetworkPeeringArgs> VirtualNetworkPeerings
        {
            get => _virtualNetworkPeerings ?? (_virtualNetworkPeerings = new InputList<Inputs.VirtualNetworkPeeringArgs>());
            set => _virtualNetworkPeerings = value;
        }

        public VirtualNetworkArgs()
        {
        }
    }
}
