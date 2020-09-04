// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AzureStackHCI.V20200301Preview.Outputs
{

    [OutputType]
    public sealed class ClusterNodeResponseResult
    {
        /// <summary>
        /// Number of physical cores on the cluster node.
        /// </summary>
        public readonly double CoreCount;
        /// <summary>
        /// Id of the node in the cluster.
        /// </summary>
        public readonly double Id;
        /// <summary>
        /// Manufacturer of the cluster node hardware.
        /// </summary>
        public readonly string Manufacturer;
        /// <summary>
        /// Total available memory on the cluster node (in GiB).
        /// </summary>
        public readonly double MemoryInGiB;
        /// <summary>
        /// Model name of the cluster node hardware.
        /// </summary>
        public readonly string Model;
        /// <summary>
        /// Name of the cluster node.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Operating system running on the cluster node.
        /// </summary>
        public readonly string OsName;
        /// <summary>
        /// Version of the operating system running on the cluster node.
        /// </summary>
        public readonly string OsVersion;
        /// <summary>
        /// Immutable id of the cluster node.
        /// </summary>
        public readonly string SerialNumber;

        [OutputConstructor]
        private ClusterNodeResponseResult(
            double coreCount,

            double id,

            string manufacturer,

            double memoryInGiB,

            string model,

            string name,

            string osName,

            string osVersion,

            string serialNumber)
        {
            CoreCount = coreCount;
            Id = id;
            Manufacturer = manufacturer;
            MemoryInGiB = memoryInGiB;
            Model = model;
            Name = name;
            OsName = osName;
            OsVersion = osVersion;
            SerialNumber = serialNumber;
        }
    }
}
