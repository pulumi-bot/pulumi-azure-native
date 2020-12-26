// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageCache.V20201001.Outputs
{

    [OutputType]
    public sealed class CacheUpgradeStatusResponse
    {
        /// <summary>
        /// Version string of the firmware currently installed on this Cache.
        /// </summary>
        public readonly string CurrentFirmwareVersion;
        /// <summary>
        /// Time at which the pending firmware update will automatically be installed on the Cache.
        /// </summary>
        public readonly string FirmwareUpdateDeadline;
        /// <summary>
        /// True if there is a firmware update ready to install on this Cache. The firmware will automatically be installed after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
        /// </summary>
        public readonly string FirmwareUpdateStatus;
        /// <summary>
        /// Time of the last successful firmware update.
        /// </summary>
        public readonly string LastFirmwareUpdate;
        /// <summary>
        /// When firmwareUpdateAvailable is true, this field holds the version string for the update.
        /// </summary>
        public readonly string PendingFirmwareVersion;

        [OutputConstructor]
        private CacheUpgradeStatusResponse(
            string currentFirmwareVersion,

            string firmwareUpdateDeadline,

            string firmwareUpdateStatus,

            string lastFirmwareUpdate,

            string pendingFirmwareVersion)
        {
            CurrentFirmwareVersion = currentFirmwareVersion;
            FirmwareUpdateDeadline = firmwareUpdateDeadline;
            FirmwareUpdateStatus = firmwareUpdateStatus;
            LastFirmwareUpdate = lastFirmwareUpdate;
            PendingFirmwareVersion = pendingFirmwareVersion;
        }
    }
}
