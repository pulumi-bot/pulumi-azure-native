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
    public sealed class CertificateConfigurationResponse
    {
        /// <summary>
        /// Certificate information.
        /// </summary>
        public readonly Outputs.CertificateInformationResponse? Certificate;
        /// <summary>
        /// Certificate Password.
        /// </summary>
        public readonly string? CertificatePassword;
        /// <summary>
        /// Base64 Encoded certificate.
        /// </summary>
        public readonly string? EncodedCertificate;
        /// <summary>
        /// The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and CertificateAuthority are valid locations.
        /// </summary>
        public readonly string StoreName;

        [OutputConstructor]
        private CertificateConfigurationResponse(
            Outputs.CertificateInformationResponse? certificate,

            string? certificatePassword,

            string? encodedCertificate,

            string storeName)
        {
            Certificate = certificate;
            CertificatePassword = certificatePassword;
            EncodedCertificate = encodedCertificate;
            StoreName = storeName;
        }
    }
}
