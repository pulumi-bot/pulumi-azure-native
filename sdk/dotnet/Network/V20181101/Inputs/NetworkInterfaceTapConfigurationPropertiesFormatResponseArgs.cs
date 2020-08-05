// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181101.Inputs
{

    /// <summary>
    /// Properties of Virtual Network Tap configuration.
    /// </summary>
    public sealed class NetworkInterfaceTapConfigurationPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The provisioning state of the network interface tap configuration. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState", required: true)]
        public string ProvisioningState { get; set; } = null!;

        /// <summary>
        /// The reference of the Virtual Network Tap resource.
        /// </summary>
        [Input("virtualNetworkTap")]
        public Inputs.VirtualNetworkTapResponseArgs? VirtualNetworkTap { get; set; }

        public NetworkInterfaceTapConfigurationPropertiesFormatResponseArgs()
        {
        }
    }
}
