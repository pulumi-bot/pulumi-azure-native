// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview
{
    public static class GetBackend
    {
        public static Task<GetBackendResult> InvokeAsync(GetBackendArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackendResult>("azurerm:apimanagement/v20191201preview:getBackend", args ?? new GetBackendArgs(), options.WithVersion());
    }


    public sealed class GetBackendArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier of the Backend entity. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("backendId", required: true)]
        public string BackendId { get; set; } = null!;

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

        public GetBackendArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackendResult
    {
        /// <summary>
        /// Backend Credentials Contract Properties
        /// </summary>
        public readonly Outputs.BackendCredentialsContractResponseResult? Credentials;
        /// <summary>
        /// Backend Description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Backend Properties contract
        /// </summary>
        public readonly Outputs.BackendPropertiesResponseResult Properties;
        /// <summary>
        /// Backend communication protocol.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Backend Proxy Contract Properties
        /// </summary>
        public readonly Outputs.BackendProxyContractResponseResult? Proxy;
        /// <summary>
        /// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// Backend Title.
        /// </summary>
        public readonly string? Title;
        /// <summary>
        /// Backend TLS Properties
        /// </summary>
        public readonly Outputs.BackendTlsPropertiesResponseResult? Tls;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Runtime Url of the Backend.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetBackendResult(
            Outputs.BackendCredentialsContractResponseResult? credentials,

            string? description,

            string name,

            Outputs.BackendPropertiesResponseResult properties,

            string protocol,

            Outputs.BackendProxyContractResponseResult? proxy,

            string? resourceId,

            string? title,

            Outputs.BackendTlsPropertiesResponseResult? tls,

            string type,

            string url)
        {
            Credentials = credentials;
            Description = description;
            Name = name;
            Properties = properties;
            Protocol = protocol;
            Proxy = proxy;
            ResourceId = resourceId;
            Title = title;
            Tls = tls;
            Type = type;
            Url = url;
        }
    }
}
