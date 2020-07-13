// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.Outputs
{

    [OutputType]
    public sealed class GatewayHostnameConfigurationContractPropertiesResponseResult
    {
        /// <summary>
        /// Identifier of Certificate entity that will be used for TLS connection establishment
        /// </summary>
        public readonly string? CertificateId;
        /// <summary>
        /// Hostname value. Supports valid domain name, partial or full wildcard
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// Determines whether gateway requests client certificate
        /// </summary>
        public readonly bool? NegotiateClientCertificate;

        [OutputConstructor]
        private GatewayHostnameConfigurationContractPropertiesResponseResult(
            string? certificateId,

            string? hostname,

            bool? negotiateClientCertificate)
        {
            CertificateId = certificateId;
            Hostname = hostname;
            NegotiateClientCertificate = negotiateClientCertificate;
        }
    }
}