// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20181201Preview.Outputs
{

    [OutputType]
    public sealed class FeedbackPropertiesResponseResult
    {
        /// <summary>
        /// The lock duration for the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        public readonly string? LockDurationAsIso8601;
        /// <summary>
        /// The number of times the IoT hub attempts to deliver a message on the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        public readonly int? MaxDeliveryCount;
        /// <summary>
        /// The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        public readonly string? TtlAsIso8601;

        [OutputConstructor]
        private FeedbackPropertiesResponseResult(
            string? lockDurationAsIso8601,

            int? maxDeliveryCount,

            string? ttlAsIso8601)
        {
            LockDurationAsIso8601 = lockDurationAsIso8601;
            MaxDeliveryCount = maxDeliveryCount;
            TtlAsIso8601 = ttlAsIso8601;
        }
    }
}
