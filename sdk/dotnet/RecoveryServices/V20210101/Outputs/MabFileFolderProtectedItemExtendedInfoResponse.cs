// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20210101.Outputs
{

    [OutputType]
    public sealed class MabFileFolderProtectedItemExtendedInfoResponse
    {
        /// <summary>
        /// Last time when the agent data synced to service.
        /// </summary>
        public readonly string? LastRefreshedAt;
        /// <summary>
        /// The oldest backup copy available.
        /// </summary>
        public readonly string? OldestRecoveryPoint;
        /// <summary>
        /// Number of backup copies associated with the backup item.
        /// </summary>
        public readonly int? RecoveryPointCount;

        [OutputConstructor]
        private MabFileFolderProtectedItemExtendedInfoResponse(
            string? lastRefreshedAt,

            string? oldestRecoveryPoint,

            int? recoveryPointCount)
        {
            LastRefreshedAt = lastRefreshedAt;
            OldestRecoveryPoint = oldestRecoveryPoint;
            RecoveryPointCount = recoveryPointCount;
        }
    }
}
