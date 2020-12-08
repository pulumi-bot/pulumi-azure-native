// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EventGrid.Latest
{
    /// <summary>
    /// Event Subscription
    /// </summary>
    public partial class EventSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// The DeadLetter destination of the event subscription.
        /// </summary>
        [Output("deadLetterDestination")]
        public Output<Outputs.StorageBlobDeadLetterDestinationResponse?> DeadLetterDestination { get; private set; } = null!;

        /// <summary>
        /// Information about the destination where events have to be delivered for the event subscription.
        /// </summary>
        [Output("destination")]
        public Output<Union<Outputs.AzureFunctionEventSubscriptionDestinationResponse, Union<Outputs.EventHubEventSubscriptionDestinationResponse, Union<Outputs.HybridConnectionEventSubscriptionDestinationResponse, Union<Outputs.ServiceBusQueueEventSubscriptionDestinationResponse, Union<Outputs.ServiceBusTopicEventSubscriptionDestinationResponse, Union<Outputs.StorageQueueEventSubscriptionDestinationResponse, Outputs.WebHookEventSubscriptionDestinationResponse>>>>>>?> Destination { get; private set; } = null!;

        /// <summary>
        /// The event delivery schema for the event subscription.
        /// </summary>
        [Output("eventDeliverySchema")]
        public Output<string?> EventDeliverySchema { get; private set; } = null!;

        /// <summary>
        /// Expiration time of the event subscription.
        /// </summary>
        [Output("expirationTimeUtc")]
        public Output<string?> ExpirationTimeUtc { get; private set; } = null!;

        /// <summary>
        /// Information about the filter for the event subscription.
        /// </summary>
        [Output("filter")]
        public Output<Outputs.EventSubscriptionFilterResponse?> Filter { get; private set; } = null!;

        /// <summary>
        /// List of user defined labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the event subscription.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
        /// </summary>
        [Output("retryPolicy")]
        public Output<Outputs.RetryPolicyResponse?> RetryPolicy { get; private set; } = null!;

        /// <summary>
        /// Name of the topic of the event subscription.
        /// </summary>
        [Output("topic")]
        public Output<string> Topic { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EventSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventSubscription(string name, EventSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:eventgrid/latest:EventSubscription", name, args ?? new EventSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:eventgrid/latest:EventSubscription", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20170615preview:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20170915preview:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20180101:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20180501preview:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20180915preview:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20190101:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20190201preview:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20190601:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20200101preview:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20200401preview:EventSubscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20200601:EventSubscription"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventSubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventSubscription(name, id, options);
        }
    }

    public sealed class EventSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DeadLetter destination of the event subscription.
        /// </summary>
        [Input("deadLetterDestination")]
        public Input<Inputs.StorageBlobDeadLetterDestinationArgs>? DeadLetterDestination { get; set; }

        /// <summary>
        /// Information about the destination where events have to be delivered for the event subscription.
        /// </summary>
        [Input("destination")]
        public InputUnion<Inputs.AzureFunctionEventSubscriptionDestinationArgs, InputUnion<Inputs.EventHubEventSubscriptionDestinationArgs, InputUnion<Inputs.HybridConnectionEventSubscriptionDestinationArgs, InputUnion<Inputs.ServiceBusQueueEventSubscriptionDestinationArgs, InputUnion<Inputs.ServiceBusTopicEventSubscriptionDestinationArgs, InputUnion<Inputs.StorageQueueEventSubscriptionDestinationArgs, Inputs.WebHookEventSubscriptionDestinationArgs>>>>>>? Destination { get; set; }

        /// <summary>
        /// The event delivery schema for the event subscription.
        /// </summary>
        [Input("eventDeliverySchema")]
        public InputUnion<string, Pulumi.AzureNextGen.EventGrid.Latest.EventDeliverySchema>? EventDeliverySchema { get; set; }

        /// <summary>
        /// Name of the event subscription. Event subscription names must be between 3 and 64 characters in length and should use alphanumeric letters only.
        /// </summary>
        [Input("eventSubscriptionName", required: true)]
        public Input<string> EventSubscriptionName { get; set; } = null!;

        /// <summary>
        /// Expiration time of the event subscription.
        /// </summary>
        [Input("expirationTimeUtc")]
        public Input<string>? ExpirationTimeUtc { get; set; }

        /// <summary>
        /// Information about the filter for the event subscription.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.EventSubscriptionFilterArgs>? Filter { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// List of user defined labels.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.RetryPolicyArgs>? RetryPolicy { get; set; }

        /// <summary>
        /// The identifier of the resource to which the event subscription needs to be created or updated. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public EventSubscriptionArgs()
        {
        }
    }
}
