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
    /// Properties of redirect configuration of the application gateway.
    /// </summary>
    public sealed class ApplicationGatewayRedirectConfigurationPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Include path in the redirected url.
        /// </summary>
        [Input("includePath")]
        public Input<bool>? IncludePath { get; set; }

        /// <summary>
        /// Include query string in the redirected url.
        /// </summary>
        [Input("includeQueryString")]
        public Input<bool>? IncludeQueryString { get; set; }

        [Input("pathRules")]
        private InputList<Inputs.SubResourceArgs>? _pathRules;

        /// <summary>
        /// Path rules specifying redirect configuration.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> PathRules
        {
            get => _pathRules ?? (_pathRules = new InputList<Inputs.SubResourceArgs>());
            set => _pathRules = value;
        }

        /// <summary>
        /// Supported http redirection types - Permanent, Temporary, Found, SeeOther.
        /// </summary>
        [Input("redirectType")]
        public Input<string>? RedirectType { get; set; }

        [Input("requestRoutingRules")]
        private InputList<Inputs.SubResourceArgs>? _requestRoutingRules;

        /// <summary>
        /// Request routing specifying redirect configuration.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> RequestRoutingRules
        {
            get => _requestRoutingRules ?? (_requestRoutingRules = new InputList<Inputs.SubResourceArgs>());
            set => _requestRoutingRules = value;
        }

        /// <summary>
        /// Reference to a listener to redirect the request to.
        /// </summary>
        [Input("targetListener")]
        public Input<Inputs.SubResourceArgs>? TargetListener { get; set; }

        /// <summary>
        /// Url to redirect the request to.
        /// </summary>
        [Input("targetUrl")]
        public Input<string>? TargetUrl { get; set; }

        [Input("urlPathMaps")]
        private InputList<Inputs.SubResourceArgs>? _urlPathMaps;

        /// <summary>
        /// Url path maps specifying default redirect configuration.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> UrlPathMaps
        {
            get => _urlPathMaps ?? (_urlPathMaps = new InputList<Inputs.SubResourceArgs>());
            set => _urlPathMaps = value;
        }

        public ApplicationGatewayRedirectConfigurationPropertiesFormatArgs()
        {
        }
    }
}