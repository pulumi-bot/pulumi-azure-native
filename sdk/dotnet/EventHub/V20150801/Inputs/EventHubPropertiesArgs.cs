// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventHub.V20150801.Inputs
{

    /// <summary>
    /// Properties supplied to the Create Or Update Event Hub operation.
    /// </summary>
    public sealed class EventHubPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of days to retain the events for this Event Hub.
        /// </summary>
        [Input("messageRetentionInDays")]
        public Input<int>? MessageRetentionInDays { get; set; }

        /// <summary>
        /// Number of partitions created for the Event Hub.
        /// </summary>
        [Input("partitionCount")]
        public Input<int>? PartitionCount { get; set; }

        /// <summary>
        /// Enumerates the possible values for the status of the Event Hub.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public EventHubPropertiesArgs()
        {
        }
    }
}