// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Maps.V20200201Preview
{
    public static class GetPrivateAtlase
    {
        public static Task<GetPrivateAtlaseResult> InvokeAsync(GetPrivateAtlaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateAtlaseResult>("azure-nextgen:maps/v20200201preview:getPrivateAtlase", args ?? new GetPrivateAtlaseArgs(), options.WithVersion());
    }


    public sealed class GetPrivateAtlaseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Maps Account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Private Atlas instance.
        /// </summary>
        [Input("privateAtlasName", required: true)]
        public string PrivateAtlasName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateAtlaseArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateAtlaseResult
    {
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Private Atlas resource properties.
        /// </summary>
        public readonly Outputs.PrivateAtlasPropertiesResponse Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateAtlaseResult(
            string location,

            string name,

            Outputs.PrivateAtlasPropertiesResponse properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
