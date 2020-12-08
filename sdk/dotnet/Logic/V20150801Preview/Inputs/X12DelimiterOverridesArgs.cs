// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.V20150801Preview.Inputs
{

    public sealed class X12DelimiterOverridesArgs : Pulumi.ResourceArgs
    {
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
        /// The message id.
        /// </summary>
        [Input("messageId")]
        public Input<string>? MessageId { get; set; }

        /// <summary>
        /// The protocol version.
        /// </summary>
        [Input("protocolVersion")]
        public Input<string>? ProtocolVersion { get; set; }

        /// <summary>
        /// The replacement character.
        /// </summary>
        [Input("replaceCharacter")]
        public Input<int>? ReplaceCharacter { get; set; }

        /// <summary>
        /// The value indicating whether to replace separators in payload.
        /// </summary>
        [Input("replaceSeparatorsInPayload")]
        public Input<bool>? ReplaceSeparatorsInPayload { get; set; }

        /// <summary>
        /// The segment terminator.
        /// </summary>
        [Input("segmentTerminator")]
        public Input<int>? SegmentTerminator { get; set; }

        /// <summary>
        /// The segment terminator suffix.
        /// </summary>
        [Input("segmentTerminatorSuffix")]
        public Input<Pulumi.AzureNextGen.Logic.V20150801Preview.SegmentTerminatorSuffix>? SegmentTerminatorSuffix { get; set; }

        /// <summary>
        /// The target namespace on which this delimiter settings has to be applied.
        /// </summary>
        [Input("targetNamespace")]
        public Input<string>? TargetNamespace { get; set; }

        public X12DelimiterOverridesArgs()
        {
        }
    }
}
