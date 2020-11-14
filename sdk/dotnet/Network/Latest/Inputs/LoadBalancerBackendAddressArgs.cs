// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest.Inputs
{

    /// <summary>
    /// Load balancer backend addresses.
    /// </summary>
    public sealed class LoadBalancerBackendAddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Address belonging to the referenced virtual network.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Reference to the frontend ip address configuration defined in regional loadbalancer.
        /// </summary>
        [Input("loadBalancerFrontendIPConfiguration")]
        public Input<Inputs.SubResourceArgs>? LoadBalancerFrontendIPConfiguration { get; set; }

        /// <summary>
        /// Name of the backend address.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Reference to an existing virtual network.
        /// </summary>
        [Input("virtualNetwork")]
        public Input<Inputs.SubResourceArgs>? VirtualNetwork { get; set; }

        public LoadBalancerBackendAddressArgs()
        {
        }
    }
}
