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
    /// Properties of IP configuration.
    /// </summary>
    public sealed class NetworkInterfaceIPConfigurationPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        [Input("applicationGatewayBackendAddressPools")]
        private List<Inputs.ApplicationGatewayBackendAddressPoolResponseArgs>? _applicationGatewayBackendAddressPools;

        /// <summary>
        /// The reference of ApplicationGatewayBackendAddressPool resource.
        /// </summary>
        public List<Inputs.ApplicationGatewayBackendAddressPoolResponseArgs> ApplicationGatewayBackendAddressPools
        {
            get => _applicationGatewayBackendAddressPools ?? (_applicationGatewayBackendAddressPools = new List<Inputs.ApplicationGatewayBackendAddressPoolResponseArgs>());
            set => _applicationGatewayBackendAddressPools = value;
        }

        [Input("applicationSecurityGroups")]
        private List<Inputs.ApplicationSecurityGroupResponseArgs>? _applicationSecurityGroups;

        /// <summary>
        /// Application security groups in which the IP configuration is included.
        /// </summary>
        public List<Inputs.ApplicationSecurityGroupResponseArgs> ApplicationSecurityGroups
        {
            get => _applicationSecurityGroups ?? (_applicationSecurityGroups = new List<Inputs.ApplicationSecurityGroupResponseArgs>());
            set => _applicationSecurityGroups = value;
        }

        [Input("loadBalancerBackendAddressPools")]
        private List<Inputs.BackendAddressPoolResponseArgs>? _loadBalancerBackendAddressPools;

        /// <summary>
        /// The reference of LoadBalancerBackendAddressPool resource.
        /// </summary>
        public List<Inputs.BackendAddressPoolResponseArgs> LoadBalancerBackendAddressPools
        {
            get => _loadBalancerBackendAddressPools ?? (_loadBalancerBackendAddressPools = new List<Inputs.BackendAddressPoolResponseArgs>());
            set => _loadBalancerBackendAddressPools = value;
        }

        [Input("loadBalancerInboundNatRules")]
        private List<Inputs.InboundNatRuleResponseArgs>? _loadBalancerInboundNatRules;

        /// <summary>
        /// A list of references of LoadBalancerInboundNatRules.
        /// </summary>
        public List<Inputs.InboundNatRuleResponseArgs> LoadBalancerInboundNatRules
        {
            get => _loadBalancerInboundNatRules ?? (_loadBalancerInboundNatRules = new List<Inputs.InboundNatRuleResponseArgs>());
            set => _loadBalancerInboundNatRules = value;
        }

        /// <summary>
        /// Gets whether this is a primary customer address on the network interface.
        /// </summary>
        [Input("primary")]
        public bool? Primary { get; set; }

        /// <summary>
        /// Private IP address of the IP configuration.
        /// </summary>
        [Input("privateIPAddress")]
        public string? PrivateIPAddress { get; set; }

        /// <summary>
        /// Available from Api-Version 2016-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: 'IPv4' and 'IPv6'.
        /// </summary>
        [Input("privateIPAddressVersion")]
        public string? PrivateIPAddressVersion { get; set; }

        /// <summary>
        /// Defines how a private IP address is assigned. Possible values are: 'Static' and 'Dynamic'.
        /// </summary>
        [Input("privateIPAllocationMethod")]
        public string? PrivateIPAllocationMethod { get; set; }

        /// <summary>
        /// The provisioning state of the network interface IP configuration. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        /// <summary>
        /// Public IP address bound to the IP configuration.
        /// </summary>
        [Input("publicIPAddress")]
        public Inputs.PublicIPAddressResponseArgs? PublicIPAddress { get; set; }

        /// <summary>
        /// Subnet bound to the IP configuration.
        /// </summary>
        [Input("subnet")]
        public Inputs.SubnetResponseArgs? Subnet { get; set; }

        [Input("virtualNetworkTaps")]
        private List<Inputs.VirtualNetworkTapResponseArgs>? _virtualNetworkTaps;

        /// <summary>
        /// The reference to Virtual Network Taps.
        /// </summary>
        public List<Inputs.VirtualNetworkTapResponseArgs> VirtualNetworkTaps
        {
            get => _virtualNetworkTaps ?? (_virtualNetworkTaps = new List<Inputs.VirtualNetworkTapResponseArgs>());
            set => _virtualNetworkTaps = value;
        }

        public NetworkInterfaceIPConfigurationPropertiesFormatResponseArgs()
        {
        }
    }
}
