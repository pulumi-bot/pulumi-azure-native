// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20170601.Inputs
{

    /// <summary>
    /// Redirect configuration of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayRedirectConfigurationArgs : Pulumi.ResourceArgs
    {
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
        /// Include path in the redirected url.
        /// </summary>
        [Input("includePath")]
        public Input<bool>? IncludePath { get; set; }

        /// <summary>
        /// Include query string in the redirected url.
        /// </summary>
        [Input("includeQueryString")]
        public Input<bool>? IncludeQueryString { get; set; }

        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20170601.ApplicationGatewayRedirectType>? RedirectType { get; set; }

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

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

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

        public ApplicationGatewayRedirectConfigurationArgs()
        {
        }
    }
}
