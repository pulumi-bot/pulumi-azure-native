// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Media.Latest.Inputs
{

    public sealed class AccountEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The properties of the key used to encrypt the account.
        /// </summary>
        [Input("keyVaultProperties")]
        public Input<Inputs.KeyVaultPropertiesArgs>? KeyVaultProperties { get; set; }

        /// <summary>
        /// The type of key used to encrypt the Account Key.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Media.Latest.AccountEncryptionKeyType> Type { get; set; } = null!;

        public AccountEncryptionArgs()
        {
        }
    }
}
