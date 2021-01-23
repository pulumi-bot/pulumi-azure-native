// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20201201.Outputs
{

    [OutputType]
    public sealed class EncryptionDetailsResponse
    {
        /// <summary>
        /// Customer Managed Key Details
        /// </summary>
        public readonly Outputs.CustomerManagedKeyDetailsResponse? Cmk;
        /// <summary>
        /// Double Encryption enabled
        /// </summary>
        public readonly bool DoubleEncryptionEnabled;

        [OutputConstructor]
        private EncryptionDetailsResponse(
            Outputs.CustomerManagedKeyDetailsResponse? cmk,

            bool doubleEncryptionEnabled)
        {
            Cmk = cmk;
            DoubleEncryptionEnabled = doubleEncryptionEnabled;
        }
    }
}
