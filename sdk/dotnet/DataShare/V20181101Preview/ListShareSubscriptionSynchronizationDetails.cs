// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataShare.V20181101Preview
{
    public static class ListShareSubscriptionSynchronizationDetails
    {
        public static Task<ListShareSubscriptionSynchronizationDetailsResult> InvokeAsync(ListShareSubscriptionSynchronizationDetailsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListShareSubscriptionSynchronizationDetailsResult>("azurerm:datashare/v20181101preview:listShareSubscriptionSynchronizationDetails", args ?? new ListShareSubscriptionSynchronizationDetailsArgs(), options.WithVersion());
    }


    public sealed class ListShareSubscriptionSynchronizationDetailsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Filters the results using OData syntax.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// Sorts the results using OData syntax.
        /// </summary>
        [Input("orderby")]
        public string? Orderby { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share subscription.
        /// </summary>
        [Input("shareSubscriptionName", required: true)]
        public string ShareSubscriptionName { get; set; } = null!;

        /// <summary>
        /// Continuation token
        /// </summary>
        [Input("skipToken")]
        public string? SkipToken { get; set; }

        /// <summary>
        /// Synchronization id
        /// </summary>
        [Input("synchronizationId", required: true)]
        public string SynchronizationId { get; set; } = null!;

        public ListShareSubscriptionSynchronizationDetailsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListShareSubscriptionSynchronizationDetailsResult
    {
        /// <summary>
        /// The Url of next result page.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// Collection of items of type DataTransferObjects.
        /// </summary>
        public readonly ImmutableArray<Outputs.SynchronizationDetailsResponseResult> Value;

        [OutputConstructor]
        private ListShareSubscriptionSynchronizationDetailsResult(
            string? nextLink,

            ImmutableArray<Outputs.SynchronizationDetailsResponseResult> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}
