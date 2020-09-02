// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventHub.V20180101Preview.Outputs
{

    [OutputType]
    public sealed class KeyVaultPropertiesResponseResult
    {
        /// <summary>
        /// Name of the Key from KeyVault
        /// </summary>
        public readonly string? KeyName;
        /// <summary>
        /// Uri of KeyVault
        /// </summary>
        public readonly string? KeyVaultUri;
        /// <summary>
        /// Key Version
        /// </summary>
        public readonly string? KeyVersion;

        [OutputConstructor]
        private KeyVaultPropertiesResponseResult(
            string? keyName,

            string? keyVaultUri,

            string? keyVersion)
        {
            KeyName = keyName;
            KeyVaultUri = keyVaultUri;
            KeyVersion = keyVersion;
        }
    }
}
