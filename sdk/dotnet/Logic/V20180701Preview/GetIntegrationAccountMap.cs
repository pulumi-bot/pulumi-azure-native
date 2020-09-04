// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview
{
    public static class GetIntegrationAccountMap
    {
        public static Task<GetIntegrationAccountMapResult> InvokeAsync(GetIntegrationAccountMapArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationAccountMapResult>("azurerm:logic/v20180701preview:getIntegrationAccountMap", args ?? new GetIntegrationAccountMapArgs(), options.WithVersion());
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
        public readonly string? Content;
        /// <summary>
        /// The content link.
        /// </summary>
        public readonly Outputs.ContentLinkResponseResult ContentLink;
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
        public readonly string MapType;
        /// <summary>
        /// The metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Metadata;
        /// <summary>
        /// Gets the resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The parameters schema of integration account map.
        /// </summary>
        public readonly Outputs.IntegrationAccountMapPropertiesResponseParametersSchemaResult? ParametersSchema;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Gets the resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIntegrationAccountMapResult(
            string changedTime,

            string? content,

            Outputs.ContentLinkResponseResult contentLink,

            string? contentType,

            string createdTime,

            string? location,

            string mapType,

            ImmutableDictionary<string, object>? metadata,

            string name,

            Outputs.IntegrationAccountMapPropertiesResponseParametersSchemaResult? parametersSchema,

            ImmutableDictionary<string, string>? tags,

            string type)
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
            ParametersSchema = parametersSchema;
            Tags = tags;
            Type = type;
        }
    }
}
