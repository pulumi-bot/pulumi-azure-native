// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190901
{
    /// <summary>
    /// Peering in an ExpressRouteCircuit resource.
    /// </summary>
    public partial class ExpressRouteCircuitPeering : Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure ASN.
        /// </summary>
        [Output("azureASN")]
        public Output<int?> AzureASN { get; private set; } = null!;

        /// <summary>
        /// The list of circuit connections associated with Azure Private Peering for this circuit.
        /// </summary>
        [Output("connections")]
        public Output<ImmutableArray<Outputs.ExpressRouteCircuitConnectionResponse>> Connections { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The ExpressRoute connection.
        /// </summary>
        [Output("expressRouteConnection")]
        public Output<Outputs.ExpressRouteConnectionIdResponse?> ExpressRouteConnection { get; private set; } = null!;

        /// <summary>
        /// The GatewayManager Etag.
        /// </summary>
        [Output("gatewayManagerEtag")]
        public Output<string?> GatewayManagerEtag { get; private set; } = null!;

        /// <summary>
        /// The IPv6 peering configuration.
        /// </summary>
        [Output("ipv6PeeringConfig")]
        public Output<Outputs.Ipv6ExpressRouteCircuitPeeringConfigResponse?> Ipv6PeeringConfig { get; private set; } = null!;

        /// <summary>
        /// Who was the last to modify the peering.
        /// </summary>
        [Output("lastModifiedBy")]
        public Output<string> LastModifiedBy { get; private set; } = null!;

        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        [Output("microsoftPeeringConfig")]
        public Output<Outputs.ExpressRouteCircuitPeeringConfigResponse?> MicrosoftPeeringConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The peer ASN.
        /// </summary>
        [Output("peerASN")]
        public Output<int?> PeerASN { get; private set; } = null!;

        /// <summary>
        /// The list of peered circuit connections associated with Azure Private Peering for this circuit.
        /// </summary>
        [Output("peeredConnections")]
        public Output<ImmutableArray<Outputs.PeerExpressRouteCircuitConnectionResponse>> PeeredConnections { get; private set; } = null!;

        /// <summary>
        /// The peering type.
        /// </summary>
        [Output("peeringType")]
        public Output<string?> PeeringType { get; private set; } = null!;

        /// <summary>
        /// The primary port.
        /// </summary>
        [Output("primaryAzurePort")]
        public Output<string?> PrimaryAzurePort { get; private set; } = null!;

        /// <summary>
        /// The primary address prefix.
        /// </summary>
        [Output("primaryPeerAddressPrefix")]
        public Output<string?> PrimaryPeerAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the express route circuit peering resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The reference of the RouteFilter resource.
        /// </summary>
        [Output("routeFilter")]
        public Output<Outputs.SubResourceResponse?> RouteFilter { get; private set; } = null!;

        /// <summary>
        /// The secondary port.
        /// </summary>
        [Output("secondaryAzurePort")]
        public Output<string?> SecondaryAzurePort { get; private set; } = null!;

        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        [Output("secondaryPeerAddressPrefix")]
        public Output<string?> SecondaryPeerAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// The shared key.
        /// </summary>
        [Output("sharedKey")]
        public Output<string?> SharedKey { get; private set; } = null!;

        /// <summary>
        /// The peering state.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// The peering stats of express route circuit.
        /// </summary>
        [Output("stats")]
        public Output<Outputs.ExpressRouteCircuitStatsResponse?> Stats { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Output("vlanId")]
        public Output<int?> VlanId { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRouteCircuitPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRouteCircuitPeering(string name, ExpressRouteCircuitPeeringArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190901:ExpressRouteCircuitPeering", name, args ?? new ExpressRouteCircuitPeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteCircuitPeering(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190901:ExpressRouteCircuitPeering", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150501preview:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150615:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160330:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160901:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20161201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170301:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170801:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170901:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171001:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171101:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180101:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180401:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180701:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181001:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190401:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:ExpressRouteCircuitPeering"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExpressRouteCircuitPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRouteCircuitPeering Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRouteCircuitPeering(name, id, options);
        }
    }

    public sealed class ExpressRouteCircuitPeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure ASN.
        /// </summary>
        [Input("azureASN")]
        public Input<int>? AzureASN { get; set; }

        /// <summary>
        /// The name of the express route circuit.
        /// </summary>
        [Input("circuitName", required: true)]
        public Input<string> CircuitName { get; set; } = null!;

        [Input("connections")]
        private InputList<Inputs.ExpressRouteCircuitConnectionArgs>? _connections;

        /// <summary>
        /// The list of circuit connections associated with Azure Private Peering for this circuit.
        /// </summary>
        public InputList<Inputs.ExpressRouteCircuitConnectionArgs> Connections
        {
            get => _connections ?? (_connections = new InputList<Inputs.ExpressRouteCircuitConnectionArgs>());
            set => _connections = value;
        }

        /// <summary>
        /// The GatewayManager Etag.
        /// </summary>
        [Input("gatewayManagerEtag")]
        public Input<string>? GatewayManagerEtag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The IPv6 peering configuration.
        /// </summary>
        [Input("ipv6PeeringConfig")]
        public Input<Inputs.Ipv6ExpressRouteCircuitPeeringConfigArgs>? Ipv6PeeringConfig { get; set; }

        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        [Input("microsoftPeeringConfig")]
        public Input<Inputs.ExpressRouteCircuitPeeringConfigArgs>? MicrosoftPeeringConfig { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The peer ASN.
        /// </summary>
        [Input("peerASN")]
        public Input<int>? PeerASN { get; set; }

        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public Input<string> PeeringName { get; set; } = null!;

        /// <summary>
        /// The peering type.
        /// </summary>
        [Input("peeringType")]
        public Input<string>? PeeringType { get; set; }

        /// <summary>
        /// The primary port.
        /// </summary>
        [Input("primaryAzurePort")]
        public Input<string>? PrimaryAzurePort { get; set; }

        /// <summary>
        /// The primary address prefix.
        /// </summary>
        [Input("primaryPeerAddressPrefix")]
        public Input<string>? PrimaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The reference of the RouteFilter resource.
        /// </summary>
        [Input("routeFilter")]
        public Input<Inputs.SubResourceArgs>? RouteFilter { get; set; }

        /// <summary>
        /// The secondary port.
        /// </summary>
        [Input("secondaryAzurePort")]
        public Input<string>? SecondaryAzurePort { get; set; }

        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        [Input("secondaryPeerAddressPrefix")]
        public Input<string>? SecondaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// The shared key.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        /// <summary>
        /// The peering state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The peering stats of express route circuit.
        /// </summary>
        [Input("stats")]
        public Input<Inputs.ExpressRouteCircuitStatsArgs>? Stats { get; set; }

        /// <summary>
        /// The VLAN ID.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public ExpressRouteCircuitPeeringArgs()
        {
        }
    }
}
