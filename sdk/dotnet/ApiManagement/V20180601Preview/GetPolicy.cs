// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20180601Preview
{
    public static class GetPolicy
    {
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("azurerm:apimanagement/v20180601preview:getPolicy", args ?? new GetPolicyArgs(), options.WithVersion());
    }


    public sealed class GetPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the Policy.
        /// </summary>
        [Input("policyId", required: true)]
        public string PolicyId { get; set; } = null!;

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

        public GetPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// Format of the policyContent.
        /// </summary>
        public readonly string? ContentFormat;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Json escaped Xml Encoded contents of the Policy.
        /// </summary>
        public readonly string PolicyContent;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPolicyResult(
            string? contentFormat,

            string name,

            string policyContent,

            string type)
        {
            ContentFormat = contentFormat;
            Name = name;
            PolicyContent = policyContent;
            Type = type;
        }
    }
}
