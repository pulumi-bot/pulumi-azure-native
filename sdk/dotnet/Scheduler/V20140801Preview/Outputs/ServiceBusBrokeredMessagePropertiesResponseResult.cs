// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Scheduler.V20140801Preview.Outputs
{

    [OutputType]
    public sealed class ServiceBusBrokeredMessagePropertiesResponseResult
    {
        /// <summary>
        /// Gets or sets the content type.
        /// </summary>
        public readonly string? ContentType;
        /// <summary>
        /// Gets or sets the correlation id.
        /// </summary>
        public readonly string? CorrelationId;
        /// <summary>
        /// Gets or sets the force persistence.
        /// </summary>
        public readonly bool? ForcePersistence;
        /// <summary>
        /// Gets or sets the label.
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// Gets or sets the message id.
        /// </summary>
        public readonly string? MessageId;
        /// <summary>
        /// Gets or sets the partition key.
        /// </summary>
        public readonly string? PartitionKey;
        /// <summary>
        /// Gets or sets the reply to.
        /// </summary>
        public readonly string? ReplyTo;
        /// <summary>
        /// Gets or sets the reply to session id.
        /// </summary>
        public readonly string? ReplyToSessionId;
        /// <summary>
        /// Gets or sets the scheduled enqueue time UTC.
        /// </summary>
        public readonly string? ScheduledEnqueueTimeUtc;
        /// <summary>
        /// Gets or sets the session id.
        /// </summary>
        public readonly string? SessionId;
        /// <summary>
        /// Gets or sets the time to live.
        /// </summary>
        public readonly string? TimeToLive;
        /// <summary>
        /// Gets or sets the to.
        /// </summary>
        public readonly string? To;
        /// <summary>
        /// Gets or sets the via partition key.
        /// </summary>
        public readonly string? ViaPartitionKey;

        [OutputConstructor]
        private ServiceBusBrokeredMessagePropertiesResponseResult(
            string? contentType,

            string? correlationId,

            bool? forcePersistence,

            string? label,

            string? messageId,

            string? partitionKey,

            string? replyTo,

            string? replyToSessionId,

            string? scheduledEnqueueTimeUtc,

            string? sessionId,

            string? timeToLive,

            string? to,

            string? viaPartitionKey)
        {
            ContentType = contentType;
            CorrelationId = correlationId;
            ForcePersistence = forcePersistence;
            Label = label;
            MessageId = messageId;
            PartitionKey = partitionKey;
            ReplyTo = replyTo;
            ReplyToSessionId = replyToSessionId;
            ScheduledEnqueueTimeUtc = scheduledEnqueueTimeUtc;
            SessionId = sessionId;
            TimeToLive = timeToLive;
            To = to;
            ViaPartitionKey = viaPartitionKey;
        }
    }
}
