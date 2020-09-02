// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview
{
    public static class GetTagByOperation
    {
        public static Task<GetTagByOperationResult> InvokeAsync(GetTagByOperationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagByOperationResult>("azurerm:apimanagement/v20191201preview:getTagByOperation", args ?? new GetTagByOperationArgs(), options.WithVersion());
    }


    public sealed class GetTagByOperationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        /// <summary>
        /// Operation identifier within an API. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("operationId", required: true)]
        public string OperationId { get; set; } = null!;

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

        /// <summary>
        /// Tag identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("tagId", required: true)]
        public string TagId { get; set; } = null!;

        public GetTagByOperationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagByOperationResult
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTagByOperationResult(
            string displayName,

            string name,

            string type)
        {
            DisplayName = displayName;
            Name = name;
            Type = type;
        }
    }
}
