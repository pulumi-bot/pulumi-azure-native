// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20191201.Inputs
{

    /// <summary>
    /// This is the OS disk image.
    /// </summary>
    public sealed class GalleryOSDiskImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The host caching of the disk. Valid values are 'None', 'ReadOnly', and 'ReadWrite'
        /// </summary>
        [Input("hostCaching")]
        public Input<Pulumi.AzureNextGen.Compute.V20191201.HostCaching>? HostCaching { get; set; }

        /// <summary>
        /// The gallery artifact version source.
        /// </summary>
        [Input("source")]
        public Input<Inputs.GalleryArtifactVersionSourceArgs>? Source { get; set; }

        public GalleryOSDiskImageArgs()
        {
        }
    }
}
