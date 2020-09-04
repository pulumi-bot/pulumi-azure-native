// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.PolicyInsights.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class RemediationDeploymentSummaryResponseResult
    {
        /// <summary>
        /// The number of deployments required by the remediation that have failed.
        /// </summary>
        public readonly int? FailedDeployments;
        /// <summary>
        /// The number of deployments required by the remediation that have succeeded.
        /// </summary>
        public readonly int? SuccessfulDeployments;
        /// <summary>
        /// The number of deployments required by the remediation.
        /// </summary>
        public readonly int? TotalDeployments;

        [OutputConstructor]
        private RemediationDeploymentSummaryResponseResult(
            int? failedDeployments,

            int? successfulDeployments,

            int? totalDeployments)
        {
            FailedDeployments = failedDeployments;
            SuccessfulDeployments = successfulDeployments;
            TotalDeployments = totalDeployments;
        }
    }
}
