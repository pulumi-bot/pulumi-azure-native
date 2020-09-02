// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Inputs
{

    /// <summary>
    /// The X12 one-way agreement.
    /// </summary>
    public sealed class X12OneWayAgreementArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The X12 protocol settings.
        /// </summary>
        [Input("protocolSettings", required: true)]
        public Input<Inputs.X12ProtocolSettingsArgs> ProtocolSettings { get; set; } = null!;

        /// <summary>
        /// The receiver business identity
        /// </summary>
        [Input("receiverBusinessIdentity", required: true)]
        public Input<Inputs.BusinessIdentityArgs> ReceiverBusinessIdentity { get; set; } = null!;

        /// <summary>
        /// The sender business identity
        /// </summary>
        [Input("senderBusinessIdentity", required: true)]
        public Input<Inputs.BusinessIdentityArgs> SenderBusinessIdentity { get; set; } = null!;

        public X12OneWayAgreementArgs()
        {
        }
    }
}
