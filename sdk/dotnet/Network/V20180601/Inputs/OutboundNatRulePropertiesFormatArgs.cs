// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180601.Inputs
{

    /// <summary>
    /// Outbound NAT pool of the load balancer.
    /// </summary>
    public sealed class OutboundNatRulePropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of outbound ports to be used for NAT.
        /// </summary>
        [Input("allocatedOutboundPorts")]
        public Input<int>? AllocatedOutboundPorts { get; set; }

        /// <summary>
        /// A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend IPs.
        /// </summary>
        [Input("backendAddressPool", required: true)]
        public Input<Inputs.SubResourceArgs> BackendAddressPool { get; set; } = null!;

        [Input("frontendIPConfigurations")]
        private InputList<Inputs.SubResourceArgs>? _frontendIPConfigurations;

        /// <summary>
        /// The Frontend IP addresses of the load balancer.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> FrontendIPConfigurations
        {
            get => _frontendIPConfigurations ?? (_frontendIPConfigurations = new InputList<Inputs.SubResourceArgs>());
            set => _frontendIPConfigurations = value;
        }

        /// <summary>
        /// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public OutboundNatRulePropertiesFormatArgs()
        {
        }
    }
}