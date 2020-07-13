// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Outputs
{

    [OutputType]
    public sealed class ContainerNetworkInterfaceConfigurationPropertiesFormatResponse
    {
        /// <summary>
        /// A list of container network interfaces created from this container network interface configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> ContainerNetworkInterfaces;
        /// <summary>
        /// A list of ip configurations of the container network interface configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.IPConfigurationProfileResponse> IpConfigurations;
        /// <summary>
        /// The provisioning state of the container network interface configuration resource.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private ContainerNetworkInterfaceConfigurationPropertiesFormatResponse(
            ImmutableArray<Outputs.SubResourceResponse> containerNetworkInterfaces,

            ImmutableArray<Outputs.IPConfigurationProfileResponse> ipConfigurations,

            string provisioningState)
        {
            ContainerNetworkInterfaces = containerNetworkInterfaces;
            IpConfigurations = ipConfigurations;
            ProvisioningState = provisioningState;
        }
    }
}