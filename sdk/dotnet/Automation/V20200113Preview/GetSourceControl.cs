// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20200113Preview
{
    public static class GetSourceControl
    {
        public static Task<GetSourceControlResult> InvokeAsync(GetSourceControlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSourceControlResult>("azure-nextgen:automation/v20200113preview:getSourceControl", args ?? new GetSourceControlArgs(), options.WithVersion());
    }


    public sealed class GetSourceControlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of source control.
        /// </summary>
        [Input("sourceControlName", required: true)]
        public string SourceControlName { get; set; } = null!;

        public GetSourceControlArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSourceControlResult
    {
        /// <summary>
        /// The auto sync of the source control. Default is false.
        /// </summary>
        public readonly bool? AutoSync;
        /// <summary>
        /// The repo branch of the source control. Include branch as empty string for VsoTfvc.
        /// </summary>
        public readonly string? Branch;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The folder path of the source control.
        /// </summary>
        public readonly string? FolderPath;
        /// <summary>
        /// Fully qualified resource Id for the resource
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last modified time.
        /// </summary>
        public readonly string? LastModifiedTime;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The auto publish of the source control. Default is true.
        /// </summary>
        public readonly bool? PublishRunbook;
        /// <summary>
        /// The repo url of the source control.
        /// </summary>
        public readonly string? RepoUrl;
        /// <summary>
        /// The source type. Must be one of VsoGit, VsoTfvc, GitHub.
        /// </summary>
        public readonly string? SourceType;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSourceControlResult(
            bool? autoSync,

            string? branch,

            string? creationTime,

            string? description,

            string? folderPath,

            string id,

            string? lastModifiedTime,

            string name,

            bool? publishRunbook,

            string? repoUrl,

            string? sourceType,

            string type)
        {
            AutoSync = autoSync;
            Branch = branch;
            CreationTime = creationTime;
            Description = description;
            FolderPath = folderPath;
            Id = id;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            PublishRunbook = publishRunbook;
            RepoUrl = repoUrl;
            SourceType = sourceType;
            Type = type;
        }
    }
}
