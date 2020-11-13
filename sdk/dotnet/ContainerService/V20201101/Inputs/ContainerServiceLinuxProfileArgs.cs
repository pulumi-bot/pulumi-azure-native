// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20201101.Inputs
{

    /// <summary>
    /// Profile for Linux VMs in the container service cluster.
    /// </summary>
    public sealed class ContainerServiceLinuxProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrator username to use for Linux VMs.
        /// </summary>
        [Input("adminUsername", required: true)]
        public Input<string> AdminUsername { get; set; } = null!;

        /// <summary>
        /// SSH configuration for Linux-based VMs running on Azure.
        /// </summary>
        [Input("ssh", required: true)]
        public Input<Inputs.ContainerServiceSshConfigurationArgs> Ssh { get; set; } = null!;

        public ContainerServiceLinuxProfileArgs()
        {
        }
    }
}
