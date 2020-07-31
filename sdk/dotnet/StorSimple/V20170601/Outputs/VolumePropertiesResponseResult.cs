// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorSimple.V20170601.Outputs
{

    [OutputType]
    public sealed class VolumePropertiesResponseResult
    {
        /// <summary>
        /// The IDs of the access control records, associated with the volume.
        /// </summary>
        public readonly ImmutableArray<string> AccessControlRecordIds;
        /// <summary>
        /// The IDs of the backup policies, in which this volume is part of.
        /// </summary>
        public readonly ImmutableArray<string> BackupPolicyIds;
        /// <summary>
        /// The backup status of the volume.
        /// </summary>
        public readonly string BackupStatus;
        /// <summary>
        /// The monitoring status of the volume.
        /// </summary>
        public readonly string MonitoringStatus;
        /// <summary>
        /// The operation status on the volume.
        /// </summary>
        public readonly string OperationStatus;
        /// <summary>
        /// The size of the volume in bytes.
        /// </summary>
        public readonly int SizeInBytes;
        /// <summary>
        /// The ID of the volume container, in which this volume is created.
        /// </summary>
        public readonly string VolumeContainerId;
        /// <summary>
        /// The volume status.
        /// </summary>
        public readonly string VolumeStatus;
        /// <summary>
        /// The type of the volume.
        /// </summary>
        public readonly string VolumeType;

        [OutputConstructor]
        private VolumePropertiesResponseResult(
            ImmutableArray<string> accessControlRecordIds,

            ImmutableArray<string> backupPolicyIds,

            string backupStatus,

            string monitoringStatus,

            string operationStatus,

            int sizeInBytes,

            string volumeContainerId,

            string volumeStatus,

            string volumeType)
        {
            AccessControlRecordIds = accessControlRecordIds;
            BackupPolicyIds = backupPolicyIds;
            BackupStatus = backupStatus;
            MonitoringStatus = monitoringStatus;
            OperationStatus = operationStatus;
            SizeInBytes = sizeInBytes;
            VolumeContainerId = volumeContainerId;
            VolumeStatus = volumeStatus;
            VolumeType = volumeType;
        }
    }
}