// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201.Outputs
{

    [OutputType]
    public sealed class CorsSettingsResponseResult
    {
        /// <summary>
        /// Gets or sets the list of origins that should be allowed to make cross-origin
        /// calls (for example: http://example.com:12345). Use "*" to allow all.
        /// </summary>
        public readonly ImmutableArray<string> AllowedOrigins;
        /// <summary>
        /// Gets or sets whether CORS requests with credentials are allowed. See 
        /// https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS#Requests_with_credentials
        /// for more details.
        /// </summary>
        public readonly bool? SupportCredentials;

        [OutputConstructor]
        private CorsSettingsResponseResult(
            ImmutableArray<string> allowedOrigins,

            bool? supportCredentials)
        {
            AllowedOrigins = allowedOrigins;
            SupportCredentials = supportCredentials;
        }
    }
}