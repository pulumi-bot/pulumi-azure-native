// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Inputs
{

    /// <summary>
    /// Encryption settings for one disk volume.
    /// </summary>
    public sealed class EncryptionSettingsElementArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key Vault Secret Url and vault id of the disk encryption key
        /// </summary>
        [Input("diskEncryptionKey")]
        public Input<Inputs.KeyVaultAndSecretReferenceArgs>? DiskEncryptionKey { get; set; }

        /// <summary>
        /// Key Vault Key Url and vault id of the key encryption key. KeyEncryptionKey is optional and when provided is used to unwrap the disk encryption key.
        /// </summary>
        [Input("keyEncryptionKey")]
        public Input<Inputs.KeyVaultAndKeyReferenceArgs>? KeyEncryptionKey { get; set; }

        public EncryptionSettingsElementArgs()
        {
        }
    }
}