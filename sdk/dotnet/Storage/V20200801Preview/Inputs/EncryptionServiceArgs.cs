// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20200801Preview.Inputs
{

    /// <summary>
    /// A service that allows server-side encryption to be used.
    /// </summary>
    public sealed class EncryptionServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean indicating whether or not the service encrypts the data as it is stored.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped encryption key will be used. 'Service' key type implies that a default service key is used.
        /// </summary>
        [Input("keyType")]
        public InputUnion<string, Pulumi.AzureNextGen.Storage.V20200801Preview.KeyType>? KeyType { get; set; }

        public EncryptionServiceArgs()
        {
        }
    }
}
