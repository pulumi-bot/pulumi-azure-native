// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Inputs
{

    /// <summary>
    /// Properties of request routing rule of the application gateway.
    /// </summary>
    public sealed class ApplicationGatewayRequestRoutingRulePropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backend address pool resource of the application gateway. 
        /// </summary>
        [Input("backendAddressPool")]
        public Input<Inputs.SubResourceArgs>? BackendAddressPool { get; set; }

        /// <summary>
        /// Frontend port resource of the application gateway.
        /// </summary>
        [Input("backendHttpSettings")]
        public Input<Inputs.SubResourceArgs>? BackendHttpSettings { get; set; }

        /// <summary>
        /// Http listener resource of the application gateway. 
        /// </summary>
        [Input("httpListener")]
        public Input<Inputs.SubResourceArgs>? HttpListener { get; set; }

        /// <summary>
        /// Provisioning state of the request routing rule resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Redirect configuration resource of the application gateway.
        /// </summary>
        [Input("redirectConfiguration")]
        public Input<Inputs.SubResourceArgs>? RedirectConfiguration { get; set; }

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

        public ApplicationGatewayRequestRoutingRulePropertiesFormatArgs()
        {
        }
    }
}