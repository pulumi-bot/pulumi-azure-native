// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20150801Preview.Outputs
{

    [OutputType]
    public sealed class AS2AgreementContentResponseResult
    {
        /// <summary>
        /// The AS2 one-way receive agreement.
        /// </summary>
        public readonly Outputs.AS2OneWayAgreementResponseResult? ReceiveAgreement;
        /// <summary>
        /// The AS2 one-way send agreement.
        /// </summary>
        public readonly Outputs.AS2OneWayAgreementResponseResult? SendAgreement;

        [OutputConstructor]
        private AS2AgreementContentResponseResult(
            Outputs.AS2OneWayAgreementResponseResult? receiveAgreement,

            Outputs.AS2OneWayAgreementResponseResult? sendAgreement)
        {
            ReceiveAgreement = receiveAgreement;
            SendAgreement = sendAgreement;
        }
    }
}
