// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.Outputs
{

    [OutputType]
    public sealed class AssetPropertiesResponseResult
    {
        /// <summary>
        /// The alternate ID of the Asset.
        /// </summary>
        public readonly string? AlternateId;
        /// <summary>
        /// The Asset ID.
        /// </summary>
        public readonly string AssetId;
        /// <summary>
        /// The name of the asset blob container.
        /// </summary>
        public readonly string? Container;
        /// <summary>
        /// The creation date of the Asset.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// The Asset description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The last modified date of the Asset.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// The name of the storage account.
        /// </summary>
        public readonly string? StorageAccountName;
        /// <summary>
        /// The Asset encryption format. One of None or MediaStorageEncryption.
        /// </summary>
        public readonly string StorageEncryptionFormat;

        [OutputConstructor]
        private AssetPropertiesResponseResult(
            string? alternateId,

            string assetId,

            string? container,

            string created,

            string? description,

            string lastModified,

            string? storageAccountName,

            string storageEncryptionFormat)
        {
            AlternateId = alternateId;
            AssetId = assetId;
            Container = container;
            Created = created;
            Description = description;
            LastModified = lastModified;
            StorageAccountName = storageAccountName;
            StorageEncryptionFormat = storageEncryptionFormat;
        }
    }
}