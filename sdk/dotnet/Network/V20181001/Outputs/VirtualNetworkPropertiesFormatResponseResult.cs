// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181001.Outputs
{

    [OutputType]
    public sealed class VirtualNetworkPropertiesFormatResponseResult
    {
        /// <summary>
        /// The AddressSpace that contains an array of IP address ranges that can be used by subnets.
        /// </summary>
        public readonly Outputs.AddressSpaceResponseResult? AddressSpace;
        /// <summary>
        /// The DDoS protection plan associated with the virtual network.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? DdosProtectionPlan;
        /// <summary>
        /// The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
        /// </summary>
        public readonly Outputs.DhcpOptionsResponseResult? DhcpOptions;
        /// <summary>
        /// Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
        /// </summary>
        public readonly bool? EnableDdosProtection;
        /// <summary>
        /// Indicates if VM protection is enabled for all the subnets in the virtual network.
        /// </summary>
        public readonly bool? EnableVmProtection;
        /// <summary>
        /// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The resourceGuid property of the Virtual Network resource.
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// A list of subnets in a Virtual Network.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetResponseResult> Subnets;
        /// <summary>
        /// A list of peerings in a Virtual Network.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualNetworkPeeringResponseResult> VirtualNetworkPeerings;

        [OutputConstructor]
        private VirtualNetworkPropertiesFormatResponseResult(
            Outputs.AddressSpaceResponseResult? addressSpace,

            Outputs.SubResourceResponseResult? ddosProtectionPlan,

            Outputs.DhcpOptionsResponseResult? dhcpOptions,

            bool? enableDdosProtection,

            bool? enableVmProtection,

            string? provisioningState,

            string? resourceGuid,

            ImmutableArray<Outputs.SubnetResponseResult> subnets,

            ImmutableArray<Outputs.VirtualNetworkPeeringResponseResult> virtualNetworkPeerings)
        {
            AddressSpace = addressSpace;
            DdosProtectionPlan = ddosProtectionPlan;
            DhcpOptions = dhcpOptions;
            EnableDdosProtection = enableDdosProtection;
            EnableVmProtection = enableVmProtection;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            Subnets = subnets;
            VirtualNetworkPeerings = virtualNetworkPeerings;
        }
    }
}