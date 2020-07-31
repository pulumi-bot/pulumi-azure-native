// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150615.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayBackendHttpSettingsPropertiesFormatResponseResult
    {
        /// <summary>
        /// Cookie based affinity. Possible values are: 'Enabled' and 'Disabled'.
        /// </summary>
        public readonly string? CookieBasedAffinity;
        /// <summary>
        /// Port
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Probe resource of an application gateway.
        /// </summary>
        public readonly Outputs.SubResourceResponseResult? Probe;
        /// <summary>
        /// Protocol. Possible values are: 'Http' and 'Https'.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Gets or sets Provisioning state of the backend http settings resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Request timeout in seconds. Application Gateway will fail the request if response is not received within RequestTimeout. Acceptable values are from 1 second to 86400 seconds.
        /// </summary>
        public readonly int? RequestTimeout;

        [OutputConstructor]
        private ApplicationGatewayBackendHttpSettingsPropertiesFormatResponseResult(
            string? cookieBasedAffinity,

            int? port,

            Outputs.SubResourceResponseResult? probe,

            string? protocol,

            string? provisioningState,

            int? requestTimeout)
        {
            CookieBasedAffinity = cookieBasedAffinity;
            Port = port;
            Probe = probe;
            Protocol = protocol;
            ProvisioningState = provisioningState;
            RequestTimeout = requestTimeout;
        }
    }
}