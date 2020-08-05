// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190401
{
    /// <summary>
    /// Private link service resource.
    /// </summary>
    public partial class PrivateLinkService : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

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
        /// Properties of the private link service.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateLinkServicePropertiesResponseResult> Properties { get; private set; } = null!;

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
        /// Create a PrivateLinkService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateLinkService(string name, PrivateLinkServiceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:PrivateLinkService", name, args ?? new PrivateLinkServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateLinkService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:PrivateLinkService", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateLinkService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateLinkService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateLinkService(name, id, options);
        }
    }

    public sealed class PrivateLinkServiceArgs : Pulumi.ResourceArgs
    {
        [Input("autoApproval")]
        private InputMap<object>? _autoApproval;

        /// <summary>
        /// The auto-approval list of the private link service.
        /// </summary>
        public InputMap<object> AutoApproval
        {
            get => _autoApproval ?? (_autoApproval = new InputMap<object>());
            set => _autoApproval = value;
        }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("fqdns")]
        private InputList<string>? _fqdns;

        /// <summary>
        /// The list of Fqdn.
        /// </summary>
        public InputList<string> Fqdns
        {
            get => _fqdns ?? (_fqdns = new InputList<string>());
            set => _fqdns = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.PrivateLinkServiceIpConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// An array of references to the private link service IP configuration.
        /// </summary>
        public InputList<Inputs.PrivateLinkServiceIpConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.PrivateLinkServiceIpConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        [Input("loadBalancerFrontendIpConfigurations")]
        private InputList<Inputs.FrontendIPConfigurationArgs>? _loadBalancerFrontendIpConfigurations;

        /// <summary>
        /// An array of references to the load balancer IP configurations.
        /// </summary>
        public InputList<Inputs.FrontendIPConfigurationArgs> LoadBalancerFrontendIpConfigurations
        {
            get => _loadBalancerFrontendIpConfigurations ?? (_loadBalancerFrontendIpConfigurations = new InputList<Inputs.FrontendIPConfigurationArgs>());
            set => _loadBalancerFrontendIpConfigurations = value;
        }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the private link service.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("privateEndpointConnections")]
        private InputList<Inputs.PrivateEndpointConnectionArgs>? _privateEndpointConnections;

        /// <summary>
        /// An array of list about connections to the private endpoint.
        /// </summary>
        public InputList<Inputs.PrivateEndpointConnectionArgs> PrivateEndpointConnections
        {
            get => _privateEndpointConnections ?? (_privateEndpointConnections = new InputList<Inputs.PrivateEndpointConnectionArgs>());
            set => _privateEndpointConnections = value;
        }

        /// <summary>
        /// The name of the resource group.
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

        [Input("visibility")]
        private InputMap<object>? _visibility;

        /// <summary>
        /// The visibility list of the private link service.
        /// </summary>
        public InputMap<object> Visibility
        {
            get => _visibility ?? (_visibility = new InputMap<object>());
            set => _visibility = value;
        }

        public PrivateLinkServiceArgs()
        {
        }
    }
}
