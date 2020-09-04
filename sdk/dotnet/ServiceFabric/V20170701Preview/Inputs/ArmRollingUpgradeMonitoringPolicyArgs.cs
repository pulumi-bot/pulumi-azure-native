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
    /// The policy used for monitoring the application upgrade
    /// </summary>
    public sealed class ArmRollingUpgradeMonitoringPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The activation Mode of the service package
        /// </summary>
        [Input("failureAction")]
        public Input<string>? FailureAction { get; set; }

        /// <summary>
        /// The amount of time to retry health evaluation when the application or cluster is unhealthy before FailureAction is executed. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        [Input("healthCheckRetryTimeout")]
        public Input<string>? HealthCheckRetryTimeout { get; set; }

        /// <summary>
        /// The amount of time that the application or cluster must remain healthy before the upgrade proceeds to the next upgrade domain. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        [Input("healthCheckStableDuration")]
        public Input<string>? HealthCheckStableDuration { get; set; }

        /// <summary>
        /// The amount of time to wait after completing an upgrade domain before applying health policies. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        [Input("healthCheckWaitDuration")]
        public Input<string>? HealthCheckWaitDuration { get; set; }

        /// <summary>
        /// The amount of time each upgrade domain has to complete before FailureAction is executed. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        [Input("upgradeDomainTimeout")]
        public Input<string>? UpgradeDomainTimeout { get; set; }

        /// <summary>
        /// The amount of time the overall upgrade has to complete before FailureAction is executed. It is first interpreted as a string representing an ISO 8601 duration. If that fails, then it is interpreted as a number representing the total number of milliseconds.
        /// </summary>
        [Input("upgradeTimeout")]
        public Input<string>? UpgradeTimeout { get; set; }

        public ArmRollingUpgradeMonitoringPolicyArgs()
        {
        }
    }
}
