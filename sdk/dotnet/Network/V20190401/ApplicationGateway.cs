// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190401
{
    /// <summary>
    /// Application gateway resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:network/v20190401:ApplicationGateway")]
    public partial class ApplicationGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("authenticationCertificates")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayAuthenticationCertificateResponse>> AuthenticationCertificates { get; private set; } = null!;

        /// <summary>
        /// Autoscale Configuration.
        /// </summary>
        [Output("autoscaleConfiguration")]
        public Output<Outputs.ApplicationGatewayAutoscaleConfigurationResponse?> AutoscaleConfiguration { get; private set; } = null!;

        /// <summary>
        /// Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("backendAddressPools")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayBackendAddressPoolResponse>> BackendAddressPools { get; private set; } = null!;

        /// <summary>
        /// Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("backendHttpSettingsCollection")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayBackendHttpSettingsResponse>> BackendHttpSettingsCollection { get; private set; } = null!;

        /// <summary>
        /// Custom error configurations of the application gateway resource.
        /// </summary>
        [Output("customErrorConfigurations")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayCustomErrorResponse>> CustomErrorConfigurations { get; private set; } = null!;

        /// <summary>
        /// Whether FIPS is enabled on the application gateway resource.
        /// </summary>
        [Output("enableFips")]
        public Output<bool?> EnableFips { get; private set; } = null!;

        /// <summary>
        /// Whether HTTP2 is enabled on the application gateway resource.
        /// </summary>
        [Output("enableHttp2")]
        public Output<bool?> EnableHttp2 { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Reference of the FirewallPolicy resource.
        /// </summary>
        [Output("firewallPolicy")]
        public Output<Outputs.SubResourceResponse?> FirewallPolicy { get; private set; } = null!;

        /// <summary>
        /// Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("frontendIPConfigurations")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayFrontendIPConfigurationResponse>> FrontendIPConfigurations { get; private set; } = null!;

        /// <summary>
        /// Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("frontendPorts")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayFrontendPortResponse>> FrontendPorts { get; private set; } = null!;

        /// <summary>
        /// Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("gatewayIPConfigurations")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayIPConfigurationResponse>> GatewayIPConfigurations { get; private set; } = null!;

        /// <summary>
        /// Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("httpListeners")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayHttpListenerResponse>> HttpListeners { get; private set; } = null!;

        /// <summary>
        /// The identity of the application gateway, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

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
        /// Operational state of the application gateway resource.
        /// </summary>
        [Output("operationalState")]
        public Output<string> OperationalState { get; private set; } = null!;

        /// <summary>
        /// Probes of the application gateway resource.
        /// </summary>
        [Output("probes")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayProbeResponse>> Probes { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("redirectConfigurations")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayRedirectConfigurationResponse>> RedirectConfigurations { get; private set; } = null!;

        /// <summary>
        /// Request routing rules of the application gateway resource.
        /// </summary>
        [Output("requestRoutingRules")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayRequestRoutingRuleResponse>> RequestRoutingRules { get; private set; } = null!;

        /// <summary>
        /// Resource GUID property of the application gateway resource.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string?> ResourceGuid { get; private set; } = null!;

        /// <summary>
        /// Rewrite rules for the application gateway resource.
        /// </summary>
        [Output("rewriteRuleSets")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayRewriteRuleSetResponse>> RewriteRuleSets { get; private set; } = null!;

        /// <summary>
        /// SKU of the application gateway resource.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ApplicationGatewaySkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("sslCertificates")]
        public Output<ImmutableArray<Outputs.ApplicationGatewaySslCertificateResponse>> SslCertificates { get; private set; } = null!;

        /// <summary>
        /// SSL policy of the application gateway resource.
        /// </summary>
        [Output("sslPolicy")]
        public Output<Outputs.ApplicationGatewaySslPolicyResponse?> SslPolicy { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("trustedRootCertificates")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayTrustedRootCertificateResponse>> TrustedRootCertificates { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        [Output("urlPathMaps")]
        public Output<ImmutableArray<Outputs.ApplicationGatewayUrlPathMapResponse>> UrlPathMaps { get; private set; } = null!;

        /// <summary>
        /// Web application firewall configuration.
        /// </summary>
        [Output("webApplicationFirewallConfiguration")]
        public Output<Outputs.ApplicationGatewayWebApplicationFirewallConfigurationResponse?> WebApplicationFirewallConfiguration { get; private set; } = null!;

        /// <summary>
        /// A list of availability zones denoting where the resource needs to come from.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationGateway(string name, ApplicationGatewayArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190401:ApplicationGateway", name, args ?? new ApplicationGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190401:ApplicationGateway", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150501preview:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150615:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160330:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160601:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160901:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20161201:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170301:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170601:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170801:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170901:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171001:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171101:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180101:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180201:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180401:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180601:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180701:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181001:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:ApplicationGateway"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200801:ApplicationGateway"},
                },
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
        /// <summary>
        /// The name of the application gateway.
        /// </summary>
        [Input("applicationGatewayName", required: true)]
        public Input<string> ApplicationGatewayName { get; set; } = null!;

        [Input("authenticationCertificates")]
        private InputList<Inputs.ApplicationGatewayAuthenticationCertificateArgs>? _authenticationCertificates;

        /// <summary>
        /// Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayAuthenticationCertificateArgs> AuthenticationCertificates
        {
            get => _authenticationCertificates ?? (_authenticationCertificates = new InputList<Inputs.ApplicationGatewayAuthenticationCertificateArgs>());
            set => _authenticationCertificates = value;
        }

        /// <summary>
        /// Autoscale Configuration.
        /// </summary>
        [Input("autoscaleConfiguration")]
        public Input<Inputs.ApplicationGatewayAutoscaleConfigurationArgs>? AutoscaleConfiguration { get; set; }

        [Input("backendAddressPools")]
        private InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs>? _backendAddressPools;

        /// <summary>
        /// Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs> BackendAddressPools
        {
            get => _backendAddressPools ?? (_backendAddressPools = new InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs>());
            set => _backendAddressPools = value;
        }

        [Input("backendHttpSettingsCollection")]
        private InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs>? _backendHttpSettingsCollection;

        /// <summary>
        /// Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs> BackendHttpSettingsCollection
        {
            get => _backendHttpSettingsCollection ?? (_backendHttpSettingsCollection = new InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs>());
            set => _backendHttpSettingsCollection = value;
        }

        [Input("customErrorConfigurations")]
        private InputList<Inputs.ApplicationGatewayCustomErrorArgs>? _customErrorConfigurations;

        /// <summary>
        /// Custom error configurations of the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayCustomErrorArgs> CustomErrorConfigurations
        {
            get => _customErrorConfigurations ?? (_customErrorConfigurations = new InputList<Inputs.ApplicationGatewayCustomErrorArgs>());
            set => _customErrorConfigurations = value;
        }

        /// <summary>
        /// Whether FIPS is enabled on the application gateway resource.
        /// </summary>
        [Input("enableFips")]
        public Input<bool>? EnableFips { get; set; }

        /// <summary>
        /// Whether HTTP2 is enabled on the application gateway resource.
        /// </summary>
        [Input("enableHttp2")]
        public Input<bool>? EnableHttp2 { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Reference of the FirewallPolicy resource.
        /// </summary>
        [Input("firewallPolicy")]
        public Input<Inputs.SubResourceArgs>? FirewallPolicy { get; set; }

        [Input("frontendIPConfigurations")]
        private InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs>? _frontendIPConfigurations;

        /// <summary>
        /// Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs> FrontendIPConfigurations
        {
            get => _frontendIPConfigurations ?? (_frontendIPConfigurations = new InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs>());
            set => _frontendIPConfigurations = value;
        }

        [Input("frontendPorts")]
        private InputList<Inputs.ApplicationGatewayFrontendPortArgs>? _frontendPorts;

        /// <summary>
        /// Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayFrontendPortArgs> FrontendPorts
        {
            get => _frontendPorts ?? (_frontendPorts = new InputList<Inputs.ApplicationGatewayFrontendPortArgs>());
            set => _frontendPorts = value;
        }

        [Input("gatewayIPConfigurations")]
        private InputList<Inputs.ApplicationGatewayIPConfigurationArgs>? _gatewayIPConfigurations;

        /// <summary>
        /// Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayIPConfigurationArgs> GatewayIPConfigurations
        {
            get => _gatewayIPConfigurations ?? (_gatewayIPConfigurations = new InputList<Inputs.ApplicationGatewayIPConfigurationArgs>());
            set => _gatewayIPConfigurations = value;
        }

        [Input("httpListeners")]
        private InputList<Inputs.ApplicationGatewayHttpListenerArgs>? _httpListeners;

        /// <summary>
        /// Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayHttpListenerArgs> HttpListeners
        {
            get => _httpListeners ?? (_httpListeners = new InputList<Inputs.ApplicationGatewayHttpListenerArgs>());
            set => _httpListeners = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The identity of the application gateway, if configured.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

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

        [Input("redirectConfigurations")]
        private InputList<Inputs.ApplicationGatewayRedirectConfigurationArgs>? _redirectConfigurations;

        /// <summary>
        /// Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayRedirectConfigurationArgs> RedirectConfigurations
        {
            get => _redirectConfigurations ?? (_redirectConfigurations = new InputList<Inputs.ApplicationGatewayRedirectConfigurationArgs>());
            set => _redirectConfigurations = value;
        }

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

        [Input("rewriteRuleSets")]
        private InputList<Inputs.ApplicationGatewayRewriteRuleSetArgs>? _rewriteRuleSets;

        /// <summary>
        /// Rewrite rules for the application gateway resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayRewriteRuleSetArgs> RewriteRuleSets
        {
            get => _rewriteRuleSets ?? (_rewriteRuleSets = new InputList<Inputs.ApplicationGatewayRewriteRuleSetArgs>());
            set => _rewriteRuleSets = value;
        }

        /// <summary>
        /// SKU of the application gateway resource.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ApplicationGatewaySkuArgs>? Sku { get; set; }

        [Input("sslCertificates")]
        private InputList<Inputs.ApplicationGatewaySslCertificateArgs>? _sslCertificates;

        /// <summary>
        /// SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewaySslCertificateArgs> SslCertificates
        {
            get => _sslCertificates ?? (_sslCertificates = new InputList<Inputs.ApplicationGatewaySslCertificateArgs>());
            set => _sslCertificates = value;
        }

        /// <summary>
        /// SSL policy of the application gateway resource.
        /// </summary>
        [Input("sslPolicy")]
        public Input<Inputs.ApplicationGatewaySslPolicyArgs>? SslPolicy { get; set; }

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

        [Input("trustedRootCertificates")]
        private InputList<Inputs.ApplicationGatewayTrustedRootCertificateArgs>? _trustedRootCertificates;

        /// <summary>
        /// Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayTrustedRootCertificateArgs> TrustedRootCertificates
        {
            get => _trustedRootCertificates ?? (_trustedRootCertificates = new InputList<Inputs.ApplicationGatewayTrustedRootCertificateArgs>());
            set => _trustedRootCertificates = value;
        }

        [Input("urlPathMaps")]
        private InputList<Inputs.ApplicationGatewayUrlPathMapArgs>? _urlPathMaps;

        /// <summary>
        /// URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
        /// </summary>
        public InputList<Inputs.ApplicationGatewayUrlPathMapArgs> UrlPathMaps
        {
            get => _urlPathMaps ?? (_urlPathMaps = new InputList<Inputs.ApplicationGatewayUrlPathMapArgs>());
            set => _urlPathMaps = value;
        }

        /// <summary>
        /// Web application firewall configuration.
        /// </summary>
        [Input("webApplicationFirewallConfiguration")]
        public Input<Inputs.ApplicationGatewayWebApplicationFirewallConfigurationArgs>? WebApplicationFirewallConfiguration { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// A list of availability zones denoting where the resource needs to come from.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public ApplicationGatewayArgs()
        {
        }
    }
}
