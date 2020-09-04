// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20190701Preview.Outputs
{

    [OutputType]
    public sealed class CertificatePropertiesResponseResult
    {
        /// <summary>
        /// The certificate content
        /// </summary>
        public readonly string? Certificate;
        /// <summary>
        /// The certificate's create date and time.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// The certificate's expiration date and time.
        /// </summary>
        public readonly string Expiry;
        /// <summary>
        /// Determines whether certificate has been verified.
        /// </summary>
        public readonly bool IsVerified;
        /// <summary>
        /// The certificate's subject name.
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// The certificate's thumbprint.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// The certificate's last update date and time.
        /// </summary>
        public readonly string Updated;

        [OutputConstructor]
        private CertificatePropertiesResponseResult(
            string? certificate,

            string created,

            string expiry,

            bool isVerified,

            string subject,

            string thumbprint,

            string updated)
        {
            Certificate = certificate;
            Created = created;
            Expiry = expiry;
            IsVerified = isVerified;
            Subject = subject;
            Thumbprint = thumbprint;
            Updated = updated;
        }
    }
}
