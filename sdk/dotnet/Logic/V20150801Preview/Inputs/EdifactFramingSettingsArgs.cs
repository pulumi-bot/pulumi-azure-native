// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20150801Preview.Inputs
{

    public sealed class EdifactFramingSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The character encoding.
        /// </summary>
        [Input("characterEncoding")]
        public Input<string>? CharacterEncoding { get; set; }

        /// <summary>
        /// The EDIFACT frame setting characterSet.
        /// </summary>
        [Input("characterSet")]
        public Input<string>? CharacterSet { get; set; }

        /// <summary>
        /// The component separator.
        /// </summary>
        [Input("componentSeparator")]
        public Input<int>? ComponentSeparator { get; set; }

        /// <summary>
        /// The data element separator.
        /// </summary>
        [Input("dataElementSeparator")]
        public Input<int>? DataElementSeparator { get; set; }

        /// <summary>
        /// The EDIFACT frame setting decimal indicator.
        /// </summary>
        [Input("decimalPointIndicator")]
        public Input<string>? DecimalPointIndicator { get; set; }

        /// <summary>
        /// The protocol version.
        /// </summary>
        [Input("protocolVersion")]
        public Input<int>? ProtocolVersion { get; set; }

        /// <summary>
        /// The release indicator.
        /// </summary>
        [Input("releaseIndicator")]
        public Input<int>? ReleaseIndicator { get; set; }

        /// <summary>
        /// The repetition separator.
        /// </summary>
        [Input("repetitionSeparator")]
        public Input<int>? RepetitionSeparator { get; set; }

        /// <summary>
        /// The segment terminator.
        /// </summary>
        [Input("segmentTerminator")]
        public Input<int>? SegmentTerminator { get; set; }

        /// <summary>
        /// The EDIFACT frame setting segment terminator suffix.
        /// </summary>
        [Input("segmentTerminatorSuffix")]
        public Input<string>? SegmentTerminatorSuffix { get; set; }

        /// <summary>
        /// The service code list directory version.
        /// </summary>
        [Input("serviceCodeListDirectoryVersion")]
        public Input<string>? ServiceCodeListDirectoryVersion { get; set; }

        public EdifactFramingSettingsArgs()
        {
        }
    }
}
