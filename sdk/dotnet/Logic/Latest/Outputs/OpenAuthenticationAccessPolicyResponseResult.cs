// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.Latest.Outputs
{

    [OutputType]
    public sealed class OpenAuthenticationAccessPolicyResponseResult
    {
        /// <summary>
        /// The access policy claims.
        /// </summary>
        public readonly ImmutableArray<Outputs.OpenAuthenticationPolicyClaimResponseResult> Claims;
        /// <summary>
        /// Type of provider for OAuth.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private OpenAuthenticationAccessPolicyResponseResult(
            ImmutableArray<Outputs.OpenAuthenticationPolicyClaimResponseResult> claims,

            string type)
        {
            Claims = claims;
            Type = type;
        }
    }
}