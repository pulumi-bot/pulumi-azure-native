// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20190801.Outputs
{

    [OutputType]
    public sealed class AzureStorageInfoValueResponseResult
    {
        /// <summary>
        /// Access key for the storage account.
        /// </summary>
        public readonly string? AccessKey;
        /// <summary>
        /// Name of the storage account.
        /// </summary>
        public readonly string? AccountName;
        /// <summary>
        /// Path to mount the storage within the site's runtime environment.
        /// </summary>
        public readonly string? MountPath;
        /// <summary>
        /// Name of the file share (container name, for Blob storage).
        /// </summary>
        public readonly string? ShareName;
        /// <summary>
        /// State of the storage account.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Type of storage.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private AzureStorageInfoValueResponseResult(
            string? accessKey,

            string? accountName,

            string? mountPath,

            string? shareName,

            string state,

            string? type)
        {
            AccessKey = accessKey;
            AccountName = accountName;
            MountPath = mountPath;
            ShareName = shareName;
            State = state;
            Type = type;
        }
    }
}
