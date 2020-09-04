// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20170907privatePreview
{
    /// <summary>
    /// Class representing an event hub connection.
    /// </summary>
    public partial class EventHubConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The event hub consumer group.
        /// </summary>
        [Output("consumerGroup")]
        public Output<string> ConsumerGroup { get; private set; } = null!;

        /// <summary>
        /// The data format of the message. Optionally the data format can be added to each message.
        /// </summary>
        [Output("dataFormat")]
        public Output<string?> DataFormat { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the event hub to be used to create a data connection.
        /// </summary>
        [Output("eventHubResourceId")]
        public Output<string> EventHubResourceId { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
        /// </summary>
        [Output("mappingRuleName")]
        public Output<string?> MappingRuleName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The table where the data should be ingested. Optionally the table information can be added to each message.
        /// </summary>
        [Output("tableName")]
        public Output<string?> TableName { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EventHubConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventHubConnection(string name, EventHubConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:kusto/v20170907privatepreview:EventHubConnection", name, args ?? new EventHubConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventHubConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:kusto/v20170907privatepreview:EventHubConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:kusto/v20180907preview:EventHubConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventHubConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventHubConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventHubConnection(name, id, options);
        }
    }

    public sealed class EventHubConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The event hub consumer group.
        /// </summary>
        [Input("consumerGroup", required: true)]
        public Input<string> ConsumerGroup { get; set; } = null!;

        /// <summary>
        /// The data format of the message. Optionally the data format can be added to each message.
        /// </summary>
        [Input("dataFormat")]
        public Input<string>? DataFormat { get; set; }

        /// <summary>
        /// The name of the database in the Kusto cluster.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the event hub connection.
        /// </summary>
        [Input("eventHubConnectionName", required: true)]
        public Input<string> EventHubConnectionName { get; set; } = null!;

        /// <summary>
        /// The resource ID of the event hub to be used to create a data connection.
        /// </summary>
        [Input("eventHubResourceId", required: true)]
        public Input<string> EventHubResourceId { get; set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
        /// </summary>
        [Input("mappingRuleName")]
        public Input<string>? MappingRuleName { get; set; }

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The table where the data should be ingested. Optionally the table information can be added to each message.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public EventHubConnectionArgs()
        {
        }
    }
}
