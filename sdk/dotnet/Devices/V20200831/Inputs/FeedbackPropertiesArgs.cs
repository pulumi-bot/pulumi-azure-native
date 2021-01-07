// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200831.Inputs
{

    /// <summary>
    /// The properties of the feedback queue for cloud-to-device messages.
    /// </summary>
    public sealed class FeedbackPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The lock duration for the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        [Input("lockDurationAsIso8601")]
        public Input<string>? LockDurationAsIso8601 { get; set; }

        /// <summary>
        /// The number of times the IoT hub attempts to deliver a message on the feedback queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        [Input("maxDeliveryCount")]
        public Input<int>? MaxDeliveryCount { get; set; }

        /// <summary>
        /// The period of time for which a message is available to consume before it is expired by the IoT hub. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        [Input("ttlAsIso8601")]
        public Input<string>? TtlAsIso8601 { get; set; }

        public FeedbackPropertiesArgs()
        {
        }
    }
}
