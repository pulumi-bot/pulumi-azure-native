// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20181101.Outputs
{

    [OutputType]
    public sealed class RequestsBasedTriggerResponseResult
    {
        /// <summary>
        /// Request Count.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// Time interval.
        /// </summary>
        public readonly string? TimeInterval;

        [OutputConstructor]
        private RequestsBasedTriggerResponseResult(
            int? count,

            string? timeInterval)
        {
            Count = count;
            TimeInterval = timeInterval;
        }
    }
}