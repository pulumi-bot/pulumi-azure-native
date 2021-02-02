// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20210101.Outputs
{

    [OutputType]
    public sealed class MonthlyRetentionScheduleResponse
    {
        /// <summary>
        /// Retention duration of retention Policy.
        /// </summary>
        public readonly Outputs.RetentionDurationResponse? RetentionDuration;
        /// <summary>
        /// Daily retention format for monthly retention policy.
        /// </summary>
        public readonly Outputs.DailyRetentionFormatResponse? RetentionScheduleDaily;
        /// <summary>
        /// Retention schedule format type for monthly retention policy.
        /// </summary>
        public readonly string? RetentionScheduleFormatType;
        /// <summary>
        /// Weekly retention format for monthly retention policy.
        /// </summary>
        public readonly Outputs.WeeklyRetentionFormatResponse? RetentionScheduleWeekly;
        /// <summary>
        /// Retention times of retention policy.
        /// </summary>
        public readonly ImmutableArray<string> RetentionTimes;

        [OutputConstructor]
        private MonthlyRetentionScheduleResponse(
            Outputs.RetentionDurationResponse? retentionDuration,

            Outputs.DailyRetentionFormatResponse? retentionScheduleDaily,

            string? retentionScheduleFormatType,

            Outputs.WeeklyRetentionFormatResponse? retentionScheduleWeekly,

            ImmutableArray<string> retentionTimes)
        {
            RetentionDuration = retentionDuration;
            RetentionScheduleDaily = retentionScheduleDaily;
            RetentionScheduleFormatType = retentionScheduleFormatType;
            RetentionScheduleWeekly = retentionScheduleWeekly;
            RetentionTimes = retentionTimes;
        }
    }
}
