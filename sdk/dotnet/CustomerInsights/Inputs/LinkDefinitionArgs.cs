// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.Inputs
{

    /// <summary>
    /// The definition of Link.
    /// </summary>
    public sealed class LinkDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        private InputMap<string>? _description;

        /// <summary>
        /// Localized descriptions for the Link.
        /// </summary>
        public InputMap<string> Description
        {
            get => _description ?? (_description = new InputMap<string>());
            set => _description = value;
        }

        [Input("displayName")]
        private InputMap<string>? _displayName;

        /// <summary>
        /// Localized display name for the Link.
        /// </summary>
        public InputMap<string> DisplayName
        {
            get => _displayName ?? (_displayName = new InputMap<string>());
            set => _displayName = value;
        }

        [Input("mappings")]
        private InputList<Inputs.TypePropertiesMappingArgs>? _mappings;

        /// <summary>
        /// The set of properties mappings between the source and target Types.
        /// </summary>
        public InputList<Inputs.TypePropertiesMappingArgs> Mappings
        {
            get => _mappings ?? (_mappings = new InputList<Inputs.TypePropertiesMappingArgs>());
            set => _mappings = value;
        }

        /// <summary>
        /// Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
        /// </summary>
        [Input("operationType")]
        public Input<string>? OperationType { get; set; }

        [Input("participantPropertyReferences", required: true)]
        private InputList<Inputs.ParticipantPropertyReferenceArgs>? _participantPropertyReferences;

        /// <summary>
        /// The properties that represent the participating profile.
        /// </summary>
        public InputList<Inputs.ParticipantPropertyReferenceArgs> ParticipantPropertyReferences
        {
            get => _participantPropertyReferences ?? (_participantPropertyReferences = new InputList<Inputs.ParticipantPropertyReferenceArgs>());
            set => _participantPropertyReferences = value;
        }

        /// <summary>
        /// Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
        /// </summary>
        [Input("referenceOnly")]
        public Input<bool>? ReferenceOnly { get; set; }

        /// <summary>
        /// Type of source entity.
        /// </summary>
        [Input("sourceEntityType", required: true)]
        public Input<string> SourceEntityType { get; set; } = null!;

        /// <summary>
        /// Name of the source Entity Type.
        /// </summary>
        [Input("sourceEntityTypeName", required: true)]
        public Input<string> SourceEntityTypeName { get; set; } = null!;

        /// <summary>
        /// Type of target entity.
        /// </summary>
        [Input("targetEntityType", required: true)]
        public Input<string> TargetEntityType { get; set; } = null!;

        /// <summary>
        /// Name of the target Entity Type.
        /// </summary>
        [Input("targetEntityTypeName", required: true)]
        public Input<string> TargetEntityTypeName { get; set; } = null!;

        public LinkDefinitionArgs()
        {
        }
    }
}