// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Latest
{
    [Obsolete(@"The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:web:getConnectionGateway'.")]
    public static class GetConnectionGateway
    {
        /// <summary>
        /// The gateway definition
        /// Latest API Version: 2016-06-01.
        /// </summary>
        public static Task<GetConnectionGatewayResult> InvokeAsync(GetConnectionGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionGatewayResult>("azure-native:web/latest:getConnectionGateway", args ?? new GetConnectionGatewayArgs(), options.WithVersion());
    }


    public sealed class GetConnectionGatewayArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The connection gateway name
        /// </summary>
        [Input("connectionGatewayName", required: true)]
        public string ConnectionGatewayName { get; set; } = null!;

        /// <summary>
        /// The resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Subscription Id
        /// </summary>
        [Input("subscriptionId")]
        public string? SubscriptionId { get; set; }

        public GetConnectionGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionGatewayResult
    {
        /// <summary>
        /// Resource ETag
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        public readonly Outputs.ConnectionGatewayDefinitionResponseProperties Properties;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConnectionGatewayResult(
            string? etag,

            string id,

            string? location,

            string name,

            Outputs.ConnectionGatewayDefinitionResponseProperties properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
