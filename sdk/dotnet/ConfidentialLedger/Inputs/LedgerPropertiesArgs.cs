// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ConfidentialLedger.Inputs
{

    /// <summary>
    /// Additional Confidential Ledger properties.
    /// </summary>
    public sealed class LedgerPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("certUsers")]
        private InputList<Inputs.ConfidentialLedgerCertUserArgs>? _certUsers;

        /// <summary>
        /// Array of all the cert based users who can access Confidential Ledger
        /// </summary>
        public InputList<Inputs.ConfidentialLedgerCertUserArgs> CertUsers
        {
            get => _certUsers ?? (_certUsers = new InputList<Inputs.ConfidentialLedgerCertUserArgs>());
            set => _certUsers = value;
        }

        /// <summary>
        /// Name of the Blob Storage Account for saving ledger files
        /// </summary>
        [Input("ledgerStorageAccount")]
        public Input<string>? LedgerStorageAccount { get; set; }

        /// <summary>
        /// Type of Confidential Ledger
        /// </summary>
        [Input("ledgerType")]
        public InputUnion<string, Pulumi.AzureNative.ConfidentialLedger.LedgerType>? LedgerType { get; set; }

        public LedgerPropertiesArgs()
        {
        }
    }
}
