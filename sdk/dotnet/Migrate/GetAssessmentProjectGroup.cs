// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate
{
    public static class GetAssessmentProjectGroup
    {
        public static Task<GetAssessmentProjectGroupResult> InvokeAsync(GetAssessmentProjectGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAssessmentProjectGroupResult>("azurerm:migrate:getAssessmentProjectGroup", args ?? new GetAssessmentProjectGroupArgs(), options.WithVersion());
    }


    public sealed class GetAssessmentProjectGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique name of a group within a project.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName", required: true)]
        public string ProjectName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAssessmentProjectGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAssessmentProjectGroupResult
    {
        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// Name of the group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the group.
        /// </summary>
        public readonly Outputs.GroupPropertiesResponseResult Properties;
        /// <summary>
        /// Type of the object = [Microsoft.Migrate/assessmentProjects/groups].
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAssessmentProjectGroupResult(
            string? eTag,

            string name,

            Outputs.GroupPropertiesResponseResult properties,

            string type)
        {
            ETag = eTag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}