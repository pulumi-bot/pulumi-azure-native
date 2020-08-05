// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceBus.V20150801
{
    /// <summary>
    /// Description of queue Resource.
    /// </summary>
    public partial class Queue : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Queue Properties definition.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.QueuePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Queue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Queue(string name, QueueArgs args, CustomResourceOptions? options = null)
            : base("azurerm:servicebus/v20150801:Queue", name, args ?? new QueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Queue(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:servicebus/v20150801:Queue", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Queue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Queue Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Queue(name, id, options);
        }
    }

    public sealed class QueueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the TimeSpan idle interval after which the queue is automatically deleted. The minimum duration is 5 minutes.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// A value that indicates whether this queue has dead letter support when a message expires.
        /// </summary>
        [Input("deadLetteringOnMessageExpiration")]
        public Input<bool>? DeadLetteringOnMessageExpiration { get; set; }

        /// <summary>
        /// The default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Input("defaultMessageTimeToLive")]
        public Input<string>? DefaultMessageTimeToLive { get; set; }

        /// <summary>
        /// TimeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
        /// </summary>
        [Input("duplicateDetectionHistoryTimeWindow")]
        public Input<string>? DuplicateDetectionHistoryTimeWindow { get; set; }

        /// <summary>
        /// A value that indicates whether server-side batched operations are enabled.
        /// </summary>
        [Input("enableBatchedOperations")]
        public Input<bool>? EnableBatchedOperations { get; set; }

        /// <summary>
        /// A value that indicates whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage.
        /// </summary>
        [Input("enableExpress")]
        public Input<bool>? EnableExpress { get; set; }

        /// <summary>
        /// A value that indicates whether the queue is to be partitioned across multiple message brokers.
        /// </summary>
        [Input("enablePartitioning")]
        public Input<bool>? EnablePartitioning { get; set; }

        /// <summary>
        /// Entity availability status for the queue.
        /// </summary>
        [Input("entityAvailabilityStatus")]
        public Input<string>? EntityAvailabilityStatus { get; set; }

        /// <summary>
        /// A value that indicates whether the message is accessible anonymously.
        /// </summary>
        [Input("isAnonymousAccessible")]
        public Input<bool>? IsAnonymousAccessible { get; set; }

        /// <summary>
        /// location of the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.
        /// </summary>
        [Input("lockDuration")]
        public Input<string>? LockDuration { get; set; }

        /// <summary>
        /// The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
        /// </summary>
        [Input("maxDeliveryCount")]
        public Input<int>? MaxDeliveryCount { get; set; }

        /// <summary>
        /// The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
        /// </summary>
        [Input("maxSizeInMegabytes")]
        public Input<int>? MaxSizeInMegabytes { get; set; }

        /// <summary>
        /// The queue name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// A value indicating if this queue requires duplicate detection.
        /// </summary>
        [Input("requiresDuplicateDetection")]
        public Input<bool>? RequiresDuplicateDetection { get; set; }

        /// <summary>
        /// A value that indicates whether the queue supports the concept of sessions.
        /// </summary>
        [Input("requiresSession")]
        public Input<bool>? RequiresSession { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Enumerates the possible values for the status of a messaging entity.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// A value that indicates whether the queue supports ordering.
        /// </summary>
        [Input("supportOrdering")]
        public Input<bool>? SupportOrdering { get; set; }

        public QueueArgs()
        {
        }
    }
}
