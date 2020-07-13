// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Outputs
{

    [OutputType]
    public sealed class VirtualHubRouteTableV2PropertiesResponseResult
    {
        /// <summary>
        /// List of all connections attached to this route table v2.
        /// </summary>
        public readonly ImmutableArray<string> AttachedConnections;
        /// <summary>
        /// The provisioning state of the virtual hub route table v2 resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// List of all routes.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualHubRouteV2ResponseResult> Routes;

        [OutputConstructor]
        private VirtualHubRouteTableV2PropertiesResponseResult(
            ImmutableArray<string> attachedConnections,

            string provisioningState,

            ImmutableArray<Outputs.VirtualHubRouteV2ResponseResult> routes)
        {
            AttachedConnections = attachedConnections;
            ProvisioningState = provisioningState;
            Routes = routes;
        }
    }
}