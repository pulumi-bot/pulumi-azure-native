// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppConfiguration.Outputs
{

    [OutputType]
    public sealed class ConfigurationStorePropertiesResponse
    {
        /// <summary>
        /// The creation date of configuration store.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The encryption settings of the configuration store.
        /// </summary>
        public readonly Outputs.EncryptionPropertiesResponse? Encryption;
        /// <summary>
        /// The DNS endpoint where the configuration store API will be available.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The list of private endpoint connections that are set up for this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionReferenceResponse> PrivateEndpointConnections;
        /// <summary>
        /// The provisioning state of the configuration store.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
        /// </summary>
        public readonly string? PublicNetworkAccess;

        [OutputConstructor]
        private ConfigurationStorePropertiesResponse(
            string creationDate,

            Outputs.EncryptionPropertiesResponse? encryption,

            string endpoint,

            ImmutableArray<Outputs.PrivateEndpointConnectionReferenceResponse> privateEndpointConnections,

            string provisioningState,

            string? publicNetworkAccess)
        {
            CreationDate = creationDate;
            Encryption = encryption;
            Endpoint = endpoint;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
        }
    }
}