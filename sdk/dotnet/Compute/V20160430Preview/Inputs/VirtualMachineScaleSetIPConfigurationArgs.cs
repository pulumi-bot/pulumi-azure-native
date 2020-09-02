// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Inputs
{

    /// <summary>
    /// Describes a virtual machine scale set network profile's IP configuration.
    /// </summary>
    public sealed class VirtualMachineScaleSetIPConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("applicationGatewayBackendAddressPools")]
        private InputList<Inputs.SubResourceArgs>? _applicationGatewayBackendAddressPools;

        /// <summary>
        /// The application gateway backend address pools.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> ApplicationGatewayBackendAddressPools
        {
            get => _applicationGatewayBackendAddressPools ?? (_applicationGatewayBackendAddressPools = new InputList<Inputs.SubResourceArgs>());
            set => _applicationGatewayBackendAddressPools = value;
        }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("loadBalancerBackendAddressPools")]
        private InputList<Inputs.SubResourceArgs>? _loadBalancerBackendAddressPools;

        /// <summary>
        /// The load balancer backend address pools.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> LoadBalancerBackendAddressPools
        {
            get => _loadBalancerBackendAddressPools ?? (_loadBalancerBackendAddressPools = new InputList<Inputs.SubResourceArgs>());
            set => _loadBalancerBackendAddressPools = value;
        }

        [Input("loadBalancerInboundNatPools")]
        private InputList<Inputs.SubResourceArgs>? _loadBalancerInboundNatPools;

        /// <summary>
        /// The load balancer inbound nat pools.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> LoadBalancerInboundNatPools
        {
            get => _loadBalancerInboundNatPools ?? (_loadBalancerInboundNatPools = new InputList<Inputs.SubResourceArgs>());
            set => _loadBalancerInboundNatPools = value;
        }

        /// <summary>
        /// The IP configuration name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The subnet.
        /// </summary>
        [Input("subnet", required: true)]
        public Input<Inputs.ApiEntityReferenceArgs> Subnet { get; set; } = null!;

        public VirtualMachineScaleSetIPConfigurationArgs()
        {
        }
    }
}
