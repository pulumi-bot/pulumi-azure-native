// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Scheduler.V20140801Preview.Inputs
{

    public sealed class HttpRequestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the http authentication.
        /// </summary>
        [Input("authentication")]
        public Input<Inputs.HttpAuthenticationArgs>? Authentication { get; set; }

        /// <summary>
        /// Gets or sets the request body.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;

        /// <summary>
        /// Gets or sets the headers.
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        /// <summary>
        /// Gets or sets the method of the request.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Gets or sets the Uri.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public HttpRequestArgs()
        {
        }
    }
}
