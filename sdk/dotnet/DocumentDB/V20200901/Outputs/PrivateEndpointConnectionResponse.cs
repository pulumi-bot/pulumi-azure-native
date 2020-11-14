// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200901.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointConnectionResponse
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private endpoint which the connection belongs to.
        /// </summary>
        public readonly Outputs.PrivateEndpointPropertyResponse? PrivateEndpoint;
        /// <summary>
        /// Connection State of the Private Endpoint Connection.
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStatePropertyResponse? PrivateLinkServiceConnectionState;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PrivateEndpointConnectionResponse(
            string id,

            string name,

            Outputs.PrivateEndpointPropertyResponse? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStatePropertyResponse? privateLinkServiceConnectionState,

            string type)
        {
            Id = id;
            Name = name;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            Type = type;
        }
    }
}
