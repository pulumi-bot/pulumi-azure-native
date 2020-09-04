// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200515Preview.Inputs
{

    /// <summary>
    /// The resource requirements for the container (cpu and memory).
    /// </summary>
    public sealed class ContainerResourceRequirementsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of CPU cores on the container.
        /// </summary>
        [Input("cpu")]
        public Input<double>? Cpu { get; set; }

        /// <summary>
        /// The number of FPGA PCIE devices exposed to the container. Must be multiple of 2.
        /// </summary>
        [Input("fpga")]
        public Input<int>? Fpga { get; set; }

        /// <summary>
        /// The number of GPU cores in the container.
        /// </summary>
        [Input("gpu")]
        public Input<int>? Gpu { get; set; }

        /// <summary>
        /// The amount of memory on the container in GB.
        /// </summary>
        [Input("memoryInGB")]
        public Input<double>? MemoryInGB { get; set; }

        public ContainerResourceRequirementsArgs()
        {
        }
    }
}
