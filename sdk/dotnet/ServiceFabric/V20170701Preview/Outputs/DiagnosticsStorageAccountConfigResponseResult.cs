// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20170701Preview.Outputs
{

    [OutputType]
    public sealed class DiagnosticsStorageAccountConfigResponseResult
    {
        /// <summary>
        /// The blob endpoint of the azure storage account.
        /// </summary>
        public readonly string BlobEndpoint;
        /// <summary>
        /// The protected diagnostics storage key name.
        /// </summary>
        public readonly string ProtectedAccountKeyName;
        /// <summary>
        /// The queue endpoint of the azure storage account.
        /// </summary>
        public readonly string QueueEndpoint;
        /// <summary>
        /// The Azure storage account name.
        /// </summary>
        public readonly string StorageAccountName;
        /// <summary>
        /// The table endpoint of the azure storage account.
        /// </summary>
        public readonly string TableEndpoint;

        [OutputConstructor]
        private DiagnosticsStorageAccountConfigResponseResult(
            string blobEndpoint,

            string protectedAccountKeyName,

            string queueEndpoint,

            string storageAccountName,

            string tableEndpoint)
        {
            BlobEndpoint = blobEndpoint;
            ProtectedAccountKeyName = protectedAccountKeyName;
            QueueEndpoint = queueEndpoint;
            StorageAccountName = storageAccountName;
            TableEndpoint = tableEndpoint;
        }
    }
}
