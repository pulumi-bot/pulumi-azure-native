// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201.Inputs
{

    /// <summary>
    /// The private link service ip configuration.
    /// </summary>
    public sealed class PrivateLinkServiceIpConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of private link service ip configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the ip configuration is primary or not.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// The private IP address of the IP configuration.
        /// </summary>
        [Input("privateIPAddress")]
        public Input<string>? PrivateIPAddress { get; set; }

        /// <summary>
        /// Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.
        /// </summary>
        [Input("privateIPAddressVersion")]
        public Input<string>? PrivateIPAddressVersion { get; set; }

        /// <summary>
        /// The private IP address allocation method.
        /// </summary>
        [Input("privateIPAllocationMethod")]
        public Input<string>? PrivateIPAllocationMethod { get; set; }

        /// <summary>
        /// The reference to the subnet resource.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.SubnetArgs>? Subnet { get; set; }

        public PrivateLinkServiceIpConfigurationArgs()
        {
        }
    }
}
