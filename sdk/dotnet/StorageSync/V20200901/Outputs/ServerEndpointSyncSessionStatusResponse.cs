// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageSync.V20200901.Outputs
{

    [OutputType]
    public sealed class ServerEndpointSyncSessionStatusResponse
    {
        /// <summary>
        /// Array of per-item errors coming from the last sync session.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerEndpointFilesNotSyncingErrorResponse> FilesNotSyncingErrors;
        /// <summary>
        /// Sync mode
        /// </summary>
        public readonly string LastSyncMode;
        /// <summary>
        /// Last sync per item error count.
        /// </summary>
        public readonly int LastSyncPerItemErrorCount;
        /// <summary>
        /// Last sync result (HResult)
        /// </summary>
        public readonly int LastSyncResult;
        /// <summary>
        /// Last sync success timestamp
        /// </summary>
        public readonly string LastSyncSuccessTimestamp;
        /// <summary>
        /// Last sync timestamp
        /// </summary>
        public readonly string LastSyncTimestamp;
        /// <summary>
        /// Count of persistent files not syncing.
        /// </summary>
        public readonly int PersistentFilesNotSyncingCount;
        /// <summary>
        /// Count of transient files not syncing.
        /// </summary>
        public readonly int TransientFilesNotSyncingCount;

        [OutputConstructor]
        private ServerEndpointSyncSessionStatusResponse(
            ImmutableArray<Outputs.ServerEndpointFilesNotSyncingErrorResponse> filesNotSyncingErrors,

            string lastSyncMode,

            int lastSyncPerItemErrorCount,

            int lastSyncResult,

            string lastSyncSuccessTimestamp,

            string lastSyncTimestamp,

            int persistentFilesNotSyncingCount,

            int transientFilesNotSyncingCount)
        {
            FilesNotSyncingErrors = filesNotSyncingErrors;
            LastSyncMode = lastSyncMode;
            LastSyncPerItemErrorCount = lastSyncPerItemErrorCount;
            LastSyncResult = lastSyncResult;
            LastSyncSuccessTimestamp = lastSyncSuccessTimestamp;
            LastSyncTimestamp = lastSyncTimestamp;
            PersistentFilesNotSyncingCount = persistentFilesNotSyncingCount;
            TransientFilesNotSyncingCount = transientFilesNotSyncingCount;
        }
    }
}
