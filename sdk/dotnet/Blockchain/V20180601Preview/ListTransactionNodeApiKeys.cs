// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Blockchain.V20180601Preview
{
    public static class ListTransactionNodeApiKeys
    {
        public static Task<ListTransactionNodeApiKeysResult> InvokeAsync(ListTransactionNodeApiKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListTransactionNodeApiKeysResult>("azurerm:blockchain/v20180601preview:listTransactionNodeApiKeys", args ?? new ListTransactionNodeApiKeysArgs(), options.WithVersion());
    }


    public sealed class ListTransactionNodeApiKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Blockchain member name.
        /// </summary>
        [Input("blockchainMemberName", required: true)]
        public string BlockchainMemberName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Transaction node name.
        /// </summary>
        [Input("transactionNodeName", required: true)]
        public string TransactionNodeName { get; set; } = null!;

        public ListTransactionNodeApiKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListTransactionNodeApiKeysResult
    {
        /// <summary>
        /// Gets or sets the collection of API key.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiKeyResponseResult> Keys;

        [OutputConstructor]
        private ListTransactionNodeApiKeysResult(ImmutableArray<Outputs.ApiKeyResponseResult> keys)
        {
            Keys = keys;
        }
    }
}
