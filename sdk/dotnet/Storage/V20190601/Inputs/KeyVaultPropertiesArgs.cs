// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20190601.Inputs
{

    /// <summary>
    /// Properties of key vault.
    /// </summary>
    public sealed class KeyVaultPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of KeyVault key.
        /// </summary>
        [Input("keyname")]
        public Input<string>? Keyname { get; set; }

        /// <summary>
        /// The Uri of KeyVault.
        /// </summary>
        [Input("keyvaulturi")]
        public Input<string>? Keyvaulturi { get; set; }

        /// <summary>
        /// The version of KeyVault key.
        /// </summary>
        [Input("keyversion")]
        public Input<string>? Keyversion { get; set; }

        public KeyVaultPropertiesArgs()
        {
        }
    }
}