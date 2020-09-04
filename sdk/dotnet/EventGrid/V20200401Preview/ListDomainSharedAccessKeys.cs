// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200401Preview
{
    public static class ListDomainSharedAccessKeys
    {
        public static Task<ListDomainSharedAccessKeysResult> InvokeAsync(ListDomainSharedAccessKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListDomainSharedAccessKeysResult>("azurerm:eventgrid/v20200401preview:listDomainSharedAccessKeys", args ?? new ListDomainSharedAccessKeysArgs(), options.WithVersion());
    }


    public sealed class ListDomainSharedAccessKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListDomainSharedAccessKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListDomainSharedAccessKeysResult
    {
        /// <summary>
        /// Shared access key1 for the domain.
        /// </summary>
        public readonly string? Key1;
        /// <summary>
        /// Shared access key2 for the domain.
        /// </summary>
        public readonly string? Key2;

        [OutputConstructor]
        private ListDomainSharedAccessKeysResult(
            string? key1,

            string? key2)
        {
            Key1 = key1;
            Key2 = key2;
        }
    }
}
