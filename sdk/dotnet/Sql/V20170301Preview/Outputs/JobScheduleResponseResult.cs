// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20170301Preview.Outputs
{

    [OutputType]
    public sealed class JobScheduleResponseResult
    {
        /// <summary>
        /// Whether or not the schedule is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Schedule end time.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// Value of the schedule's recurring interval, if the schedule type is recurring. ISO8601 duration format.
        /// </summary>
        public readonly string? Interval;
        /// <summary>
        /// Schedule start time.
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// Schedule interval type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private JobScheduleResponseResult(
            bool? enabled,

            string? endTime,

            string? interval,

            string? startTime,

            string? type)
        {
            Enabled = enabled;
            EndTime = endTime;
            Interval = interval;
            StartTime = startTime;
            Type = type;
        }
    }
}
