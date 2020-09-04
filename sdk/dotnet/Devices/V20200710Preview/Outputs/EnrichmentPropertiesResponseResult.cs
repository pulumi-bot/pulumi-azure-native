// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200710Preview.Outputs
{

    [OutputType]
    public sealed class EnrichmentPropertiesResponseResult
    {
        /// <summary>
        /// The list of endpoints for which the enrichment is applied to the message.
        /// </summary>
        public readonly ImmutableArray<string> EndpointNames;
        /// <summary>
        /// The key or name for the enrichment property.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value for the enrichment property.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private EnrichmentPropertiesResponseResult(
            ImmutableArray<string> endpointNames,

            string key,

            string value)
        {
            EndpointNames = endpointNames;
            Key = key;
            Value = value;
        }
    }
}
