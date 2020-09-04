// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200501Preview
{
    public static class GetLinkedWorkspace
    {
        public static Task<GetLinkedWorkspaceResult> InvokeAsync(GetLinkedWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLinkedWorkspaceResult>("azurerm:machinelearningservices/v20200501preview:getLinkedWorkspace", args ?? new GetLinkedWorkspaceArgs(), options.WithVersion());
    }


    public sealed class GetLinkedWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Friendly name of the linked workspace
        /// </summary>
        [Input("linkName", required: true)]
        public string LinkName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group in which workspace is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Azure Machine Learning workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetLinkedWorkspaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLinkedWorkspaceResult
    {
        /// <summary>
        /// Friendly name of the linked workspace.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// LinkedWorkspace specific properties.
        /// </summary>
        public readonly Outputs.LinkedWorkspacePropsResponseResult Properties;
        /// <summary>
        /// Resource type of linked workspace.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetLinkedWorkspaceResult(
            string name,

            Outputs.LinkedWorkspacePropsResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
