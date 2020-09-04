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
    public sealed class EdifactAcknowledgementSettingsResponseResult
    {
        /// <summary>
        /// The acknowledgement control number lower bound.
        /// </summary>
        public readonly int AcknowledgementControlNumberLowerBound;
        /// <summary>
        /// The acknowledgement control number prefix.
        /// </summary>
        public readonly string? AcknowledgementControlNumberPrefix;
        /// <summary>
        /// The acknowledgement control number suffix.
        /// </summary>
        public readonly string? AcknowledgementControlNumberSuffix;
        /// <summary>
        /// The acknowledgement control number upper bound.
        /// </summary>
        public readonly int AcknowledgementControlNumberUpperBound;
        /// <summary>
        /// The value indicating whether to batch functional acknowledgements.
        /// </summary>
        public readonly bool BatchFunctionalAcknowledgements;
        /// <summary>
        /// The value indicating whether to batch the technical acknowledgements.
        /// </summary>
        public readonly bool BatchTechnicalAcknowledgements;
        /// <summary>
        /// The value indicating whether functional acknowledgement is needed.
        /// </summary>
        public readonly bool NeedFunctionalAcknowledgement;
        /// <summary>
        /// The value indicating whether a loop is needed for valid messages.
        /// </summary>
        public readonly bool NeedLoopForValidMessages;
        /// <summary>
        /// The value indicating whether technical acknowledgement is needed.
        /// </summary>
        public readonly bool NeedTechnicalAcknowledgement;
        /// <summary>
        /// The value indicating whether to rollover acknowledgement control number.
        /// </summary>
        public readonly bool RolloverAcknowledgementControlNumber;
        /// <summary>
        /// The value indicating whether to send synchronous acknowledgement.
        /// </summary>
        public readonly bool SendSynchronousAcknowledgement;

        [OutputConstructor]
        private EdifactAcknowledgementSettingsResponseResult(
            int acknowledgementControlNumberLowerBound,

            string? acknowledgementControlNumberPrefix,

            string? acknowledgementControlNumberSuffix,

            int acknowledgementControlNumberUpperBound,

            bool batchFunctionalAcknowledgements,

            bool batchTechnicalAcknowledgements,

            bool needFunctionalAcknowledgement,

            bool needLoopForValidMessages,

            bool needTechnicalAcknowledgement,

            bool rolloverAcknowledgementControlNumber,

            bool sendSynchronousAcknowledgement)
        {
            AcknowledgementControlNumberLowerBound = acknowledgementControlNumberLowerBound;
            AcknowledgementControlNumberPrefix = acknowledgementControlNumberPrefix;
            AcknowledgementControlNumberSuffix = acknowledgementControlNumberSuffix;
            AcknowledgementControlNumberUpperBound = acknowledgementControlNumberUpperBound;
            BatchFunctionalAcknowledgements = batchFunctionalAcknowledgements;
            BatchTechnicalAcknowledgements = batchTechnicalAcknowledgements;
            NeedFunctionalAcknowledgement = needFunctionalAcknowledgement;
            NeedLoopForValidMessages = needLoopForValidMessages;
            NeedTechnicalAcknowledgement = needTechnicalAcknowledgement;
            RolloverAcknowledgementControlNumber = rolloverAcknowledgementControlNumber;
            SendSynchronousAcknowledgement = sendSynchronousAcknowledgement;
        }
    }
}
