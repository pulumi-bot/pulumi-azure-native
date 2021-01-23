// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.Latest.Inputs
{

    /// <summary>
    /// Virtual Network Profile
    /// </summary>
    public sealed class VirtualNetworkProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Subnet ID used for computes in workspace
        /// </summary>
        [Input("computeSubnetId")]
        public Input<string>? ComputeSubnetId { get; set; }

        public VirtualNetworkProfileArgs()
        {
        }
    }
}
