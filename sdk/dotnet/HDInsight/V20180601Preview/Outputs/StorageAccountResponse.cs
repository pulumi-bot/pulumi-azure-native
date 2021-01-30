// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HDInsight.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class StorageAccountResponse
    {
        /// <summary>
        /// The container in the storage account, only to be specified for WASB storage accounts.
        /// </summary>
        public readonly string? Container;
        /// <summary>
        /// The filesystem, only to be specified for Azure Data Lake Storage Gen 2.
        /// </summary>
        public readonly string? FileSystem;
        /// <summary>
        /// The file share name.
        /// </summary>
        public readonly string? Fileshare;
        /// <summary>
        /// Whether or not the storage account is the default storage account.
        /// </summary>
        public readonly bool? IsDefault;
        /// <summary>
        /// The storage account access key.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The managed identity (MSI) that is allowed to access the storage account, only to be specified for Azure Data Lake Storage Gen 2.
        /// </summary>
        public readonly string? MsiResourceId;
        /// <summary>
        /// The name of the storage account.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The resource ID of storage account, only to be specified for Azure Data Lake Storage Gen 2.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// The shared access signature key.
        /// </summary>
        public readonly string? Saskey;

        [OutputConstructor]
        private StorageAccountResponse(
            string? container,

            string? fileSystem,

            string? fileshare,

            bool? isDefault,

            string? key,

            string? msiResourceId,

            string? name,

            string? resourceId,

            string? saskey)
        {
            Container = container;
            FileSystem = fileSystem;
            Fileshare = fileshare;
            IsDefault = isDefault;
            Key = key;
            MsiResourceId = msiResourceId;
            Name = name;
            ResourceId = resourceId;
            Saskey = saskey;
        }
    }
}
