// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101.Inputs
{

    /// <summary>
    /// IP configuration for virtual network gateway.
    /// </summary>
    public sealed class VirtualNetworkGatewayIPConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The private IP address allocation method.
        /// </summary>
        [Input("privateIPAllocationMethod")]
        public Input<string>? PrivateIPAllocationMethod { get; set; }

        /// <summary>
        /// The reference to the public IP resource.
        /// </summary>
        [Input("publicIPAddress")]
        public Input<Inputs.SubResourceArgs>? PublicIPAddress { get; set; }

        /// <summary>
        /// The reference to the subnet resource.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.SubResourceArgs>? Subnet { get; set; }

        public VirtualNetworkGatewayIPConfigurationArgs()
        {
        }
    }
}
