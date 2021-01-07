// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20191101
{
    /// <summary>
    /// VirtualHub Resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:network/v20191101:VirtualHub")]
    public partial class VirtualHub : Pulumi.CustomResource
    {
        /// <summary>
        /// Address-prefix for this VirtualHub.
        /// </summary>
        [Output("addressPrefix")]
        public Output<string?> AddressPrefix { get; private set; } = null!;

        /// <summary>
        /// The azureFirewall associated with this VirtualHub.
        /// </summary>
        [Output("azureFirewall")]
        public Output<Outputs.SubResourceResponse?> AzureFirewall { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The expressRouteGateway associated with this VirtualHub.
        /// </summary>
        [Output("expressRouteGateway")]
        public Output<Outputs.SubResourceResponse?> ExpressRouteGateway { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The P2SVpnGateway associated with this VirtualHub.
        /// </summary>
        [Output("p2SVpnGateway")]
        public Output<Outputs.SubResourceResponse?> P2SVpnGateway { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the virtual hub resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The routeTable associated with this virtual hub.
        /// </summary>
        [Output("routeTable")]
        public Output<Outputs.VirtualHubRouteTableResponse?> RouteTable { get; private set; } = null!;

        /// <summary>
        /// The Security Provider name.
        /// </summary>
        [Output("securityProviderName")]
        public Output<string?> SecurityProviderName { get; private set; } = null!;

        /// <summary>
        /// The sku of this VirtualHub.
        /// </summary>
        [Output("sku")]
        public Output<string?> Sku { get; private set; } = null!;

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
        /// List of all virtual hub route table v2s associated with this VirtualHub.
        /// </summary>
        [Output("virtualHubRouteTableV2s")]
        public Output<ImmutableArray<Outputs.VirtualHubRouteTableV2Response>> VirtualHubRouteTableV2s { get; private set; } = null!;

        /// <summary>
        /// List of all vnet connections with this VirtualHub.
        /// </summary>
        [Output("virtualNetworkConnections")]
        public Output<ImmutableArray<Outputs.HubVirtualNetworkConnectionResponse>> VirtualNetworkConnections { get; private set; } = null!;

        /// <summary>
        /// The VirtualWAN to which the VirtualHub belongs.
        /// </summary>
        [Output("virtualWan")]
        public Output<Outputs.SubResourceResponse?> VirtualWan { get; private set; } = null!;

        /// <summary>
        /// The VpnGateway associated with this VirtualHub.
        /// </summary>
        [Output("vpnGateway")]
        public Output<Outputs.SubResourceResponse?> VpnGateway { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualHub(string name, VirtualHubArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20191101:VirtualHub", name, args ?? new VirtualHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualHub(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20191101:VirtualHub", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180401:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180601:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180701:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181001:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190401:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:VirtualHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200801:VirtualHub"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualHub Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualHub(name, id, options);
        }
    }

    public sealed class VirtualHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address-prefix for this VirtualHub.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// The azureFirewall associated with this VirtualHub.
        /// </summary>
        [Input("azureFirewall")]
        public Input<Inputs.SubResourceArgs>? AzureFirewall { get; set; }

        /// <summary>
        /// The expressRouteGateway associated with this VirtualHub.
        /// </summary>
        [Input("expressRouteGateway")]
        public Input<Inputs.SubResourceArgs>? ExpressRouteGateway { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The P2SVpnGateway associated with this VirtualHub.
        /// </summary>
        [Input("p2SVpnGateway")]
        public Input<Inputs.SubResourceArgs>? P2SVpnGateway { get; set; }

        /// <summary>
        /// The resource group name of the VirtualHub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The routeTable associated with this virtual hub.
        /// </summary>
        [Input("routeTable")]
        public Input<Inputs.VirtualHubRouteTableArgs>? RouteTable { get; set; }

        /// <summary>
        /// The Security Provider name.
        /// </summary>
        [Input("securityProviderName")]
        public Input<string>? SecurityProviderName { get; set; }

        /// <summary>
        /// The sku of this VirtualHub.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

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
        /// The name of the VirtualHub.
        /// </summary>
        [Input("virtualHubName", required: true)]
        public Input<string> VirtualHubName { get; set; } = null!;

        [Input("virtualHubRouteTableV2s")]
        private InputList<Inputs.VirtualHubRouteTableV2Args>? _virtualHubRouteTableV2s;

        /// <summary>
        /// List of all virtual hub route table v2s associated with this VirtualHub.
        /// </summary>
        public InputList<Inputs.VirtualHubRouteTableV2Args> VirtualHubRouteTableV2s
        {
            get => _virtualHubRouteTableV2s ?? (_virtualHubRouteTableV2s = new InputList<Inputs.VirtualHubRouteTableV2Args>());
            set => _virtualHubRouteTableV2s = value;
        }

        [Input("virtualNetworkConnections")]
        private InputList<Inputs.HubVirtualNetworkConnectionArgs>? _virtualNetworkConnections;

        /// <summary>
        /// List of all vnet connections with this VirtualHub.
        /// </summary>
        public InputList<Inputs.HubVirtualNetworkConnectionArgs> VirtualNetworkConnections
        {
            get => _virtualNetworkConnections ?? (_virtualNetworkConnections = new InputList<Inputs.HubVirtualNetworkConnectionArgs>());
            set => _virtualNetworkConnections = value;
        }

        /// <summary>
        /// The VirtualWAN to which the VirtualHub belongs.
        /// </summary>
        [Input("virtualWan")]
        public Input<Inputs.SubResourceArgs>? VirtualWan { get; set; }

        /// <summary>
        /// The VpnGateway associated with this VirtualHub.
        /// </summary>
        [Input("vpnGateway")]
        public Input<Inputs.SubResourceArgs>? VpnGateway { get; set; }

        public VirtualHubArgs()
        {
        }
    }
}
