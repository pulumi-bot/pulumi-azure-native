// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801.Inputs
{

    /// <summary>
    /// Qos Traffic Profiler Port range properties.
    /// </summary>
    public sealed class QosPortRangeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Qos Port Range end.
        /// </summary>
        [Input("end")]
        public Input<int>? End { get; set; }

        /// <summary>
        /// Qos Port Range start.
        /// </summary>
        [Input("start")]
        public Input<int>? Start { get; set; }

        public QosPortRangeArgs()
        {
        }
    }
}
