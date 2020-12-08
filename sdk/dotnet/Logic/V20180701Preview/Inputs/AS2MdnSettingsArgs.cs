// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.V20180701Preview.Inputs
{

    /// <summary>
    /// The AS2 agreement mdn settings.
    /// </summary>
    public sealed class AS2MdnSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The disposition notification to header value.
        /// </summary>
        [Input("dispositionNotificationTo")]
        public Input<string>? DispositionNotificationTo { get; set; }

        /// <summary>
        /// The MDN text.
        /// </summary>
        [Input("mdnText")]
        public Input<string>? MdnText { get; set; }

        /// <summary>
        /// The signing or hashing algorithm.
        /// </summary>
        [Input("micHashingAlgorithm", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Logic.V20180701Preview.HashingAlgorithm> MicHashingAlgorithm { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to send or request a MDN.
        /// </summary>
        [Input("needMDN", required: true)]
        public Input<bool> NeedMDN { get; set; } = null!;

        /// <summary>
        /// The receipt delivery URL.
        /// </summary>
        [Input("receiptDeliveryUrl")]
        public Input<string>? ReceiptDeliveryUrl { get; set; }

        /// <summary>
        /// The value indicating whether to send inbound MDN to message box.
        /// </summary>
        [Input("sendInboundMDNToMessageBox", required: true)]
        public Input<bool> SendInboundMDNToMessageBox { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to send the asynchronous MDN.
        /// </summary>
        [Input("sendMDNAsynchronously", required: true)]
        public Input<bool> SendMDNAsynchronously { get; set; } = null!;

        /// <summary>
        /// The value indicating whether the MDN needs to be signed or not.
        /// </summary>
        [Input("signMDN", required: true)]
        public Input<bool> SignMDN { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to sign the outbound MDN if optional.
        /// </summary>
        [Input("signOutboundMDNIfOptional", required: true)]
        public Input<bool> SignOutboundMDNIfOptional { get; set; } = null!;

        public AS2MdnSettingsArgs()
        {
        }
    }
}
