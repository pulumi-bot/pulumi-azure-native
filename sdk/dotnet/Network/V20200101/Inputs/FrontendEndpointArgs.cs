// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200101.Inputs
{

    /// <summary>
    /// A frontend endpoint used for routing.
    /// </summary>
    public sealed class FrontendEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The host name of the frontendEndpoint. Must be a domain name.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource status.
        /// </summary>
        [Input("resourceState")]
        public Input<string>? ResourceState { get; set; }

        /// <summary>
        /// Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
        /// </summary>
        [Input("sessionAffinityEnabledState")]
        public Input<string>? SessionAffinityEnabledState { get; set; }

        /// <summary>
        /// UNUSED. This field will be ignored. The TTL to use in seconds for session affinity, if applicable.
        /// </summary>
        [Input("sessionAffinityTtlSeconds")]
        public Input<int>? SessionAffinityTtlSeconds { get; set; }

        /// <summary>
        /// Defines the Web Application Firewall policy for each host (if applicable)
        /// </summary>
        [Input("webApplicationFirewallPolicyLink")]
        public Input<Inputs.FrontendEndpointUpdateParametersPropertiesArgs>? WebApplicationFirewallPolicyLink { get; set; }

        public FrontendEndpointArgs()
        {
        }
    }
}
