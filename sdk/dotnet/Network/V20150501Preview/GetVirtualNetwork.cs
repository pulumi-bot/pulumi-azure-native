// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150501Preview
{
    public static class GetVirtualNetwork
    {
        public static Task<GetVirtualNetworkResult> InvokeAsync(GetVirtualNetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkResult>("azurerm:network/v20150501preview:getVirtualNetwork", args ?? new GetVirtualNetworkArgs(), options.WithVersion());
    }


    public sealed class GetVirtualNetworkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network.
        /// </summary>
        [Input("virtualNetworkName", required: true)]
        public string VirtualNetworkName { get; set; } = null!;

        public GetVirtualNetworkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualNetworkResult
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
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets resource guid property of the VirtualNetwork resource
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// Gets or sets List of subnets in a VirtualNetwork
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetResponseResult> Subnets;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVirtualNetworkResult(
            Outputs.AddressSpaceResponseResult? addressSpace,

            Outputs.DhcpOptionsResponseResult? dhcpOptions,

            string? etag,

            string location,

            string name,

            string? provisioningState,

            string? resourceGuid,

            ImmutableArray<Outputs.SubnetResponseResult> subnets,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AddressSpace = addressSpace;
            DhcpOptions = dhcpOptions;
            Etag = etag;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            Subnets = subnets;
            Tags = tags;
            Type = type;
        }
    }
}
