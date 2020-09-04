// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.KeyVault.V20180214Preview.Outputs
{

    [OutputType]
    public sealed class SkuResponseResult
    {
        /// <summary>
        /// SKU family name
        /// </summary>
        public readonly string Family;
        /// <summary>
        /// SKU name to specify whether the key vault is a standard vault or a premium vault.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private SkuResponseResult(
            string family,

            string name)
        {
            Family = family;
            Name = name;
        }
    }
}
