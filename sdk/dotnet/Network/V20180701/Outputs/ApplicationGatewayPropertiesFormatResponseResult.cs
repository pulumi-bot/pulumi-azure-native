// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180701.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayPropertiesFormatResponseResult
    {
        /// <summary>
        /// Authentication certificates of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayAuthenticationCertificateResponseResult> AuthenticationCertificates;
        /// <summary>
        /// Autoscale Configuration.
        /// </summary>
        public readonly Outputs.ApplicationGatewayAutoscaleConfigurationResponseResult? AutoscaleConfiguration;
        /// <summary>
        /// Backend address pool of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayBackendAddressPoolResponseResult> BackendAddressPools;
        /// <summary>
        /// Backend http settings of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayBackendHttpSettingsResponseResult> BackendHttpSettingsCollection;
        /// <summary>
        /// Whether FIPS is enabled on the application gateway resource.
        /// </summary>
        public readonly bool? EnableFips;
        /// <summary>
        /// Whether HTTP2 is enabled on the application gateway resource.
        /// </summary>
        public readonly bool? EnableHttp2;
        /// <summary>
        /// Frontend IP addresses of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayFrontendIPConfigurationResponseResult> FrontendIPConfigurations;
        /// <summary>
        /// Frontend ports of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayFrontendPortResponseResult> FrontendPorts;
        /// <summary>
        /// Subnets of application the gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayIPConfigurationResponseResult> GatewayIPConfigurations;
        /// <summary>
        /// Http listeners of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayHttpListenerResponseResult> HttpListeners;
        /// <summary>
        /// Operational state of the application gateway resource.
        /// </summary>
        public readonly string OperationalState;
        /// <summary>
        /// Probes of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayProbeResponseResult> Probes;
        /// <summary>
        /// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Redirect configurations of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayRedirectConfigurationResponseResult> RedirectConfigurations;
        /// <summary>
        /// Request routing rules of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayRequestRoutingRuleResponseResult> RequestRoutingRules;
        /// <summary>
        /// Resource GUID property of the application gateway resource.
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// SKU of the application gateway resource.
        /// </summary>
        public readonly Outputs.ApplicationGatewaySkuResponseResult? Sku;
        /// <summary>
        /// SSL certificates of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewaySslCertificateResponseResult> SslCertificates;
        /// <summary>
        /// SSL policy of the application gateway resource.
        /// </summary>
        public readonly Outputs.ApplicationGatewaySslPolicyResponseResult? SslPolicy;
        /// <summary>
        /// URL path map of the application gateway resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayUrlPathMapResponseResult> UrlPathMaps;
        /// <summary>
        /// Web application firewall configuration.
        /// </summary>
        public readonly Outputs.ApplicationGatewayWebApplicationFirewallConfigurationResponseResult? WebApplicationFirewallConfiguration;

        [OutputConstructor]
        private ApplicationGatewayPropertiesFormatResponseResult(
            ImmutableArray<Outputs.ApplicationGatewayAuthenticationCertificateResponseResult> authenticationCertificates,

            Outputs.ApplicationGatewayAutoscaleConfigurationResponseResult? autoscaleConfiguration,

            ImmutableArray<Outputs.ApplicationGatewayBackendAddressPoolResponseResult> backendAddressPools,

            ImmutableArray<Outputs.ApplicationGatewayBackendHttpSettingsResponseResult> backendHttpSettingsCollection,

            bool? enableFips,

            bool? enableHttp2,

            ImmutableArray<Outputs.ApplicationGatewayFrontendIPConfigurationResponseResult> frontendIPConfigurations,

            ImmutableArray<Outputs.ApplicationGatewayFrontendPortResponseResult> frontendPorts,

            ImmutableArray<Outputs.ApplicationGatewayIPConfigurationResponseResult> gatewayIPConfigurations,

            ImmutableArray<Outputs.ApplicationGatewayHttpListenerResponseResult> httpListeners,

            string operationalState,

            ImmutableArray<Outputs.ApplicationGatewayProbeResponseResult> probes,

            string? provisioningState,

            ImmutableArray<Outputs.ApplicationGatewayRedirectConfigurationResponseResult> redirectConfigurations,

            ImmutableArray<Outputs.ApplicationGatewayRequestRoutingRuleResponseResult> requestRoutingRules,

            string? resourceGuid,

            Outputs.ApplicationGatewaySkuResponseResult? sku,

            ImmutableArray<Outputs.ApplicationGatewaySslCertificateResponseResult> sslCertificates,

            Outputs.ApplicationGatewaySslPolicyResponseResult? sslPolicy,

            ImmutableArray<Outputs.ApplicationGatewayUrlPathMapResponseResult> urlPathMaps,

            Outputs.ApplicationGatewayWebApplicationFirewallConfigurationResponseResult? webApplicationFirewallConfiguration)
        {
            AuthenticationCertificates = authenticationCertificates;
            AutoscaleConfiguration = autoscaleConfiguration;
            BackendAddressPools = backendAddressPools;
            BackendHttpSettingsCollection = backendHttpSettingsCollection;
            EnableFips = enableFips;
            EnableHttp2 = enableHttp2;
            FrontendIPConfigurations = frontendIPConfigurations;
            FrontendPorts = frontendPorts;
            GatewayIPConfigurations = gatewayIPConfigurations;
            HttpListeners = httpListeners;
            OperationalState = operationalState;
            Probes = probes;
            ProvisioningState = provisioningState;
            RedirectConfigurations = redirectConfigurations;
            RequestRoutingRules = requestRoutingRules;
            ResourceGuid = resourceGuid;
            Sku = sku;
            SslCertificates = sslCertificates;
            SslPolicy = sslPolicy;
            UrlPathMaps = urlPathMaps;
            WebApplicationFirewallConfiguration = webApplicationFirewallConfiguration;
        }
    }
}