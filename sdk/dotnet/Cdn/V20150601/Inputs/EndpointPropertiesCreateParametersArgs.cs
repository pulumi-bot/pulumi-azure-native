// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20150601.Inputs
{

    public sealed class EndpointPropertiesCreateParametersArgs : Pulumi.ResourceArgs
    {
        [Input("contentTypesToCompress")]
        private InputList<string>? _contentTypesToCompress;

        /// <summary>
        /// List of content types on which compression will be applied. The value for the elements should be a valid MIME type.
        /// </summary>
        public InputList<string> ContentTypesToCompress
        {
            get => _contentTypesToCompress ?? (_contentTypesToCompress = new InputList<string>());
            set => _contentTypesToCompress = value;
        }

        /// <summary>
        /// Indicates whether content compression is enabled. Default value is false. If compression is enabled, the content transferred from the CDN endpoint to the end user will be compressed. The requested content must be larger than 1 byte and smaller than 1 MB.
        /// </summary>
        [Input("isCompressionEnabled")]
        public Input<bool>? IsCompressionEnabled { get; set; }

        /// <summary>
        /// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        /// </summary>
        [Input("isHttpAllowed")]
        public Input<bool>? IsHttpAllowed { get; set; }

        /// <summary>
        /// Indicates whether https traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        /// </summary>
        [Input("isHttpsAllowed")]
        public Input<bool>? IsHttpsAllowed { get; set; }

        /// <summary>
        /// The host header CDN provider will send along with content requests to origins. The default value is the host name of the origin.
        /// </summary>
        [Input("originHostHeader")]
        public Input<string>? OriginHostHeader { get; set; }

        /// <summary>
        /// The path used for origin requests.
        /// </summary>
        [Input("originPath")]
        public Input<string>? OriginPath { get; set; }

        [Input("origins", required: true)]
        private InputList<Inputs.DeepCreatedOriginArgs>? _origins;

        /// <summary>
        /// The set of origins for the CDN endpoint. When multiple origins exist, the first origin will be used as primary and rest will be used as failover options.
        /// </summary>
        public InputList<Inputs.DeepCreatedOriginArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.DeepCreatedOriginArgs>());
            set => _origins = value;
        }

        /// <summary>
        /// Defines the query string caching behavior.
        /// </summary>
        [Input("queryStringCachingBehavior")]
        public Input<string>? QueryStringCachingBehavior { get; set; }

        public EndpointPropertiesCreateParametersArgs()
        {
        }
    }
}