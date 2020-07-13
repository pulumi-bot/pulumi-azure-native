// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventHub.Inputs
{

    /// <summary>
    /// Namespace properties supplied for create namespace operation.
    /// </summary>
    public sealed class EHNamespacePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value that indicates whether AutoInflate is enabled for eventhub namespace.
        /// </summary>
        [Input("isAutoInflateEnabled")]
        public Input<bool>? IsAutoInflateEnabled { get; set; }

        /// <summary>
        /// Value that indicates whether Kafka is enabled for eventhub namespace.
        /// </summary>
        [Input("kafkaEnabled")]
        public Input<bool>? KafkaEnabled { get; set; }

        /// <summary>
        /// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if AutoInflateEnabled = true)
        /// </summary>
        [Input("maximumThroughputUnits")]
        public Input<int>? MaximumThroughputUnits { get; set; }

        public EHNamespacePropertiesArgs()
        {
        }
    }
}