// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190601Preview.Outputs
{

    [OutputType]
    public sealed class ServerCertificateCommonNameResponseResult
    {
        /// <summary>
        /// The common name of the server certificate.
        /// </summary>
        public readonly string CertificateCommonName;
        /// <summary>
        /// The issuer thumbprint of the server certificate.
        /// </summary>
        public readonly string CertificateIssuerThumbprint;

        [OutputConstructor]
        private ServerCertificateCommonNameResponseResult(
            string certificateCommonName,

            string certificateIssuerThumbprint)
        {
            CertificateCommonName = certificateCommonName;
            CertificateIssuerThumbprint = certificateIssuerThumbprint;
        }
    }
}
