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
    public sealed class ArmRollingUpgradeMonitoringPolicyResponseResult
    {
        /// <summary>
        /// The activation Mode of the service package
        /// </summary>
        public readonly string? FailureAction;
        /// <summary>
        /// The amount of time to retry health evaluation when the application or cluster is unhealthy before FailureAction is executed. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        public readonly string? HealthCheckRetryTimeout;
        /// <summary>
        /// The amount of time that the application or cluster must remain healthy before the upgrade proceeds to the next upgrade domain. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        public readonly string? HealthCheckStableDuration;
        /// <summary>
        /// The amount of time to wait after completing an upgrade domain before applying health policies. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        public readonly string? HealthCheckWaitDuration;
        /// <summary>
        /// The amount of time each upgrade domain has to complete before FailureAction is executed. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        public readonly string? UpgradeDomainTimeout;
        /// <summary>
        /// The amount of time the overall upgrade has to complete before FailureAction is executed. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        public readonly string? UpgradeTimeout;

        [OutputConstructor]
        private ArmRollingUpgradeMonitoringPolicyResponseResult(
            string? failureAction,

            string? healthCheckRetryTimeout,

            string? healthCheckStableDuration,

            string? healthCheckWaitDuration,

            string? upgradeDomainTimeout,

            string? upgradeTimeout)
        {
            FailureAction = failureAction;
            HealthCheckRetryTimeout = healthCheckRetryTimeout;
            HealthCheckStableDuration = healthCheckStableDuration;
            HealthCheckWaitDuration = healthCheckWaitDuration;
            UpgradeDomainTimeout = upgradeDomainTimeout;
            UpgradeTimeout = upgradeTimeout;
        }
    }
}
