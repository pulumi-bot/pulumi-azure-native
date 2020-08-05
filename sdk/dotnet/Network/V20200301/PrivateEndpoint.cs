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
    /// Private endpoint resource.
    /// </summary>
    public partial class PrivateEndpoint : Pulumi.CustomResource
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
        /// Properties of the private endpoint.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateEndpointPropertiesResponseResult> Properties { get; private set; } = null!;

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
        /// Create a PrivateEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateEndpoint(string name, PrivateEndpointArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200301:PrivateEndpoint", name, args ?? new PrivateEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200301:PrivateEndpoint", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateEndpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateEndpoint(name, id, options);
        }
    }

    public sealed class PrivateEndpointArgs : Pulumi.ResourceArgs
    {
        [Input("customDnsConfigs")]
        private InputList<Inputs.CustomDnsConfigPropertiesFormatArgs>? _customDnsConfigs;

        /// <summary>
        /// An array of custom dns configurations.
        /// </summary>
        public InputList<Inputs.CustomDnsConfigPropertiesFormatArgs> CustomDnsConfigs
        {
            get => _customDnsConfigs ?? (_customDnsConfigs = new InputList<Inputs.CustomDnsConfigPropertiesFormatArgs>());
            set => _customDnsConfigs = value;
        }

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

        [Input("manualPrivateLinkServiceConnections")]
        private InputList<Inputs.PrivateLinkServiceConnectionArgs>? _manualPrivateLinkServiceConnections;

        /// <summary>
        /// A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
        /// </summary>
        public InputList<Inputs.PrivateLinkServiceConnectionArgs> ManualPrivateLinkServiceConnections
        {
            get => _manualPrivateLinkServiceConnections ?? (_manualPrivateLinkServiceConnections = new InputList<Inputs.PrivateLinkServiceConnectionArgs>());
            set => _manualPrivateLinkServiceConnections = value;
        }

        /// <summary>
        /// The name of the private endpoint.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("privateLinkServiceConnections")]
        private InputList<Inputs.PrivateLinkServiceConnectionArgs>? _privateLinkServiceConnections;

        /// <summary>
        /// A grouping of information about the connection to the remote resource.
        /// </summary>
        public InputList<Inputs.PrivateLinkServiceConnectionArgs> PrivateLinkServiceConnections
        {
            get => _privateLinkServiceConnections ?? (_privateLinkServiceConnections = new InputList<Inputs.PrivateLinkServiceConnectionArgs>());
            set => _privateLinkServiceConnections = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ID of the subnet from which the private IP will be allocated.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.SubnetArgs>? Subnet { get; set; }

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

        public PrivateEndpointArgs()
        {
        }
    }
}
