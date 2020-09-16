// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20200901.Outputs
{

    [OutputType]
    public sealed class AzureBlobFileSystemConfigurationResponse
    {
        /// <summary>
        /// This property is mutually exclusive with sasKey and one must be specified.
        /// </summary>
        public readonly string? AccountKey;
        public readonly string AccountName;
        /// <summary>
        /// These are 'net use' options in Windows and 'mount' options in Linux.
        /// </summary>
        public readonly string? BlobfuseOptions;
        public readonly string ContainerName;
        /// <summary>
        /// All file systems are mounted relative to the Batch mounts directory, accessible via the AZ_BATCH_NODE_MOUNTS_DIR environment variable.
        /// </summary>
        public readonly string RelativeMountPath;
        /// <summary>
        /// This property is mutually exclusive with accountKey and one must be specified.
        /// </summary>
        public readonly string? SasKey;

        [OutputConstructor]
        private AzureBlobFileSystemConfigurationResponse(
            string? accountKey,

            string accountName,

            string? blobfuseOptions,

            string containerName,

            string relativeMountPath,

            string? sasKey)
        {
            AccountKey = accountKey;
            AccountName = accountName;
            BlobfuseOptions = blobfuseOptions;
            ContainerName = containerName;
            RelativeMountPath = relativeMountPath;
            SasKey = sasKey;
        }
    }
}
