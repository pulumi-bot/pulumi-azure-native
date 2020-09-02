// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20181201Preview.Inputs
{

    /// <summary>
    /// The IoT hub cloud-to-device messaging properties.
    /// </summary>
    public sealed class CloudToDevicePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default time to live for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        [Input("defaultTtlAsIso8601")]
        public Input<string>? DefaultTtlAsIso8601 { get; set; }

        /// <summary>
        /// The properties of the feedback queue for cloud-to-device messages.
        /// </summary>
        [Input("feedback")]
        public Input<Inputs.FeedbackPropertiesArgs>? Feedback { get; set; }

        /// <summary>
        /// The max delivery count for cloud-to-device messages in the device queue. See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging#cloud-to-device-messages.
        /// </summary>
        [Input("maxDeliveryCount")]
        public Input<int>? MaxDeliveryCount { get; set; }

        public CloudToDevicePropertiesArgs()
        {
        }
    }
}
