// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200901.Inputs
{

    /// <summary>
    /// The CORS policy for the Cosmos DB database account.
    /// </summary>
    public sealed class CorsPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The request headers that the origin domain may specify on the CORS request.
        /// </summary>
        [Input("allowedHeaders")]
        public Input<string>? AllowedHeaders { get; set; }

        /// <summary>
        /// The methods (HTTP request verbs) that the origin domain may use for a CORS request.
        /// </summary>
        [Input("allowedMethods")]
        public Input<string>? AllowedMethods { get; set; }

        /// <summary>
        /// The origin domains that are permitted to make a request against the service via CORS.
        /// </summary>
        [Input("allowedOrigins", required: true)]
        public Input<string> AllowedOrigins { get; set; } = null!;

        /// <summary>
        /// The response headers that may be sent in the response to the CORS request and exposed by the browser to the request issuer.
        /// </summary>
        [Input("exposedHeaders")]
        public Input<string>? ExposedHeaders { get; set; }

        /// <summary>
        /// The maximum amount time that a browser should cache the preflight OPTIONS request.
        /// </summary>
        [Input("maxAgeInSeconds")]
        public Input<int>? MaxAgeInSeconds { get; set; }

        public CorsPolicyArgs()
        {
        }
    }
}
