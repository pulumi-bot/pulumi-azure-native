// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventHub.V20180101Preview
{
    /// <summary>
    /// Properties of the PrivateEndpointConnection.
    /// </summary>
    public partial class PrivateEndpointConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
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
        /// Resource type.
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
            : base("azurerm:eventhub/v20180101preview:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:eventhub/v20180101preview:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
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
        /// <summary>
        /// The Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The Private Endpoint resource for this Connection.
        /// </summary>
        [Input("privateEndpoint")]
        public Input<Inputs.PrivateEndpointArgs>? PrivateEndpoint { get; set; }

        /// <summary>
        /// The PrivateEndpointConnection name
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
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
    }
}
