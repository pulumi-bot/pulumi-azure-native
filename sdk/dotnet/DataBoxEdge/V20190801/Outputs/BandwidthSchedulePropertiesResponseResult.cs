// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge.V20190801.Outputs
{

    [OutputType]
    public sealed class BandwidthSchedulePropertiesResponseResult
    {
        /// <summary>
        /// The days of the week when this schedule is applicable.
        /// </summary>
        public readonly ImmutableArray<string> Days;
        /// <summary>
        /// The bandwidth rate in Mbps.
        /// </summary>
        public readonly int RateInMbps;
        /// <summary>
        /// The start time of the schedule in UTC.
        /// </summary>
        public readonly string Start;
        /// <summary>
        /// The stop time of the schedule in UTC.
        /// </summary>
        public readonly string Stop;

        [OutputConstructor]
        private BandwidthSchedulePropertiesResponseResult(
            ImmutableArray<string> days,

            int rateInMbps,

            string start,

            string stop)
        {
            Days = days;
            RateInMbps = rateInMbps;
            Start = start;
            Stop = stop;
        }
    }
}