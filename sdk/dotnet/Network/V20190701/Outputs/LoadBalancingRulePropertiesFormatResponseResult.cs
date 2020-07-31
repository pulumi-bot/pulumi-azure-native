// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Outputs
{

    [OutputType]
    public sealed class LoadBalancingRulePropertiesFormatResponseResult
    {
        /// <summary>
        /// A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend IPs.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? BackendAddressPool;
        /// <summary>
        /// The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note that value 0 enables "Any Port".
        /// </summary>
        public readonly int? BackendPort;
        /// <summary>
        /// Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the frontend of the load balancing rule.
        /// </summary>
        public readonly bool? DisableOutboundSnat;
        /// <summary>
        /// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
        /// </summary>
        public readonly bool? EnableFloatingIP;
        /// <summary>
        /// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
        /// </summary>
        public readonly bool? EnableTcpReset;
        /// <summary>
        /// A reference to frontend IP addresses.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? FrontendIPConfiguration;
        /// <summary>
        /// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values are between 0 and 65534. Note that value 0 enables "Any Port".
        /// </summary>
        public readonly int FrontendPort;
        /// <summary>
        /// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        /// <summary>
        /// The load distribution policy for this rule.
        /// </summary>
        public readonly string? LoadDistribution;
        /// <summary>
        /// The reference of the load balancer probe used by the load balancing rule.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? Probe;
        /// <summary>
        /// The reference to the transport protocol used by the load balancing rule.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The provisioning state of the load balancing rule resource.
        /// </summary>
        public readonly string? ProvisioningState;

        [OutputConstructor]
        private LoadBalancingRulePropertiesFormatResponseResult(
            Outputs.SubResourceResponseResult? backendAddressPool,

            int? backendPort,

            bool? disableOutboundSnat,

            bool? enableFloatingIP,

            bool? enableTcpReset,

            Outputs.SubResourceResponseResult? frontendIPConfiguration,

            int frontendPort,

            int? idleTimeoutInMinutes,

            string? loadDistribution,

            Outputs.SubResourceResponseResult? probe,

            string protocol,

            string? provisioningState)
        {
            BackendAddressPool = backendAddressPool;
            BackendPort = backendPort;
            DisableOutboundSnat = disableOutboundSnat;
            EnableFloatingIP = enableFloatingIP;
            EnableTcpReset = enableTcpReset;
            FrontendIPConfiguration = frontendIPConfiguration;
            FrontendPort = frontendPort;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            LoadDistribution = loadDistribution;
            Probe = probe;
            Protocol = protocol;
            ProvisioningState = provisioningState;
        }
    }
}