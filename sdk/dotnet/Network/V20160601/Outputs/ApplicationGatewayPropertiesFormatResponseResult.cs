// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayPropertiesFormatResponseResult
    {
        /// <summary>
        /// Authentication certificates of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayAuthenticationCertificateResponseResult> AuthenticationCertificates;
        /// <summary>
        /// Backend address pool of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayBackendAddressPoolResponseResult> BackendAddressPools;
        /// <summary>
        /// Backend http settings of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayBackendHttpSettingsResponseResult> BackendHttpSettingsCollection;
        /// <summary>
        /// Frontend IP addresses of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayFrontendIPConfigurationResponseResult> FrontendIPConfigurations;
        /// <summary>
        /// Frontend ports of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayFrontendPortResponseResult> FrontendPorts;
        /// <summary>
        /// Subnets of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayIPConfigurationResponseResult> GatewayIPConfigurations;
        /// <summary>
        /// HTTP listeners of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayHttpListenerResponseResult> HttpListeners;
        /// <summary>
        /// Operational state of application gateway resource
        /// </summary>
        public readonly string OperationalState;
        /// <summary>
        /// Probes of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayProbeResponseResult> Probes;
        /// <summary>
        /// Provisioning state of the ApplicationGateway resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Request routing rules of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayRequestRoutingRuleResponseResult> RequestRoutingRules;
        /// <summary>
        /// Resource guid property of the ApplicationGateway resource
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// Sku of application gateway resource
        /// </summary>
        public readonly Outputs.ApplicationGatewaySkuResponseResult? Sku;
        /// <summary>
        /// SSL certificates of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewaySslCertificateResponseResult> SslCertificates;
        /// <summary>
        /// SSL policy of application gateway resource
        /// </summary>
        public readonly Outputs.ApplicationGatewaySslPolicyResponseResult? SslPolicy;
        /// <summary>
        /// URL path map of application gateway resource
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayUrlPathMapResponseResult> UrlPathMaps;

        [OutputConstructor]
        private ApplicationGatewayPropertiesFormatResponseResult(
            ImmutableArray<Outputs.ApplicationGatewayAuthenticationCertificateResponseResult> authenticationCertificates,

            ImmutableArray<Outputs.ApplicationGatewayBackendAddressPoolResponseResult> backendAddressPools,

            ImmutableArray<Outputs.ApplicationGatewayBackendHttpSettingsResponseResult> backendHttpSettingsCollection,

            ImmutableArray<Outputs.ApplicationGatewayFrontendIPConfigurationResponseResult> frontendIPConfigurations,

            ImmutableArray<Outputs.ApplicationGatewayFrontendPortResponseResult> frontendPorts,

            ImmutableArray<Outputs.ApplicationGatewayIPConfigurationResponseResult> gatewayIPConfigurations,

            ImmutableArray<Outputs.ApplicationGatewayHttpListenerResponseResult> httpListeners,

            string operationalState,

            ImmutableArray<Outputs.ApplicationGatewayProbeResponseResult> probes,

            string? provisioningState,

            ImmutableArray<Outputs.ApplicationGatewayRequestRoutingRuleResponseResult> requestRoutingRules,

            string? resourceGuid,

            Outputs.ApplicationGatewaySkuResponseResult? sku,

            ImmutableArray<Outputs.ApplicationGatewaySslCertificateResponseResult> sslCertificates,

            Outputs.ApplicationGatewaySslPolicyResponseResult? sslPolicy,

            ImmutableArray<Outputs.ApplicationGatewayUrlPathMapResponseResult> urlPathMaps)
        {
            AuthenticationCertificates = authenticationCertificates;
            BackendAddressPools = backendAddressPools;
            BackendHttpSettingsCollection = backendHttpSettingsCollection;
            FrontendIPConfigurations = frontendIPConfigurations;
            FrontendPorts = frontendPorts;
            GatewayIPConfigurations = gatewayIPConfigurations;
            HttpListeners = httpListeners;
            OperationalState = operationalState;
            Probes = probes;
            ProvisioningState = provisioningState;
            RequestRoutingRules = requestRoutingRules;
            ResourceGuid = resourceGuid;
            Sku = sku;
            SslCertificates = sslCertificates;
            SslPolicy = sslPolicy;
            UrlPathMaps = urlPathMaps;
        }
    }
}