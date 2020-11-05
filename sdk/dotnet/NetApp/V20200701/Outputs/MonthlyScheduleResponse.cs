// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.NetApp.V20200701.Outputs
{

    [OutputType]
    public sealed class MonthlyScheduleResponse
    {
        /// <summary>
        /// Indicates which days of the month snapshot should be taken. A comma delimited string.
        /// </summary>
        public readonly string? DaysOfMonth;
        /// <summary>
        /// Indicates which hour in UTC timezone a snapshot should be taken
        /// </summary>
        public readonly int? Hour;
        /// <summary>
        /// Indicates which minute snapshot should be taken
        /// </summary>
        public readonly int? Minute;
        /// <summary>
        /// Monthly snapshot count to keep
        /// </summary>
        public readonly int? SnapshotsToKeep;
        /// <summary>
        /// Resource size in bytes, current storage usage for the volume in bytes
        /// </summary>
        public readonly int? UsedBytes;

        [OutputConstructor]
        private MonthlyScheduleResponse(
            string? daysOfMonth,

            int? hour,

            int? minute,

            int? snapshotsToKeep,

            int? usedBytes)
        {
            DaysOfMonth = daysOfMonth;
            Hour = hour;
            Minute = minute;
            SnapshotsToKeep = snapshotsToKeep;
            UsedBytes = usedBytes;
        }
    }
}
