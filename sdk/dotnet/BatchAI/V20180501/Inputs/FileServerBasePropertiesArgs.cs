// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20180501.Inputs
{

    /// <summary>
    /// The properties of a file server.
    /// </summary>
    public sealed class FileServerBasePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for the data disks which will be created for the File Server.
        /// </summary>
        [Input("dataDisks", required: true)]
        public Input<Inputs.DataDisksArgs> DataDisks { get; set; } = null!;

        /// <summary>
        /// SSH configuration for the File Server node.
        /// </summary>
        [Input("sshConfiguration", required: true)]
        public Input<Inputs.SshConfigurationArgs> SshConfiguration { get; set; } = null!;

        /// <summary>
        /// Identifier of an existing virtual network subnet to put the File Server in. If not provided, a new virtual network and subnet will be created.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.ResourceIdArgs>? Subnet { get; set; }

        /// <summary>
        /// The size of the virtual machine for the File Server. For information about available VM sizes from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
        /// </summary>
        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public FileServerBasePropertiesArgs()
        {
        }
    }
}