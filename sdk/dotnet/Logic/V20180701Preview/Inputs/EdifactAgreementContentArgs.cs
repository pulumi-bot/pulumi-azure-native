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
    /// The Edifact agreement content.
    /// </summary>
    public sealed class EdifactAgreementContentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The EDIFACT one-way receive agreement.
        /// </summary>
        [Input("receiveAgreement", required: true)]
        public Input<Inputs.EdifactOneWayAgreementArgs> ReceiveAgreement { get; set; } = null!;

        /// <summary>
        /// The EDIFACT one-way send agreement.
        /// </summary>
        [Input("sendAgreement", required: true)]
        public Input<Inputs.EdifactOneWayAgreementArgs> SendAgreement { get; set; } = null!;

        public EdifactAgreementContentArgs()
        {
        }
    }
}
