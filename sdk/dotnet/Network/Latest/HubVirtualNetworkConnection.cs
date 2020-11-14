// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest
{
    /// <summary>
    /// HubVirtualNetworkConnection Resource.
    /// </summary>
    public partial class HubVirtualNetworkConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
        /// </summary>
        [Output("allowHubToRemoteVnetTransit")]
        public Output<bool?> AllowHubToRemoteVnetTransit { get; private set; } = null!;

        /// <summary>
        /// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
        /// </summary>
        [Output("allowRemoteVnetToUseHubVnetGateways")]
        public Output<bool?> AllowRemoteVnetToUseHubVnetGateways { get; private set; } = null!;

        /// <summary>
        /// Enable internet security.
        /// </summary>
        [Output("enableInternetSecurity")]
        public Output<bool?> EnableInternetSecurity { get; private set; } = null!;

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
        /// The provisioning state of the hub virtual network connection resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Reference to the remote virtual network.
        /// </summary>
        [Output("remoteVirtualNetwork")]
        public Output<Outputs.SubResourceResponse?> RemoteVirtualNetwork { get; private set; } = null!;

        /// <summary>
        /// The Routing Configuration indicating the associated and propagated route tables on this connection.
        /// </summary>
        [Output("routingConfiguration")]
        public Output<Outputs.RoutingConfigurationResponse?> RoutingConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a HubVirtualNetworkConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HubVirtualNetworkConnection(string name, HubVirtualNetworkConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:HubVirtualNetworkConnection", name, args ?? new HubVirtualNetworkConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HubVirtualNetworkConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:HubVirtualNetworkConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:HubVirtualNetworkConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:HubVirtualNetworkConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:HubVirtualNetworkConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HubVirtualNetworkConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HubVirtualNetworkConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HubVirtualNetworkConnection(name, id, options);
        }
    }

    public sealed class HubVirtualNetworkConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
        /// </summary>
        [Input("allowHubToRemoteVnetTransit")]
        public Input<bool>? AllowHubToRemoteVnetTransit { get; set; }

        /// <summary>
        /// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
        /// </summary>
        [Input("allowRemoteVnetToUseHubVnetGateways")]
        public Input<bool>? AllowRemoteVnetToUseHubVnetGateways { get; set; }

        /// <summary>
        /// The name of the HubVirtualNetworkConnection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Enable internet security.
        /// </summary>
        [Input("enableInternetSecurity")]
        public Input<bool>? EnableInternetSecurity { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Reference to the remote virtual network.
        /// </summary>
        [Input("remoteVirtualNetwork")]
        public Input<Inputs.SubResourceArgs>? RemoteVirtualNetwork { get; set; }

        /// <summary>
        /// The resource group name of the HubVirtualNetworkConnection.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Routing Configuration indicating the associated and propagated route tables on this connection.
        /// </summary>
        [Input("routingConfiguration")]
        public Input<Inputs.RoutingConfigurationArgs>? RoutingConfiguration { get; set; }

        /// <summary>
        /// The name of the VirtualHub.
        /// </summary>
        [Input("virtualHubName", required: true)]
        public Input<string> VirtualHubName { get; set; } = null!;

        public HubVirtualNetworkConnectionArgs()
        {
        }
    }
}
