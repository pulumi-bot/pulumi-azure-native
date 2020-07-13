// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network
{
    public static class GetApplicationGatewayWebApplicationFirewallPolicy
    {
        public static Task<GetApplicationGatewayWebApplicationFirewallPolicyResult> InvokeAsync(GetApplicationGatewayWebApplicationFirewallPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationGatewayWebApplicationFirewallPolicyResult>("azurerm:network:getApplicationGatewayWebApplicationFirewallPolicy", args ?? new GetApplicationGatewayWebApplicationFirewallPolicyArgs(), options.WithVersion());
    }


    public sealed class GetApplicationGatewayWebApplicationFirewallPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationGatewayWebApplicationFirewallPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationGatewayWebApplicationFirewallPolicyResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the web application firewall policy.
        /// </summary>
        public readonly Outputs.WebApplicationFirewallPolicyPropertiesFormatResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationGatewayWebApplicationFirewallPolicyResult(
            string etag,

            string? location,

            string name,

            Outputs.WebApplicationFirewallPolicyPropertiesFormatResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}