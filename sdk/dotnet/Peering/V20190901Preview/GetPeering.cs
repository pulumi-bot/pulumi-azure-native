// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.V20190901Preview
{
    public static class GetPeering
    {
        public static Task<GetPeeringResult> InvokeAsync(GetPeeringArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPeeringResult>("azurerm:peering/v20190901preview:getPeering", args ?? new GetPeeringArgs(), options.WithVersion());
    }


    public sealed class GetPeeringArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public string PeeringName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPeeringArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPeeringResult
    {
        /// <summary>
        /// The properties that define a direct peering.
        /// </summary>
        public readonly Outputs.PeeringPropertiesDirectResponseResult? Direct;
        /// <summary>
        /// The properties that define an exchange peering.
        /// </summary>
        public readonly Outputs.PeeringPropertiesExchangeResponseResult? Exchange;
        /// <summary>
        /// The kind of the peering.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The location of the resource.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The location of the peering.
        /// </summary>
        public readonly string? PeeringLocation;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The SKU that defines the tier and kind of the peering.
        /// </summary>
        public readonly Outputs.PeeringSkuResponseResult Sku;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPeeringResult(
            Outputs.PeeringPropertiesDirectResponseResult? direct,

            Outputs.PeeringPropertiesExchangeResponseResult? exchange,

            string kind,

            string location,

            string name,

            string? peeringLocation,

            string provisioningState,

            Outputs.PeeringSkuResponseResult sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Direct = direct;
            Exchange = exchange;
            Kind = kind;
            Location = location;
            Name = name;
            PeeringLocation = peeringLocation;
            ProvisioningState = provisioningState;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}
