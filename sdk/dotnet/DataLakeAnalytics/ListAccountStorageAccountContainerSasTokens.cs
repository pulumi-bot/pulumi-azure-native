// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataLakeAnalytics
{
    public static class ListAccountStorageAccountContainerSasTokens
    {
        public static Task<ListAccountStorageAccountContainerSasTokensResult> InvokeAsync(ListAccountStorageAccountContainerSasTokensArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListAccountStorageAccountContainerSasTokensResult>("azurerm:datalakeanalytics:listAccountStorageAccountContainerSasTokens", args ?? new ListAccountStorageAccountContainerSasTokensArgs(), options.WithVersion());
    }


    public sealed class ListAccountStorageAccountContainerSasTokensArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Data Lake Analytics account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure storage container for which the SAS token is being requested.
        /// </summary>
        [Input("containerName", required: true)]
        public string ContainerName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure storage account for which the SAS token is being requested.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public string StorageAccountName { get; set; } = null!;

        public ListAccountStorageAccountContainerSasTokensArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListAccountStorageAccountContainerSasTokensResult
    {
        /// <summary>
        /// The link (url) to the next page of results.
        /// </summary>
        public readonly string NextLink;
        /// <summary>
        /// The results of the list operation.
        /// </summary>
        public readonly ImmutableArray<Outputs.SasTokenInformationResponseResult> Value;

        [OutputConstructor]
        private ListAccountStorageAccountContainerSasTokensResult(
            string nextLink,

            ImmutableArray<Outputs.SasTokenInformationResponseResult> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}