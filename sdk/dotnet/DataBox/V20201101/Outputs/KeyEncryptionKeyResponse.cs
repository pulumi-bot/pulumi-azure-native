// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBox.V20201101.Outputs
{

    [OutputType]
    public sealed class KeyEncryptionKeyResponse
    {
        /// <summary>
        /// Managed identity properties used for key encryption.
        /// </summary>
        public readonly Outputs.IdentityPropertiesResponse? IdentityProperties;
        /// <summary>
        /// Type of encryption key used for key encryption.
        /// </summary>
        public readonly string KekType;
        /// <summary>
        /// Key encryption key. It is required in case of Customer managed KekType.
        /// </summary>
        public readonly string? KekUrl;
        /// <summary>
        /// Kek vault resource id. It is required in case of Customer managed KekType.
        /// </summary>
        public readonly string? KekVaultResourceID;

        [OutputConstructor]
        private KeyEncryptionKeyResponse(
            Outputs.IdentityPropertiesResponse? identityProperties,

            string kekType,

            string? kekUrl,

            string? kekVaultResourceID)
        {
            IdentityProperties = identityProperties;
            KekType = kekType;
            KekUrl = kekUrl;
            KekVaultResourceID = kekVaultResourceID;
        }
    }
}
