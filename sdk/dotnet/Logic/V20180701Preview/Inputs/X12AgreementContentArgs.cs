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
    /// The X12 agreement content.
    /// </summary>
    public sealed class X12AgreementContentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The X12 one-way receive agreement.
        /// </summary>
        [Input("receiveAgreement", required: true)]
        public Input<Inputs.X12OneWayAgreementArgs> ReceiveAgreement { get; set; } = null!;

        /// <summary>
        /// The X12 one-way send agreement.
        /// </summary>
        [Input("sendAgreement", required: true)]
        public Input<Inputs.X12OneWayAgreementArgs> SendAgreement { get; set; } = null!;

        public X12AgreementContentArgs()
        {
        }
    }
}
