// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150615
{
    /// <summary>
    /// Application gateway resource
    /// </summary>
    public partial class ApplicationGateway : Pulumi.CustomResource
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
        /// Properties of the application gateway.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ApplicationGatewayPropertiesFormatResponseResult> Properties { get; private set; } = null!;

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
        /// Create a ApplicationGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationGateway(string name, ApplicationGatewayArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20150615:ApplicationGateway", name, args ?? new ApplicationGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20150615:ApplicationGateway", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApplicationGateway(name, id, options);
        }
    }

    public sealed class ApplicationGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("backendAddressPools")]
        private InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs>? _backendAddressPools;

        /// <summary>
        /// Backend address pool of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs> BackendAddressPools
        {
            get => _backendAddressPools ?? (_backendAddressPools = new InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs>());
            set => _backendAddressPools = value;
        }

        [Input("backendHttpSettingsCollection")]
        private InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs>? _backendHttpSettingsCollection;

        /// <summary>
        /// Backend http settings of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs> BackendHttpSettingsCollection
        {
            get => _backendHttpSettingsCollection ?? (_backendHttpSettingsCollection = new InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs>());
            set => _backendHttpSettingsCollection = value;
        }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("frontendIPConfigurations")]
        private InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs>? _frontendIPConfigurations;

        /// <summary>
        /// Frontend IP addresses of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs> FrontendIPConfigurations
        {
            get => _frontendIPConfigurations ?? (_frontendIPConfigurations = new InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs>());
            set => _frontendIPConfigurations = value;
        }

        [Input("frontendPorts")]
        private InputList<Inputs.ApplicationGatewayFrontendPortArgs>? _frontendPorts;

        /// <summary>
        /// Frontend ports of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayFrontendPortArgs> FrontendPorts
        {
            get => _frontendPorts ?? (_frontendPorts = new InputList<Inputs.ApplicationGatewayFrontendPortArgs>());
            set => _frontendPorts = value;
        }

        [Input("gatewayIPConfigurations")]
        private InputList<Inputs.ApplicationGatewayIPConfigurationArgs>? _gatewayIPConfigurations;

        /// <summary>
        /// Gets or sets subnets of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayIPConfigurationArgs> GatewayIPConfigurations
        {
            get => _gatewayIPConfigurations ?? (_gatewayIPConfigurations = new InputList<Inputs.ApplicationGatewayIPConfigurationArgs>());
            set => _gatewayIPConfigurations = value;
        }

        [Input("httpListeners")]
        private InputList<Inputs.ApplicationGatewayHttpListenerArgs>? _httpListeners;

        /// <summary>
        /// Http listeners of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayHttpListenerArgs> HttpListeners
        {
            get => _httpListeners ?? (_httpListeners = new InputList<Inputs.ApplicationGatewayHttpListenerArgs>());
            set => _httpListeners = value;
        }

        /// <summary>
        /// Resource Identifier.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the application gateway.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("probes")]
        private InputList<Inputs.ApplicationGatewayProbeArgs>? _probes;

        /// <summary>
        /// Probes of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayProbeArgs> Probes
        {
            get => _probes ?? (_probes = new InputList<Inputs.ApplicationGatewayProbeArgs>());
            set => _probes = value;
        }

        /// <summary>
        /// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        [Input("requestRoutingRules")]
        private InputList<Inputs.ApplicationGatewayRequestRoutingRuleArgs>? _requestRoutingRules;

        /// <summary>
        /// Request routing rules of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayRequestRoutingRuleArgs> RequestRoutingRules
        {
            get => _requestRoutingRules ?? (_requestRoutingRules = new InputList<Inputs.ApplicationGatewayRequestRoutingRuleArgs>());
            set => _requestRoutingRules = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Resource GUID property of the application gateway resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        /// <summary>
        /// SKU of the application gateway resource.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ApplicationGatewaySkuArgs>? Sku { get; set; }

        [Input("sslCertificates")]
        private InputList<Inputs.ApplicationGatewaySslCertificateArgs>? _sslCertificates;

        /// <summary>
        /// SSL certificates of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewaySslCertificateArgs> SslCertificates
        {
            get => _sslCertificates ?? (_sslCertificates = new InputList<Inputs.ApplicationGatewaySslCertificateArgs>());
            set => _sslCertificates = value;
        }

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

        [Input("urlPathMaps")]
        private InputList<Inputs.ApplicationGatewayUrlPathMapArgs>? _urlPathMaps;

        /// <summary>
        /// URL path map of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayUrlPathMapArgs> UrlPathMaps
        {
            get => _urlPathMaps ?? (_urlPathMaps = new InputList<Inputs.ApplicationGatewayUrlPathMapArgs>());
            set => _urlPathMaps = value;
        }

        public ApplicationGatewayArgs()
        {
        }
    }
}
