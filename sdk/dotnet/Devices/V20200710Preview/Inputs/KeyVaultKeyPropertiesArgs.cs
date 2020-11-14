// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200710Preview.Inputs
{

    /// <summary>
    /// The properties of the KeyVault key.
    /// </summary>
    public sealed class KeyVaultKeyPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.KEKIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The identifier of the key.
        /// </summary>
        [Input("keyIdentifier")]
        public Input<string>? KeyIdentifier { get; set; }

        public KeyVaultKeyPropertiesArgs()
        {
        }
    }
}
