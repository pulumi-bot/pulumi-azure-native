// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.V20150801Preview.Inputs
{

    public sealed class X12ValidationSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value indicating whether to allow leading and trailing spaces and zeroes.
        /// </summary>
        [Input("allowLeadingAndTrailingSpacesAndZeroes")]
        public Input<bool>? AllowLeadingAndTrailingSpacesAndZeroes { get; set; }

        /// <summary>
        /// The value indicating whether to check for duplicate group control number.
        /// </summary>
        [Input("checkDuplicateGroupControlNumber")]
        public Input<bool>? CheckDuplicateGroupControlNumber { get; set; }

        /// <summary>
        /// The value indicating whether to check for duplicate interchange control number.
        /// </summary>
        [Input("checkDuplicateInterchangeControlNumber")]
        public Input<bool>? CheckDuplicateInterchangeControlNumber { get; set; }

        /// <summary>
        /// The value indicating whether to check for duplicate transaction set control number.
        /// </summary>
        [Input("checkDuplicateTransactionSetControlNumber")]
        public Input<bool>? CheckDuplicateTransactionSetControlNumber { get; set; }

        /// <summary>
        /// The validity period of interchange control number.
        /// </summary>
        [Input("interchangeControlNumberValidityDays")]
        public Input<int>? InterchangeControlNumberValidityDays { get; set; }

        /// <summary>
        /// The trailing separator policy.
        /// </summary>
        [Input("trailingSeparatorPolicy")]
        public Input<Pulumi.AzureNextGen.Logic.V20150801Preview.TrailingSeparatorPolicy>? TrailingSeparatorPolicy { get; set; }

        /// <summary>
        /// The value indicating whether to trim leading and trailing spaces and zeroes.
        /// </summary>
        [Input("trimLeadingAndTrailingSpacesAndZeroes")]
        public Input<bool>? TrimLeadingAndTrailingSpacesAndZeroes { get; set; }

        /// <summary>
        /// The value indicating whether to validate character set in the message.
        /// </summary>
        [Input("validateCharacterSet")]
        public Input<bool>? ValidateCharacterSet { get; set; }

        /// <summary>
        /// The value indicating whether to Whether to validate EDI types.
        /// </summary>
        [Input("validateEDITypes")]
        public Input<bool>? ValidateEDITypes { get; set; }

        /// <summary>
        /// The value indicating whether to Whether to validate XSD types.
        /// </summary>
        [Input("validateXSDTypes")]
        public Input<bool>? ValidateXSDTypes { get; set; }

        public X12ValidationSettingsArgs()
        {
        }
    }
}
