// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201.Outputs
{

    [OutputType]
    public sealed class PrivateLinkServicePropertiesResponseResult
    {
        /// <summary>
        /// The alias of the private link service.
        /// </summary>
        public readonly string Alias;
        /// <summary>
        /// The auto-approval list of the private link service.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AutoApproval;
        /// <summary>
        /// Whether the private link service is enabled for proxy protocol or not.
        /// </summary>
        public readonly bool? EnableProxyProtocol;
        /// <summary>
        /// The list of Fqdn.
        /// </summary>
        public readonly ImmutableArray<string> Fqdns;
        /// <summary>
        /// An array of private link service IP configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateLinkServiceIpConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// An array of references to the load balancer IP configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.FrontendIPConfigurationResponseResult> LoadBalancerFrontendIpConfigurations;
        /// <summary>
        /// An array of references to the network interfaces created for this private link service.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceResponseResult> NetworkInterfaces;
        /// <summary>
        /// An array of list about connections to the private endpoint.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult> PrivateEndpointConnections;
        /// <summary>
        /// The provisioning state of the private link service resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The visibility list of the private link service.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Visibility;

        [OutputConstructor]
        private PrivateLinkServicePropertiesResponseResult(
            string alias,

            ImmutableDictionary<string, object>? autoApproval,

            bool? enableProxyProtocol,

            ImmutableArray<string> fqdns,

            ImmutableArray<Outputs.PrivateLinkServiceIpConfigurationResponseResult> ipConfigurations,

            ImmutableArray<Outputs.FrontendIPConfigurationResponseResult> loadBalancerFrontendIpConfigurations,

            ImmutableArray<Outputs.NetworkInterfaceResponseResult> networkInterfaces,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult> privateEndpointConnections,

            string provisioningState,

            ImmutableDictionary<string, object>? visibility)
        {
            Alias = alias;
            AutoApproval = autoApproval;
            EnableProxyProtocol = enableProxyProtocol;
            Fqdns = fqdns;
            IpConfigurations = ipConfigurations;
            LoadBalancerFrontendIpConfigurations = loadBalancerFrontendIpConfigurations;
            NetworkInterfaces = networkInterfaces;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            Visibility = visibility;
        }
    }
}