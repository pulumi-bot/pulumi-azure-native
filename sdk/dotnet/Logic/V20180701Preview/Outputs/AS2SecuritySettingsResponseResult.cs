// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class AS2SecuritySettingsResponseResult
    {
        /// <summary>
        /// The value indicating whether to enable NRR for inbound decoded messages.
        /// </summary>
        public readonly bool EnableNRRForInboundDecodedMessages;
        /// <summary>
        /// The value indicating whether to enable NRR for inbound encoded messages.
        /// </summary>
        public readonly bool EnableNRRForInboundEncodedMessages;
        /// <summary>
        /// The value indicating whether to enable NRR for inbound MDN.
        /// </summary>
        public readonly bool EnableNRRForInboundMDN;
        /// <summary>
        /// The value indicating whether to enable NRR for outbound decoded messages.
        /// </summary>
        public readonly bool EnableNRRForOutboundDecodedMessages;
        /// <summary>
        /// The value indicating whether to enable NRR for outbound encoded messages.
        /// </summary>
        public readonly bool EnableNRRForOutboundEncodedMessages;
        /// <summary>
        /// The value indicating whether to enable NRR for outbound MDN.
        /// </summary>
        public readonly bool EnableNRRForOutboundMDN;
        /// <summary>
        /// The name of the encryption certificate.
        /// </summary>
        public readonly string? EncryptionCertificateName;
        /// <summary>
        /// The value indicating whether to send or request a MDN.
        /// </summary>
        public readonly bool OverrideGroupSigningCertificate;
        /// <summary>
        /// The Sha2 algorithm format. Valid values are Sha2, ShaHashSize, ShaHyphenHashSize, Sha2UnderscoreHashSize.
        /// </summary>
        public readonly string? Sha2AlgorithmFormat;
        /// <summary>
        /// The name of the signing certificate.
        /// </summary>
        public readonly string? SigningCertificateName;

        [OutputConstructor]
        private AS2SecuritySettingsResponseResult(
            bool enableNRRForInboundDecodedMessages,

            bool enableNRRForInboundEncodedMessages,

            bool enableNRRForInboundMDN,

            bool enableNRRForOutboundDecodedMessages,

            bool enableNRRForOutboundEncodedMessages,

            bool enableNRRForOutboundMDN,

            string? encryptionCertificateName,

            bool overrideGroupSigningCertificate,

            string? sha2AlgorithmFormat,

            string? signingCertificateName)
        {
            EnableNRRForInboundDecodedMessages = enableNRRForInboundDecodedMessages;
            EnableNRRForInboundEncodedMessages = enableNRRForInboundEncodedMessages;
            EnableNRRForInboundMDN = enableNRRForInboundMDN;
            EnableNRRForOutboundDecodedMessages = enableNRRForOutboundDecodedMessages;
            EnableNRRForOutboundEncodedMessages = enableNRRForOutboundEncodedMessages;
            EnableNRRForOutboundMDN = enableNRRForOutboundMDN;
            EncryptionCertificateName = encryptionCertificateName;
            OverrideGroupSigningCertificate = overrideGroupSigningCertificate;
            Sha2AlgorithmFormat = sha2AlgorithmFormat;
            SigningCertificateName = signingCertificateName;
        }
    }
}
