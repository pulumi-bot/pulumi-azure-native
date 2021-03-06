// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CustomerInsights.Latest
{
    [Obsolete(@"The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:customerinsights:getRelationship'.")]
    public static class GetRelationship
    {
        /// <summary>
        /// The relationship resource format.
        /// Latest API Version: 2017-04-26.
        /// </summary>
        public static Task<GetRelationshipResult> InvokeAsync(GetRelationshipArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRelationshipResult>("azure-native:customerinsights/latest:getRelationship", args ?? new GetRelationshipArgs(), options.WithVersion());
    }


    public sealed class GetRelationshipArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the hub.
        /// </summary>
        [Input("hubName", required: true)]
        public string HubName { get; set; } = null!;

        /// <summary>
        /// The name of the relationship.
        /// </summary>
        [Input("relationshipName", required: true)]
        public string RelationshipName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRelationshipArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRelationshipResult
    {
        /// <summary>
        /// The Relationship Cardinality.
        /// </summary>
        public readonly string? Cardinality;
        /// <summary>
        /// Localized descriptions for the Relationship.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Description;
        /// <summary>
        /// Localized display name for the Relationship.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DisplayName;
        /// <summary>
        /// The expiry date time in UTC.
        /// </summary>
        public readonly string? ExpiryDateTimeUtc;
        /// <summary>
        /// The properties of the Relationship.
        /// </summary>
        public readonly ImmutableArray<Outputs.PropertyDefinitionResponse> Fields;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Optional property to be used to map fields in profile to their strong ids in related profile.
        /// </summary>
        public readonly ImmutableArray<Outputs.RelationshipTypeMappingResponse> LookupMappings;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Profile type.
        /// </summary>
        public readonly string ProfileType;
        /// <summary>
        /// Provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Related profile being referenced.
        /// </summary>
        public readonly string RelatedProfileType;
        /// <summary>
        /// The relationship guid id.
        /// </summary>
        public readonly string RelationshipGuidId;
        /// <summary>
        /// The Relationship name.
        /// </summary>
        public readonly string RelationshipName;
        /// <summary>
        /// The hub name.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRelationshipResult(
            string? cardinality,

            ImmutableDictionary<string, string>? description,

            ImmutableDictionary<string, string>? displayName,

            string? expiryDateTimeUtc,

            ImmutableArray<Outputs.PropertyDefinitionResponse> fields,

            string id,

            ImmutableArray<Outputs.RelationshipTypeMappingResponse> lookupMappings,

            string name,

            string profileType,

            string provisioningState,

            string relatedProfileType,

            string relationshipGuidId,

            string relationshipName,

            string tenantId,

            string type)
        {
            Cardinality = cardinality;
            Description = description;
            DisplayName = displayName;
            ExpiryDateTimeUtc = expiryDateTimeUtc;
            Fields = fields;
            Id = id;
            LookupMappings = lookupMappings;
            Name = name;
            ProfileType = profileType;
            ProvisioningState = provisioningState;
            RelatedProfileType = relatedProfileType;
            RelationshipGuidId = relationshipGuidId;
            RelationshipName = relationshipName;
            TenantId = tenantId;
            Type = type;
        }
    }
}
