// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Blueprint.V20181101Preview
{
    public static class GetBlueprint
    {
        public static Task<GetBlueprintResult> InvokeAsync(GetBlueprintArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBlueprintResult>("azurerm:blueprint/v20181101preview:getBlueprint", args ?? new GetBlueprintArgs(), options.WithVersion());
    }


    public sealed class GetBlueprintArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the blueprint definition.
        /// </summary>
        [Input("blueprintName", required: true)]
        public string BlueprintName { get; set; } = null!;

        /// <summary>
        /// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
        /// </summary>
        [Input("resourceScope", required: true)]
        public string ResourceScope { get; set; } = null!;

        public GetBlueprintArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBlueprintResult
    {
        /// <summary>
        /// Multi-line explain this resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// One-liner string explain this resource.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Layout view of the blueprint definition for UI reference.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Layout;
        /// <summary>
        /// Name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Parameters required by this blueprint definition.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterDefinitionResponseResult>? Parameters;
        /// <summary>
        /// Resource group placeholders defined by this blueprint definition.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ResourceGroupDefinitionResponseResult>? ResourceGroups;
        /// <summary>
        /// Status of the blueprint. This field is readonly.
        /// </summary>
        public readonly Outputs.BlueprintStatusResponseResult Status;
        /// <summary>
        /// The scope where this blueprint definition can be assigned.
        /// </summary>
        public readonly string TargetScope;
        /// <summary>
        /// Type of this resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Published versions of this blueprint definition.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Versions;

        [OutputConstructor]
        private GetBlueprintResult(
            string? description,

            string? displayName,

            ImmutableDictionary<string, object>? layout,

            string name,

            ImmutableDictionary<string, Outputs.ParameterDefinitionResponseResult>? parameters,

            ImmutableDictionary<string, Outputs.ResourceGroupDefinitionResponseResult>? resourceGroups,

            Outputs.BlueprintStatusResponseResult status,

            string targetScope,

            string type,

            ImmutableDictionary<string, object>? versions)
        {
            Description = description;
            DisplayName = displayName;
            Layout = layout;
            Name = name;
            Parameters = parameters;
            ResourceGroups = resourceGroups;
            Status = status;
            TargetScope = targetScope;
            Type = type;
            Versions = versions;
        }
    }
}
