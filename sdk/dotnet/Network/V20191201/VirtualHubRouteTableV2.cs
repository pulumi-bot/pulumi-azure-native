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
    /// VirtualHubRouteTableV2 Resource.
    /// </summary>
    public partial class VirtualHubRouteTableV2 : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the virtual hub route table v2.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualHubRouteTableV2PropertiesResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualHubRouteTableV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualHubRouteTableV2(string name, VirtualHubRouteTableV2Args args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191201:VirtualHubRouteTableV2", name, args ?? new VirtualHubRouteTableV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualHubRouteTableV2(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191201:VirtualHubRouteTableV2", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualHubRouteTableV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualHubRouteTableV2 Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualHubRouteTableV2(name, id, options);
        }
    }

    public sealed class VirtualHubRouteTableV2Args : Pulumi.ResourceArgs
    {
        [Input("attachedConnections")]
        private InputList<string>? _attachedConnections;

        /// <summary>
        /// List of all connections attached to this route table v2.
        /// </summary>
        public InputList<string> AttachedConnections
        {
            get => _attachedConnections ?? (_attachedConnections = new InputList<string>());
            set => _attachedConnections = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the VirtualHubRouteTableV2.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The resource group name of the VirtualHub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("routes")]
        private InputList<Inputs.VirtualHubRouteV2Args>? _routes;

        /// <summary>
        /// List of all routes.
        /// </summary>
        public InputList<Inputs.VirtualHubRouteV2Args> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.VirtualHubRouteV2Args>());
            set => _routes = value;
        }

        /// <summary>
        /// The name of the VirtualHub.
        /// </summary>
        [Input("virtualHubName", required: true)]
        public Input<string> VirtualHubName { get; set; } = null!;

        public VirtualHubRouteTableV2Args()
        {
        }
    }
}
