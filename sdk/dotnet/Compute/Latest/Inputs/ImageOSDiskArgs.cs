// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.Latest.Inputs
{

    /// <summary>
    /// Describes an Operating System disk.
    /// </summary>
    public sealed class ImageOSDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Virtual Hard Disk.
        /// </summary>
        [Input("blobUri")]
        public Input<string>? BlobUri { get; set; }

        /// <summary>
        /// Specifies the caching requirements. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **None** &lt;br&gt;&lt;br&gt; **ReadOnly** &lt;br&gt;&lt;br&gt; **ReadWrite** &lt;br&gt;&lt;br&gt; Default: **None for Standard storage. ReadOnly for Premium storage**
        /// </summary>
        [Input("caching")]
        public Input<Pulumi.AzureNextGen.Compute.Latest.CachingTypes>? Caching { get; set; }

        /// <summary>
        /// Specifies the customer managed disk encryption set resource id for the managed image disk.
        /// </summary>
        [Input("diskEncryptionSet")]
        public Input<Inputs.DiskEncryptionSetParametersArgs>? DiskEncryptionSet { get; set; }

        /// <summary>
        /// Specifies the size of empty data disks in gigabytes. This element can be used to overwrite the name of the disk in a virtual machine image. &lt;br&gt;&lt;br&gt; This value cannot be larger than 1023 GB
        /// </summary>
        [Input("diskSizeGB")]
        public Input<int>? DiskSizeGB { get; set; }

        /// <summary>
        /// The managedDisk.
        /// </summary>
        [Input("managedDisk")]
        public Input<Inputs.SubResourceArgs>? ManagedDisk { get; set; }

        /// <summary>
        /// The OS State.
        /// </summary>
        [Input("osState", required: true)]
        public Input<Pulumi.AzureNextGen.Compute.Latest.OperatingSystemStateTypes> OsState { get; set; } = null!;

        /// <summary>
        /// This property allows you to specify the type of the OS that is included in the disk if creating a VM from a custom image. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;&lt;br&gt; **Windows** &lt;br&gt;&lt;br&gt; **Linux**
        /// </summary>
        [Input("osType", required: true)]
        public Input<Pulumi.AzureNextGen.Compute.Latest.OperatingSystemTypes> OsType { get; set; } = null!;

        /// <summary>
        /// The snapshot.
        /// </summary>
        [Input("snapshot")]
        public Input<Inputs.SubResourceArgs>? Snapshot { get; set; }

        /// <summary>
        /// Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with OS Disk.
        /// </summary>
        [Input("storageAccountType")]
        public InputUnion<string, Pulumi.AzureNextGen.Compute.Latest.StorageAccountTypes>? StorageAccountType { get; set; }

        public ImageOSDiskArgs()
        {
        }
    }
}
