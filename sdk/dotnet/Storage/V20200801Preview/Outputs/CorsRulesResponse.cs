// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20200801Preview.Outputs
{

    [OutputType]
    public sealed class CorsRulesResponse
    {
        /// <summary>
        /// The List of CORS rules. You can include up to five CorsRule elements in the request. 
        /// </summary>
        public readonly ImmutableArray<Outputs.CorsRuleResponse> CorsRules;

        [OutputConstructor]
        private CorsRulesResponse(ImmutableArray<Outputs.CorsRuleResponse> corsRules)
        {
            CorsRules = corsRules;
        }
    }
}
