// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Inputs
{

    /// <summary>
    /// The X12 processing settings.
    /// </summary>
    public sealed class X12ProcessingSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value indicating whether to convert numerical type to implied decimal.
        /// </summary>
        [Input("convertImpliedDecimal", required: true)]
        public Input<bool> ConvertImpliedDecimal { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to create empty xml tags for trailing separators.
        /// </summary>
        [Input("createEmptyXmlTagsForTrailingSeparators", required: true)]
        public Input<bool> CreateEmptyXmlTagsForTrailingSeparators { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to mask security information.
        /// </summary>
        [Input("maskSecurityInfo", required: true)]
        public Input<bool> MaskSecurityInfo { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to preserve interchange.
        /// </summary>
        [Input("preserveInterchange", required: true)]
        public Input<bool> PreserveInterchange { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to suspend interchange on error.
        /// </summary>
        [Input("suspendInterchangeOnError", required: true)]
        public Input<bool> SuspendInterchangeOnError { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to use dot as decimal separator.
        /// </summary>
        [Input("useDotAsDecimalSeparator", required: true)]
        public Input<bool> UseDotAsDecimalSeparator { get; set; } = null!;

        public X12ProcessingSettingsArgs()
        {
        }
    }
}
