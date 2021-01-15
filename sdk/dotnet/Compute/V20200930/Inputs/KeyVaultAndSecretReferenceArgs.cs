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
    /// Key Vault Secret Url and vault id of the encryption key 
    /// </summary>
    public sealed class KeyVaultAndSecretReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Url pointing to a key or secret in KeyVault
        /// </summary>
        [Input("secretUrl", required: true)]
        public Input<string> SecretUrl { get; set; } = null!;

        /// <summary>
        /// Resource id of the KeyVault containing the key or secret
        /// </summary>
        [Input("sourceVault", required: true)]
        public Input<Inputs.SourceVaultArgs> SourceVault { get; set; } = null!;

        public KeyVaultAndSecretReferenceArgs()
        {
        }
    }
}
