// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190601Preview.Inputs
{

    /// <summary>
    /// Describes the policy for a monitored application upgrade.
    /// </summary>
    public sealed class ApplicationUpgradePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines a health policy used to evaluate the health of an application or one of its children entities.
        /// </summary>
        [Input("applicationHealthPolicy")]
        public Input<Inputs.ArmApplicationHealthPolicyArgs>? ApplicationHealthPolicy { get; set; }

        /// <summary>
        /// If true, then processes are forcefully restarted during upgrade even when the code version has not changed (the upgrade only changes configuration or data).
        /// </summary>
        [Input("forceRestart")]
        public Input<bool>? ForceRestart { get; set; }

        /// <summary>
        /// The policy used for monitoring the application upgrade
        /// </summary>
        [Input("rollingUpgradeMonitoringPolicy")]
        public Input<Inputs.ArmRollingUpgradeMonitoringPolicyArgs>? RollingUpgradeMonitoringPolicy { get; set; }

        /// <summary>
        /// The maximum amount of time to block processing of an upgrade domain and prevent loss of availability when there are unexpected issues. When this timeout expires, processing of the upgrade domain will proceed regardless of availability loss issues. The timeout is reset at the start of each upgrade domain. Valid values are between 0 and 42949672925 inclusive. (unsigned 32-bit integer).
        /// </summary>
        [Input("upgradeReplicaSetCheckTimeout")]
        public Input<string>? UpgradeReplicaSetCheckTimeout { get; set; }

        public ApplicationUpgradePolicyArgs()
        {
        }
    }
}
