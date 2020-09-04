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
    public sealed class X12ValidationSettingsResponseResult
    {
        /// <summary>
        /// The value indicating whether to allow leading and trailing spaces and zeroes.
        /// </summary>
        public readonly bool AllowLeadingAndTrailingSpacesAndZeroes;
        /// <summary>
        /// The value indicating whether to check for duplicate group control number.
        /// </summary>
        public readonly bool CheckDuplicateGroupControlNumber;
        /// <summary>
        /// The value indicating whether to check for duplicate interchange control number.
        /// </summary>
        public readonly bool CheckDuplicateInterchangeControlNumber;
        /// <summary>
        /// The value indicating whether to check for duplicate transaction set control number.
        /// </summary>
        public readonly bool CheckDuplicateTransactionSetControlNumber;
        /// <summary>
        /// The validity period of interchange control number.
        /// </summary>
        public readonly int InterchangeControlNumberValidityDays;
        /// <summary>
        /// The trailing separator policy.
        /// </summary>
        public readonly string TrailingSeparatorPolicy;
        /// <summary>
        /// The value indicating whether to trim leading and trailing spaces and zeroes.
        /// </summary>
        public readonly bool TrimLeadingAndTrailingSpacesAndZeroes;
        /// <summary>
        /// The value indicating whether to validate character set in the message.
        /// </summary>
        public readonly bool ValidateCharacterSet;
        /// <summary>
        /// The value indicating whether to Whether to validate EDI types.
        /// </summary>
        public readonly bool ValidateEDITypes;
        /// <summary>
        /// The value indicating whether to Whether to validate XSD types.
        /// </summary>
        public readonly bool ValidateXSDTypes;

        [OutputConstructor]
        private X12ValidationSettingsResponseResult(
            bool allowLeadingAndTrailingSpacesAndZeroes,

            bool checkDuplicateGroupControlNumber,

            bool checkDuplicateInterchangeControlNumber,

            bool checkDuplicateTransactionSetControlNumber,

            int interchangeControlNumberValidityDays,

            string trailingSeparatorPolicy,

            bool trimLeadingAndTrailingSpacesAndZeroes,

            bool validateCharacterSet,

            bool validateEDITypes,

            bool validateXSDTypes)
        {
            AllowLeadingAndTrailingSpacesAndZeroes = allowLeadingAndTrailingSpacesAndZeroes;
            CheckDuplicateGroupControlNumber = checkDuplicateGroupControlNumber;
            CheckDuplicateInterchangeControlNumber = checkDuplicateInterchangeControlNumber;
            CheckDuplicateTransactionSetControlNumber = checkDuplicateTransactionSetControlNumber;
            InterchangeControlNumberValidityDays = interchangeControlNumberValidityDays;
            TrailingSeparatorPolicy = trailingSeparatorPolicy;
            TrimLeadingAndTrailingSpacesAndZeroes = trimLeadingAndTrailingSpacesAndZeroes;
            ValidateCharacterSet = validateCharacterSet;
            ValidateEDITypes = validateEDITypes;
            ValidateXSDTypes = validateXSDTypes;
        }
    }
}
