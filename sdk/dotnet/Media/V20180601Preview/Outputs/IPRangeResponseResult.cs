// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class IPRangeResponseResult
    {
        /// <summary>
        /// The IP address.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// The friendly name for the IP address range.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The subnet mask prefix length (see CIDR notation).
        /// </summary>
        public readonly int? SubnetPrefixLength;

        [OutputConstructor]
        private IPRangeResponseResult(
            string? address,

            string? name,

            int? subnetPrefixLength)
        {
            Address = address;
            Name = name;
            SubnetPrefixLength = subnetPrefixLength;
        }
    }
}
