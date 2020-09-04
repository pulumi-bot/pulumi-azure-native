// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class ApplicationHealthPolicyResponseResult
    {
        /// <summary>
        /// The health policy used by default to evaluate the health of a service type.
        /// </summary>
        public readonly Outputs.ServiceTypeHealthPolicyResponseResult? DefaultServiceTypeHealthPolicy;
        /// <summary>
        /// The map with service type health policy per service type name. The map is empty by default.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ServiceTypeHealthPolicyResponseResult>? ServiceTypeHealthPolicies;

        [OutputConstructor]
        private ApplicationHealthPolicyResponseResult(
            Outputs.ServiceTypeHealthPolicyResponseResult? defaultServiceTypeHealthPolicy,

            ImmutableDictionary<string, Outputs.ServiceTypeHealthPolicyResponseResult>? serviceTypeHealthPolicies)
        {
            DefaultServiceTypeHealthPolicy = defaultServiceTypeHealthPolicy;
            ServiceTypeHealthPolicies = serviceTypeHealthPolicies;
        }
    }
}
