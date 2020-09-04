// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class HttpMessageDiagnosticResponseResult
    {
        /// <summary>
        /// Body logging settings.
        /// </summary>
        public readonly Outputs.BodyDiagnosticSettingsResponseResult? Body;
        /// <summary>
        /// Array of HTTP Headers to log.
        /// </summary>
        public readonly ImmutableArray<string> Headers;

        [OutputConstructor]
        private HttpMessageDiagnosticResponseResult(
            Outputs.BodyDiagnosticSettingsResponseResult? body,

            ImmutableArray<string> headers)
        {
            Body = body;
            Headers = headers;
        }
    }
}
