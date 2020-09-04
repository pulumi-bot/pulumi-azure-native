// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.V20190801Preview.Inputs
{

    /// <summary>
    /// The properties that define a direct connection.
    /// </summary>
    public sealed class DirectConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bandwidth of the connection.
        /// </summary>
        [Input("bandwidthInMbps")]
        public Input<int>? BandwidthInMbps { get; set; }

        /// <summary>
        /// The BGP session associated with the connection.
        /// </summary>
        [Input("bgpSession")]
        public Input<Inputs.BgpSessionArgs>? BgpSession { get; set; }

        /// <summary>
        /// The unique identifier (GUID) for the connection.
        /// </summary>
        [Input("connectionIdentifier")]
        public Input<string>? ConnectionIdentifier { get; set; }

        /// <summary>
        /// The PeeringDB.com ID of the facility at which the connection has to be set up.
        /// </summary>
        [Input("peeringDBFacilityId")]
        public Input<int>? PeeringDBFacilityId { get; set; }

        /// <summary>
        /// The bandwidth that is actually provisioned.
        /// </summary>
        [Input("provisionedBandwidthInMbps")]
        public Input<int>? ProvisionedBandwidthInMbps { get; set; }

        /// <summary>
        /// The field indicating if Microsoft provides session ip addresses.
        /// </summary>
        [Input("sessionAddressProvider")]
        public Input<string>? SessionAddressProvider { get; set; }

        /// <summary>
        /// The flag that indicates whether or not the connection is used for peering service.
        /// </summary>
        [Input("useForPeeringService")]
        public Input<bool>? UseForPeeringService { get; set; }

        public DirectConnectionArgs()
        {
        }
    }
}
