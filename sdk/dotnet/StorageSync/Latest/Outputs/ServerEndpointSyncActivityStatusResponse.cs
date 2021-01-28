// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageSync.Latest.Outputs
{

    [OutputType]
    public sealed class ServerEndpointSyncActivityStatusResponse
    {
        /// <summary>
        /// Applied bytes
        /// </summary>
        public readonly double AppliedBytes;
        /// <summary>
        /// Applied item count.
        /// </summary>
        public readonly double AppliedItemCount;
        /// <summary>
        /// Per item error count
        /// </summary>
        public readonly double PerItemErrorCount;
        /// <summary>
        /// Timestamp when properties were updated
        /// </summary>
        public readonly string Timestamp;
        /// <summary>
        /// Total bytes (if available)
        /// </summary>
        public readonly double TotalBytes;
        /// <summary>
        /// Total item count (if available)
        /// </summary>
        public readonly double TotalItemCount;

        [OutputConstructor]
        private ServerEndpointSyncActivityStatusResponse(
            double appliedBytes,

            double appliedItemCount,

            double perItemErrorCount,

            string timestamp,

            double totalBytes,

            double totalItemCount)
        {
            AppliedBytes = appliedBytes;
            AppliedItemCount = appliedItemCount;
            PerItemErrorCount = perItemErrorCount;
            Timestamp = timestamp;
            TotalBytes = totalBytes;
            TotalItemCount = totalItemCount;
        }
    }
}
