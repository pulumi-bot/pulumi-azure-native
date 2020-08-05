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
    /// VirtualRouter Resource.
    /// </summary>
    public partial class VirtualRouter : Pulumi.CustomResource
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
        /// Properties of the Virtual Router.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualRouterPropertiesFormatResponseResult> Properties { get; private set; } = null!;

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
        /// Create a VirtualRouter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualRouter(string name, VirtualRouterArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191201:VirtualRouter", name, args ?? new VirtualRouterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualRouter(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191201:VirtualRouter", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualRouter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualRouter Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualRouter(name, id, options);
        }
    }

    public sealed class VirtualRouterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Gateway on which VirtualRouter is hosted.
        /// </summary>
        [Input("hostedGateway")]
        public Input<Inputs.SubResourceArgs>? HostedGateway { get; set; }

        /// <summary>
        /// The Subnet on which VirtualRouter is hosted.
        /// </summary>
        [Input("hostedSubnet")]
        public Input<Inputs.SubResourceArgs>? HostedSubnet { get; set; }

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
        /// The name of the Virtual Router.
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

        /// <summary>
        /// VirtualRouter ASN.
        /// </summary>
        [Input("virtualRouterAsn")]
        public Input<int>? VirtualRouterAsn { get; set; }

        [Input("virtualRouterIps")]
        private InputList<string>? _virtualRouterIps;

        /// <summary>
        /// VirtualRouter IPs.
        /// </summary>
        public InputList<string> VirtualRouterIps
        {
            get => _virtualRouterIps ?? (_virtualRouterIps = new InputList<string>());
            set => _virtualRouterIps = value;
        }

        public VirtualRouterArgs()
        {
        }
    }
}
