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
    /// PrivateEndpointConnection resource.
    /// </summary>
    public partial class PrivateLinkServicePrivateEndpointConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The consumer link id.
        /// </summary>
        [Output("linkIdentifier")]
        public Output<string> LinkIdentifier { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The resource of private end point.
        /// </summary>
        [Output("privateEndpoint")]
        public Output<Outputs.PrivateEndpointResponse> PrivateEndpoint { get; private set; } = null!;

        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        [Output("privateLinkServiceConnectionState")]
        public Output<Outputs.PrivateLinkServiceConnectionStateResponse?> PrivateLinkServiceConnectionState { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the private endpoint connection resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateLinkServicePrivateEndpointConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateLinkServicePrivateEndpointConnection(string name, PrivateLinkServicePrivateEndpointConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190901:PrivateLinkServicePrivateEndpointConnection", name, args ?? new PrivateLinkServicePrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateLinkServicePrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190901:PrivateLinkServicePrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:PrivateLinkServicePrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:PrivateLinkServicePrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:PrivateLinkServicePrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:PrivateLinkServicePrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:PrivateLinkServicePrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:PrivateLinkServicePrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:PrivateLinkServicePrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:PrivateLinkServicePrivateEndpointConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateLinkServicePrivateEndpointConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateLinkServicePrivateEndpointConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateLinkServicePrivateEndpointConnection(name, id, options);
        }
    }

    public sealed class PrivateLinkServicePrivateEndpointConnectionArgs : Pulumi.ResourceArgs
    {
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
        /// The name of the private end point connection.
        /// </summary>
        [Input("peConnectionName", required: true)]
        public Input<string> PeConnectionName { get; set; } = null!;

        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.PrivateLinkServiceConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the private link service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public PrivateLinkServicePrivateEndpointConnectionArgs()
        {
        }
    }
}
