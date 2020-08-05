// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181201
{
    /// <summary>
    /// Virtual Network Tap resource
    /// </summary>
    public partial class VirtualNetworkTap : Pulumi.CustomResource
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
        /// Virtual Network Tap Properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualNetworkTapPropertiesFormatResponseResult> Properties { get; private set; } = null!;

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
        /// Create a VirtualNetworkTap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetworkTap(string name, VirtualNetworkTapArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20181201:VirtualNetworkTap", name, args ?? new VirtualNetworkTapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetworkTap(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20181201:VirtualNetworkTap", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualNetworkTap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetworkTap Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualNetworkTap(name, id, options);
        }
    }

    public sealed class VirtualNetworkTapArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reference to the private IP address on the internal Load Balancer that will receive the tap
        /// </summary>
        [Input("destinationLoadBalancerFrontEndIPConfiguration")]
        public Input<Inputs.FrontendIPConfigurationArgs>? DestinationLoadBalancerFrontEndIPConfiguration { get; set; }

        /// <summary>
        /// The reference to the private IP Address of the collector nic that will receive the tap
        /// </summary>
        [Input("destinationNetworkInterfaceIPConfiguration")]
        public Input<Inputs.NetworkInterfaceIPConfigurationArgs>? DestinationNetworkInterfaceIPConfiguration { get; set; }

        /// <summary>
        /// The VXLAN destination port that will receive the tapped traffic.
        /// </summary>
        [Input("destinationPort")]
        public Input<int>? DestinationPort { get; set; }

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
        /// The name of the virtual network tap.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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

        public VirtualNetworkTapArgs()
        {
        }
    }
}
