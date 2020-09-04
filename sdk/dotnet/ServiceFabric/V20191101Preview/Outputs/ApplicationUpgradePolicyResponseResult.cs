// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class ApplicationUpgradePolicyResponseResult
    {
        /// <summary>
        /// Defines a health policy used to evaluate the health of an application or one of its children entities.
        /// </summary>
        public readonly Outputs.ArmApplicationHealthPolicyResponseResult? ApplicationHealthPolicy;
        /// <summary>
        /// If true, then processes are forcefully restarted during upgrade even when the code version has not changed (the upgrade only changes configuration or data).
        /// </summary>
        public readonly bool? ForceRestart;
        /// <summary>
        /// The policy used for monitoring the application upgrade
        /// </summary>
        public readonly Outputs.ArmRollingUpgradeMonitoringPolicyResponseResult? RollingUpgradeMonitoringPolicy;
        /// <summary>
        /// The mode used to monitor health during a rolling upgrade. The values are UnmonitoredAuto, UnmonitoredManual, and Monitored.
        /// </summary>
        public readonly string? UpgradeMode;
        /// <summary>
        /// The maximum amount of time to block processing of an upgrade domain and prevent loss of availability when there are unexpected issues. When this timeout expires, processing of the upgrade domain will proceed regardless of availability loss issues. The timeout is reset at the start of each upgrade domain. Valid values are between 0 and 42949672925 inclusive. (unsigned 32-bit integer).
        /// </summary>
        public readonly string? UpgradeReplicaSetCheckTimeout;

        [OutputConstructor]
        private ApplicationUpgradePolicyResponseResult(
            Outputs.ArmApplicationHealthPolicyResponseResult? applicationHealthPolicy,

            bool? forceRestart,

            Outputs.ArmRollingUpgradeMonitoringPolicyResponseResult? rollingUpgradeMonitoringPolicy,

            string? upgradeMode,

            string? upgradeReplicaSetCheckTimeout)
        {
            ApplicationHealthPolicy = applicationHealthPolicy;
            ForceRestart = forceRestart;
            RollingUpgradeMonitoringPolicy = rollingUpgradeMonitoringPolicy;
            UpgradeMode = upgradeMode;
            UpgradeReplicaSetCheckTimeout = upgradeReplicaSetCheckTimeout;
        }
    }
}
