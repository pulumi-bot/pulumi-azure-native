// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200601
{
    public partial class PrivateEndpointConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// GroupIds from the private link service resource.
        /// </summary>
        [Output("groupIds")]
        public Output<ImmutableArray<string>> GroupIds { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Private Endpoint resource for this Connection.
        /// </summary>
        [Output("privateEndpoint")]
        public Output<Outputs.PrivateEndpointResponseResult?> PrivateEndpoint { get; private set; } = null!;

        /// <summary>
        /// Details about the state of the connection.
        /// </summary>
        [Output("privateLinkServiceConnectionState")]
        public Output<Outputs.ConnectionStateResponseResult?> PrivateLinkServiceConnectionState { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the Private Endpoint Connection.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateEndpointConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateEndpointConnection(string name, PrivateEndpointConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:eventgrid/v20200601:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:eventgrid/v20200601:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:eventgrid/latest:PrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azurerm:eventgrid/v20200401preview:PrivateEndpointConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateEndpointConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateEndpointConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateEndpointConnection(name, id, options);
        }
    }

    public sealed class PrivateEndpointConnectionArgs : Pulumi.ResourceArgs
    {
        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// GroupIds from the private link service resource.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The name of the parent resource (namely, either, the topic name or domain name).
        /// </summary>
        [Input("parentName", required: true)]
        public Input<string> ParentName { get; set; } = null!;

        /// <summary>
        /// The type of the parent resource. This can be either \'topics\' or \'domains\'.
        /// </summary>
        [Input("parentType", required: true)]
        public Input<string> ParentType { get; set; } = null!;

        /// <summary>
        /// The Private Endpoint resource for this Connection.
        /// </summary>
        [Input("privateEndpoint")]
        public Input<Inputs.PrivateEndpointArgs>? PrivateEndpoint { get; set; }

        /// <summary>
        /// The name of the private endpoint connection connection.
        /// </summary>
        [Input("privateEndpointConnectionName", required: true)]
        public Input<string> PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// Details about the state of the connection.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.ConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// Provisioning state of the Private Endpoint Connection.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
    }
}
