// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.V20170801Preview.Outputs
{

    [OutputType]
    public sealed class ContainerEventResponseResult
    {
        /// <summary>
        /// The count of the event.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// The date-time of the earliest logged event.
        /// </summary>
        public readonly string? FirstTimestamp;
        /// <summary>
        /// The date-time of the latest logged event.
        /// </summary>
        public readonly string? LastTimestamp;
        /// <summary>
        /// The event message.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// The event type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ContainerEventResponseResult(
            int? count,

            string? firstTimestamp,

            string? lastTimestamp,

            string? message,

            string? type)
        {
            Count = count;
            FirstTimestamp = firstTimestamp;
            LastTimestamp = lastTimestamp;
            Message = message;
            Type = type;
        }
    }
}
