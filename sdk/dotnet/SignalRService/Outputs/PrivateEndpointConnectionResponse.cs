// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SignalRService.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointConnectionResponse
    {
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the private endpoint connection
        /// </summary>
        public readonly Outputs.PrivateEndpointConnectionPropertiesResponse? Properties;
        /// <summary>
        /// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PrivateEndpointConnectionResponse(
            string id,

            string name,

            Outputs.PrivateEndpointConnectionPropertiesResponse? properties,

            string type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}