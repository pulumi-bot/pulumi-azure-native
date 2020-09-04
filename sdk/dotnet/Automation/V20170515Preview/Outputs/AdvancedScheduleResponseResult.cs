// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.V20170515Preview.Outputs
{

    [OutputType]
    public sealed class AdvancedScheduleResponseResult
    {
        /// <summary>
        /// Days of the month that the job should execute on. Must be between 1 and 31.
        /// </summary>
        public readonly ImmutableArray<int> MonthDays;
        /// <summary>
        /// Occurrences of days within a month.
        /// </summary>
        public readonly ImmutableArray<Outputs.AdvancedScheduleMonthlyOccurrenceResponseResult> MonthlyOccurrences;
        /// <summary>
        /// Days of the week that the job should execute on.
        /// </summary>
        public readonly ImmutableArray<string> WeekDays;

        [OutputConstructor]
        private AdvancedScheduleResponseResult(
            ImmutableArray<int> monthDays,

            ImmutableArray<Outputs.AdvancedScheduleMonthlyOccurrenceResponseResult> monthlyOccurrences,

            ImmutableArray<string> weekDays)
        {
            MonthDays = monthDays;
            MonthlyOccurrences = monthlyOccurrences;
            WeekDays = weekDays;
        }
    }
}
