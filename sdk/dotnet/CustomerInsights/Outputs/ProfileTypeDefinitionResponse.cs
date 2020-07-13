// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CustomerInsights.Outputs
{

    [OutputType]
    public sealed class ProfileTypeDefinitionResponse
    {
        /// <summary>
        /// The api entity set name. This becomes the odata entity set name for the entity Type being referred in this object.
        /// </summary>
        public readonly string? ApiEntitySetName;
        /// <summary>
        /// The attributes for the Type.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Attributes;
        /// <summary>
        /// Localized descriptions for the property.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Description;
        /// <summary>
        /// Localized display names for the property.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DisplayName;
        /// <summary>
        /// Type of entity.
        /// </summary>
        public readonly string? EntityType;
        /// <summary>
        /// The properties of the Profile.
        /// </summary>
        public readonly ImmutableArray<Outputs.PropertyDefinitionResponse> Fields;
        /// <summary>
        /// The instance count.
        /// </summary>
        public readonly int? InstancesCount;
        /// <summary>
        /// Large Image associated with the Property or EntityType.
        /// </summary>
        public readonly string? LargeImage;
        /// <summary>
        /// The last changed time for the type definition.
        /// </summary>
        public readonly string LastChangedUtc;
        /// <summary>
        /// Any custom localized attributes for the Type.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? LocalizedAttributes;
        /// <summary>
        /// Medium Image associated with the Property or EntityType.
        /// </summary>
        public readonly string? MediumImage;
        /// <summary>
        /// Provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The schema org link. This helps ACI identify and suggest semantic models.
        /// </summary>
        public readonly string? SchemaItemTypeLink;
        /// <summary>
        /// Small Image associated with the Property or EntityType.
        /// </summary>
        public readonly string? SmallImage;
        /// <summary>
        /// The strong IDs.
        /// </summary>
        public readonly ImmutableArray<Outputs.StrongIdResponse> StrongIds;
        /// <summary>
        /// The hub name.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The timestamp property name. Represents the time when the interaction or profile update happened.
        /// </summary>
        public readonly string? TimestampFieldName;
        /// <summary>
        /// The name of the entity.
        /// </summary>
        public readonly string? TypeName;

        [OutputConstructor]
        private ProfileTypeDefinitionResponse(
            string? apiEntitySetName,

            ImmutableDictionary<string, string>? attributes,

            ImmutableDictionary<string, string>? description,

            ImmutableDictionary<string, string>? displayName,

            string? entityType,

            ImmutableArray<Outputs.PropertyDefinitionResponse> fields,

            int? instancesCount,

            string? largeImage,

            string lastChangedUtc,

            ImmutableDictionary<string, string>? localizedAttributes,

            string? mediumImage,

            string provisioningState,

            string? schemaItemTypeLink,

            string? smallImage,

            ImmutableArray<Outputs.StrongIdResponse> strongIds,

            string tenantId,

            string? timestampFieldName,

            string? typeName)
        {
            ApiEntitySetName = apiEntitySetName;
            Attributes = attributes;
            Description = description;
            DisplayName = displayName;
            EntityType = entityType;
            Fields = fields;
            InstancesCount = instancesCount;
            LargeImage = largeImage;
            LastChangedUtc = lastChangedUtc;
            LocalizedAttributes = localizedAttributes;
            MediumImage = mediumImage;
            ProvisioningState = provisioningState;
            SchemaItemTypeLink = schemaItemTypeLink;
            SmallImage = smallImage;
            StrongIds = strongIds;
            TenantId = tenantId;
            TimestampFieldName = timestampFieldName;
            TypeName = typeName;
        }
    }
}