// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20190930Preview.Outputs
{

    [OutputType]
    public sealed class OpenShiftManagedClusterMasterPoolProfileResponseResult
    {
        /// <summary>
        /// Number of masters (VMs) to host docker containers. The default value is 3.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// Unique name of the master pool profile in the context of the subscription and resource group.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// Subnet CIDR for the peering.
        /// </summary>
        public readonly string? SubnetCidr;
        /// <summary>
        /// Size of agent VMs.
        /// </summary>
        public readonly string VmSize;

        [OutputConstructor]
        private OpenShiftManagedClusterMasterPoolProfileResponseResult(
            int count,

            string? name,

            string? osType,

            string? subnetCidr,

            string vmSize)
        {
            Count = count;
            Name = name;
            OsType = osType;
            SubnetCidr = subnetCidr;
            VmSize = vmSize;
        }
    }
}
