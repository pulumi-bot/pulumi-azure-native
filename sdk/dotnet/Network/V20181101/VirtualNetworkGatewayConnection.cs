// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20181101
{
    /// <summary>
    /// A common class for general resource information
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:network/v20181101:VirtualNetworkGatewayConnection")]
    public partial class VirtualNetworkGatewayConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The authorizationKey.
        /// </summary>
        [Output("authorizationKey")]
        public Output<string?> AuthorizationKey { get; private set; } = null!;

        /// <summary>
        /// Connection protocol used for this connection
        /// </summary>
        [Output("connectionProtocol")]
        public Output<string?> ConnectionProtocol { get; private set; } = null!;

        /// <summary>
        /// Virtual network Gateway connection status. Possible values are 'Unknown', 'Connecting', 'Connected' and 'NotConnected'.
        /// </summary>
        [Output("connectionStatus")]
        public Output<string> ConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// Gateway connection type. Possible values are: 'Ipsec','Vnet2Vnet','ExpressRoute', and 'VPNClient.
        /// </summary>
        [Output("connectionType")]
        public Output<string> ConnectionType { get; private set; } = null!;

        /// <summary>
        /// The egress bytes transferred in this connection.
        /// </summary>
        [Output("egressBytesTransferred")]
        public Output<double> EgressBytesTransferred { get; private set; } = null!;

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Output("enableBgp")]
        public Output<bool?> EnableBgp { get; private set; } = null!;

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Bypass ExpressRoute Gateway for data forwarding
        /// </summary>
        [Output("expressRouteGatewayBypass")]
        public Output<bool?> ExpressRouteGatewayBypass { get; private set; } = null!;

        /// <summary>
        /// The ingress bytes transferred in this connection.
        /// </summary>
        [Output("ingressBytesTransferred")]
        public Output<double> IngressBytesTransferred { get; private set; } = null!;

        /// <summary>
        /// The IPSec Policies to be considered by this connection.
        /// </summary>
        [Output("ipsecPolicies")]
        public Output<ImmutableArray<Outputs.IpsecPolicyResponse>> IpsecPolicies { get; private set; } = null!;

        /// <summary>
        /// The reference to local network gateway resource.
        /// </summary>
        [Output("localNetworkGateway2")]
        public Output<Outputs.LocalNetworkGatewayResponse?> LocalNetworkGateway2 { get; private set; } = null!;

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
        /// The reference to peerings resource.
        /// </summary>
        [Output("peer")]
        public Output<Outputs.SubResourceResponse?> Peer { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the VirtualNetworkGatewayConnection resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource GUID property of the VirtualNetworkGatewayConnection resource.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string?> ResourceGuid { get; private set; } = null!;

        /// <summary>
        /// The routing weight.
        /// </summary>
        [Output("routingWeight")]
        public Output<int?> RoutingWeight { get; private set; } = null!;

        /// <summary>
        /// The IPSec shared key.
        /// </summary>
        [Output("sharedKey")]
        public Output<string?> SharedKey { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Collection of all tunnels' connection health status.
        /// </summary>
        [Output("tunnelConnectionStatus")]
        public Output<ImmutableArray<Outputs.TunnelConnectionHealthResponse>> TunnelConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Enable policy-based traffic selectors.
        /// </summary>
        [Output("usePolicyBasedTrafficSelectors")]
        public Output<bool?> UsePolicyBasedTrafficSelectors { get; private set; } = null!;

        /// <summary>
        /// The reference to virtual network gateway resource.
        /// </summary>
        [Output("virtualNetworkGateway1")]
        public Output<Outputs.VirtualNetworkGatewayResponse> VirtualNetworkGateway1 { get; private set; } = null!;

        /// <summary>
        /// The reference to virtual network gateway resource.
        /// </summary>
        [Output("virtualNetworkGateway2")]
        public Output<Outputs.VirtualNetworkGatewayResponse?> VirtualNetworkGateway2 { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualNetworkGatewayConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualNetworkGatewayConnection(string name, VirtualNetworkGatewayConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20181101:VirtualNetworkGatewayConnection", name, args ?? new VirtualNetworkGatewayConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualNetworkGatewayConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20181101:VirtualNetworkGatewayConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150615:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160330:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160601:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160901:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20161201:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170301:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170601:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170801:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170901:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171001:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171101:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180101:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180201:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180401:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180601:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180701:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181001:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190401:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:VirtualNetworkGatewayConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200801:VirtualNetworkGatewayConnection"},
                },
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
        /// Connection protocol used for this connection
        /// </summary>
        [Input("connectionProtocol")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20181101.VirtualNetworkGatewayConnectionProtocol>? ConnectionProtocol { get; set; }

        /// <summary>
        /// Gateway connection type. Possible values are: 'Ipsec','Vnet2Vnet','ExpressRoute', and 'VPNClient.
        /// </summary>
        [Input("connectionType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20181101.VirtualNetworkGatewayConnectionType> ConnectionType { get; set; } = null!;

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Bypass ExpressRoute Gateway for data forwarding
        /// </summary>
        [Input("expressRouteGatewayBypass")]
        public Input<bool>? ExpressRouteGatewayBypass { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipsecPolicies")]
        private InputList<Inputs.IpsecPolicyArgs>? _ipsecPolicies;

        /// <summary>
        /// The IPSec Policies to be considered by this connection.
        /// </summary>
        public InputList<Inputs.IpsecPolicyArgs> IpsecPolicies
        {
            get => _ipsecPolicies ?? (_ipsecPolicies = new InputList<Inputs.IpsecPolicyArgs>());
            set => _ipsecPolicies = value;
        }

        /// <summary>
        /// The reference to local network gateway resource.
        /// </summary>
        [Input("localNetworkGateway2")]
        public Input<Inputs.LocalNetworkGatewayArgs>? LocalNetworkGateway2 { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The reference to peerings resource.
        /// </summary>
        [Input("peer")]
        public Input<Inputs.SubResourceArgs>? Peer { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource GUID property of the VirtualNetworkGatewayConnection resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        /// <summary>
        /// The routing weight.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        /// <summary>
        /// The IPSec shared key.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

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
        /// Enable policy-based traffic selectors.
        /// </summary>
        [Input("usePolicyBasedTrafficSelectors")]
        public Input<bool>? UsePolicyBasedTrafficSelectors { get; set; }

        /// <summary>
        /// The reference to virtual network gateway resource.
        /// </summary>
        [Input("virtualNetworkGateway1", required: true)]
        public Input<Inputs.VirtualNetworkGatewayArgs> VirtualNetworkGateway1 { get; set; } = null!;

        /// <summary>
        /// The reference to virtual network gateway resource.
        /// </summary>
        [Input("virtualNetworkGateway2")]
        public Input<Inputs.VirtualNetworkGatewayArgs>? VirtualNetworkGateway2 { get; set; }

        /// <summary>
        /// The name of the virtual network gateway connection.
        /// </summary>
        [Input("virtualNetworkGatewayConnectionName", required: true)]
        public Input<string> VirtualNetworkGatewayConnectionName { get; set; } = null!;

        public VirtualNetworkGatewayConnectionArgs()
        {
        }
    }
}
