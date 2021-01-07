// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801.Outputs
{

    [OutputType]
    public sealed class ExpressRouteCircuitPeeringIdResponse
    {
        /// <summary>
        /// The ID of the ExpressRoute circuit peering.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ExpressRouteCircuitPeeringIdResponse(string? id)
        {
            Id = id;
        }
    }
}
