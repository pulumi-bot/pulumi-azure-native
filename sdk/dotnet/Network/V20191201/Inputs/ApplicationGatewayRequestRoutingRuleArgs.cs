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
    /// Request routing rule of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayRequestRoutingRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backend address pool resource of the application gateway.
        /// </summary>
        [Input("backendAddressPool")]
        public Input<Inputs.SubResourceArgs>? BackendAddressPool { get; set; }

        /// <summary>
        /// Backend http settings resource of the application gateway.
        /// </summary>
        [Input("backendHttpSettings")]
        public Input<Inputs.SubResourceArgs>? BackendHttpSettings { get; set; }

        /// <summary>
        /// Http listener resource of the application gateway.
        /// </summary>
        [Input("httpListener")]
        public Input<Inputs.SubResourceArgs>? HttpListener { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the request routing rule that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority of the request routing rule.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Redirect configuration resource of the application gateway.
        /// </summary>
        [Input("redirectConfiguration")]
        public Input<Inputs.SubResourceArgs>? RedirectConfiguration { get; set; }

        /// <summary>
        /// Rewrite Rule Set resource in Basic rule of the application gateway.
        /// </summary>
        [Input("rewriteRuleSet")]
        public Input<Inputs.SubResourceArgs>? RewriteRuleSet { get; set; }

        /// <summary>
        /// Rule type.
        /// </summary>
        [Input("ruleType")]
        public Input<string>? RuleType { get; set; }

        /// <summary>
        /// URL path map resource of the application gateway.
        /// </summary>
        [Input("urlPathMap")]
        public Input<Inputs.SubResourceArgs>? UrlPathMap { get; set; }

        public ApplicationGatewayRequestRoutingRuleArgs()
        {
        }
    }
}
