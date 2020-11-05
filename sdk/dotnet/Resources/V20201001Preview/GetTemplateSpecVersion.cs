// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Resources.V20201001Preview
{
    public static class GetTemplateSpecVersion
    {
        public static Task<GetTemplateSpecVersionResult> InvokeAsync(GetTemplateSpecVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTemplateSpecVersionResult>("azure-nextgen:resources/v20201001preview:getTemplateSpecVersion", args ?? new GetTemplateSpecVersionArgs(), options.WithVersion());
    }


    public sealed class GetTemplateSpecVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the Template Spec.
        /// </summary>
        [Input("templateSpecName", required: true)]
        public string TemplateSpecName { get; set; } = null!;

        /// <summary>
        /// The version of the Template Spec.
        /// </summary>
        [Input("templateSpecVersion", required: true)]
        public string TemplateSpecVersion { get; set; } = null!;

        public GetTemplateSpecVersionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTemplateSpecVersionResult
    {
        /// <summary>
        /// An array of Template Spec artifacts.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateSpecTemplateArtifactResponse> Artifacts;
        /// <summary>
        /// Template Spec version description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The location of the Template Spec Version. It must match the location of the parent Template Spec.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The Azure Resource Manager template content.
        /// </summary>
        public readonly object? Template;
        /// <summary>
        /// Type of this resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The Azure Resource Manager template UI definition content
        /// </summary>
        public readonly object? UiDefinition;

        [OutputConstructor]
        private GetTemplateSpecVersionResult(
            ImmutableArray<Outputs.TemplateSpecTemplateArtifactResponse> artifacts,

            string? description,

            string location,

            string name,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            object? template,

            string type,

            object? uiDefinition)
        {
            Artifacts = artifacts;
            Description = description;
            Location = location;
            Name = name;
            SystemData = systemData;
            Tags = tags;
            Template = template;
            Type = type;
            UiDefinition = uiDefinition;
        }
    }
}
