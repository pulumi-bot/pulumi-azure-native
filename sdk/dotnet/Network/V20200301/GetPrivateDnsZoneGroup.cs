// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200301
{
    public static class GetPrivateDnsZoneGroup
    {
        public static Task<GetPrivateDnsZoneGroupResult> InvokeAsync(GetPrivateDnsZoneGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateDnsZoneGroupResult>("azurerm:network/v20200301:getPrivateDnsZoneGroup", args ?? new GetPrivateDnsZoneGroupArgs(), options.WithVersion());
    }


    public sealed class GetPrivateDnsZoneGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the private dns zone group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the private endpoint.
        /// </summary>
        [Input("privateEndpointName", required: true)]
        public string PrivateEndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateDnsZoneGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateDnsZoneGroupResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Properties of the private dns zone group.
        /// </summary>
        public readonly Outputs.PrivateDnsZoneGroupPropertiesFormatResponseResult Properties;

        [OutputConstructor]
        private GetPrivateDnsZoneGroupResult(
            string etag,

            string? name,

            Outputs.PrivateDnsZoneGroupPropertiesFormatResponseResult properties)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
        }
    }
}