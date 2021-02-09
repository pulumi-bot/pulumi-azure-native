// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Migrate.Latest.Outputs
{

    [OutputType]
    public sealed class VirtualNetworkResourceSettingsResponse
    {
        /// <summary>
        /// Gets or sets the address prefixes for the virtual network.
        /// </summary>
        public readonly ImmutableArray<string> AddressSpace;
        /// <summary>
        /// Gets or sets DHCPOptions that contains an array of DNS servers available to VMs
        /// deployed in the virtual network.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// Gets or sets a value indicating whether gets or sets whether the
        /// DDOS protection should be switched on.
        /// </summary>
        public readonly bool? EnableDdosProtection;
        /// <summary>
        /// The resource type. For example, the value can be Microsoft.Compute/virtualMachines.
        /// Expected value is 'Microsoft.Network/virtualNetworks'.
        /// </summary>
        public readonly string ResourceType;
        /// <summary>
        /// Gets or sets List of subnets in a VirtualNetwork.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetResourceSettingsResponse> Subnets;
        /// <summary>
        /// Gets or sets the target Resource name.
        /// </summary>
        public readonly string TargetResourceName;

        [OutputConstructor]
        private VirtualNetworkResourceSettingsResponse(
            ImmutableArray<string> addressSpace,

            ImmutableArray<string> dnsServers,

            bool? enableDdosProtection,

            string resourceType,

            ImmutableArray<Outputs.SubnetResourceSettingsResponse> subnets,

            string targetResourceName)
        {
            AddressSpace = addressSpace;
            DnsServers = dnsServers;
            EnableDdosProtection = enableDdosProtection;
            ResourceType = resourceType;
            Subnets = subnets;
            TargetResourceName = targetResourceName;
        }
    }
}
