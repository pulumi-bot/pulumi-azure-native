// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200501
{
    /// <summary>
    /// Virtual Appliance Site resource.
    /// </summary>
    public partial class VirtualHubBgpConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The current state of the VirtualHub to Peer.
        /// </summary>
        [Output("connectionState")]
        public Output<string> ConnectionState { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Name of the connection.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Peer ASN.
        /// </summary>
        [Output("peerAsn")]
        public Output<int?> PeerAsn { get; private set; } = null!;

        /// <summary>
        /// Peer IP.
        /// </summary>
        [Output("peerIp")]
        public Output<string?> PeerIp { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Connection type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualHubBgpConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualHubBgpConnection(string name, VirtualHubBgpConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20200501:VirtualHubBgpConnection", name, args ?? new VirtualHubBgpConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualHubBgpConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20200501:VirtualHubBgpConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:VirtualHubBgpConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:VirtualHubBgpConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:VirtualHubBgpConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualHubBgpConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualHubBgpConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualHubBgpConnection(name, id, options);
        }
    }

    public sealed class VirtualHubBgpConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Peer ASN.
        /// </summary>
        [Input("peerAsn")]
        public Input<int>? PeerAsn { get; set; }

        /// <summary>
        /// Peer IP.
        /// </summary>
        [Input("peerIp")]
        public Input<string>? PeerIp { get; set; }

        /// <summary>
        /// The resource group name of the VirtualHub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the VirtualHub.
        /// </summary>
        [Input("virtualHubName", required: true)]
        public Input<string> VirtualHubName { get; set; } = null!;

        public VirtualHubBgpConnectionArgs()
        {
        }
    }
}
