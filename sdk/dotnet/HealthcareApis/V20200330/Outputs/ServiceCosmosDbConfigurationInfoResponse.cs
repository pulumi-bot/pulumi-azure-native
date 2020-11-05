// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HealthcareApis.V20200330.Outputs
{

    [OutputType]
    public sealed class ServiceCosmosDbConfigurationInfoResponse
    {
        /// <summary>
        /// The URI of the customer-managed key for the backing database.
        /// </summary>
        public readonly string? KeyVaultKeyUri;
        /// <summary>
        /// The provisioned throughput for the backing database.
        /// </summary>
        public readonly int? OfferThroughput;

        [OutputConstructor]
        private ServiceCosmosDbConfigurationInfoResponse(
            string? keyVaultKeyUri,

            int? offerThroughput)
        {
            KeyVaultKeyUri = keyVaultKeyUri;
            OfferThroughput = offerThroughput;
        }
    }
}
