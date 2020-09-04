// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20150801Preview
{
    public static class GetIntegrationAccountMap
    {
        public static Task<GetIntegrationAccountMapResult> InvokeAsync(GetIntegrationAccountMapArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationAccountMapResult>("azurerm:logic/v20150801preview:getIntegrationAccountMap", args ?? new GetIntegrationAccountMapArgs(), options.WithVersion());
    }


    public sealed class GetIntegrationAccountMapArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public string IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The integration account map name.
        /// </summary>
        [Input("mapName", required: true)]
        public string MapName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetIntegrationAccountMapArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntegrationAccountMapResult
    {
        /// <summary>
        /// The changed time.
        /// </summary>
        public readonly string ChangedTime;
        /// <summary>
        /// The content.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Content;
        /// <summary>
        /// The content link.
        /// </summary>
        public readonly Outputs.IntegrationAccountContentLinkResponseResult ContentLink;
        /// <summary>
        /// The content type.
        /// </summary>
        public readonly string? ContentType;
        /// <summary>
        /// The created time.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The map type.
        /// </summary>
        public readonly string? MapType;
        /// <summary>
        /// The metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Metadata;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetIntegrationAccountMapResult(
            string changedTime,

            ImmutableDictionary<string, object>? content,

            Outputs.IntegrationAccountContentLinkResponseResult contentLink,

            string? contentType,

            string createdTime,

            string? location,

            string? mapType,

            ImmutableDictionary<string, object>? metadata,

            string? name,

            ImmutableDictionary<string, string>? tags,

            string? type)
        {
            ChangedTime = changedTime;
            Content = content;
            ContentLink = contentLink;
            ContentType = contentType;
            CreatedTime = createdTime;
            Location = location;
            MapType = mapType;
            Metadata = metadata;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
}
