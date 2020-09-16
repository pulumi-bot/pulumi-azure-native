// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20180710.Outputs
{

    [OutputType]
    public sealed class InMageRcmMobilityAgentDetailsResponse
    {
        /// <summary>
        /// The agent version expiry date.
        /// </summary>
        public readonly string AgentVersionExpiryDate;
        /// <summary>
        /// The driver version.
        /// </summary>
        public readonly string DriverVersion;
        /// <summary>
        /// The driver version expiry date.
        /// </summary>
        public readonly string DriverVersionExpiryDate;
        /// <summary>
        /// A value indicating whether agent is upgradeable or not.
        /// </summary>
        public readonly string IsUpgradeable;
        /// <summary>
        /// The time of the last heartbeat received from the agent.
        /// </summary>
        public readonly string LastHeartbeatUtc;
        /// <summary>
        /// The latest upgradeable version available without reboot.
        /// </summary>
        public readonly string LatestUpgradableVersionWithoutReboot;
        /// <summary>
        /// The latest agent version available.
        /// </summary>
        public readonly string LatestVersion;
        /// <summary>
        /// The whether update is possible or not.
        /// </summary>
        public readonly ImmutableArray<string> ReasonsBlockingUpgrade;
        /// <summary>
        /// The agent version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private InMageRcmMobilityAgentDetailsResponse(
            string agentVersionExpiryDate,

            string driverVersion,

            string driverVersionExpiryDate,

            string isUpgradeable,

            string lastHeartbeatUtc,

            string latestUpgradableVersionWithoutReboot,

            string latestVersion,

            ImmutableArray<string> reasonsBlockingUpgrade,

            string version)
        {
            AgentVersionExpiryDate = agentVersionExpiryDate;
            DriverVersion = driverVersion;
            DriverVersionExpiryDate = driverVersionExpiryDate;
            IsUpgradeable = isUpgradeable;
            LastHeartbeatUtc = lastHeartbeatUtc;
            LatestUpgradableVersionWithoutReboot = latestUpgradableVersionWithoutReboot;
            LatestVersion = latestVersion;
            ReasonsBlockingUpgrade = reasonsBlockingUpgrade;
            Version = version;
        }
    }
}
