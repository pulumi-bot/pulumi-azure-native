// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview
{
    public static class GetNotebookWorkspace
    {
        public static Task<GetNotebookWorkspaceResult> InvokeAsync(GetNotebookWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNotebookWorkspaceResult>("azurerm:documentdb/v20200601preview:getNotebookWorkspace", args ?? new GetNotebookWorkspaceArgs(), options.WithVersion());
    }


    public sealed class GetNotebookWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the notebook workspace resource.
        /// </summary>
        [Input("notebookWorkspaceName", required: true)]
        public string NotebookWorkspaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNotebookWorkspaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNotebookWorkspaceResult
    {
        /// <summary>
        /// The name of the database account.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the endpoint of Notebook server.
        /// </summary>
        public readonly string NotebookServerEndpoint;
        /// <summary>
        /// Status of the notebook workspace. Possible values are: Creating, Online, Deleting, Failed, Updating.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNotebookWorkspaceResult(
            string name,

            string notebookServerEndpoint,

            string status,

            string type)
        {
            Name = name;
            NotebookServerEndpoint = notebookServerEndpoint;
            Status = status;
            Type = type;
        }
    }
}
