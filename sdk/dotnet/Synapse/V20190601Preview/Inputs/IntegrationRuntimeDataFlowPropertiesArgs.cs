// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20190601Preview.Inputs
{

    /// <summary>
    /// Data flow properties for managed integration runtime.
    /// </summary>
    public sealed class IntegrationRuntimeDataFlowPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Compute type of the cluster which will execute data flow job.
        /// </summary>
        [Input("computeType")]
        public InputUnion<string, Pulumi.AzureNextGen.Synapse.V20190601Preview.DataFlowComputeType>? ComputeType { get; set; }

        /// <summary>
        /// Core count of the cluster which will execute data flow job. Supported values are: 8, 16, 32, 48, 80, 144 and 272.
        /// </summary>
        [Input("coreCount")]
        public Input<int>? CoreCount { get; set; }

        /// <summary>
        /// Time to live (in minutes) setting of the cluster which will execute data flow job.
        /// </summary>
        [Input("timeToLive")]
        public Input<int>? TimeToLive { get; set; }

        public IntegrationRuntimeDataFlowPropertiesArgs()
        {
        }
    }
}
