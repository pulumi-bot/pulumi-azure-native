// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200831Preview.Outputs
{

    [OutputType]
    public sealed class EncryptionPropertiesDescriptionResponse
    {
        /// <summary>
        /// The source of the key.
        /// </summary>
        public readonly string? KeySource;
        /// <summary>
        /// The properties of the KeyVault key.
        /// </summary>
        public readonly ImmutableArray<Outputs.KeyVaultKeyPropertiesResponse> KeyVaultProperties;

        [OutputConstructor]
        private EncryptionPropertiesDescriptionResponse(
            string? keySource,

            ImmutableArray<Outputs.KeyVaultKeyPropertiesResponse> keyVaultProperties)
        {
            KeySource = keySource;
            KeyVaultProperties = keyVaultProperties;
        }
    }
}
