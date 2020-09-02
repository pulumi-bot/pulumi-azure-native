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
    public sealed class ServiceBusQueueMessageResponseResult
    {
        /// <summary>
        /// Gets or sets the authentication.
        /// </summary>
        public readonly Outputs.ServiceBusAuthenticationResponseResult? Authentication;
        /// <summary>
        /// Gets or sets the brokered message properties.
        /// </summary>
        public readonly Outputs.ServiceBusBrokeredMessagePropertiesResponseResult? BrokeredMessageProperties;
        /// <summary>
        /// Gets or sets the custom message properties.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomMessageProperties;
        /// <summary>
        /// Gets or sets the message.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Gets or sets the namespace.
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// Gets or sets the queue name.
        /// </summary>
        public readonly string? QueueName;
        /// <summary>
        /// Gets or sets the transport type.
        /// </summary>
        public readonly string? TransportType;

        [OutputConstructor]
        private ServiceBusQueueMessageResponseResult(
            Outputs.ServiceBusAuthenticationResponseResult? authentication,

            Outputs.ServiceBusBrokeredMessagePropertiesResponseResult? brokeredMessageProperties,

            ImmutableDictionary<string, string>? customMessageProperties,

            string? message,

            string? @namespace,

            string? queueName,

            string? transportType)
        {
            Authentication = authentication;
            BrokeredMessageProperties = brokeredMessageProperties;
            CustomMessageProperties = customMessageProperties;
            Message = message;
            Namespace = @namespace;
            QueueName = queueName;
            TransportType = transportType;
        }
    }
}
