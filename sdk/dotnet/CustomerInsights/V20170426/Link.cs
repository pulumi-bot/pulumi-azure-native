// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CustomerInsights.V20170426
{
    /// <summary>
    /// The link resource format.
    /// </summary>
    public partial class Link : Pulumi.CustomResource
    {
        /// <summary>
        /// Localized descriptions for the Link.
        /// </summary>
        [Output("description")]
        public Output<ImmutableDictionary<string, string>?> Description { get; private set; } = null!;

        /// <summary>
        /// Localized display name for the Link.
        /// </summary>
        [Output("displayName")]
        public Output<ImmutableDictionary<string, string>?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The link name.
        /// </summary>
        [Output("linkName")]
        public Output<string> LinkName { get; private set; } = null!;

        /// <summary>
        /// The set of properties mappings between the source and target Types.
        /// </summary>
        [Output("mappings")]
        public Output<ImmutableArray<Outputs.TypePropertiesMappingResponse>> Mappings { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
        /// </summary>
        [Output("operationType")]
        public Output<string?> OperationType { get; private set; } = null!;

        /// <summary>
        /// The properties that represent the participating profile.
        /// </summary>
        [Output("participantPropertyReferences")]
        public Output<ImmutableArray<Outputs.ParticipantPropertyReferenceResponse>> ParticipantPropertyReferences { get; private set; } = null!;

        /// <summary>
        /// Provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
        /// </summary>
        [Output("referenceOnly")]
        public Output<bool?> ReferenceOnly { get; private set; } = null!;

        /// <summary>
        /// Type of source entity.
        /// </summary>
        [Output("sourceEntityType")]
        public Output<string> SourceEntityType { get; private set; } = null!;

        /// <summary>
        /// Name of the source Entity Type.
        /// </summary>
        [Output("sourceEntityTypeName")]
        public Output<string> SourceEntityTypeName { get; private set; } = null!;

        /// <summary>
        /// Type of target entity.
        /// </summary>
        [Output("targetEntityType")]
        public Output<string> TargetEntityType { get; private set; } = null!;

        /// <summary>
        /// Name of the target Entity Type.
        /// </summary>
        [Output("targetEntityTypeName")]
        public Output<string> TargetEntityTypeName { get; private set; } = null!;

        /// <summary>
        /// The hub name.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Link resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Link(string name, LinkArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:customerinsights/v20170426:Link", name, args ?? new LinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Link(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:customerinsights/v20170426:Link", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:customerinsights/latest:Link"},
                    new Pulumi.Alias { Type = "azure-nextgen:customerinsights/v20170101:Link"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Link resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Link Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Link(name, id, options);
        }
    }

    public sealed class LinkArgs : Pulumi.ResourceArgs
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

        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public Input<string> HubName { get; set; } = null!;

        /// <summary>
        /// The name of the link.
        /// </summary>
        [Input("linkName", required: true)]
        public Input<string> LinkName { get; set; } = null!;

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
        public Input<Pulumi.AzureNextGen.CustomerInsights.V20170426.InstanceOperationType>? OperationType { get; set; }

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
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Type of source entity.
        /// </summary>
        [Input("sourceEntityType", required: true)]
        public Input<Pulumi.AzureNextGen.CustomerInsights.V20170426.EntityType> SourceEntityType { get; set; } = null!;

        /// <summary>
        /// Name of the source Entity Type.
        /// </summary>
        [Input("sourceEntityTypeName", required: true)]
        public Input<string> SourceEntityTypeName { get; set; } = null!;

        /// <summary>
        /// Type of target entity.
        /// </summary>
        [Input("targetEntityType", required: true)]
        public Input<Pulumi.AzureNextGen.CustomerInsights.V20170426.EntityType> TargetEntityType { get; set; } = null!;

        /// <summary>
        /// Name of the target Entity Type.
        /// </summary>
        [Input("targetEntityTypeName", required: true)]
        public Input<string> TargetEntityTypeName { get; set; } = null!;

        public LinkArgs()
        {
        }
    }
}
