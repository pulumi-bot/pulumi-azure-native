// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Inputs
{

    /// <summary>
    /// Describes a reference to Key Vault Secret
    /// </summary>
    public sealed class KeyVaultSecretReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL referencing a secret in a Key Vault.
        /// </summary>
        [Input("secretUrl", required: true)]
        public Input<string> SecretUrl { get; set; } = null!;

        /// <summary>
        /// The relative URL of the Key Vault containing the secret.
        /// </summary>
        [Input("sourceVault", required: true)]
        public Input<Inputs.SubResourceArgs> SourceVault { get; set; } = null!;

        public KeyVaultSecretReferenceArgs()
        {
        }
    }
}
