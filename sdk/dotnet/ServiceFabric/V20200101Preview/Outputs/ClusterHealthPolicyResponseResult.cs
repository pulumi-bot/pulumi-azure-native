// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class ClusterHealthPolicyResponseResult
    {
        /// <summary>
        /// Defines the application health policy map used to evaluate the health of an application or one of its children entities.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ApplicationHealthPolicyResponseResult>? ApplicationHealthPolicies;
        /// <summary>
        /// The maximum allowed percentage of unhealthy applications before reporting an error. For example, to allow 10% of applications to be unhealthy, this value would be 10.
        /// 
        /// The percentage represents the maximum tolerated percentage of applications that can be unhealthy before the cluster is considered in error.
        /// If the percentage is respected but there is at least one unhealthy application, the health is evaluated as Warning.
        /// This is calculated by dividing the number of unhealthy applications over the total number of application instances in the cluster, excluding applications of application types that are included in the ApplicationTypeHealthPolicyMap.
        /// The computation rounds up to tolerate one failure on small numbers of applications. Default percentage is zero.
        /// </summary>
        public readonly int? MaxPercentUnhealthyApplications;
        /// <summary>
        /// The maximum allowed percentage of unhealthy nodes before reporting an error. For example, to allow 10% of nodes to be unhealthy, this value would be 10.
        /// 
        /// The percentage represents the maximum tolerated percentage of nodes that can be unhealthy before the cluster is considered in error.
        /// If the percentage is respected but there is at least one unhealthy node, the health is evaluated as Warning.
        /// The percentage is calculated by dividing the number of unhealthy nodes over the total number of nodes in the cluster.
        /// The computation rounds up to tolerate one failure on small numbers of nodes. Default percentage is zero.
        /// 
        /// In large clusters, some nodes will always be down or out for repairs, so this percentage should be configured to tolerate that.
        /// </summary>
        public readonly int? MaxPercentUnhealthyNodes;

        [OutputConstructor]
        private ClusterHealthPolicyResponseResult(
            ImmutableDictionary<string, Outputs.ApplicationHealthPolicyResponseResult>? applicationHealthPolicies,

            int? maxPercentUnhealthyApplications,

            int? maxPercentUnhealthyNodes)
        {
            ApplicationHealthPolicies = applicationHealthPolicies;
            MaxPercentUnhealthyApplications = maxPercentUnhealthyApplications;
            MaxPercentUnhealthyNodes = maxPercentUnhealthyNodes;
        }
    }
}
