// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network
{
    /// <summary>
    /// HubVirtualNetworkConnection Resource.
    /// </summary>
    public partial class VirtualHubHubVirtualNetworkConnection : Pulumi.CustomResource
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
        /// Properties of the hub virtual network connection.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.HubVirtualNetworkConnectionPropertiesResponse> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualHubHubVirtualNetworkConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualHubHubVirtualNetworkConnection(string name, VirtualHubHubVirtualNetworkConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network:VirtualHubHubVirtualNetworkConnection", name, args ?? new VirtualHubHubVirtualNetworkConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualHubHubVirtualNetworkConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network:VirtualHubHubVirtualNetworkConnection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualHubHubVirtualNetworkConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualHubHubVirtualNetworkConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualHubHubVirtualNetworkConnection(name, id, options);
        }
    }

    public sealed class VirtualHubHubVirtualNetworkConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the HubVirtualNetworkConnection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

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
        /// Properties of the hub virtual network connection.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.HubVirtualNetworkConnectionPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The resource group name of the HubVirtualNetworkConnection.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the VirtualHub.
        /// </summary>
        [Input("virtualHubName", required: true)]
        public Input<string> VirtualHubName { get; set; } = null!;

        public VirtualHubHubVirtualNetworkConnectionArgs()
        {
        }
    }
}
