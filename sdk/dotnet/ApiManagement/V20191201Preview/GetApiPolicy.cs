// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview
{
    public static class GetApiPolicy
    {
        public static Task<GetApiPolicyResult> InvokeAsync(GetApiPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiPolicyResult>("azurerm:apimanagement/v20191201preview:getApiPolicy", args ?? new GetApiPolicyArgs(), options.WithVersion());
    }


    public sealed class GetApiPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        /// <summary>
        /// Policy Export Format.
        /// </summary>
        [Input("format")]
        public string? Format { get; set; }

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

        public GetApiPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiPolicyResult
    {
        /// <summary>
        /// Format of the policyContent.
        /// </summary>
        public readonly string? Format;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Contents of the Policy as defined by the format.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetApiPolicyResult(
            string? format,

            string name,

            string type,

            string value)
        {
            Format = format;
            Name = name;
            Type = type;
            Value = value;
        }
    }
}
