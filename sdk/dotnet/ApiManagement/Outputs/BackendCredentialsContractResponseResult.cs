// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.Outputs
{

    [OutputType]
    public sealed class BackendCredentialsContractResponseResult
    {
        /// <summary>
        /// Authorization header authentication
        /// </summary>
        public readonly Outputs.BackendAuthorizationHeaderCredentialsResponseResult? Authorization;
        /// <summary>
        /// List of Client Certificate Thumbprint.
        /// </summary>
        public readonly ImmutableArray<string> Certificate;
        /// <summary>
        /// Header Parameter description.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Header;
        /// <summary>
        /// Query Parameter description.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Query;

        [OutputConstructor]
        private BackendCredentialsContractResponseResult(
            Outputs.BackendAuthorizationHeaderCredentialsResponseResult? authorization,

            ImmutableArray<string> certificate,

            ImmutableDictionary<string, string>? header,

            ImmutableDictionary<string, string>? query)
        {
            Authorization = authorization;
            Certificate = certificate;
            Header = header;
            Query = query;
        }
    }
}