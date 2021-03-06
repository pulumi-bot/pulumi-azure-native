// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.Outputs
{

    [OutputType]
    public sealed class DataDisksResponse
    {
        /// <summary>
        /// Caching type for the disks. Available values are none (default), readonly, readwrite. Caching type can be set only for VM sizes supporting premium storage.
        /// </summary>
        public readonly string? CachingType;
        /// <summary>
        /// Number of data disks attached to the File Server. If multiple disks attached, they will be configured in RAID level 0.
        /// </summary>
        public readonly int DiskCount;
        /// <summary>
        /// Disk size in GB for the blank data disks.
        /// </summary>
        public readonly int DiskSizeInGB;
        /// <summary>
        /// Type of storage account to be used on the disk. Possible values are: Standard_LRS or Premium_LRS. Premium storage account type can only be used with VM sizes supporting premium storage.
        /// </summary>
        public readonly string StorageAccountType;

        [OutputConstructor]
        private DataDisksResponse(
            string? cachingType,

            int diskCount,

            int diskSizeInGB,

            string storageAccountType)
        {
            CachingType = cachingType;
            DiskCount = diskCount;
            DiskSizeInGB = diskSizeInGB;
            StorageAccountType = storageAccountType;
        }
    }
}
