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
    public sealed class EdifactEnvelopeSettingsResponseResult
    {
        /// <summary>
        /// The application reference id.
        /// </summary>
        public readonly string? ApplicationReferenceId;
        /// <summary>
        /// The value indicating whether to apply delimiter string advice.
        /// </summary>
        public readonly bool ApplyDelimiterStringAdvice;
        /// <summary>
        /// The communication agreement id.
        /// </summary>
        public readonly string? CommunicationAgreementId;
        /// <summary>
        /// The value indicating whether to create grouping segments.
        /// </summary>
        public readonly bool CreateGroupingSegments;
        /// <summary>
        /// The value indicating whether to enable default group headers.
        /// </summary>
        public readonly bool EnableDefaultGroupHeaders;
        /// <summary>
        /// The functional group id.
        /// </summary>
        public readonly string? FunctionalGroupId;
        /// <summary>
        /// The group application password.
        /// </summary>
        public readonly string? GroupApplicationPassword;
        /// <summary>
        /// The group application receiver id.
        /// </summary>
        public readonly string? GroupApplicationReceiverId;
        /// <summary>
        /// The group application receiver qualifier.
        /// </summary>
        public readonly string? GroupApplicationReceiverQualifier;
        /// <summary>
        /// The group application sender id.
        /// </summary>
        public readonly string? GroupApplicationSenderId;
        /// <summary>
        /// The group application sender qualifier.
        /// </summary>
        public readonly string? GroupApplicationSenderQualifier;
        /// <summary>
        /// The group association assigned code.
        /// </summary>
        public readonly string? GroupAssociationAssignedCode;
        /// <summary>
        /// The group control number lower bound.
        /// </summary>
        public readonly int GroupControlNumberLowerBound;
        /// <summary>
        /// The group control number prefix.
        /// </summary>
        public readonly string? GroupControlNumberPrefix;
        /// <summary>
        /// The group control number suffix.
        /// </summary>
        public readonly string? GroupControlNumberSuffix;
        /// <summary>
        /// The group control number upper bound.
        /// </summary>
        public readonly int GroupControlNumberUpperBound;
        /// <summary>
        /// The group controlling agency code.
        /// </summary>
        public readonly string? GroupControllingAgencyCode;
        /// <summary>
        /// The group message release.
        /// </summary>
        public readonly string? GroupMessageRelease;
        /// <summary>
        /// The group message version.
        /// </summary>
        public readonly string? GroupMessageVersion;
        /// <summary>
        /// The interchange control number lower bound.
        /// </summary>
        public readonly int InterchangeControlNumberLowerBound;
        /// <summary>
        /// The interchange control number prefix.
        /// </summary>
        public readonly string? InterchangeControlNumberPrefix;
        /// <summary>
        /// The interchange control number suffix.
        /// </summary>
        public readonly string? InterchangeControlNumberSuffix;
        /// <summary>
        /// The interchange control number upper bound.
        /// </summary>
        public readonly int InterchangeControlNumberUpperBound;
        /// <summary>
        /// The value indicating whether the message is a test interchange.
        /// </summary>
        public readonly bool IsTestInterchange;
        /// <summary>
        /// The value indicating whether to overwrite existing transaction set control number.
        /// </summary>
        public readonly bool OverwriteExistingTransactionSetControlNumber;
        /// <summary>
        /// The processing priority code.
        /// </summary>
        public readonly string? ProcessingPriorityCode;
        /// <summary>
        /// The receiver internal identification.
        /// </summary>
        public readonly string? ReceiverInternalIdentification;
        /// <summary>
        /// The receiver internal sub identification.
        /// </summary>
        public readonly string? ReceiverInternalSubIdentification;
        /// <summary>
        /// The receiver reverse routing address.
        /// </summary>
        public readonly string? ReceiverReverseRoutingAddress;
        /// <summary>
        /// The recipient reference password qualifier.
        /// </summary>
        public readonly string? RecipientReferencePasswordQualifier;
        /// <summary>
        /// The recipient reference password value.
        /// </summary>
        public readonly string? RecipientReferencePasswordValue;
        /// <summary>
        /// The value indicating whether to rollover group control number.
        /// </summary>
        public readonly bool RolloverGroupControlNumber;
        /// <summary>
        /// The value indicating whether to rollover interchange control number.
        /// </summary>
        public readonly bool RolloverInterchangeControlNumber;
        /// <summary>
        /// The value indicating whether to rollover transaction set control number.
        /// </summary>
        public readonly bool RolloverTransactionSetControlNumber;
        /// <summary>
        /// The sender internal identification.
        /// </summary>
        public readonly string? SenderInternalIdentification;
        /// <summary>
        /// The sender internal sub identification.
        /// </summary>
        public readonly string? SenderInternalSubIdentification;
        /// <summary>
        /// The sender reverse routing address.
        /// </summary>
        public readonly string? SenderReverseRoutingAddress;
        /// <summary>
        /// The transaction set control number lower bound.
        /// </summary>
        public readonly int TransactionSetControlNumberLowerBound;
        /// <summary>
        /// The transaction set control number prefix.
        /// </summary>
        public readonly string? TransactionSetControlNumberPrefix;
        /// <summary>
        /// The transaction set control number suffix.
        /// </summary>
        public readonly string? TransactionSetControlNumberSuffix;
        /// <summary>
        /// The transaction set control number upper bound.
        /// </summary>
        public readonly int TransactionSetControlNumberUpperBound;

        [OutputConstructor]
        private EdifactEnvelopeSettingsResponseResult(
            string? applicationReferenceId,

            bool applyDelimiterStringAdvice,

            string? communicationAgreementId,

            bool createGroupingSegments,

            bool enableDefaultGroupHeaders,

            string? functionalGroupId,

            string? groupApplicationPassword,

            string? groupApplicationReceiverId,

            string? groupApplicationReceiverQualifier,

            string? groupApplicationSenderId,

            string? groupApplicationSenderQualifier,

            string? groupAssociationAssignedCode,

            int groupControlNumberLowerBound,

            string? groupControlNumberPrefix,

            string? groupControlNumberSuffix,

            int groupControlNumberUpperBound,

            string? groupControllingAgencyCode,

            string? groupMessageRelease,

            string? groupMessageVersion,

            int interchangeControlNumberLowerBound,

            string? interchangeControlNumberPrefix,

            string? interchangeControlNumberSuffix,

            int interchangeControlNumberUpperBound,

            bool isTestInterchange,

            bool overwriteExistingTransactionSetControlNumber,

            string? processingPriorityCode,

            string? receiverInternalIdentification,

            string? receiverInternalSubIdentification,

            string? receiverReverseRoutingAddress,

            string? recipientReferencePasswordQualifier,

            string? recipientReferencePasswordValue,

            bool rolloverGroupControlNumber,

            bool rolloverInterchangeControlNumber,

            bool rolloverTransactionSetControlNumber,

            string? senderInternalIdentification,

            string? senderInternalSubIdentification,

            string? senderReverseRoutingAddress,

            int transactionSetControlNumberLowerBound,

            string? transactionSetControlNumberPrefix,

            string? transactionSetControlNumberSuffix,

            int transactionSetControlNumberUpperBound)
        {
            ApplicationReferenceId = applicationReferenceId;
            ApplyDelimiterStringAdvice = applyDelimiterStringAdvice;
            CommunicationAgreementId = communicationAgreementId;
            CreateGroupingSegments = createGroupingSegments;
            EnableDefaultGroupHeaders = enableDefaultGroupHeaders;
            FunctionalGroupId = functionalGroupId;
            GroupApplicationPassword = groupApplicationPassword;
            GroupApplicationReceiverId = groupApplicationReceiverId;
            GroupApplicationReceiverQualifier = groupApplicationReceiverQualifier;
            GroupApplicationSenderId = groupApplicationSenderId;
            GroupApplicationSenderQualifier = groupApplicationSenderQualifier;
            GroupAssociationAssignedCode = groupAssociationAssignedCode;
            GroupControlNumberLowerBound = groupControlNumberLowerBound;
            GroupControlNumberPrefix = groupControlNumberPrefix;
            GroupControlNumberSuffix = groupControlNumberSuffix;
            GroupControlNumberUpperBound = groupControlNumberUpperBound;
            GroupControllingAgencyCode = groupControllingAgencyCode;
            GroupMessageRelease = groupMessageRelease;
            GroupMessageVersion = groupMessageVersion;
            InterchangeControlNumberLowerBound = interchangeControlNumberLowerBound;
            InterchangeControlNumberPrefix = interchangeControlNumberPrefix;
            InterchangeControlNumberSuffix = interchangeControlNumberSuffix;
            InterchangeControlNumberUpperBound = interchangeControlNumberUpperBound;
            IsTestInterchange = isTestInterchange;
            OverwriteExistingTransactionSetControlNumber = overwriteExistingTransactionSetControlNumber;
            ProcessingPriorityCode = processingPriorityCode;
            ReceiverInternalIdentification = receiverInternalIdentification;
            ReceiverInternalSubIdentification = receiverInternalSubIdentification;
            ReceiverReverseRoutingAddress = receiverReverseRoutingAddress;
            RecipientReferencePasswordQualifier = recipientReferencePasswordQualifier;
            RecipientReferencePasswordValue = recipientReferencePasswordValue;
            RolloverGroupControlNumber = rolloverGroupControlNumber;
            RolloverInterchangeControlNumber = rolloverInterchangeControlNumber;
            RolloverTransactionSetControlNumber = rolloverTransactionSetControlNumber;
            SenderInternalIdentification = senderInternalIdentification;
            SenderInternalSubIdentification = senderInternalSubIdentification;
            SenderReverseRoutingAddress = senderReverseRoutingAddress;
            TransactionSetControlNumberLowerBound = transactionSetControlNumberLowerBound;
            TransactionSetControlNumberPrefix = transactionSetControlNumberPrefix;
            TransactionSetControlNumberSuffix = transactionSetControlNumberSuffix;
            TransactionSetControlNumberUpperBound = transactionSetControlNumberUpperBound;
        }
    }
}
