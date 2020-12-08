// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20171001.Inputs
{

    /// <summary>
    /// The encryption settings on the storage account.
    /// </summary>
    public sealed class EncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage, Microsoft.Keyvault
        /// </summary>
        [Input("keySource", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Storage.V20171001.KeySource> KeySource { get; set; } = null!;

        /// <summary>
        /// Properties provided by key vault.
        /// </summary>
        [Input("keyVaultProperties")]
        public Input<Inputs.KeyVaultPropertiesArgs>? KeyVaultProperties { get; set; }

        /// <summary>
        /// List of services which support encryption.
        /// </summary>
        [Input("services")]
        public Input<Inputs.EncryptionServicesArgs>? Services { get; set; }

        public EncryptionArgs()
        {
        }
    }
}
