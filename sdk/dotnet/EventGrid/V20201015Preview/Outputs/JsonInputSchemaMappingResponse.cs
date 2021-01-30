// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EventGrid.V20201015Preview.Outputs
{

    [OutputType]
    public sealed class JsonInputSchemaMappingResponse
    {
        /// <summary>
        /// The mapping information for the DataVersion property of the Event Grid Event.
        /// </summary>
        public readonly Outputs.JsonFieldWithDefaultResponse? DataVersion;
        /// <summary>
        /// The mapping information for the EventTime property of the Event Grid Event.
        /// </summary>
        public readonly Outputs.JsonFieldResponse? EventTime;
        /// <summary>
        /// The mapping information for the EventType property of the Event Grid Event.
        /// </summary>
        public readonly Outputs.JsonFieldWithDefaultResponse? EventType;
        /// <summary>
        /// The mapping information for the Id property of the Event Grid Event.
        /// </summary>
        public readonly Outputs.JsonFieldResponse? Id;
        /// <summary>
        /// Type of the custom mapping
        /// Expected value is 'Json'.
        /// </summary>
        public readonly string InputSchemaMappingType;
        /// <summary>
        /// The mapping information for the Subject property of the Event Grid Event.
        /// </summary>
        public readonly Outputs.JsonFieldWithDefaultResponse? Subject;
        /// <summary>
        /// The mapping information for the Topic property of the Event Grid Event.
        /// </summary>
        public readonly Outputs.JsonFieldResponse? Topic;

        [OutputConstructor]
        private JsonInputSchemaMappingResponse(
            Outputs.JsonFieldWithDefaultResponse? dataVersion,

            Outputs.JsonFieldResponse? eventTime,

            Outputs.JsonFieldWithDefaultResponse? eventType,

            Outputs.JsonFieldResponse? id,

            string inputSchemaMappingType,

            Outputs.JsonFieldWithDefaultResponse? subject,

            Outputs.JsonFieldResponse? topic)
        {
            DataVersion = dataVersion;
            EventTime = eventTime;
            EventType = eventType;
            Id = id;
            InputSchemaMappingType = inputSchemaMappingType;
            Subject = subject;
            Topic = topic;
        }
    }
}
