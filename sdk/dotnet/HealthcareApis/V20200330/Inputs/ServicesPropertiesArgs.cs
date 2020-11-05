// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HealthcareApis.V20200330.Inputs
{

    /// <summary>
    /// The properties of a service instance.
    /// </summary>
    public sealed class ServicesPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("accessPolicies")]
        private InputList<Inputs.ServiceAccessPolicyEntryArgs>? _accessPolicies;

        /// <summary>
        /// The access policies of the service instance.
        /// </summary>
        public InputList<Inputs.ServiceAccessPolicyEntryArgs> AccessPolicies
        {
            get => _accessPolicies ?? (_accessPolicies = new InputList<Inputs.ServiceAccessPolicyEntryArgs>());
            set => _accessPolicies = value;
        }

        /// <summary>
        /// The authentication configuration for the service instance.
        /// </summary>
        [Input("authenticationConfiguration")]
        public Input<Inputs.ServiceAuthenticationConfigurationInfoArgs>? AuthenticationConfiguration { get; set; }

        /// <summary>
        /// The settings for the CORS configuration of the service instance.
        /// </summary>
        [Input("corsConfiguration")]
        public Input<Inputs.ServiceCorsConfigurationInfoArgs>? CorsConfiguration { get; set; }

        /// <summary>
        /// The settings for the Cosmos DB database backing the service.
        /// </summary>
        [Input("cosmosDbConfiguration")]
        public Input<Inputs.ServiceCosmosDbConfigurationInfoArgs>? CosmosDbConfiguration { get; set; }

        /// <summary>
        /// The settings for the export operation of the service instance.
        /// </summary>
        [Input("exportConfiguration")]
        public Input<Inputs.ServiceExportConfigurationInfoArgs>? ExportConfiguration { get; set; }

        [Input("privateEndpointConnections")]
        private InputList<Inputs.PrivateEndpointConnectionArgs>? _privateEndpointConnections;

        /// <summary>
        /// The list of private endpoint connections that are set up for this resource.
        /// </summary>
        public InputList<Inputs.PrivateEndpointConnectionArgs> PrivateEndpointConnections
        {
            get => _privateEndpointConnections ?? (_privateEndpointConnections = new InputList<Inputs.PrivateEndpointConnectionArgs>());
            set => _privateEndpointConnections = value;
        }

        /// <summary>
        /// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
        /// </summary>
        [Input("publicNetworkAccess")]
        public Input<string>? PublicNetworkAccess { get; set; }

        public ServicesPropertiesArgs()
        {
        }
    }
}
