// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview.Inputs
{

    /// <summary>
    /// Certificate configuration which consist of non-trusted intermediates and root certificates.
    /// </summary>
    public sealed class CertificateConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate information.
        /// </summary>
        [Input("certificate")]
        public Input<Inputs.CertificateInformationArgs>? Certificate { get; set; }

        /// <summary>
        /// Certificate Password.
        /// </summary>
        [Input("certificatePassword")]
        public Input<string>? CertificatePassword { get; set; }

        /// <summary>
        /// Base64 Encoded certificate.
        /// </summary>
        [Input("encodedCertificate")]
        public Input<string>? EncodedCertificate { get; set; }

        /// <summary>
        /// The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and CertificateAuthority are valid locations.
        /// </summary>
        [Input("storeName", required: true)]
        public Input<string> StoreName { get; set; } = null!;

        public CertificateConfigurationArgs()
        {
        }
    }
}
