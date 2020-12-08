// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Media.V20200501.Inputs
{

    /// <summary>
    /// The live event input.
    /// </summary>
    public sealed class LiveEventInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access control for live event input.
        /// </summary>
        [Input("accessControl")]
        public Input<Inputs.LiveEventInputAccessControlArgs>? AccessControl { get; set; }

        /// <summary>
        /// A UUID in string form to uniquely identify the stream. This can be specified at creation time but cannot be updated. If omitted, the service will generate a unique value.
        /// </summary>
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.LiveEventEndpointArgs>? _endpoints;

        /// <summary>
        /// The input endpoints for the live event.
        /// </summary>
        public InputList<Inputs.LiveEventEndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.LiveEventEndpointArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// ISO 8601 time duration of the key frame interval duration of the input. This value sets the EXT-X-TARGETDURATION property in the HLS output. For example, use PT2S to indicate 2 seconds. Leave the value empty for encoding live events.
        /// </summary>
        [Input("keyFrameIntervalDuration")]
        public Input<string>? KeyFrameIntervalDuration { get; set; }

        /// <summary>
        /// The input protocol for the live event. This is specified at creation time and cannot be updated.
        /// </summary>
        [Input("streamingProtocol", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Media.V20200501.LiveEventInputProtocol> StreamingProtocol { get; set; } = null!;

        public LiveEventInputArgs()
        {
        }
    }
}
