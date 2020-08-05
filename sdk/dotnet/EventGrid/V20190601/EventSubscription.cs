// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20190601
{
    /// <summary>
    /// Event Subscription
    /// </summary>
    public partial class EventSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the event subscription
        /// </summary>
        [Output("properties")]
        public Output<Outputs.EventSubscriptionPropertiesResponseResult> Properties { get; private set; } = null!;

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
            : base("azurerm:eventgrid/v20190601:EventSubscription", name, args ?? new EventSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:eventgrid/v20190601:EventSubscription", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
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
        public Input<Inputs.DeadLetterDestinationArgs>? DeadLetterDestination { get; set; }

        /// <summary>
        /// Information about the destination where events have to be delivered for the event subscription.
        /// </summary>
        [Input("destination")]
        public Input<Inputs.EventSubscriptionDestinationArgs>? Destination { get; set; }

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
        /// Name of the event subscription. Event subscription names must be between 3 and 64 characters in length and should use alphanumeric letters only.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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
