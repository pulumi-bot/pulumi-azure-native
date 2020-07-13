// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions.Outputs
{

    [OutputType]
    public sealed class JitSchedulingPolicyResponse
    {
        public readonly string Duration;
        /// <summary>
        /// The start time of the request.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The type of JIT schedule.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private JitSchedulingPolicyResponse(
            string duration,

            string startTime,

            string type)
        {
            Duration = duration;
            StartTime = startTime;
            Type = type;
        }
    }
}