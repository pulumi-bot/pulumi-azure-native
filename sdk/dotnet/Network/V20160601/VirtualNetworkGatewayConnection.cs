// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601
{
    /// <summary>
    /// A common class for general resource information
    /// </summary>
    public partial class VirtualNetworkGatewayConnection : Pulumi.CustomResource
    {
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
        /// VirtualNetworkGatewayConnection properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualNetworkGatewayConnectionPropertiesFormatResponseResult> Properties { get; private set; } = null!;

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
        /// Create a VirtualNetworkGatewayConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetworkGatewayConnection(string name, VirtualNetworkGatewayConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160601:VirtualNetworkGatewayConnection", name, args ?? new VirtualNetworkGatewayConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetworkGatewayConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160601:VirtualNetworkGatewayConnection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualNetworkGatewayConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualNetworkGatewayConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualNetworkGatewayConnection(name, id, options);
        }
    }

    public sealed class VirtualNetworkGatewayConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorizationKey.
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// Virtual network Gateway connection status
        /// </summary>
        [Input("connectionStatus")]
        public Input<string>? ConnectionStatus { get; set; }

        /// <summary>
        /// Gateway connection type IPsec/Dedicated/VpnClient/Vnet2Vnet
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// The Egress Bytes Transferred in this connection
        /// </summary>
        [Input("egressBytesTransferred")]
        public Input<int>? EgressBytesTransferred { get; set; }

        /// <summary>
        /// EnableBgp Flag
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

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
        /// The Ingress Bytes Transferred in this connection
        /// </summary>
        [Input("ingressBytesTransferred")]
        public Input<int>? IngressBytesTransferred { get; set; }

        /// <summary>
        /// A common class for general resource information
        /// </summary>
        [Input("localNetworkGateway2")]
        public Input<Inputs.LocalNetworkGatewayArgs>? LocalNetworkGateway2 { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the virtual network gateway connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The reference to peerings resource.
        /// </summary>
        [Input("peer")]
        public Input<Inputs.SubResourceArgs>? Peer { get; set; }

        /// <summary>
        /// Gets provisioning state of the VirtualNetworkGatewayConnection resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets resource guid property of the VirtualNetworkGatewayConnection resource
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        /// <summary>
        /// The Routing weight.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        /// <summary>
        /// The IPsec share key.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

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
        /// A common class for general resource information
        /// </summary>
        [Input("virtualNetworkGateway1")]
        public Input<Inputs.VirtualNetworkGatewayArgs>? VirtualNetworkGateway1 { get; set; }

        /// <summary>
        /// A common class for general resource information
        /// </summary>
        [Input("virtualNetworkGateway2")]
        public Input<Inputs.VirtualNetworkGatewayArgs>? VirtualNetworkGateway2 { get; set; }

        public VirtualNetworkGatewayConnectionArgs()
        {
        }
    }
}
