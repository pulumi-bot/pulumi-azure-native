// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.Outputs
{

    [OutputType]
    public sealed class KeyVaultSecretReferenceResponse
    {
        /// <summary>
        /// The URL referencing a secret in the Key Vault.
        /// </summary>
        public readonly string SecretUrl;
        /// <summary>
        /// Fully qualified resource identifier of the Key Vault.
        /// </summary>
        public readonly Outputs.ResourceIdResponse SourceVault;

        [OutputConstructor]
        private KeyVaultSecretReferenceResponse(
            string secretUrl,

            Outputs.ResourceIdResponse sourceVault)
        {
            SecretUrl = secretUrl;
            SourceVault = sourceVault;
        }
    }
}