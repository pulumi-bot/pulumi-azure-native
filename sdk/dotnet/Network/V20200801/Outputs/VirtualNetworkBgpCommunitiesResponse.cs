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
    public sealed class VirtualNetworkBgpCommunitiesResponse
    {
        /// <summary>
        /// The BGP community associated with the region of the virtual network.
        /// </summary>
        public readonly string RegionalCommunity;
        /// <summary>
        /// The BGP community associated with the virtual network.
        /// </summary>
        public readonly string VirtualNetworkCommunity;

        [OutputConstructor]
        private VirtualNetworkBgpCommunitiesResponse(
            string regionalCommunity,

            string virtualNetworkCommunity)
        {
            RegionalCommunity = regionalCommunity;
            VirtualNetworkCommunity = virtualNetworkCommunity;
        }
    }
}
