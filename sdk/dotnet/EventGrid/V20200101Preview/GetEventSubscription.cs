// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200101Preview
{
    public static class GetEventSubscription
    {
        public static Task<GetEventSubscriptionResult> InvokeAsync(GetEventSubscriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventSubscriptionResult>("azurerm:eventgrid/v20200101preview:getEventSubscription", args ?? new GetEventSubscriptionArgs(), options.WithVersion());
    }


    public sealed class GetEventSubscriptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the event subscription
        /// </summary>
        [Input("eventSubscriptionName", required: true)]
        public string EventSubscriptionName { get; set; } = null!;

        /// <summary>
        /// The scope of the event subscription. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
        /// </summary>
        [Input("scope", required: true)]
        public string Scope { get; set; } = null!;

        public GetEventSubscriptionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEventSubscriptionResult
    {
        /// <summary>
        /// The DeadLetter destination of the event subscription.
        /// </summary>
        public readonly Outputs.DeadLetterDestinationResponseResult? DeadLetterDestination;
        /// <summary>
        /// Information about the destination where events have to be delivered for the event subscription.
        /// </summary>
        public readonly Outputs.EventSubscriptionDestinationResponseResult? Destination;
        /// <summary>
        /// The event delivery schema for the event subscription.
        /// </summary>
        public readonly string? EventDeliverySchema;
        /// <summary>
        /// Expiration time of the event subscription.
        /// </summary>
        public readonly string? ExpirationTimeUtc;
        /// <summary>
        /// Information about the filter for the event subscription.
        /// </summary>
        public readonly Outputs.EventSubscriptionFilterResponseResult? Filter;
        /// <summary>
        /// List of user defined labels.
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// Name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the event subscription.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
        /// </summary>
        public readonly Outputs.RetryPolicyResponseResult? RetryPolicy;
        /// <summary>
        /// Name of the topic of the event subscription.
        /// </summary>
        public readonly string Topic;
        /// <summary>
        /// Type of the resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetEventSubscriptionResult(
            Outputs.DeadLetterDestinationResponseResult? deadLetterDestination,

            Outputs.EventSubscriptionDestinationResponseResult? destination,

            string? eventDeliverySchema,

            string? expirationTimeUtc,

            Outputs.EventSubscriptionFilterResponseResult? filter,

            ImmutableArray<string> labels,

            string name,

            string provisioningState,

            Outputs.RetryPolicyResponseResult? retryPolicy,

            string topic,

            string type)
        {
            DeadLetterDestination = deadLetterDestination;
            Destination = destination;
            EventDeliverySchema = eventDeliverySchema;
            ExpirationTimeUtc = expirationTimeUtc;
            Filter = filter;
            Labels = labels;
            Name = name;
            ProvisioningState = provisioningState;
            RetryPolicy = retryPolicy;
            Topic = topic;
            Type = type;
        }
    }
}
