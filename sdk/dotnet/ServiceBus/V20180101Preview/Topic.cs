// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceBus.V20180101Preview
{
    /// <summary>
    /// Description of topic resource.
    /// </summary>
    public partial class Topic : Pulumi.CustomResource
    {
        /// <summary>
        /// Last time the message was sent, or a request was received, for this topic.
        /// </summary>
        [Output("accessedAt")]
        public Output<string> AccessedAt { get; private set; } = null!;

        /// <summary>
        /// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
        /// </summary>
        [Output("autoDeleteOnIdle")]
        public Output<string?> AutoDeleteOnIdle { get; private set; } = null!;

        /// <summary>
        /// Message count details
        /// </summary>
        [Output("countDetails")]
        public Output<Outputs.MessageCountDetailsResponse> CountDetails { get; private set; } = null!;

        /// <summary>
        /// Exact time the message was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Output("defaultMessageTimeToLive")]
        public Output<string?> DefaultMessageTimeToLive { get; private set; } = null!;

        /// <summary>
        /// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
        /// </summary>
        [Output("duplicateDetectionHistoryTimeWindow")]
        public Output<string?> DuplicateDetectionHistoryTimeWindow { get; private set; } = null!;

        /// <summary>
        /// Value that indicates whether server-side batched operations are enabled.
        /// </summary>
        [Output("enableBatchedOperations")]
        public Output<bool?> EnableBatchedOperations { get; private set; } = null!;

        /// <summary>
        /// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
        /// </summary>
        [Output("enableExpress")]
        public Output<bool?> EnableExpress { get; private set; } = null!;

        /// <summary>
        /// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
        /// </summary>
        [Output("enablePartitioning")]
        public Output<bool?> EnablePartitioning { get; private set; } = null!;

        /// <summary>
        /// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
        /// </summary>
        [Output("maxSizeInMegabytes")]
        public Output<int?> MaxSizeInMegabytes { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Value indicating if this topic requires duplicate detection.
        /// </summary>
        [Output("requiresDuplicateDetection")]
        public Output<bool?> RequiresDuplicateDetection { get; private set; } = null!;

        /// <summary>
        /// Size of the topic, in bytes.
        /// </summary>
        [Output("sizeInBytes")]
        public Output<int> SizeInBytes { get; private set; } = null!;

        /// <summary>
        /// Enumerates the possible values for the status of a messaging entity.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Number of subscriptions.
        /// </summary>
        [Output("subscriptionCount")]
        public Output<int> SubscriptionCount { get; private set; } = null!;

        /// <summary>
        /// Value that indicates whether the topic supports ordering.
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
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/v20180101preview:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/v20180101preview:Topic", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/latest:Topic"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20140901:Topic"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20150801:Topic"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20170401:Topic"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, options);
        }
    }

    public sealed class TopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Input("defaultMessageTimeToLive")]
        public Input<string>? DefaultMessageTimeToLive { get; set; }

        /// <summary>
        /// ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
        /// </summary>
        [Input("duplicateDetectionHistoryTimeWindow")]
        public Input<string>? DuplicateDetectionHistoryTimeWindow { get; set; }

        /// <summary>
        /// Value that indicates whether server-side batched operations are enabled.
        /// </summary>
        [Input("enableBatchedOperations")]
        public Input<bool>? EnableBatchedOperations { get; set; }

        /// <summary>
        /// Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
        /// </summary>
        [Input("enableExpress")]
        public Input<bool>? EnableExpress { get; set; }

        /// <summary>
        /// Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
        /// </summary>
        [Input("enablePartitioning")]
        public Input<bool>? EnablePartitioning { get; set; }

        /// <summary>
        /// Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
        /// </summary>
        [Input("maxSizeInMegabytes")]
        public Input<int>? MaxSizeInMegabytes { get; set; }

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Value indicating if this topic requires duplicate detection.
        /// </summary>
        [Input("requiresDuplicateDetection")]
        public Input<bool>? RequiresDuplicateDetection { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Enumerates the possible values for the status of a messaging entity.
        /// </summary>
        [Input("status")]
        public Input<Pulumi.AzureNextGen.ServiceBus.V20180101Preview.EntityStatus>? Status { get; set; }

        /// <summary>
        /// Value that indicates whether the topic supports ordering.
        /// </summary>
        [Input("supportOrdering")]
        public Input<bool>? SupportOrdering { get; set; }

        /// <summary>
        /// The topic name.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public TopicArgs()
        {
        }
    }
}
