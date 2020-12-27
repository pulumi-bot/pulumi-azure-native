// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageCache.Latest.Outputs
{

    [OutputType]
    public sealed class NfsAccessPolicyResponse
    {
        /// <summary>
        /// The set of rules describing client accesses allowed under this policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.NfsAccessRuleResponse> AccessRules;
        /// <summary>
        /// Name identifying this policy. Access Policy names are not case sensitive.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private NfsAccessPolicyResponse(
            ImmutableArray<Outputs.NfsAccessRuleResponse> accessRules,

            string name)
        {
            AccessRules = accessRules;
            Name = name;
        }
    }
}
