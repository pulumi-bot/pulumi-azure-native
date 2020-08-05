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
    /// Properties of Frontend IP Configuration of the load balancer.
    /// </summary>
    public sealed class FrontendIPConfigurationPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        [Input("inboundNatPools", required: true)]
        private List<Inputs.SubResourceResponseArgs>? _inboundNatPools;

        /// <summary>
        /// Read only. Inbound pools URIs that use this frontend IP.
        /// </summary>
        public List<Inputs.SubResourceResponseArgs> InboundNatPools
        {
            get => _inboundNatPools ?? (_inboundNatPools = new List<Inputs.SubResourceResponseArgs>());
            set => _inboundNatPools = value;
        }

        [Input("inboundNatRules", required: true)]
        private List<Inputs.SubResourceResponseArgs>? _inboundNatRules;

        /// <summary>
        /// Read only. Inbound rules URIs that use this frontend IP.
        /// </summary>
        public List<Inputs.SubResourceResponseArgs> InboundNatRules
        {
            get => _inboundNatRules ?? (_inboundNatRules = new List<Inputs.SubResourceResponseArgs>());
            set => _inboundNatRules = value;
        }

        [Input("loadBalancingRules", required: true)]
        private List<Inputs.SubResourceResponseArgs>? _loadBalancingRules;

        /// <summary>
        /// Gets load balancing rules URIs that use this frontend IP.
        /// </summary>
        public List<Inputs.SubResourceResponseArgs> LoadBalancingRules
        {
            get => _loadBalancingRules ?? (_loadBalancingRules = new List<Inputs.SubResourceResponseArgs>());
            set => _loadBalancingRules = value;
        }

        [Input("outboundRules", required: true)]
        private List<Inputs.SubResourceResponseArgs>? _outboundRules;

        /// <summary>
        /// Read only. Outbound rules URIs that use this frontend IP.
        /// </summary>
        public List<Inputs.SubResourceResponseArgs> OutboundRules
        {
            get => _outboundRules ?? (_outboundRules = new List<Inputs.SubResourceResponseArgs>());
            set => _outboundRules = value;
        }

        /// <summary>
        /// The private IP address of the IP configuration.
        /// </summary>
        [Input("privateIPAddress")]
        public string? PrivateIPAddress { get; set; }

        /// <summary>
        /// The Private IP allocation method. Possible values are: 'Static' and 'Dynamic'.
        /// </summary>
        [Input("privateIPAllocationMethod")]
        public string? PrivateIPAllocationMethod { get; set; }

        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        /// <summary>
        /// The reference of the Public IP resource.
        /// </summary>
        [Input("publicIPAddress")]
        public Inputs.PublicIPAddressResponseArgs? PublicIPAddress { get; set; }

        /// <summary>
        /// The reference of the Public IP Prefix resource.
        /// </summary>
        [Input("publicIPPrefix")]
        public Inputs.SubResourceResponseArgs? PublicIPPrefix { get; set; }

        /// <summary>
        /// The reference of the subnet resource.
        /// </summary>
        [Input("subnet")]
        public Inputs.SubnetResponseArgs? Subnet { get; set; }

        public FrontendIPConfigurationPropertiesFormatResponseArgs()
        {
        }
    }
}
