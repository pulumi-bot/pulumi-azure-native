// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181101
{
    /// <summary>
    /// Network profile resource.
    /// </summary>
    public partial class NetworkProfile : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
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
        /// Network profile properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.NetworkProfilePropertiesFormatResponseResult> Properties { get; private set; } = null!;

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
        /// Create a NetworkProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkProfile(string name, NetworkProfileArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20181101:NetworkProfile", name, args ?? new NetworkProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20181101:NetworkProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkProfile(name, id, options);
        }
    }

    public sealed class NetworkProfileArgs : Pulumi.ResourceArgs
    {
        [Input("containerNetworkInterfaceConfigurations")]
        private InputList<Inputs.ContainerNetworkInterfaceConfigurationArgs>? _containerNetworkInterfaceConfigurations;

        /// <summary>
        /// List of chid container network interface configurations.
        /// </summary>
        public InputList<Inputs.ContainerNetworkInterfaceConfigurationArgs> ContainerNetworkInterfaceConfigurations
        {
            get => _containerNetworkInterfaceConfigurations ?? (_containerNetworkInterfaceConfigurations = new InputList<Inputs.ContainerNetworkInterfaceConfigurationArgs>());
            set => _containerNetworkInterfaceConfigurations = value;
        }

        [Input("containerNetworkInterfaces")]
        private InputList<Inputs.ContainerNetworkInterfaceArgs>? _containerNetworkInterfaces;

        /// <summary>
        /// List of child container network interfaces.
        /// </summary>
        public InputList<Inputs.ContainerNetworkInterfaceArgs> ContainerNetworkInterfaces
        {
            get => _containerNetworkInterfaces ?? (_containerNetworkInterfaces = new InputList<Inputs.ContainerNetworkInterfaceArgs>());
            set => _containerNetworkInterfaces = value;
        }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
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
        /// The name of the network profile.
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

        public NetworkProfileArgs()
        {
        }
    }
}
