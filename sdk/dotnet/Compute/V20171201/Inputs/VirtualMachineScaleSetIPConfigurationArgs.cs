// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20171201.Inputs
{

    /// <summary>
    /// Describes a virtual machine scale set network profile's IP configuration.
    /// </summary>
    public sealed class VirtualMachineScaleSetIPConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("applicationGatewayBackendAddressPools")]
        private InputList<Inputs.SubResourceArgs>? _applicationGatewayBackendAddressPools;

        /// <summary>
        /// Specifies an array of references to backend address pools of application gateways. A scale set can reference backend address pools of multiple application gateways. Multiple scale sets cannot use the same application gateway.
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
        /// Specifies an array of references to backend address pools of load balancers. A scale set can reference backend address pools of one public and one internal load balancer. Multiple scale sets cannot use the same load balancer.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> LoadBalancerBackendAddressPools
        {
            get => _loadBalancerBackendAddressPools ?? (_loadBalancerBackendAddressPools = new InputList<Inputs.SubResourceArgs>());
            set => _loadBalancerBackendAddressPools = value;
        }

        [Input("loadBalancerInboundNatPools")]
        private InputList<Inputs.SubResourceArgs>? _loadBalancerInboundNatPools;

        /// <summary>
        /// Specifies an array of references to inbound Nat pools of the load balancers. A scale set can reference inbound nat pools of one public and one internal load balancer. Multiple scale sets cannot use the same load balancer
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
        /// Specifies the primary network interface in case the virtual machine has more than 1 network interface.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// Available from Api-Version 2017-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: 'IPv4' and 'IPv6'.
        /// </summary>
        [Input("privateIPAddressVersion")]
        public InputUnion<string, Pulumi.AzureNextGen.Compute.V20171201.IPVersion>? PrivateIPAddressVersion { get; set; }

        /// <summary>
        /// The publicIPAddressConfiguration.
        /// </summary>
        [Input("publicIPAddressConfiguration")]
        public Input<Inputs.VirtualMachineScaleSetPublicIPAddressConfigurationArgs>? PublicIPAddressConfiguration { get; set; }

        /// <summary>
        /// Specifies the identifier of the subnet.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.ApiEntityReferenceArgs>? Subnet { get; set; }

        public VirtualMachineScaleSetIPConfigurationArgs()
        {
        }
    }
}
