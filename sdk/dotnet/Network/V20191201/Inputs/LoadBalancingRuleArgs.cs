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
    /// A load balancing rule for a load balancer.
    /// </summary>
    public sealed class LoadBalancingRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend IPs.
        /// </summary>
        [Input("backendAddressPool")]
        public Input<Inputs.SubResourceArgs>? BackendAddressPool { get; set; }

        /// <summary>
        /// The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note that value 0 enables "Any Port".
        /// </summary>
        [Input("backendPort")]
        public Input<int>? BackendPort { get; set; }

        /// <summary>
        /// Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the frontend of the load balancing rule.
        /// </summary>
        [Input("disableOutboundSnat")]
        public Input<bool>? DisableOutboundSnat { get; set; }

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
        /// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values are between 0 and 65534. Note that value 0 enables "Any Port".
        /// </summary>
        [Input("frontendPort", required: true)]
        public Input<int> FrontendPort { get; set; } = null!;

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
        /// The load distribution policy for this rule.
        /// </summary>
        [Input("loadDistribution")]
        public Input<string>? LoadDistribution { get; set; }

        /// <summary>
        /// The name of the resource that is unique within the set of load balancing rules used by the load balancer. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The reference to the load balancer probe used by the load balancing rule.
        /// </summary>
        [Input("probe")]
        public Input<Inputs.SubResourceArgs>? Probe { get; set; }

        /// <summary>
        /// The reference to the transport protocol used by the load balancing rule.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public LoadBalancingRuleArgs()
        {
        }
    }
}
