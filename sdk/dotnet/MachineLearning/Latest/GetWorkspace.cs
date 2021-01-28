// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.MachineLearning.Latest
{
    public static class GetWorkspace
    {
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("azure-nextgen:machinelearning/latest:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithVersion());
    }


    public sealed class GetWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group to which the machine learning workspace belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the machine learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// The creation time for this workspace resource.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The key vault identifier used for encrypted workspaces.
        /// </summary>
        public readonly string? KeyVaultIdentifierId;
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The email id of the owner for this workspace.
        /// </summary>
        public readonly string OwnerEmail;
        /// <summary>
        /// The regional endpoint for the machine learning studio service which hosts this workspace.
        /// </summary>
        public readonly string StudioEndpoint;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The fully qualified arm id of the storage account associated with this workspace.
        /// </summary>
        public readonly string UserStorageAccountId;
        /// <summary>
        /// The immutable id associated with this workspace.
        /// </summary>
        public readonly string WorkspaceId;
        /// <summary>
        /// The current state of workspace resource.
        /// </summary>
        public readonly string WorkspaceState;
        /// <summary>
        /// The type of this workspace.
        /// </summary>
        public readonly string WorkspaceType;

        [OutputConstructor]
        private GetWorkspaceResult(
            string creationTime,

            string id,

            string? keyVaultIdentifierId,

            string location,

            string name,

            string ownerEmail,

            string studioEndpoint,

            ImmutableDictionary<string, string>? tags,

            string type,

            string userStorageAccountId,

            string workspaceId,

            string workspaceState,

            string workspaceType)
        {
            CreationTime = creationTime;
            Id = id;
            KeyVaultIdentifierId = keyVaultIdentifierId;
            Location = location;
            Name = name;
            OwnerEmail = ownerEmail;
            StudioEndpoint = studioEndpoint;
            Tags = tags;
            Type = type;
            UserStorageAccountId = userStorageAccountId;
            WorkspaceId = workspaceId;
            WorkspaceState = workspaceState;
            WorkspaceType = workspaceType;
        }
    }
}
