// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview
{
    public static class GetApi
    {
        public static Task<GetApiResult> InvokeAsync(GetApiArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiResult>("azurerm:apimanagement/v20191201preview:getApi", args ?? new GetApiArgs(), options.WithVersion());
    }


    public sealed class GetApiArgs : Pulumi.InvokeArgs
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

        public GetApiArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiResult
    {
        /// <summary>
        /// Describes the Revision of the Api. If no value is provided, default revision 1 is created
        /// </summary>
        public readonly string? ApiRevision;
        /// <summary>
        /// Description of the Api Revision.
        /// </summary>
        public readonly string? ApiRevisionDescription;
        /// <summary>
        /// Type of API.
        /// </summary>
        public readonly string? ApiType;
        /// <summary>
        /// Indicates the Version identifier of the API if the API is versioned
        /// </summary>
        public readonly string? ApiVersion;
        /// <summary>
        /// Description of the Api Version.
        /// </summary>
        public readonly string? ApiVersionDescription;
        /// <summary>
        /// Version set details
        /// </summary>
        public readonly Outputs.ApiVersionSetContractDetailsResponseResult? ApiVersionSet;
        /// <summary>
        /// A resource identifier for the related ApiVersionSet.
        /// </summary>
        public readonly string? ApiVersionSetId;
        /// <summary>
        /// Collection of authentication settings included into this API.
        /// </summary>
        public readonly Outputs.AuthenticationSettingsContractResponseResult? AuthenticationSettings;
        /// <summary>
        /// Description of the API. May include HTML formatting tags.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// API name. Must be 1 to 300 characters long.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Indicates if API revision is current api revision.
        /// </summary>
        public readonly bool? IsCurrent;
        /// <summary>
        /// Indicates if API revision is accessible via the gateway.
        /// </summary>
        public readonly bool IsOnline;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Describes on which protocols the operations in this API can be invoked.
        /// </summary>
        public readonly ImmutableArray<string> Protocols;
        /// <summary>
        /// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
        /// </summary>
        public readonly string? ServiceUrl;
        /// <summary>
        /// API identifier of the source API.
        /// </summary>
        public readonly string? SourceApiId;
        /// <summary>
        /// Protocols over which API is made available.
        /// </summary>
        public readonly Outputs.SubscriptionKeyParameterNamesContractResponseResult? SubscriptionKeyParameterNames;
        /// <summary>
        /// Specifies whether an API or Product subscription is required for accessing the API.
        /// </summary>
        public readonly bool? SubscriptionRequired;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApiResult(
            string? apiRevision,

            string? apiRevisionDescription,

            string? apiType,

            string? apiVersion,

            string? apiVersionDescription,

            Outputs.ApiVersionSetContractDetailsResponseResult? apiVersionSet,

            string? apiVersionSetId,

            Outputs.AuthenticationSettingsContractResponseResult? authenticationSettings,

            string? description,

            string? displayName,

            bool? isCurrent,

            bool isOnline,

            string name,

            string path,

            ImmutableArray<string> protocols,

            string? serviceUrl,

            string? sourceApiId,

            Outputs.SubscriptionKeyParameterNamesContractResponseResult? subscriptionKeyParameterNames,

            bool? subscriptionRequired,

            string type)
        {
            ApiRevision = apiRevision;
            ApiRevisionDescription = apiRevisionDescription;
            ApiType = apiType;
            ApiVersion = apiVersion;
            ApiVersionDescription = apiVersionDescription;
            ApiVersionSet = apiVersionSet;
            ApiVersionSetId = apiVersionSetId;
            AuthenticationSettings = authenticationSettings;
            Description = description;
            DisplayName = displayName;
            IsCurrent = isCurrent;
            IsOnline = isOnline;
            Name = name;
            Path = path;
            Protocols = protocols;
            ServiceUrl = serviceUrl;
            SourceApiId = sourceApiId;
            SubscriptionKeyParameterNames = subscriptionKeyParameterNames;
            SubscriptionRequired = subscriptionRequired;
            Type = type;
        }
    }
}
