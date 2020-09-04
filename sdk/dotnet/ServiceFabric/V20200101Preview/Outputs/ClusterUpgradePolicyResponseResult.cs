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
    public sealed class ClusterUpgradePolicyResponseResult
    {
        /// <summary>
        /// The cluster delta health policy used when upgrading the cluster.
        /// </summary>
        public readonly Outputs.ClusterUpgradeDeltaHealthPolicyResponseResult? DeltaHealthPolicy;
        /// <summary>
        /// If true, then processes are forcefully restarted during upgrade even when the code version has not changed (the upgrade only changes configuration or data).
        /// </summary>
        public readonly bool? ForceRestart;
        /// <summary>
        /// The amount of time to retry health evaluation when the application or cluster is unhealthy before the upgrade rolls back. The timeout can be in either hh:mm:ss or in d.hh:mm:ss.ms format.
        /// </summary>
        public readonly string HealthCheckRetryTimeout;
        /// <summary>
        /// The amount of time that the application or cluster must remain healthy before the upgrade proceeds to the next upgrade domain. The duration can be in either hh:mm:ss or in d.hh:mm:ss.ms format.
        /// </summary>
        public readonly string HealthCheckStableDuration;
        /// <summary>
        /// The length of time to wait after completing an upgrade domain before performing health checks. The duration can be in either hh:mm:ss or in d.hh:mm:ss.ms format.
        /// </summary>
        public readonly string HealthCheckWaitDuration;
        /// <summary>
        /// The cluster health policy used when upgrading the cluster.
        /// </summary>
        public readonly Outputs.ClusterHealthPolicyResponseResult HealthPolicy;
        /// <summary>
        /// The amount of time each upgrade domain has to complete before the upgrade rolls back. The timeout can be in either hh:mm:ss or in d.hh:mm:ss.ms format.
        /// </summary>
        public readonly string UpgradeDomainTimeout;
        /// <summary>
        /// The maximum amount of time to block processing of an upgrade domain and prevent loss of availability when there are unexpected issues. When this timeout expires, processing of the upgrade domain will proceed regardless of availability loss issues. The timeout is reset at the start of each upgrade domain. The timeout can be in either hh:mm:ss or in d.hh:mm:ss.ms format.
        /// </summary>
        public readonly string UpgradeReplicaSetCheckTimeout;
        /// <summary>
        /// The amount of time the overall upgrade has to complete before the upgrade rolls back. The timeout can be in either hh:mm:ss or in d.hh:mm:ss.ms format.
        /// </summary>
        public readonly string UpgradeTimeout;

        [OutputConstructor]
        private ClusterUpgradePolicyResponseResult(
            Outputs.ClusterUpgradeDeltaHealthPolicyResponseResult? deltaHealthPolicy,

            bool? forceRestart,

            string healthCheckRetryTimeout,

            string healthCheckStableDuration,

            string healthCheckWaitDuration,

            Outputs.ClusterHealthPolicyResponseResult healthPolicy,

            string upgradeDomainTimeout,

            string upgradeReplicaSetCheckTimeout,

            string upgradeTimeout)
        {
            DeltaHealthPolicy = deltaHealthPolicy;
            ForceRestart = forceRestart;
            HealthCheckRetryTimeout = healthCheckRetryTimeout;
            HealthCheckStableDuration = healthCheckStableDuration;
            HealthCheckWaitDuration = healthCheckWaitDuration;
            HealthPolicy = healthPolicy;
            UpgradeDomainTimeout = upgradeDomainTimeout;
            UpgradeReplicaSetCheckTimeout = upgradeReplicaSetCheckTimeout;
            UpgradeTimeout = upgradeTimeout;
        }
    }
}
