// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Inputs
{

    /// <summary>
    /// The reference to the key vault key.
    /// </summary>
    public sealed class KeyVaultKeyReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The private key name in key vault.
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        /// <summary>
        /// The key vault reference.
        /// </summary>
        [Input("keyVault", required: true)]
        public Input<Inputs.KeyVaultKeyReferenceKeyVaultArgs> KeyVault { get; set; } = null!;

        /// <summary>
        /// The private key version in key vault.
        /// </summary>
        [Input("keyVersion")]
        public Input<string>? KeyVersion { get; set; }

        public KeyVaultKeyReferenceArgs()
        {
        }
    }
}
