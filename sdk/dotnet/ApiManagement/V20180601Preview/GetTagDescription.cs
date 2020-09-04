// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20180601Preview
{
    public static class GetTagDescription
    {
        public static Task<GetTagDescriptionResult> InvokeAsync(GetTagDescriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagDescriptionResult>("azurerm:apimanagement/v20180601preview:getTagDescription", args ?? new GetTagDescriptionArgs(), options.WithVersion());
    }


    public sealed class GetTagDescriptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

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

        public GetTagDescriptionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagDescriptionResult
    {
        /// <summary>
        /// Description of the Tag.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Description of the external resources describing the tag.
        /// </summary>
        public readonly string? ExternalDocsDescription;
        /// <summary>
        /// Absolute URL of external resources describing the tag.
        /// </summary>
        public readonly string? ExternalDocsUrl;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTagDescriptionResult(
            string? description,

            string? displayName,

            string? externalDocsDescription,

            string? externalDocsUrl,

            string name,

            string type)
        {
            Description = description;
            DisplayName = displayName;
            ExternalDocsDescription = externalDocsDescription;
            ExternalDocsUrl = externalDocsUrl;
            Name = name;
            Type = type;
        }
    }
}
