// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20170701Preview.Inputs
{

    /// <summary>
    /// Represents the health policy used to evaluate the health of services belonging to a service type.
    /// </summary>
    public sealed class ArmServiceTypeHealthPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum percentage of partitions per service allowed to be unhealthy before your application is considered in error.
        /// </summary>
        [Input("maxPercentUnhealthyPartitionsPerService")]
        public Input<int>? MaxPercentUnhealthyPartitionsPerService { get; set; }

        /// <summary>
        /// The maximum percentage of replicas per partition allowed to be unhealthy before your application is considered in error.
        /// </summary>
        [Input("maxPercentUnhealthyReplicasPerPartition")]
        public Input<int>? MaxPercentUnhealthyReplicasPerPartition { get; set; }

        /// <summary>
        /// The maximum percentage of services allowed to be unhealthy before your application is considered in error.
        /// </summary>
        [Input("maxPercentUnhealthyServices")]
        public Input<int>? MaxPercentUnhealthyServices { get; set; }

        public ArmServiceTypeHealthPolicyArgs()
        {
        }
    }
}
