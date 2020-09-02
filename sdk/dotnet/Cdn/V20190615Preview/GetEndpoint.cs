// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20190615Preview
{
    public static class GetEndpoint
    {
        public static Task<GetEndpointResult> InvokeAsync(GetEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEndpointResult>("azurerm:cdn/v20190615preview:getEndpoint", args ?? new GetEndpointArgs(), options.WithVersion());
    }


    public sealed class GetEndpointArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetEndpointArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEndpointResult
    {
        /// <summary>
        /// List of content types on which compression applies. The value should be a valid MIME type.
        /// </summary>
        public readonly ImmutableArray<string> ContentTypesToCompress;
        /// <summary>
        /// A policy that specifies the delivery rules to be used for an endpoint.
        /// </summary>
        public readonly Outputs.EndpointPropertiesUpdateParametersResponseDeliveryPolicyResult? DeliveryPolicy;
        /// <summary>
        /// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
        /// </summary>
        public readonly ImmutableArray<Outputs.GeoFilterResponseResult> GeoFilters;
        /// <summary>
        /// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
        /// </summary>
        public readonly bool? IsCompressionEnabled;
        /// <summary>
        /// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        /// </summary>
        public readonly bool? IsHttpAllowed;
        /// <summary>
        /// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        /// </summary>
        public readonly bool? IsHttpsAllowed;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
        /// </summary>
        public readonly string? OptimizationType;
        /// <summary>
        /// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
        /// </summary>
        public readonly string? OriginHostHeader;
        /// <summary>
        /// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
        /// </summary>
        public readonly string? OriginPath;
        /// <summary>
        /// The source of the content being delivered via CDN.
        /// </summary>
        public readonly ImmutableArray<Outputs.DeepCreatedOriginResponseResult> Origins;
        /// <summary>
        /// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path.
        /// </summary>
        public readonly string? ProbePath;
        /// <summary>
        /// Provisioning status of the endpoint.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
        /// </summary>
        public readonly string? QueryStringCachingBehavior;
        /// <summary>
        /// Resource status of the endpoint.
        /// </summary>
        public readonly string ResourceState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Defines the Web Application Firewall policy for the endpoint (if applicable)
        /// </summary>
        public readonly Outputs.EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkResult? WebApplicationFirewallPolicyLink;

        [OutputConstructor]
        private GetEndpointResult(
            ImmutableArray<string> contentTypesToCompress,

            Outputs.EndpointPropertiesUpdateParametersResponseDeliveryPolicyResult? deliveryPolicy,

            ImmutableArray<Outputs.GeoFilterResponseResult> geoFilters,

            string hostName,

            bool? isCompressionEnabled,

            bool? isHttpAllowed,

            bool? isHttpsAllowed,

            string location,

            string name,

            string? optimizationType,

            string? originHostHeader,

            string? originPath,

            ImmutableArray<Outputs.DeepCreatedOriginResponseResult> origins,

            string? probePath,

            string provisioningState,

            string? queryStringCachingBehavior,

            string resourceState,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkResult? webApplicationFirewallPolicyLink)
        {
            ContentTypesToCompress = contentTypesToCompress;
            DeliveryPolicy = deliveryPolicy;
            GeoFilters = geoFilters;
            HostName = hostName;
            IsCompressionEnabled = isCompressionEnabled;
            IsHttpAllowed = isHttpAllowed;
            IsHttpsAllowed = isHttpsAllowed;
            Location = location;
            Name = name;
            OptimizationType = optimizationType;
            OriginHostHeader = originHostHeader;
            OriginPath = originPath;
            Origins = origins;
            ProbePath = probePath;
            ProvisioningState = provisioningState;
            QueryStringCachingBehavior = queryStringCachingBehavior;
            ResourceState = resourceState;
            Tags = tags;
            Type = type;
            WebApplicationFirewallPolicyLink = webApplicationFirewallPolicyLink;
        }
    }
}
