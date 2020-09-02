// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.Latest.Outputs
{

    [OutputType]
    public sealed class ResourceRequestsResponseResult
    {
        /// <summary>
        /// The CPU request of this container instance.
        /// </summary>
        public readonly double Cpu;
        /// <summary>
        /// The GPU request of this container instance.
        /// </summary>
        public readonly Outputs.GpuResourceResponseResult? Gpu;
        /// <summary>
        /// The memory request in GB of this container instance.
        /// </summary>
        public readonly double MemoryInGB;

        [OutputConstructor]
        private ResourceRequestsResponseResult(
            double cpu,

            Outputs.GpuResourceResponseResult? gpu,

            double memoryInGB)
        {
            Cpu = cpu;
            Gpu = gpu;
            MemoryInGB = memoryInGB;
        }
    }
}