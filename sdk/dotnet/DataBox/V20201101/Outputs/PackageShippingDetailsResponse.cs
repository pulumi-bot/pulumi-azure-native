// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBox.V20201101.Outputs
{

    [OutputType]
    public sealed class PackageShippingDetailsResponse
    {
        /// <summary>
        /// Name of the carrier.
        /// </summary>
        public readonly string CarrierName;
        /// <summary>
        /// Tracking Id of shipment.
        /// </summary>
        public readonly string TrackingId;
        /// <summary>
        /// Url where shipment can be tracked.
        /// </summary>
        public readonly string TrackingUrl;

        [OutputConstructor]
        private PackageShippingDetailsResponse(
            string carrierName,

            string trackingId,

            string trackingUrl)
        {
            CarrierName = carrierName;
            TrackingId = trackingId;
            TrackingUrl = trackingUrl;
        }
    }
}
