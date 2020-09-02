// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20150801Preview.Outputs
{

    [OutputType]
    public sealed class X12EnvelopeOverrideResponseResult
    {
        /// <summary>
        /// The date format.
        /// </summary>
        public readonly string? DateFormat;
        /// <summary>
        /// The functional identifier code.
        /// </summary>
        public readonly string? FunctionalIdentifierCode;
        /// <summary>
        /// The header version.
        /// </summary>
        public readonly string? HeaderVersion;
        /// <summary>
        /// The message id on which this envelope settings has to be applied.
        /// </summary>
        public readonly string? MessageId;
        /// <summary>
        /// The protocol version on which this envelope settings has to be applied.
        /// </summary>
        public readonly string? ProtocolVersion;
        /// <summary>
        /// The receiver application id.
        /// </summary>
        public readonly string? ReceiverApplicationId;
        /// <summary>
        /// The responsible agency code.
        /// </summary>
        public readonly int? ResponsibleAgencyCode;
        /// <summary>
        /// The sender application id.
        /// </summary>
        public readonly string? SenderApplicationId;
        /// <summary>
        /// The target namespace on which this envelope settings has to be applied.
        /// </summary>
        public readonly string? TargetNamespace;
        /// <summary>
        /// The time format.
        /// </summary>
        public readonly string? TimeFormat;

        [OutputConstructor]
        private X12EnvelopeOverrideResponseResult(
            string? dateFormat,

            string? functionalIdentifierCode,

            string? headerVersion,

            string? messageId,

            string? protocolVersion,

            string? receiverApplicationId,

            int? responsibleAgencyCode,

            string? senderApplicationId,

            string? targetNamespace,

            string? timeFormat)
        {
            DateFormat = dateFormat;
            FunctionalIdentifierCode = functionalIdentifierCode;
            HeaderVersion = headerVersion;
            MessageId = messageId;
            ProtocolVersion = protocolVersion;
            ReceiverApplicationId = receiverApplicationId;
            ResponsibleAgencyCode = responsibleAgencyCode;
            SenderApplicationId = senderApplicationId;
            TargetNamespace = targetNamespace;
            TimeFormat = timeFormat;
        }
    }
}
