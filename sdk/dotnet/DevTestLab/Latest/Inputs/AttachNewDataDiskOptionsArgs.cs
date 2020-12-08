// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DevTestLab.Latest.Inputs
{

    /// <summary>
    /// Properties to attach new disk to the Virtual Machine.
    /// </summary>
    public sealed class AttachNewDataDiskOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the disk to be attached.
        /// </summary>
        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        /// <summary>
        /// Size of the disk to be attached in GibiBytes.
        /// </summary>
        [Input("diskSizeGiB")]
        public Input<int>? DiskSizeGiB { get; set; }

        /// <summary>
        /// The storage type for the disk (i.e. Standard, Premium).
        /// </summary>
        [Input("diskType")]
        public InputUnion<string, Pulumi.AzureNextGen.DevTestLab.Latest.StorageType>? DiskType { get; set; }

        public AttachNewDataDiskOptionsArgs()
        {
        }
    }
}
