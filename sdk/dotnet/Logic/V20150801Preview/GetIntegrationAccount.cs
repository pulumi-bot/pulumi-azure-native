// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20150801Preview
{
    public static class GetIntegrationAccount
    {
        public static Task<GetIntegrationAccountResult> InvokeAsync(GetIntegrationAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationAccountResult>("azurerm:logic/v20150801preview:getIntegrationAccount", args ?? new GetIntegrationAccountArgs(), options.WithVersion());
    }


    public sealed class GetIntegrationAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public string IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetIntegrationAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntegrationAccountResult
    {
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The sku.
        /// </summary>
        public readonly Outputs.IntegrationAccountSkuResponseResult? Sku;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetIntegrationAccountResult(
            string? location,

            string? name,

            Outputs.IntegrationAccountSkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string? type)
        {
            Location = location;
            Name = name;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
