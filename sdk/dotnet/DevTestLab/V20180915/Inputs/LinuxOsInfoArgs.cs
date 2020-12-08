// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DevTestLab.V20180915.Inputs
{

    /// <summary>
    /// Information about a Linux OS.
    /// </summary>
    public sealed class LinuxOsInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The state of the Linux OS (i.e. NonDeprovisioned, DeprovisionRequested, DeprovisionApplied).
        /// </summary>
        [Input("linuxOsState")]
        public InputUnion<string, Pulumi.AzureNextGen.DevTestLab.V20180915.LinuxOsState>? LinuxOsState { get; set; }

        public LinuxOsInfoArgs()
        {
        }
    }
}
