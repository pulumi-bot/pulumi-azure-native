// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageSync.Latest.Outputs
{

    [OutputType]
    public sealed class ServerEndpointCloudTieringStatusResponseResult
    {
        /// <summary>
        /// Information regarding how well the local cache on the server is performing.
        /// </summary>
        public readonly Outputs.CloudTieringCachePerformanceResponseResult CachePerformance;
        /// <summary>
        /// Status of the date policy
        /// </summary>
        public readonly Outputs.CloudTieringDatePolicyStatusResponseResult DatePolicyStatus;
        /// <summary>
        /// Information regarding files that failed to be tiered
        /// </summary>
        public readonly Outputs.CloudTieringFilesNotTieringResponseResult FilesNotTiering;
        /// <summary>
        /// Cloud tiering health state.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// The last updated timestamp of health state
        /// </summary>
        public readonly string HealthLastUpdatedTimestamp;
        /// <summary>
        /// Last cloud tiering result (HResult)
        /// </summary>
        public readonly int LastCloudTieringResult;
        /// <summary>
        /// Last cloud tiering success timestamp
        /// </summary>
        public readonly string LastSuccessTimestamp;
        /// <summary>
        /// Last updated timestamp
        /// </summary>
        public readonly string LastUpdatedTimestamp;
        /// <summary>
        /// Information regarding how much local space cloud tiering is saving.
        /// </summary>
        public readonly Outputs.CloudTieringSpaceSavingsResponseResult SpaceSavings;
        /// <summary>
        /// Status of the volume free space policy
        /// </summary>
        public readonly Outputs.CloudTieringVolumeFreeSpacePolicyStatusResponseResult VolumeFreeSpacePolicyStatus;

        [OutputConstructor]
        private ServerEndpointCloudTieringStatusResponseResult(
            Outputs.CloudTieringCachePerformanceResponseResult cachePerformance,

            Outputs.CloudTieringDatePolicyStatusResponseResult datePolicyStatus,

            Outputs.CloudTieringFilesNotTieringResponseResult filesNotTiering,

            string health,

            string healthLastUpdatedTimestamp,

            int lastCloudTieringResult,

            string lastSuccessTimestamp,

            string lastUpdatedTimestamp,

            Outputs.CloudTieringSpaceSavingsResponseResult spaceSavings,

            Outputs.CloudTieringVolumeFreeSpacePolicyStatusResponseResult volumeFreeSpacePolicyStatus)
        {
            CachePerformance = cachePerformance;
            DatePolicyStatus = datePolicyStatus;
            FilesNotTiering = filesNotTiering;
            Health = health;
            HealthLastUpdatedTimestamp = healthLastUpdatedTimestamp;
            LastCloudTieringResult = lastCloudTieringResult;
            LastSuccessTimestamp = lastSuccessTimestamp;
            LastUpdatedTimestamp = lastUpdatedTimestamp;
            SpaceSavings = spaceSavings;
            VolumeFreeSpacePolicyStatus = volumeFreeSpacePolicyStatus;
        }
    }
}