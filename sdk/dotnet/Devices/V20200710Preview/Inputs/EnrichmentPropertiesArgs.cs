// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200710Preview.Inputs
{

    /// <summary>
    /// The properties of an enrichment that your IoT hub applies to messages delivered to endpoints.
    /// </summary>
    public sealed class EnrichmentPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("endpointNames", required: true)]
        private InputList<string>? _endpointNames;

        /// <summary>
        /// The list of endpoints for which the enrichment is applied to the message.
        /// </summary>
        public InputList<string> EndpointNames
        {
            get => _endpointNames ?? (_endpointNames = new InputList<string>());
            set => _endpointNames = value;
        }

        /// <summary>
        /// The key or name for the enrichment property.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value for the enrichment property.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public EnrichmentPropertiesArgs()
        {
        }
    }
}
