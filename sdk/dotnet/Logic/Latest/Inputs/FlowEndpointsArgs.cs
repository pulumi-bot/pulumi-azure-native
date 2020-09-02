// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.Latest.Inputs
{

    /// <summary>
    /// The flow endpoints configuration.
    /// </summary>
    public sealed class FlowEndpointsArgs : Pulumi.ResourceArgs
    {
        [Input("accessEndpointIpAddresses")]
        private InputList<Inputs.IpAddressArgs>? _accessEndpointIpAddresses;

        /// <summary>
        /// The access endpoint ip address.
        /// </summary>
        public InputList<Inputs.IpAddressArgs> AccessEndpointIpAddresses
        {
            get => _accessEndpointIpAddresses ?? (_accessEndpointIpAddresses = new InputList<Inputs.IpAddressArgs>());
            set => _accessEndpointIpAddresses = value;
        }

        [Input("outgoingIpAddresses")]
        private InputList<Inputs.IpAddressArgs>? _outgoingIpAddresses;

        /// <summary>
        /// The outgoing ip address.
        /// </summary>
        public InputList<Inputs.IpAddressArgs> OutgoingIpAddresses
        {
            get => _outgoingIpAddresses ?? (_outgoingIpAddresses = new InputList<Inputs.IpAddressArgs>());
            set => _outgoingIpAddresses = value;
        }

        public FlowEndpointsArgs()
        {
        }
    }
}