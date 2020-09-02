// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class EdifactOneWayAgreementResponseResult
    {
        /// <summary>
        /// The EDIFACT protocol settings.
        /// </summary>
        public readonly Outputs.EdifactProtocolSettingsResponseResult ProtocolSettings;
        /// <summary>
        /// The receiver business identity
        /// </summary>
        public readonly Outputs.BusinessIdentityResponseResult ReceiverBusinessIdentity;
        /// <summary>
        /// The sender business identity
        /// </summary>
        public readonly Outputs.BusinessIdentityResponseResult SenderBusinessIdentity;

        [OutputConstructor]
        private EdifactOneWayAgreementResponseResult(
            Outputs.EdifactProtocolSettingsResponseResult protocolSettings,

            Outputs.BusinessIdentityResponseResult receiverBusinessIdentity,

            Outputs.BusinessIdentityResponseResult senderBusinessIdentity)
        {
            ProtocolSettings = protocolSettings;
            ReceiverBusinessIdentity = receiverBusinessIdentity;
            SenderBusinessIdentity = senderBusinessIdentity;
        }
    }
}
