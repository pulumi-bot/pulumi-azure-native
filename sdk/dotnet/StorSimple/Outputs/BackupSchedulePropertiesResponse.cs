// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorSimple.Outputs
{

    [OutputType]
    public sealed class BackupSchedulePropertiesResponse
    {
        /// <summary>
        /// The type of backup which needs to be taken.
        /// </summary>
        public readonly string BackupType;
        /// <summary>
        /// The last successful backup run which was triggered for the schedule.
        /// </summary>
        public readonly string LastSuccessfulRun;
        /// <summary>
        /// The number of backups to be retained.
        /// </summary>
        public readonly int RetentionCount;
        /// <summary>
        /// The schedule recurrence.
        /// </summary>
        public readonly Outputs.ScheduleRecurrenceResponse ScheduleRecurrence;
        /// <summary>
        /// The schedule status.
        /// </summary>
        public readonly string ScheduleStatus;
        /// <summary>
        /// The start time of the schedule.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private BackupSchedulePropertiesResponse(
            string backupType,

            string lastSuccessfulRun,

            int retentionCount,

            Outputs.ScheduleRecurrenceResponse scheduleRecurrence,

            string scheduleStatus,

            string startTime)
        {
            BackupType = backupType;
            LastSuccessfulRun = lastSuccessfulRun;
            RetentionCount = retentionCount;
            ScheduleRecurrence = scheduleRecurrence;
            ScheduleStatus = scheduleStatus;
            StartTime = startTime;
        }
    }
}