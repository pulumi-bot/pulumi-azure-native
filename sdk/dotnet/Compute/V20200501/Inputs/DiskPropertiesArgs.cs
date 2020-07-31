// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20200501.Inputs
{

    /// <summary>
    /// Disk resource properties.
    /// </summary>
    public sealed class DiskPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        [Input("creationData", required: true)]
        public Input<Inputs.CreationDataArgs> CreationData { get; set; } = null!;

        /// <summary>
        /// ARM id of the DiskAccess resource for using private endpoints on disks.
        /// </summary>
        [Input("diskAccessId")]
        public Input<string>? DiskAccessId { get; set; }

        /// <summary>
        /// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Input("diskIOPSReadOnly")]
        public Input<int>? DiskIOPSReadOnly { get; set; }

        /// <summary>
        /// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Input("diskIOPSReadWrite")]
        public Input<int>? DiskIOPSReadWrite { get; set; }

        /// <summary>
        /// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        /// </summary>
        [Input("diskMBpsReadOnly")]
        public Input<int>? DiskMBpsReadOnly { get; set; }

        /// <summary>
        /// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        /// </summary>
        [Input("diskMBpsReadWrite")]
        public Input<int>? DiskMBpsReadWrite { get; set; }

        /// <summary>
        /// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        /// </summary>
        [Input("diskSizeGB")]
        public Input<int>? DiskSizeGB { get; set; }

        /// <summary>
        /// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.EncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
        /// </summary>
        [Input("encryptionSettingsCollection")]
        public Input<Inputs.EncryptionSettingsCollectionArgs>? EncryptionSettingsCollection { get; set; }

        /// <summary>
        /// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        /// </summary>
        [Input("hyperVGeneration")]
        public Input<string>? HyperVGeneration { get; set; }

        /// <summary>
        /// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
        /// </summary>
        [Input("maxShares")]
        public Input<int>? MaxShares { get; set; }

        /// <summary>
        /// Policy for accessing the disk via network.
        /// </summary>
        [Input("networkAccessPolicy")]
        public Input<string>? NetworkAccessPolicy { get; set; }

        /// <summary>
        /// The Operating System type.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        public DiskPropertiesArgs()
        {
        }
    }
}