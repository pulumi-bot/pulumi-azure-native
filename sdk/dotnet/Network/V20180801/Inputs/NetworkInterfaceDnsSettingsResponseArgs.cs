// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180801.Inputs
{

    /// <summary>
    /// DNS settings of a network interface.
    /// </summary>
    public sealed class NetworkInterfaceDnsSettingsResponseArgs : Pulumi.InvokeArgs
    {
        [Input("appliedDnsServers")]
        private List<string>? _appliedDnsServers;

        /// <summary>
        /// If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set. This property is what is configured on each of those VMs.
        /// </summary>
        public List<string> AppliedDnsServers
        {
            get => _appliedDnsServers ?? (_appliedDnsServers = new List<string>());
            set => _appliedDnsServers = value;
        }

        [Input("dnsServers")]
        private List<string>? _dnsServers;

        /// <summary>
        /// List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch to azure provided DNS resolution. 'AzureProvidedDNS' value cannot be combined with other IPs, it must be the only value in dnsServers collection.
        /// </summary>
        public List<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new List<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// Relative DNS name for this NIC used for internal communications between VMs in the same virtual network.
        /// </summary>
        [Input("internalDnsNameLabel")]
        public string? InternalDnsNameLabel { get; set; }

        /// <summary>
        /// Even if internalDnsNameLabel is not specified, a DNS entry is created for the primary NIC of the VM. This DNS name can be constructed by concatenating the VM name with the value of internalDomainNameSuffix.
        /// </summary>
        [Input("internalDomainNameSuffix")]
        public string? InternalDomainNameSuffix { get; set; }

        /// <summary>
        /// Fully qualified DNS name supporting internal communications between VMs in the same virtual network.
        /// </summary>
        [Input("internalFqdn")]
        public string? InternalFqdn { get; set; }

        public NetworkInterfaceDnsSettingsResponseArgs()
        {
        }
    }
}
