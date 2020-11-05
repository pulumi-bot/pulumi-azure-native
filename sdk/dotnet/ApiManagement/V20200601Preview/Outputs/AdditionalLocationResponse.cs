// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ApiManagement.V20200601Preview.Outputs
{

    [OutputType]
    public sealed class AdditionalLocationResponse
    {
        /// <summary>
        /// Property only valid for an Api Management service deployed in multiple locations. This can be used to disable the gateway in this additional location.
        /// </summary>
        public readonly bool? DisableGateway;
        /// <summary>
        /// Gateway URL of the API Management service in the Region.
        /// </summary>
        public readonly string GatewayRegionalUrl;
        /// <summary>
        /// The location name of the additional region among Azure Data center regions.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Private Static Load Balanced IP addresses of the API Management service which is deployed in an Internal Virtual Network in a particular additional location. Available only for Basic, Standard, Premium and Isolated SKU.
        /// </summary>
        public readonly ImmutableArray<string> PrivateIPAddresses;
        /// <summary>
        /// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard, Premium and Isolated SKU.
        /// </summary>
        public readonly ImmutableArray<string> PublicIPAddresses;
        /// <summary>
        /// SKU properties of the API Management service.
        /// </summary>
        public readonly Outputs.ApiManagementServiceSkuPropertiesResponse Sku;
        /// <summary>
        /// Virtual network configuration for the location.
        /// </summary>
        public readonly Outputs.VirtualNetworkConfigurationResponse? VirtualNetworkConfiguration;
        /// <summary>
        /// A list of availability zones denoting where the resource needs to come from.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private AdditionalLocationResponse(
            bool? disableGateway,

            string gatewayRegionalUrl,

            string location,

            ImmutableArray<string> privateIPAddresses,

            ImmutableArray<string> publicIPAddresses,

            Outputs.ApiManagementServiceSkuPropertiesResponse sku,

            Outputs.VirtualNetworkConfigurationResponse? virtualNetworkConfiguration,

            ImmutableArray<string> zones)
        {
            DisableGateway = disableGateway;
            GatewayRegionalUrl = gatewayRegionalUrl;
            Location = location;
            PrivateIPAddresses = privateIPAddresses;
            PublicIPAddresses = publicIPAddresses;
            Sku = sku;
            VirtualNetworkConfiguration = virtualNetworkConfiguration;
            Zones = zones;
        }
    }
}
