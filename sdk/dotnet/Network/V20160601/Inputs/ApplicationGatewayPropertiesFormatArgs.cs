// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601.Inputs
{

    /// <summary>
    /// Properties of Application Gateway
    /// </summary>
    public sealed class ApplicationGatewayPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationCertificates")]
        private InputList<Inputs.ApplicationGatewayAuthenticationCertificateArgs>? _authenticationCertificates;

        /// <summary>
        /// Authentication certificates of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayAuthenticationCertificateArgs> AuthenticationCertificates
        {
            get => _authenticationCertificates ?? (_authenticationCertificates = new InputList<Inputs.ApplicationGatewayAuthenticationCertificateArgs>());
            set => _authenticationCertificates = value;
        }

        [Input("backendAddressPools")]
        private InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs>? _backendAddressPools;

        /// <summary>
        /// Backend address pool of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs> BackendAddressPools
        {
            get => _backendAddressPools ?? (_backendAddressPools = new InputList<Inputs.ApplicationGatewayBackendAddressPoolArgs>());
            set => _backendAddressPools = value;
        }

        [Input("backendHttpSettingsCollection")]
        private InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs>? _backendHttpSettingsCollection;

        /// <summary>
        /// Backend http settings of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs> BackendHttpSettingsCollection
        {
            get => _backendHttpSettingsCollection ?? (_backendHttpSettingsCollection = new InputList<Inputs.ApplicationGatewayBackendHttpSettingsArgs>());
            set => _backendHttpSettingsCollection = value;
        }

        [Input("frontendIPConfigurations")]
        private InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs>? _frontendIPConfigurations;

        /// <summary>
        /// Frontend IP addresses of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs> FrontendIPConfigurations
        {
            get => _frontendIPConfigurations ?? (_frontendIPConfigurations = new InputList<Inputs.ApplicationGatewayFrontendIPConfigurationArgs>());
            set => _frontendIPConfigurations = value;
        }

        [Input("frontendPorts")]
        private InputList<Inputs.ApplicationGatewayFrontendPortArgs>? _frontendPorts;

        /// <summary>
        /// Frontend ports of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayFrontendPortArgs> FrontendPorts
        {
            get => _frontendPorts ?? (_frontendPorts = new InputList<Inputs.ApplicationGatewayFrontendPortArgs>());
            set => _frontendPorts = value;
        }

        [Input("gatewayIPConfigurations")]
        private InputList<Inputs.ApplicationGatewayIPConfigurationArgs>? _gatewayIPConfigurations;

        /// <summary>
        /// Subnets of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayIPConfigurationArgs> GatewayIPConfigurations
        {
            get => _gatewayIPConfigurations ?? (_gatewayIPConfigurations = new InputList<Inputs.ApplicationGatewayIPConfigurationArgs>());
            set => _gatewayIPConfigurations = value;
        }

        [Input("httpListeners")]
        private InputList<Inputs.ApplicationGatewayHttpListenerArgs>? _httpListeners;

        /// <summary>
        /// HTTP listeners of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayHttpListenerArgs> HttpListeners
        {
            get => _httpListeners ?? (_httpListeners = new InputList<Inputs.ApplicationGatewayHttpListenerArgs>());
            set => _httpListeners = value;
        }

        [Input("probes")]
        private InputList<Inputs.ApplicationGatewayProbeArgs>? _probes;

        /// <summary>
        /// Probes of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayProbeArgs> Probes
        {
            get => _probes ?? (_probes = new InputList<Inputs.ApplicationGatewayProbeArgs>());
            set => _probes = value;
        }

        /// <summary>
        /// Provisioning state of the ApplicationGateway resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        [Input("requestRoutingRules")]
        private InputList<Inputs.ApplicationGatewayRequestRoutingRuleArgs>? _requestRoutingRules;

        /// <summary>
        /// Request routing rules of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayRequestRoutingRuleArgs> RequestRoutingRules
        {
            get => _requestRoutingRules ?? (_requestRoutingRules = new InputList<Inputs.ApplicationGatewayRequestRoutingRuleArgs>());
            set => _requestRoutingRules = value;
        }

        /// <summary>
        /// Resource guid property of the ApplicationGateway resource
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        /// <summary>
        /// Sku of application gateway resource
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ApplicationGatewaySkuArgs>? Sku { get; set; }

        [Input("sslCertificates")]
        private InputList<Inputs.ApplicationGatewaySslCertificateArgs>? _sslCertificates;

        /// <summary>
        /// SSL certificates of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewaySslCertificateArgs> SslCertificates
        {
            get => _sslCertificates ?? (_sslCertificates = new InputList<Inputs.ApplicationGatewaySslCertificateArgs>());
            set => _sslCertificates = value;
        }

        /// <summary>
        /// SSL policy of application gateway resource
        /// </summary>
        [Input("sslPolicy")]
        public Input<Inputs.ApplicationGatewaySslPolicyArgs>? SslPolicy { get; set; }

        [Input("urlPathMaps")]
        private InputList<Inputs.ApplicationGatewayUrlPathMapArgs>? _urlPathMaps;

        /// <summary>
        /// URL path map of application gateway resource
        /// </summary>
        public InputList<Inputs.ApplicationGatewayUrlPathMapArgs> UrlPathMaps
        {
            get => _urlPathMaps ?? (_urlPathMaps = new InputList<Inputs.ApplicationGatewayUrlPathMapArgs>());
            set => _urlPathMaps = value;
        }

        public ApplicationGatewayPropertiesFormatArgs()
        {
        }
    }
}