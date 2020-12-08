// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EventHub.V20150801
{
    /// <summary>
    /// Single item in List or Get Event Hub operation
    /// </summary>
    public partial class EventHub : Pulumi.CustomResource
    {
        /// <summary>
        /// Exact time the Event Hub was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Number of days to retain the events for this Event Hub.
        /// </summary>
        [Output("messageRetentionInDays")]
        public Output<int?> MessageRetentionInDays { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of partitions created for the Event Hub.
        /// </summary>
        [Output("partitionCount")]
        public Output<int?> PartitionCount { get; private set; } = null!;

        /// <summary>
        /// Current number of shards on the Event Hub.
        /// </summary>
        [Output("partitionIds")]
        public Output<ImmutableArray<string>> PartitionIds { get; private set; } = null!;

        /// <summary>
        /// Enumerates the possible values for the status of the Event Hub.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

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
        /// Create a EventHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventHub(string name, EventHubArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:eventhub/v20150801:EventHub", name, args ?? new EventHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventHub(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:eventhub/v20150801:EventHub", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:eventhub/latest:EventHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventhub/v20140901:EventHub"},
                    new Pulumi.Alias { Type = "azure-nextgen:eventhub/v20170401:EventHub"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventHub Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventHub(name, id, options);
        }
    }

    public sealed class EventHubArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Event Hub name
        /// </summary>
        [Input("eventHubName", required: true)]
        public Input<string> EventHubName { get; set; } = null!;

        /// <summary>
        /// Location of the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Number of days to retain the events for this Event Hub.
        /// </summary>
        [Input("messageRetentionInDays")]
        public Input<int>? MessageRetentionInDays { get; set; }

        /// <summary>
        /// Name of the Event Hub.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Number of partitions created for the Event Hub.
        /// </summary>
        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        /// <summary>
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Enumerates the possible values for the status of the Event Hub.
        /// </summary>
        [Input("status")]
        public Input<Pulumi.AzureNextGen.EventHub.V20150801.EntityStatus>? Status { get; set; }

        /// <summary>
        /// ARM type of the Namespace.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EventHubArgs()
        {
        }
    }
}
