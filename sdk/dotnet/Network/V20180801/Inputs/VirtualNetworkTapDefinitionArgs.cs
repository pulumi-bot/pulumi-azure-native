// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180801.Inputs
{

    /// <summary>
    /// Virtual Network Tap resource
    /// </summary>
    public sealed class VirtualNetworkTapDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reference to the private IP address on the internal Load Balancer that will receive the tap
        /// </summary>
        [Input("destinationLoadBalancerFrontEndIPConfiguration")]
        public Input<Inputs.FrontendIPConfigurationArgs>? DestinationLoadBalancerFrontEndIPConfiguration { get; set; }

        /// <summary>
        /// The reference to the private IP Address of the collector nic that will receive the tap
        /// </summary>
        [Input("destinationNetworkInterfaceIPConfiguration")]
        public Input<Inputs.NetworkInterfaceIPConfigurationArgs>? DestinationNetworkInterfaceIPConfiguration { get; set; }

        /// <summary>
        /// The VXLAN destination port that will receive the tapped traffic.
        /// </summary>
        [Input("destinationPort")]
        public Input<int>? DestinationPort { get; set; }

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public VirtualNetworkTapDefinitionArgs()
        {
        }
    }
}
