// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Outputs
{

    [OutputType]
    public sealed class DiskAccessPropertiesResponse
    {
        /// <summary>
        /// A readonly collection of private endpoint connections created on the disk. Currently only one endpoint connection is supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// The disk access resource provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The time when the disk access was created.
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private DiskAccessPropertiesResponse(
            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            string timeCreated)
        {
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            TimeCreated = timeCreated;
        }
    }
}