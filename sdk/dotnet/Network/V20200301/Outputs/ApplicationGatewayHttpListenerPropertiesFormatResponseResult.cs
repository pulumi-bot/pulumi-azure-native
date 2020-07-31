// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200301.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayHttpListenerPropertiesFormatResponseResult
    {
        /// <summary>
        /// Custom error configurations of the HTTP listener.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayCustomErrorResponseResult> CustomErrorConfigurations;
        /// <summary>
        /// Reference to the FirewallPolicy resource.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? FirewallPolicy;
        /// <summary>
        /// Frontend IP configuration resource of an application gateway.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? FrontendIPConfiguration;
        /// <summary>
        /// Frontend port resource of an application gateway.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? FrontendPort;
        /// <summary>
        /// Host name of HTTP listener.
        /// </summary>
        public readonly string? HostName;
        /// <summary>
        /// List of Host names for HTTP Listener that allows special wildcard characters as well.
        /// </summary>
        public readonly ImmutableArray<string> HostNames;
        /// <summary>
        /// Protocol of the HTTP listener.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// The provisioning state of the HTTP listener resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Applicable only if protocol is https. Enables SNI for multi-hosting.
        /// </summary>
        public readonly bool? RequireServerNameIndication;
        /// <summary>
        /// SSL certificate resource of an application gateway.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? SslCertificate;

        [OutputConstructor]
        private ApplicationGatewayHttpListenerPropertiesFormatResponseResult(
            ImmutableArray<Outputs.ApplicationGatewayCustomErrorResponseResult> customErrorConfigurations,

            Outputs.SubResourceResponseResult? firewallPolicy,

            Outputs.SubResourceResponseResult? frontendIPConfiguration,

            Outputs.SubResourceResponseResult? frontendPort,

            string? hostName,

            ImmutableArray<string> hostNames,

            string? protocol,

            string provisioningState,

            bool? requireServerNameIndication,

            Outputs.SubResourceResponseResult? sslCertificate)
        {
            CustomErrorConfigurations = customErrorConfigurations;
            FirewallPolicy = firewallPolicy;
            FrontendIPConfiguration = frontendIPConfiguration;
            FrontendPort = frontendPort;
            HostName = hostName;
            HostNames = hostNames;
            Protocol = protocol;
            ProvisioningState = provisioningState;
            RequireServerNameIndication = requireServerNameIndication;
            SslCertificate = sslCertificate;
        }
    }
}