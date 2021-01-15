// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20200930.Inputs
{

    /// <summary>
    /// Key Vault Key Url and vault id of KeK, KeK is optional and when provided is used to unwrap the encryptionKey
    /// </summary>
    public sealed class KeyVaultAndKeyReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Url pointing to a key or secret in KeyVault
        /// </summary>
        [Input("keyUrl", required: true)]
        public Input<string> KeyUrl { get; set; } = null!;

        /// <summary>
        /// Resource id of the KeyVault containing the key or secret
        /// </summary>
        [Input("sourceVault", required: true)]
        public Input<Inputs.SourceVaultArgs> SourceVault { get; set; } = null!;

        public KeyVaultAndKeyReferenceArgs()
        {
        }
    }
}
