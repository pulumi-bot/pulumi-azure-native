// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101.Inputs
{

    /// <summary>
    /// Inbound NAT rule of the load balancer.
    /// </summary>
    public sealed class InboundNatRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
        /// </summary>
        [Input("backendPort")]
        public Input<int>? BackendPort { get; set; }

        /// <summary>
        /// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
        /// </summary>
        [Input("enableFloatingIP")]
        public Input<bool>? EnableFloatingIP { get; set; }

        /// <summary>
        /// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        /// </summary>
        [Input("enableTcpReset")]
        public Input<bool>? EnableTcpReset { get; set; }

        /// <summary>
        /// A reference to frontend IP addresses.
        /// </summary>
        [Input("frontendIPConfiguration")]
        public Input<Inputs.SubResourceArgs>? FrontendIPConfiguration { get; set; }

        /// <summary>
        /// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
        /// </summary>
        [Input("frontendPort")]
        public Input<int>? FrontendPort { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        /// <summary>
        /// The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The reference to the transport protocol used by the load balancing rule.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public InboundNatRuleArgs()
        {
        }
    }
}
