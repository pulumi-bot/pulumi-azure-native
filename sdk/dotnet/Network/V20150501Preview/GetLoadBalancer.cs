// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150501Preview
{
    public static class GetLoadBalancer
    {
        public static Task<GetLoadBalancerResult> InvokeAsync(GetLoadBalancerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerResult>("azurerm:network/v20150501preview:getLoadBalancer", args ?? new GetLoadBalancerArgs(), options.WithVersion());
    }


    public sealed class GetLoadBalancerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the loadBalancer.
        /// </summary>
        [Input("loadBalancerName", required: true)]
        public string LoadBalancerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetLoadBalancerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLoadBalancerResult
    {
        /// <summary>
        /// Gets or sets Pools of backend IP addresses
        /// </summary>
        public readonly ImmutableArray<Outputs.BackendAddressPoolResponseResult> BackendAddressPools;
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Gets or sets frontend IP addresses of the load balancer
        /// </summary>
        public readonly ImmutableArray<Outputs.FrontendIpConfigurationResponseResult> FrontendIPConfigurations;
        /// <summary>
        /// Gets or sets inbound NAT pools
        /// </summary>
        public readonly ImmutableArray<Outputs.InboundNatPoolResponseResult> InboundNatPools;
        /// <summary>
        /// Gets or sets list of inbound rules
        /// </summary>
        public readonly ImmutableArray<Outputs.InboundNatRuleResponseResult> InboundNatRules;
        /// <summary>
        /// Gets or sets load balancing rules
        /// </summary>
        public readonly ImmutableArray<Outputs.LoadBalancingRuleResponseResult> LoadBalancingRules;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets outbound NAT rules
        /// </summary>
        public readonly ImmutableArray<Outputs.OutboundNatRuleResponseResult> OutboundNatRules;
        /// <summary>
        /// Gets or sets list of Load balancer probes
        /// </summary>
        public readonly ImmutableArray<Outputs.ProbeResponseResult> Probes;
        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets resource guid property of the Load balancer resource
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetLoadBalancerResult(
            ImmutableArray<Outputs.BackendAddressPoolResponseResult> backendAddressPools,

            string? etag,

            ImmutableArray<Outputs.FrontendIpConfigurationResponseResult> frontendIPConfigurations,

            ImmutableArray<Outputs.InboundNatPoolResponseResult> inboundNatPools,

            ImmutableArray<Outputs.InboundNatRuleResponseResult> inboundNatRules,

            ImmutableArray<Outputs.LoadBalancingRuleResponseResult> loadBalancingRules,

            string location,

            string name,

            ImmutableArray<Outputs.OutboundNatRuleResponseResult> outboundNatRules,

            ImmutableArray<Outputs.ProbeResponseResult> probes,

            string? provisioningState,

            string? resourceGuid,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            BackendAddressPools = backendAddressPools;
            Etag = etag;
            FrontendIPConfigurations = frontendIPConfigurations;
            InboundNatPools = inboundNatPools;
            InboundNatRules = inboundNatRules;
            LoadBalancingRules = loadBalancingRules;
            Location = location;
            Name = name;
            OutboundNatRules = outboundNatRules;
            Probes = probes;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            Tags = tags;
            Type = type;
        }
    }
}
