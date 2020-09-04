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
    public sealed class X12ProtocolSettingsResponseResult
    {
        /// <summary>
        /// The X12 acknowledgment settings.
        /// </summary>
        public readonly Outputs.X12AcknowledgementSettingsResponseResult? AcknowledgementSettings;
        /// <summary>
        /// The X12 envelope override settings.
        /// </summary>
        public readonly ImmutableArray<Outputs.X12EnvelopeOverrideResponseResult> EnvelopeOverrides;
        /// <summary>
        /// The X12 envelope settings.
        /// </summary>
        public readonly Outputs.X12EnvelopeSettingsResponseResult? EnvelopeSettings;
        /// <summary>
        /// The X12 framing settings.
        /// </summary>
        public readonly Outputs.X12FramingSettingsResponseResult? FramingSettings;
        /// <summary>
        /// The X12 message filter.
        /// </summary>
        public readonly Outputs.X12MessageFilterResponseResult? MessageFilter;
        /// <summary>
        /// The X12 message filter list.
        /// </summary>
        public readonly ImmutableArray<Outputs.X12MessageIdentifierResponseResult> MessageFilterList;
        /// <summary>
        /// The X12 processing settings.
        /// </summary>
        public readonly Outputs.X12ProcessingSettingsResponseResult? ProcessingSettings;
        /// <summary>
        /// The X12 schema references.
        /// </summary>
        public readonly ImmutableArray<Outputs.X12SchemaReferenceResponseResult> SchemaReferences;
        /// <summary>
        /// The X12 security settings.
        /// </summary>
        public readonly Outputs.X12SecuritySettingsResponseResult? SecuritySettings;
        /// <summary>
        /// The X12 validation override settings.
        /// </summary>
        public readonly ImmutableArray<Outputs.X12ValidationOverrideResponseResult> ValidationOverrides;
        /// <summary>
        /// The X12 validation settings.
        /// </summary>
        public readonly Outputs.X12ValidationSettingsResponseResult? ValidationSettings;
        /// <summary>
        /// The X12 delimiter override settings.
        /// </summary>
        public readonly ImmutableArray<Outputs.X12DelimiterOverridesResponseResult> X12DelimiterOverrides;

        [OutputConstructor]
        private X12ProtocolSettingsResponseResult(
            Outputs.X12AcknowledgementSettingsResponseResult? acknowledgementSettings,

            ImmutableArray<Outputs.X12EnvelopeOverrideResponseResult> envelopeOverrides,

            Outputs.X12EnvelopeSettingsResponseResult? envelopeSettings,

            Outputs.X12FramingSettingsResponseResult? framingSettings,

            Outputs.X12MessageFilterResponseResult? messageFilter,

            ImmutableArray<Outputs.X12MessageIdentifierResponseResult> messageFilterList,

            Outputs.X12ProcessingSettingsResponseResult? processingSettings,

            ImmutableArray<Outputs.X12SchemaReferenceResponseResult> schemaReferences,

            Outputs.X12SecuritySettingsResponseResult? securitySettings,

            ImmutableArray<Outputs.X12ValidationOverrideResponseResult> validationOverrides,

            Outputs.X12ValidationSettingsResponseResult? validationSettings,

            ImmutableArray<Outputs.X12DelimiterOverridesResponseResult> x12DelimiterOverrides)
        {
            AcknowledgementSettings = acknowledgementSettings;
            EnvelopeOverrides = envelopeOverrides;
            EnvelopeSettings = envelopeSettings;
            FramingSettings = framingSettings;
            MessageFilter = messageFilter;
            MessageFilterList = messageFilterList;
            ProcessingSettings = processingSettings;
            SchemaReferences = schemaReferences;
            SecuritySettings = securitySettings;
            ValidationOverrides = validationOverrides;
            ValidationSettings = validationSettings;
            X12DelimiterOverrides = x12DelimiterOverrides;
        }
    }
}
