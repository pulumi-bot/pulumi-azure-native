// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20200301.Outputs
{

    [OutputType]
    public sealed class ManagedClusterLoadBalancerProfileResponseResult
    {
        /// <summary>
        /// Desired number of allocated SNAT ports per VM. Allowed values must be in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports.
        /// </summary>
        public readonly int? AllocatedOutboundPorts;
        /// <summary>
        /// The effective outbound IP resources of the cluster load balancer.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceReferenceResponseResult> EffectiveOutboundIPs;
        /// <summary>
        /// Desired outbound flow idle timeout in minutes. Allowed values must be in the range of 4 to 120 (inclusive). The default value is 30 minutes.
        /// </summary>
        public readonly int? IdleTimeoutInMinutes;
        /// <summary>
        /// Desired managed outbound IPs for the cluster load balancer.
        /// </summary>
        public readonly Outputs.ManagedClusterLoadBalancerProfileResponsePropertiesResult? ManagedOutboundIPs;
        /// <summary>
        /// Desired outbound IP Prefix resources for the cluster load balancer.
        /// </summary>
        public readonly Outputs.ManagedClusterLoadBalancerProfileResponsePropertiesResult? OutboundIPPrefixes;
        /// <summary>
        /// Desired outbound IP resources for the cluster load balancer.
        /// </summary>
        public readonly Outputs.ManagedClusterLoadBalancerProfileResponsePropertiesResult? OutboundIPs;

        [OutputConstructor]
        private ManagedClusterLoadBalancerProfileResponseResult(
            int? allocatedOutboundPorts,

            ImmutableArray<Outputs.ResourceReferenceResponseResult> effectiveOutboundIPs,

            int? idleTimeoutInMinutes,

            Outputs.ManagedClusterLoadBalancerProfileResponsePropertiesResult? managedOutboundIPs,

            Outputs.ManagedClusterLoadBalancerProfileResponsePropertiesResult? outboundIPPrefixes,

            Outputs.ManagedClusterLoadBalancerProfileResponsePropertiesResult? outboundIPs)
        {
            AllocatedOutboundPorts = allocatedOutboundPorts;
            EffectiveOutboundIPs = effectiveOutboundIPs;
            IdleTimeoutInMinutes = idleTimeoutInMinutes;
            ManagedOutboundIPs = managedOutboundIPs;
            OutboundIPPrefixes = outboundIPPrefixes;
            OutboundIPs = outboundIPs;
        }
    }
}