// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceBus.V20150801
{
    /// <summary>
    /// Description of queue Resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:servicebus/v20150801:Queue")]
    public partial class Queue : Pulumi.CustomResource
    {
        /// <summary>
        /// Last time a message was sent, or the last time there was a receive request to this queue.
        /// </summary>
        [Output("accessedAt")]
        public Output<string> AccessedAt { get; private set; } = null!;

        /// <summary>
        /// the TimeSpan idle interval after which the queue is automatically deleted. The minimum duration is 5 minutes.
        /// </summary>
        [Output("autoDeleteOnIdle")]
        public Output<string?> AutoDeleteOnIdle { get; private set; } = null!;

        /// <summary>
        /// Message Count Details.
        /// </summary>
        [Output("countDetails")]
        public Output<Outputs.MessageCountDetailsResponse> CountDetails { get; private set; } = null!;

        /// <summary>
        /// The exact time the message was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether this queue has dead letter support when a message expires.
        /// </summary>
        [Output("deadLetteringOnMessageExpiration")]
        public Output<bool?> DeadLetteringOnMessageExpiration { get; private set; } = null!;

        /// <summary>
        /// The default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Output("defaultMessageTimeToLive")]
        public Output<string?> DefaultMessageTimeToLive { get; private set; } = null!;

        /// <summary>
        /// TimeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
        /// </summary>
        [Output("duplicateDetectionHistoryTimeWindow")]
        public Output<string?> DuplicateDetectionHistoryTimeWindow { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether server-side batched operations are enabled.
        /// </summary>
        [Output("enableBatchedOperations")]
        public Output<bool?> EnableBatchedOperations { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage.
        /// </summary>
        [Output("enableExpress")]
        public Output<bool?> EnableExpress { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether the queue is to be partitioned across multiple message brokers.
        /// </summary>
        [Output("enablePartitioning")]
        public Output<bool?> EnablePartitioning { get; private set; } = null!;

        /// <summary>
        /// Entity availability status for the queue.
        /// </summary>
        [Output("entityAvailabilityStatus")]
        public Output<string?> EntityAvailabilityStatus { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether the message is accessible anonymously.
        /// </summary>
        [Output("isAnonymousAccessible")]
        public Output<bool?> IsAnonymousAccessible { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.
        /// </summary>
        [Output("lockDuration")]
        public Output<string?> LockDuration { get; private set; } = null!;

        /// <summary>
        /// The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
        /// </summary>
        [Output("maxDeliveryCount")]
        public Output<int?> MaxDeliveryCount { get; private set; } = null!;

        /// <summary>
        /// The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
        /// </summary>
        [Output("maxSizeInMegabytes")]
        public Output<double?> MaxSizeInMegabytes { get; private set; } = null!;

        /// <summary>
        /// The number of messages in the queue.
        /// </summary>
        [Output("messageCount")]
        public Output<double> MessageCount { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A value indicating if this queue requires duplicate detection.
        /// </summary>
        [Output("requiresDuplicateDetection")]
        public Output<bool?> RequiresDuplicateDetection { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether the queue supports the concept of sessions.
        /// </summary>
        [Output("requiresSession")]
        public Output<bool?> RequiresSession { get; private set; } = null!;

        /// <summary>
        /// The size of the queue, in bytes.
        /// </summary>
        [Output("sizeInBytes")]
        public Output<double> SizeInBytes { get; private set; } = null!;

        /// <summary>
        /// Enumerates the possible values for the status of a messaging entity.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether the queue supports ordering.
        /// </summary>
        [Output("supportOrdering")]
        public Output<bool?> SupportOrdering { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The exact time the message was updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Queue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Queue(string name, QueueArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/v20150801:Queue", name, args ?? new QueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Queue(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/v20150801:Queue", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/latest:Queue"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20140901:Queue"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20170401:Queue"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20180101preview:Queue"},
                },
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
        public Input<Pulumi.AzureNextGen.ServiceBus.V20150801.EntityAvailabilityStatus>? EntityAvailabilityStatus { get; set; }

        /// <summary>
        /// A value that indicates whether the message is accessible anonymously.
        /// </summary>
        [Input("isAnonymousAccessible")]
        public Input<bool>? IsAnonymousAccessible { get; set; }

        /// <summary>
        /// location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

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
        public Input<double>? MaxSizeInMegabytes { get; set; }

        /// <summary>
        /// Queue name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The queue name.
        /// </summary>
        [Input("queueName", required: true)]
        public Input<string> QueueName { get; set; } = null!;

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
        public Input<Pulumi.AzureNextGen.ServiceBus.V20150801.EntityStatus>? Status { get; set; }

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
