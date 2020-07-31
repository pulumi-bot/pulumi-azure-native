// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190301.Outputs
{

    [OutputType]
    public sealed class DiskPropertiesResponseResult
    {
        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        public readonly Outputs.CreationDataResponseResult CreationData;
        /// <summary>
        /// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        public readonly int? DiskIOPSReadWrite;
        /// <summary>
        /// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        /// </summary>
        public readonly int? DiskMBpsReadWrite;
        /// <summary>
        /// The size of the disk in bytes. This field is read only.
        /// </summary>
        public readonly int DiskSizeBytes;
        /// <summary>
        /// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        /// </summary>
        public readonly int? DiskSizeGB;
        /// <summary>
        /// The state of the disk.
        /// </summary>
        public readonly string DiskState;
        /// <summary>
        /// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
        /// </summary>
        public readonly Outputs.EncryptionSettingsCollectionResponseResult? EncryptionSettingsCollection;
        /// <summary>
        /// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        /// </summary>
        public readonly string? HyperVGeneration;
        /// <summary>
        /// The Operating System type.
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// The disk provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The time when the disk was created.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// Unique Guid identifying the resource.
        /// </summary>
        public readonly string UniqueId;

        [OutputConstructor]
        private DiskPropertiesResponseResult(
            Outputs.CreationDataResponseResult creationData,

            int? diskIOPSReadWrite,

            int? diskMBpsReadWrite,

            int diskSizeBytes,

            int? diskSizeGB,

            string diskState,

            Outputs.EncryptionSettingsCollectionResponseResult? encryptionSettingsCollection,

            string? hyperVGeneration,

            string? osType,

            string provisioningState,

            string timeCreated,

            string uniqueId)
        {
            CreationData = creationData;
            DiskIOPSReadWrite = diskIOPSReadWrite;
            DiskMBpsReadWrite = diskMBpsReadWrite;
            DiskSizeBytes = diskSizeBytes;
            DiskSizeGB = diskSizeGB;
            DiskState = diskState;
            EncryptionSettingsCollection = encryptionSettingsCollection;
            HyperVGeneration = hyperVGeneration;
            OsType = osType;
            ProvisioningState = provisioningState;
            TimeCreated = timeCreated;
            UniqueId = uniqueId;
        }
    }
}