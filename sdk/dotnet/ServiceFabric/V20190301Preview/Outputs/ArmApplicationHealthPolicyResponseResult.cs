// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301Preview.Outputs
{

    [OutputType]
    public sealed class ArmApplicationHealthPolicyResponseResult
    {
        /// <summary>
        /// Indicates whether warnings are treated with the same severity as errors.
        /// </summary>
        public readonly bool? ConsiderWarningAsError;
        /// <summary>
        /// The health policy used by default to evaluate the health of a service type.
        /// </summary>
        public readonly Outputs.ArmServiceTypeHealthPolicyResponseResult? DefaultServiceTypeHealthPolicy;
        /// <summary>
        /// The maximum allowed percentage of unhealthy deployed applications. Allowed values are Byte values from zero to 100.
        /// The percentage represents the maximum tolerated percentage of deployed applications that can be unhealthy before the application is considered in error.
        /// This is calculated by dividing the number of unhealthy deployed applications over the number of nodes where the application is currently deployed on in the cluster.
        /// The computation rounds up to tolerate one failure on small numbers of nodes. Default percentage is zero.
        /// </summary>
        public readonly int? MaxPercentUnhealthyDeployedApplications;
        /// <summary>
        /// The map with service type health policy per service type name. The map is empty by default.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ArmServiceTypeHealthPolicyResponseResult>? ServiceTypeHealthPolicyMap;

        [OutputConstructor]
        private ArmApplicationHealthPolicyResponseResult(
            bool? considerWarningAsError,

            Outputs.ArmServiceTypeHealthPolicyResponseResult? defaultServiceTypeHealthPolicy,

            int? maxPercentUnhealthyDeployedApplications,

            ImmutableDictionary<string, Outputs.ArmServiceTypeHealthPolicyResponseResult>? serviceTypeHealthPolicyMap)
        {
            ConsiderWarningAsError = considerWarningAsError;
            DefaultServiceTypeHealthPolicy = defaultServiceTypeHealthPolicy;
            MaxPercentUnhealthyDeployedApplications = maxPercentUnhealthyDeployedApplications;
            ServiceTypeHealthPolicyMap = serviceTypeHealthPolicyMap;
        }
    }
}
