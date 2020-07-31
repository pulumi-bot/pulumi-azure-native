// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401.Inputs
{

    /// <summary>
    /// Properties of VirtualNetworkGatewayIPConfiguration.
    /// </summary>
    public sealed class VirtualNetworkGatewayIPConfigurationPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Private IP Address for this gateway.
        /// </summary>
        [Input("privateIPAddress", required: true)]
        public string PrivateIPAddress { get; set; } = null!;

        /// <summary>
        /// The private IP address allocation method.
        /// </summary>
        [Input("privateIPAllocationMethod")]
        public string? PrivateIPAllocationMethod { get; set; }

        /// <summary>
        /// The provisioning state of the virtual network gateway IP configuration resource.
        /// </summary>
        [Input("provisioningState", required: true)]
        public string ProvisioningState { get; set; } = null!;

        /// <summary>
        /// The reference to the public IP resource.
        /// </summary>
        [Input("publicIPAddress")]
        public Inputs.SubResourceResponseArgs? PublicIPAddress { get; set; }

        /// <summary>
        /// The reference to the subnet resource.
        /// </summary>
        [Input("subnet")]
        public Inputs.SubResourceResponseArgs? Subnet { get; set; }

        public VirtualNetworkGatewayIPConfigurationPropertiesFormatResponseArgs()
        {
        }
    }
}