// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview.Outputs
{

    [OutputType]
    public sealed class BackendTlsPropertiesResponseResult
    {
        /// <summary>
        /// Flag indicating whether SSL certificate chain validation should be done when using self-signed certificates for this backend host.
        /// </summary>
        public readonly bool? ValidateCertificateChain;
        /// <summary>
        /// Flag indicating whether SSL certificate name validation should be done when using self-signed certificates for this backend host.
        /// </summary>
        public readonly bool? ValidateCertificateName;

        [OutputConstructor]
        private BackendTlsPropertiesResponseResult(
            bool? validateCertificateChain,

            bool? validateCertificateName)
        {
            ValidateCertificateChain = validateCertificateChain;
            ValidateCertificateName = validateCertificateName;
        }
    }
}
