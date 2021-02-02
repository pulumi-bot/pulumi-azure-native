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
    public sealed class DPMProtectedItemResponse
    {
        /// <summary>
        /// Backup Management server protecting this backup item
        /// </summary>
        public readonly string? BackupEngineName;
        /// <summary>
        /// Type of backup management for the backed up item.
        /// </summary>
        public readonly string? BackupManagementType;
        /// <summary>
        /// Name of the backup set the backup item belongs to
        /// </summary>
        public readonly string? BackupSetName;
        /// <summary>
        /// Unique name of container
        /// </summary>
        public readonly string? ContainerName;
        /// <summary>
        /// Create mode to indicate recovery of existing soft deleted data source or creation of new data source.
        /// </summary>
        public readonly string? CreateMode;
        /// <summary>
        /// Time for deferred deletion in UTC
        /// </summary>
        public readonly string? DeferredDeleteTimeInUTC;
        /// <summary>
        /// Time remaining before the DS marked for deferred delete is permanently deleted
        /// </summary>
        public readonly string? DeferredDeleteTimeRemaining;
        /// <summary>
        /// Extended info of the backup item.
        /// </summary>
        public readonly Outputs.DPMProtectedItemExtendedInfoResponse? ExtendedInfo;
        /// <summary>
        /// Friendly name of the managed item
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// Flag to identify whether the deferred deleted DS is to be purged soon
        /// </summary>
        public readonly bool? IsDeferredDeleteScheduleUpcoming;
        /// <summary>
        /// Flag to identify that deferred deleted DS is to be moved into Pause state
        /// </summary>
        public readonly bool? IsRehydrate;
        /// <summary>
        /// Flag to identify whether the DS is scheduled for deferred delete
        /// </summary>
        public readonly bool? IsScheduledForDeferredDelete;
        /// <summary>
        /// Timestamp when the last (latest) backup copy was created for this backup item.
        /// </summary>
        public readonly string? LastRecoveryPoint;
        /// <summary>
        /// ID of the backup policy with which this item is backed up.
        /// </summary>
        public readonly string? PolicyId;
        /// <summary>
        /// backup item type.
        /// Expected value is 'DPMProtectedItem'.
        /// </summary>
        public readonly string ProtectedItemType;
        /// <summary>
        /// Protection state of the backup engine
        /// </summary>
        public readonly string? ProtectionState;
        /// <summary>
        /// ARM ID of the resource to be backed up.
        /// </summary>
        public readonly string? SourceResourceId;
        /// <summary>
        /// Type of workload this item represents.
        /// </summary>
        public readonly string? WorkloadType;

        [OutputConstructor]
        private DPMProtectedItemResponse(
            string? backupEngineName,

            string? backupManagementType,

            string? backupSetName,

            string? containerName,

            string? createMode,

            string? deferredDeleteTimeInUTC,

            string? deferredDeleteTimeRemaining,

            Outputs.DPMProtectedItemExtendedInfoResponse? extendedInfo,

            string? friendlyName,

            bool? isDeferredDeleteScheduleUpcoming,

            bool? isRehydrate,

            bool? isScheduledForDeferredDelete,

            string? lastRecoveryPoint,

            string? policyId,

            string protectedItemType,

            string? protectionState,

            string? sourceResourceId,

            string? workloadType)
        {
            BackupEngineName = backupEngineName;
            BackupManagementType = backupManagementType;
            BackupSetName = backupSetName;
            ContainerName = containerName;
            CreateMode = createMode;
            DeferredDeleteTimeInUTC = deferredDeleteTimeInUTC;
            DeferredDeleteTimeRemaining = deferredDeleteTimeRemaining;
            ExtendedInfo = extendedInfo;
            FriendlyName = friendlyName;
            IsDeferredDeleteScheduleUpcoming = isDeferredDeleteScheduleUpcoming;
            IsRehydrate = isRehydrate;
            IsScheduledForDeferredDelete = isScheduledForDeferredDelete;
            LastRecoveryPoint = lastRecoveryPoint;
            PolicyId = policyId;
            ProtectedItemType = protectedItemType;
            ProtectionState = protectionState;
            SourceResourceId = sourceResourceId;
            WorkloadType = workloadType;
        }
    }
}
