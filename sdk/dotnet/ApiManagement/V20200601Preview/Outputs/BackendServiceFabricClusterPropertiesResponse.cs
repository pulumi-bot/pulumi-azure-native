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
    public sealed class BackendServiceFabricClusterPropertiesResponse
    {
        /// <summary>
        /// The client certificate thumbprint for the management endpoint.
        /// </summary>
        public readonly string ClientCertificatethumbprint;
        /// <summary>
        /// The cluster management endpoint.
        /// </summary>
        public readonly ImmutableArray<string> ManagementEndpoints;
        /// <summary>
        /// Maximum number of retries while attempting resolve the partition.
        /// </summary>
        public readonly int? MaxPartitionResolutionRetries;
        /// <summary>
        /// Thumbprints of certificates cluster management service uses for tls communication
        /// </summary>
        public readonly ImmutableArray<string> ServerCertificateThumbprints;
        /// <summary>
        /// Server X509 Certificate Names Collection
        /// </summary>
        public readonly ImmutableArray<Outputs.X509CertificateNameResponse> ServerX509Names;

        [OutputConstructor]
        private BackendServiceFabricClusterPropertiesResponse(
            string clientCertificatethumbprint,

            ImmutableArray<string> managementEndpoints,

            int? maxPartitionResolutionRetries,

            ImmutableArray<string> serverCertificateThumbprints,

            ImmutableArray<Outputs.X509CertificateNameResponse> serverX509Names)
        {
            ClientCertificatethumbprint = clientCertificatethumbprint;
            ManagementEndpoints = managementEndpoints;
            MaxPartitionResolutionRetries = maxPartitionResolutionRetries;
            ServerCertificateThumbprints = serverCertificateThumbprints;
            ServerX509Names = serverX509Names;
        }
    }
}
