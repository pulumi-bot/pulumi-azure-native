// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge.V20190301.Outputs
{

    [OutputType]
    public sealed class OrderPropertiesResponseResult
    {
        /// <summary>
        /// The contact details.
        /// </summary>
        public readonly Outputs.ContactDetailsResponseResult ContactInformation;
        /// <summary>
        /// Current status of the order.
        /// </summary>
        public readonly Outputs.OrderStatusResponseResult? CurrentStatus;
        /// <summary>
        /// Tracking information for the package delivered to the customer whether it has an original or a replacement device.
        /// </summary>
        public readonly ImmutableArray<Outputs.TrackingInfoResponseResult> DeliveryTrackingInfo;
        /// <summary>
        /// List of status changes in the order.
        /// </summary>
        public readonly ImmutableArray<Outputs.OrderStatusResponseResult> OrderHistory;
        /// <summary>
        /// Tracking information for the package returned from the customer whether it has an original or a replacement device.
        /// </summary>
        public readonly ImmutableArray<Outputs.TrackingInfoResponseResult> ReturnTrackingInfo;
        /// <summary>
        /// Serial number of the device.
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// The shipping address.
        /// </summary>
        public readonly Outputs.AddressResponseResult ShippingAddress;

        [OutputConstructor]
        private OrderPropertiesResponseResult(
            Outputs.ContactDetailsResponseResult contactInformation,

            Outputs.OrderStatusResponseResult? currentStatus,

            ImmutableArray<Outputs.TrackingInfoResponseResult> deliveryTrackingInfo,

            ImmutableArray<Outputs.OrderStatusResponseResult> orderHistory,

            ImmutableArray<Outputs.TrackingInfoResponseResult> returnTrackingInfo,

            string serialNumber,

            Outputs.AddressResponseResult shippingAddress)
        {
            ContactInformation = contactInformation;
            CurrentStatus = currentStatus;
            DeliveryTrackingInfo = deliveryTrackingInfo;
            OrderHistory = orderHistory;
            ReturnTrackingInfo = returnTrackingInfo;
            SerialNumber = serialNumber;
            ShippingAddress = shippingAddress;
        }
    }
}