// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330.Outputs
{

    [OutputType]
    public sealed class VirtualNetworkPropertiesFormatResponseResult
    {
        /// <summary>
        /// Gets or sets AddressSpace that contains an array of IP address ranges that can be used by subnets
        /// </summary>
        public readonly Outputs.AddressSpaceResponseResult? AddressSpace;
        /// <summary>
        /// Gets or sets DHCPOptions that contains an array of DNS servers available to VMs deployed in the virtual network
        /// </summary>
        public readonly Outputs.DhcpOptionsResponseResult? DhcpOptions;
        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets resource GUID property of the VirtualNetwork resource
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// Gets or sets List of subnets in a VirtualNetwork
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetResponseResult> Subnets;

        [OutputConstructor]
        private VirtualNetworkPropertiesFormatResponseResult(
            Outputs.AddressSpaceResponseResult? addressSpace,

            Outputs.DhcpOptionsResponseResult? dhcpOptions,

            string? provisioningState,

            string? resourceGuid,

            ImmutableArray<Outputs.SubnetResponseResult> subnets)
        {
            AddressSpace = addressSpace;
            DhcpOptions = dhcpOptions;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            Subnets = subnets;
        }
    }
}