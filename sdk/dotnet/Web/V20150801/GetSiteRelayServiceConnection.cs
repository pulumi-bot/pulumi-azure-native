// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801
{
    public static class GetSiteRelayServiceConnection
    {
        public static Task<GetSiteRelayServiceConnectionResult> InvokeAsync(GetSiteRelayServiceConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSiteRelayServiceConnectionResult>("azurerm:web/v20150801:getSiteRelayServiceConnection", args ?? new GetSiteRelayServiceConnectionArgs(), options.WithVersion());
    }


    public sealed class GetSiteRelayServiceConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name by which the Hybrid Connection is identified
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSiteRelayServiceConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSiteRelayServiceConnectionResult
    {
        /// <summary>
        /// Kind of resource
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string? Name;
        public readonly Outputs.RelayServiceConnectionEntityResponsePropertiesResult Properties;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetSiteRelayServiceConnectionResult(
            string? kind,

            string location,

            string? name,

            Outputs.RelayServiceConnectionEntityResponsePropertiesResult properties,

            ImmutableDictionary<string, string>? tags,

            string? type)
        {
            Kind = kind;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}