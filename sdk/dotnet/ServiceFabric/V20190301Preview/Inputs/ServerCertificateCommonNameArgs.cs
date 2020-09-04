// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301Preview.Inputs
{

    /// <summary>
    /// Describes the server certificate details using common name.
    /// </summary>
    public sealed class ServerCertificateCommonNameArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The common name of the server certificate.
        /// </summary>
        [Input("certificateCommonName", required: true)]
        public Input<string> CertificateCommonName { get; set; } = null!;

        /// <summary>
        /// The issuer thumbprint of the server certificate.
        /// </summary>
        [Input("certificateIssuerThumbprint", required: true)]
        public Input<string> CertificateIssuerThumbprint { get; set; } = null!;

        public ServerCertificateCommonNameArgs()
        {
        }
    }
}
