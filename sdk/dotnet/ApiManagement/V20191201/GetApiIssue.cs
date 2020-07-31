// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201
{
    public static class GetApiIssue
    {
        public static Task<GetApiIssueResult> InvokeAsync(GetApiIssueArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiIssueResult>("azurerm:apimanagement/v20191201:getApiIssue", args ?? new GetApiIssueArgs(), options.WithVersion());
    }


    public sealed class GetApiIssueArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// API identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        /// <summary>
        /// Expand the comment attachments. 
        /// </summary>
        [Input("expandCommentsAttachments")]
        public bool? ExpandCommentsAttachments { get; set; }

        /// <summary>
        /// Issue identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetApiIssueArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiIssueResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the Issue.
        /// </summary>
        public readonly Outputs.IssueContractPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApiIssueResult(
            string name,

            Outputs.IssueContractPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}