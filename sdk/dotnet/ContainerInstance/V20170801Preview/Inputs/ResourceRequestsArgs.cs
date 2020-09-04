// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.V20170801Preview.Inputs
{

    /// <summary>
    /// The resource requests.
    /// </summary>
    public sealed class ResourceRequestsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU request of this container instance.
        /// </summary>
        [Input("cpu", required: true)]
        public Input<double> Cpu { get; set; } = null!;

        /// <summary>
        /// The memory request in GB of this container instance.
        /// </summary>
        [Input("memoryInGB", required: true)]
        public Input<double> MemoryInGB { get; set; } = null!;

        public ResourceRequestsArgs()
        {
        }
    }
}
