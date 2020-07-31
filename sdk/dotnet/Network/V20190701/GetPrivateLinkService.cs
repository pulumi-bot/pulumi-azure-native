// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701
{
    public static class GetPrivateLinkService
    {
        public static Task<GetPrivateLinkServiceResult> InvokeAsync(GetPrivateLinkServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateLinkServiceResult>("azurerm:network/v20190701:getPrivateLinkService", args ?? new GetPrivateLinkServiceArgs(), options.WithVersion());
    }


    public sealed class GetPrivateLinkServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the private link service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateLinkServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateLinkServiceResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the private link service.
        /// </summary>
        public readonly Outputs.PrivateLinkServicePropertiesResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateLinkServiceResult(
            string? etag,

            string? location,

            string name,

            Outputs.PrivateLinkServicePropertiesResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}