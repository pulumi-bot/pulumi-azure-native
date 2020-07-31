// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180701
{
    public static class GetPublicIPAddress
    {
        public static Task<GetPublicIPAddressResult> InvokeAsync(GetPublicIPAddressArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPublicIPAddressResult>("azurerm:network/v20180701:getPublicIPAddress", args ?? new GetPublicIPAddressArgs(), options.WithVersion());
    }


    public sealed class GetPublicIPAddressArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the subnet.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPublicIPAddressArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPublicIPAddressResult
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
        /// Public IP address properties.
        /// </summary>
        public readonly Outputs.PublicIPAddressPropertiesFormatResponseResult Properties;
        /// <summary>
        /// The public IP address SKU.
        /// </summary>
        public readonly Outputs.PublicIPAddressSkuResponseResult? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A list of availability zones denoting the IP allocated for the resource needs to come from.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetPublicIPAddressResult(
            string? etag,

            string? location,

            string name,

            Outputs.PublicIPAddressPropertiesFormatResponseResult properties,

            Outputs.PublicIPAddressSkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<string> zones)
        {
            Etag = etag;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
            Zones = zones;
        }
    }
}