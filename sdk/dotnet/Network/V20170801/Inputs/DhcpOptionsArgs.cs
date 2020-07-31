// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Inputs
{

    /// <summary>
    /// DhcpOptions contains an array of DNS servers available to VMs deployed in the virtual network. Standard DHCP option for a subnet overrides VNET DHCP options.
    /// </summary>
    public sealed class DhcpOptionsArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// The list of DNS servers IP addresses.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        public DhcpOptionsArgs()
        {
        }
    }
}