// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MixedReality.V20191202Preview
{
    public static class GetSpatialAnchorsAccount
    {
        public static Task<GetSpatialAnchorsAccountResult> InvokeAsync(GetSpatialAnchorsAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSpatialAnchorsAccountResult>("azurerm:mixedreality/v20191202preview:getSpatialAnchorsAccount", args ?? new GetSpatialAnchorsAccountArgs(), options.WithVersion());
    }


    public sealed class GetSpatialAnchorsAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of an Mixed Reality Account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSpatialAnchorsAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSpatialAnchorsAccountResult
    {
        /// <summary>
        /// Correspond domain name of certain Spatial Anchors Account
        /// </summary>
        public readonly string AccountDomain;
        /// <summary>
        /// unique id of certain account.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSpatialAnchorsAccountResult(
            string accountDomain,

            string accountId,

            string location,

            string name,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AccountDomain = accountDomain;
            AccountId = accountId;
            Location = location;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
}
