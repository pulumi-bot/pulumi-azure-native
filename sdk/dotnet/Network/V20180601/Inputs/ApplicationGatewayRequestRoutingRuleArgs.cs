// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20180601.Inputs
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
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

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
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20180601.ApplicationGatewayRequestRoutingRuleType>? RuleType { get; set; }

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

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
