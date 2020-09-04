// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301Preview.Outputs
{

    [OutputType]
    public sealed class ArmServiceTypeHealthPolicyResponseResult
    {
        /// <summary>
        /// The maximum percentage of partitions per service allowed to be unhealthy before your application is considered in error.
        /// </summary>
        public readonly int? MaxPercentUnhealthyPartitionsPerService;
        /// <summary>
        /// The maximum percentage of replicas per partition allowed to be unhealthy before your application is considered in error.
        /// </summary>
        public readonly int? MaxPercentUnhealthyReplicasPerPartition;
        /// <summary>
        /// The maximum percentage of services allowed to be unhealthy before your application is considered in error.
        /// </summary>
        public readonly int? MaxPercentUnhealthyServices;

        [OutputConstructor]
        private ArmServiceTypeHealthPolicyResponseResult(
            int? maxPercentUnhealthyPartitionsPerService,

            int? maxPercentUnhealthyReplicasPerPartition,

            int? maxPercentUnhealthyServices)
        {
            MaxPercentUnhealthyPartitionsPerService = maxPercentUnhealthyPartitionsPerService;
            MaxPercentUnhealthyReplicasPerPartition = maxPercentUnhealthyReplicasPerPartition;
            MaxPercentUnhealthyServices = maxPercentUnhealthyServices;
        }
    }
}
