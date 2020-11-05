// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200901.Outputs
{

    [OutputType]
    public sealed class CorsPolicyResponse
    {
        /// <summary>
        /// The request headers that the origin domain may specify on the CORS request.
        /// </summary>
        public readonly string? AllowedHeaders;
        /// <summary>
        /// The methods (HTTP request verbs) that the origin domain may use for a CORS request.
        /// </summary>
        public readonly string? AllowedMethods;
        /// <summary>
        /// The origin domains that are permitted to make a request against the service via CORS.
        /// </summary>
        public readonly string AllowedOrigins;
        /// <summary>
        /// The response headers that may be sent in the response to the CORS request and exposed by the browser to the request issuer.
        /// </summary>
        public readonly string? ExposedHeaders;
        /// <summary>
        /// The maximum amount time that a browser should cache the preflight OPTIONS request.
        /// </summary>
        public readonly int? MaxAgeInSeconds;

        [OutputConstructor]
        private CorsPolicyResponse(
            string? allowedHeaders,

            string? allowedMethods,

            string allowedOrigins,

            string? exposedHeaders,

            int? maxAgeInSeconds)
        {
            AllowedHeaders = allowedHeaders;
            AllowedMethods = allowedMethods;
            AllowedOrigins = allowedOrigins;
            ExposedHeaders = exposedHeaders;
            MaxAgeInSeconds = maxAgeInSeconds;
        }
    }
}
