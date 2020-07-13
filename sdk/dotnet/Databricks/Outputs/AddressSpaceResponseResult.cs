// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Databricks.Outputs
{

    [OutputType]
    public sealed class AddressSpaceResponseResult
    {
        /// <summary>
        /// A list of address blocks reserved for this virtual network in CIDR notation.
        /// </summary>
        public readonly ImmutableArray<string> AddressPrefixes;

        [OutputConstructor]
        private AddressSpaceResponseResult(ImmutableArray<string> addressPrefixes)
        {
            AddressPrefixes = addressPrefixes;
        }
    }
}