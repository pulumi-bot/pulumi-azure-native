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
    public sealed class AzureIaaSComputeVMProtectedItemResponse
    {
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
        /// Additional information for this backup item.
        /// </summary>
        public readonly Outputs.AzureIaaSVMProtectedItemExtendedInfoResponse? ExtendedInfo;
        /// <summary>
        /// Extended Properties for Azure IaasVM Backup.
        /// </summary>
        public readonly Outputs.ExtendedPropertiesResponse? ExtendedProperties;
        /// <summary>
        /// Friendly name of the VM represented by this backup item.
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// Health details on this backup item.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureIaaSVMHealthDetailsResponse> HealthDetails;
        /// <summary>
        /// Health status of protected item.
        /// </summary>
        public readonly string? HealthStatus;
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
        /// Health details of different KPIs
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.KPIResourceHealthDetailsResponse>? KpisHealths;
        /// <summary>
        /// Last backup operation status.
        /// </summary>
        public readonly string? LastBackupStatus;
        /// <summary>
        /// Timestamp of the last backup operation on this backup item.
        /// </summary>
        public readonly string? LastBackupTime;
        /// <summary>
        /// Timestamp when the last (latest) backup copy was created for this backup item.
        /// </summary>
        public readonly string? LastRecoveryPoint;
        /// <summary>
        /// ID of the backup policy with which this item is backed up.
        /// </summary>
        public readonly string? PolicyId;
        /// <summary>
        /// Data ID of the protected item.
        /// </summary>
        public readonly string? ProtectedItemDataId;
        /// <summary>
        /// backup item type.
        /// Expected value is 'AzureIaaSVMProtectedItem'.
        /// </summary>
        public readonly string ProtectedItemType;
        /// <summary>
        /// Backup state of this backup item.
        /// </summary>
        public readonly string? ProtectionState;
        /// <summary>
        /// Backup status of this backup item.
        /// </summary>
        public readonly string? ProtectionStatus;
        /// <summary>
        /// ARM ID of the resource to be backed up.
        /// </summary>
        public readonly string? SourceResourceId;
        /// <summary>
        /// Fully qualified ARM ID of the virtual machine represented by this item.
        /// </summary>
        public readonly string? VirtualMachineId;
        /// <summary>
        /// Type of workload this item represents.
        /// </summary>
        public readonly string? WorkloadType;

        [OutputConstructor]
        private AzureIaaSComputeVMProtectedItemResponse(
            string? backupManagementType,

            string? backupSetName,

            string? containerName,

            string? createMode,

            string? deferredDeleteTimeInUTC,

            string? deferredDeleteTimeRemaining,

            Outputs.AzureIaaSVMProtectedItemExtendedInfoResponse? extendedInfo,

            Outputs.ExtendedPropertiesResponse? extendedProperties,

            string? friendlyName,

            ImmutableArray<Outputs.AzureIaaSVMHealthDetailsResponse> healthDetails,

            string? healthStatus,

            bool? isDeferredDeleteScheduleUpcoming,

            bool? isRehydrate,

            bool? isScheduledForDeferredDelete,

            ImmutableDictionary<string, Outputs.KPIResourceHealthDetailsResponse>? kpisHealths,

            string? lastBackupStatus,

            string? lastBackupTime,

            string? lastRecoveryPoint,

            string? policyId,

            string? protectedItemDataId,

            string protectedItemType,

            string? protectionState,

            string? protectionStatus,

            string? sourceResourceId,

            string? virtualMachineId,

            string? workloadType)
        {
            BackupManagementType = backupManagementType;
            BackupSetName = backupSetName;
            ContainerName = containerName;
            CreateMode = createMode;
            DeferredDeleteTimeInUTC = deferredDeleteTimeInUTC;
            DeferredDeleteTimeRemaining = deferredDeleteTimeRemaining;
            ExtendedInfo = extendedInfo;
            ExtendedProperties = extendedProperties;
            FriendlyName = friendlyName;
            HealthDetails = healthDetails;
            HealthStatus = healthStatus;
            IsDeferredDeleteScheduleUpcoming = isDeferredDeleteScheduleUpcoming;
            IsRehydrate = isRehydrate;
            IsScheduledForDeferredDelete = isScheduledForDeferredDelete;
            KpisHealths = kpisHealths;
            LastBackupStatus = lastBackupStatus;
            LastBackupTime = lastBackupTime;
            LastRecoveryPoint = lastRecoveryPoint;
            PolicyId = policyId;
            ProtectedItemDataId = protectedItemDataId;
            ProtectedItemType = protectedItemType;
            ProtectionState = protectionState;
            ProtectionStatus = protectionStatus;
            SourceResourceId = sourceResourceId;
            VirtualMachineId = virtualMachineId;
            WorkloadType = workloadType;
        }
    }
}
