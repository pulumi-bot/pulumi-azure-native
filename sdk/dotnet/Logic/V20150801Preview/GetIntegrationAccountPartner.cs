// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20150801Preview
{
    public static class GetIntegrationAccountPartner
    {
        public static Task<GetIntegrationAccountPartnerResult> InvokeAsync(GetIntegrationAccountPartnerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationAccountPartnerResult>("azurerm:logic/v20150801preview:getIntegrationAccountPartner", args ?? new GetIntegrationAccountPartnerArgs(), options.WithVersion());
    }


    public sealed class GetIntegrationAccountPartnerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public string IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The integration account partner name.
        /// </summary>
        [Input("partnerName", required: true)]
        public string PartnerName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetIntegrationAccountPartnerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntegrationAccountPartnerResult
    {
        /// <summary>
        /// The changed time.
        /// </summary>
        public readonly string ChangedTime;
        /// <summary>
        /// The partner content.
        /// </summary>
        public readonly Outputs.PartnerContentResponseResult? Content;
        /// <summary>
        /// The created time.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Metadata;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The partner type.
        /// </summary>
        public readonly string? PartnerType;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetIntegrationAccountPartnerResult(
            string changedTime,

            Outputs.PartnerContentResponseResult? content,

            string createdTime,

            string? location,

            ImmutableDictionary<string, object>? metadata,

            string? name,

            string? partnerType,

            ImmutableDictionary<string, string>? tags,

            string? type)
        {
            ChangedTime = changedTime;
            Content = content;
            CreatedTime = createdTime;
            Location = location;
            Metadata = metadata;
            Name = name;
            PartnerType = partnerType;
            Tags = tags;
            Type = type;
        }
    }
}
