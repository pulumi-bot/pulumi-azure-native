// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20180201.Outputs
{

    [OutputType]
    public sealed class ClientCertificateThumbprintResponseResult
    {
        /// <summary>
        /// The thumbprint of the client certificate.
        /// </summary>
        public readonly string CertificateThumbprint;
        /// <summary>
        /// Indicates if the client certificate has admin access to the cluster. Non admin clients can perform only read only operations on the cluster.
        /// </summary>
        public readonly bool IsAdmin;

        [OutputConstructor]
        private ClientCertificateThumbprintResponseResult(
            string certificateThumbprint,

            bool isAdmin)
        {
            CertificateThumbprint = certificateThumbprint;
            IsAdmin = isAdmin;
        }
    }
}