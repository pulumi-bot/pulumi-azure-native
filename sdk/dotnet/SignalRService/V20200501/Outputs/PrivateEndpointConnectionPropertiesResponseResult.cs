// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SignalRService.V20200501.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointConnectionPropertiesResponseResult
    {
        /// <summary>
        /// Private endpoint associated with the private endpoint connection
        /// </summary>
        public readonly Outputs.PrivateEndpointResponseResult? PrivateEndpoint;
        /// <summary>
        /// Connection state
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStateResponseResult? PrivateLinkServiceConnectionState;
        /// <summary>
        /// Provisioning state of the private endpoint connection
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private PrivateEndpointConnectionPropertiesResponseResult(
            Outputs.PrivateEndpointResponseResult? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStateResponseResult? privateLinkServiceConnectionState,

            string provisioningState)
        {
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
        }
    }
}