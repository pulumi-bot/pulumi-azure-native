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
    /// Outbound rule of the load balancer.
    /// </summary>
    public sealed class OutboundRuleArgs : Pulumi.ResourceArgs
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

        /// <summary>
        /// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        /// </summary>
        [Input("enableTcpReset")]
        public Input<bool>? EnableTcpReset { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("frontendIPConfigurations", required: true)]
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
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The timeout for the TCP idle connection
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Protocol - TCP, UDP or All
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public OutboundRuleArgs()
        {
        }
    }
}
