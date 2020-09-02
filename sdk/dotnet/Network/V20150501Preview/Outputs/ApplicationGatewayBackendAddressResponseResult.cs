// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150501Preview.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayBackendAddressResponseResult
    {
        /// <summary>
        /// Gets or sets the dns name
        /// </summary>
        public readonly string? Fqdn;
        /// <summary>
        /// Gets or sets the ip address
        /// </summary>
        public readonly string? IpAddress;

        [OutputConstructor]
        private ApplicationGatewayBackendAddressResponseResult(
            string? fqdn,

            string? ipAddress)
        {
            Fqdn = fqdn;
            IpAddress = ipAddress;
        }
    }
}
