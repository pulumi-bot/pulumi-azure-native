// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.Outputs
{

    [OutputType]
    public sealed class MediaServicePropertiesResponse
    {
        /// <summary>
        /// The account encryption properties.
        /// </summary>
        public readonly Outputs.AccountEncryptionResponse? Encryption;
        /// <summary>
        /// The Media Services account ID.
        /// </summary>
        public readonly string MediaServiceId;
        /// <summary>
        /// The storage accounts for this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageAccountResponse> StorageAccounts;
        public readonly string? StorageAuthentication;

        [OutputConstructor]
        private MediaServicePropertiesResponse(
            Outputs.AccountEncryptionResponse? encryption,

            string mediaServiceId,

            ImmutableArray<Outputs.StorageAccountResponse> storageAccounts,

            string? storageAuthentication)
        {
            Encryption = encryption;
            MediaServiceId = mediaServiceId;
            StorageAccounts = storageAccounts;
            StorageAuthentication = storageAuthentication;
        }
    }
}