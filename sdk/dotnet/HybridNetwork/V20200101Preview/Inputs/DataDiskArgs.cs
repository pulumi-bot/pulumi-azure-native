// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HybridNetwork.V20200101Preview.Inputs
{

    /// <summary>
    /// Specifies information about the operating system disk used by the virtual machine. &lt;br&gt;&lt;br&gt; For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
    /// </summary>
    public sealed class DataDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how the virtual machine should be created.
        /// </summary>
        [Input("createOption")]
        public InputUnion<string, Pulumi.AzureNextGen.HybridNetwork.V20200101Preview.DiskCreateOptionTypes>? CreateOption { get; set; }

        /// <summary>
        /// Specifies the size of an empty disk in gigabytes. This element can be used to overwrite the size of the disk in a virtual machine image.
        /// </summary>
        [Input("diskSizeGB")]
        public Input<int>? DiskSizeGB { get; set; }

        /// <summary>
        /// The name of data disk.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DataDiskArgs()
        {
        }
    }
}
