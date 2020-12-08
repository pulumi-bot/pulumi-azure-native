// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.Latest.Inputs
{

    /// <summary>
    /// The X12 agreement validation settings.
    /// </summary>
    public sealed class X12ValidationSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value indicating whether to allow leading and trailing spaces and zeroes.
        /// </summary>
        [Input("allowLeadingAndTrailingSpacesAndZeroes", required: true)]
        public Input<bool> AllowLeadingAndTrailingSpacesAndZeroes { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to check for duplicate group control number.
        /// </summary>
        [Input("checkDuplicateGroupControlNumber", required: true)]
        public Input<bool> CheckDuplicateGroupControlNumber { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to check for duplicate interchange control number.
        /// </summary>
        [Input("checkDuplicateInterchangeControlNumber", required: true)]
        public Input<bool> CheckDuplicateInterchangeControlNumber { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to check for duplicate transaction set control number.
        /// </summary>
        [Input("checkDuplicateTransactionSetControlNumber", required: true)]
        public Input<bool> CheckDuplicateTransactionSetControlNumber { get; set; } = null!;

        /// <summary>
        /// The validity period of interchange control number.
        /// </summary>
        [Input("interchangeControlNumberValidityDays", required: true)]
        public Input<int> InterchangeControlNumberValidityDays { get; set; } = null!;

        /// <summary>
        /// The trailing separator policy.
        /// </summary>
        [Input("trailingSeparatorPolicy", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Logic.Latest.TrailingSeparatorPolicy> TrailingSeparatorPolicy { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to trim leading and trailing spaces and zeroes.
        /// </summary>
        [Input("trimLeadingAndTrailingSpacesAndZeroes", required: true)]
        public Input<bool> TrimLeadingAndTrailingSpacesAndZeroes { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to validate character set in the message.
        /// </summary>
        [Input("validateCharacterSet", required: true)]
        public Input<bool> ValidateCharacterSet { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to Whether to validate EDI types.
        /// </summary>
        [Input("validateEDITypes", required: true)]
        public Input<bool> ValidateEDITypes { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to Whether to validate XSD types.
        /// </summary>
        [Input("validateXSDTypes", required: true)]
        public Input<bool> ValidateXSDTypes { get; set; } = null!;

        public X12ValidationSettingsArgs()
        {
        }
    }
}
