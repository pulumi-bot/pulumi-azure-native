// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190601.Inputs
{

    /// <summary>
    /// UrlPathMaps give a url path to the backend mapping information for PathBasedRouting.
    /// </summary>
    public sealed class ApplicationGatewayUrlPathMapArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default backend address pool resource of URL path map.
        /// </summary>
        [Input("defaultBackendAddressPool")]
        public Input<Inputs.SubResourceArgs>? DefaultBackendAddressPool { get; set; }

        /// <summary>
        /// Default backend http settings resource of URL path map.
        /// </summary>
        [Input("defaultBackendHttpSettings")]
        public Input<Inputs.SubResourceArgs>? DefaultBackendHttpSettings { get; set; }

        /// <summary>
        /// Default redirect configuration resource of URL path map.
        /// </summary>
        [Input("defaultRedirectConfiguration")]
        public Input<Inputs.SubResourceArgs>? DefaultRedirectConfiguration { get; set; }

        /// <summary>
        /// Default Rewrite rule set resource of URL path map.
        /// </summary>
        [Input("defaultRewriteRuleSet")]
        public Input<Inputs.SubResourceArgs>? DefaultRewriteRuleSet { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the URL path map that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pathRules")]
        private InputList<Inputs.ApplicationGatewayPathRuleArgs>? _pathRules;

        /// <summary>
        /// Path rule of URL path map resource.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayPathRuleArgs> PathRules
        {
            get => _pathRules ?? (_pathRules = new InputList<Inputs.ApplicationGatewayPathRuleArgs>());
            set => _pathRules = value;
        }

        /// <summary>
        /// Provisioning state of the backend http settings resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ApplicationGatewayUrlPathMapArgs()
        {
        }
    }
}
