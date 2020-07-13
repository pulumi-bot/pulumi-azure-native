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
    public sealed class NetworkInterfaceTapConfigurationPropertiesFormatResponse
    {
        /// <summary>
        /// The provisioning state of the network interface tap configuration resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The reference to the Virtual Network Tap resource.
        /// </summary>
        public readonly Outputs.VirtualNetworkTapResponse? VirtualNetworkTap;

        [OutputConstructor]
        private NetworkInterfaceTapConfigurationPropertiesFormatResponse(
            string provisioningState,

            Outputs.VirtualNetworkTapResponse? virtualNetworkTap)
        {
            ProvisioningState = provisioningState;
            VirtualNetworkTap = virtualNetworkTap;
        }
    }
}