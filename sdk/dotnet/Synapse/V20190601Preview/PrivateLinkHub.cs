// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20190601Preview
{
    /// <summary>
    /// A privateLinkHub
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:synapse/v20190601preview:PrivateLinkHub")]
    public partial class PrivateLinkHub : Pulumi.CustomResource
    {
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of private endpoint connections
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.PrivateEndpointConnectionForPrivateLinkHubBasicResponse>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// PrivateLinkHub provisioning state
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateLinkHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateLinkHub(string name, PrivateLinkHubArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:synapse/v20190601preview:PrivateLinkHub", name, args ?? new PrivateLinkHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateLinkHub(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:synapse/v20190601preview:PrivateLinkHub", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:synapse/v20201201:PrivateLinkHub"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateLinkHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateLinkHub Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateLinkHub(name, id, options);
        }
    }

    public sealed class PrivateLinkHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Name of the privateLinkHub
        /// </summary>
        [Input("privateLinkHubName", required: true)]
        public Input<string> PrivateLinkHubName { get; set; } = null!;

        /// <summary>
        /// PrivateLinkHub provisioning state
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public PrivateLinkHubArgs()
        {
        }
    }
}
