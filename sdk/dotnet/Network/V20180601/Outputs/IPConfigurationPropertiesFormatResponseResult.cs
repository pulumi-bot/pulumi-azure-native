// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180601.Outputs
{

    [OutputType]
    public sealed class IPConfigurationPropertiesFormatResponseResult
    {
        /// <summary>
        /// The private IP address of the IP configuration.
        /// </summary>
        public readonly string? PrivateIPAddress;
        /// <summary>
        /// The private IP allocation method. Possible values are 'Static' and 'Dynamic'.
        /// </summary>
        public readonly string? PrivateIPAllocationMethod;
        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The reference of the public IP resource.
        /// </summary>
        public readonly Outputs.PublicIPAddressResponseResult? PublicIPAddress;
        /// <summary>
        /// The reference of the subnet resource.
        /// </summary>
        public readonly Outputs.SubnetResponseResult? Subnet;

        [OutputConstructor]
        private IPConfigurationPropertiesFormatResponseResult(
            string? privateIPAddress,

            string? privateIPAllocationMethod,

            string? provisioningState,

            Outputs.PublicIPAddressResponseResult? publicIPAddress,

            Outputs.SubnetResponseResult? subnet)
        {
            PrivateIPAddress = privateIPAddress;
            PrivateIPAllocationMethod = privateIPAllocationMethod;
            ProvisioningState = provisioningState;
            PublicIPAddress = publicIPAddress;
            Subnet = subnet;
        }
    }
}