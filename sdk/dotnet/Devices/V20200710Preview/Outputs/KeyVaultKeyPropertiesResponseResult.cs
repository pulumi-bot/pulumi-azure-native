// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200710Preview.Outputs
{

    [OutputType]
    public sealed class KeyVaultKeyPropertiesResponseResult
    {
        /// <summary>
        /// The identifier of the key.
        /// </summary>
        public readonly string? KeyIdentifier;

        [OutputConstructor]
        private KeyVaultKeyPropertiesResponseResult(string? keyIdentifier)
        {
            KeyIdentifier = keyIdentifier;
        }
    }
}
