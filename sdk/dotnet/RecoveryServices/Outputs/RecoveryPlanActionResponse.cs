// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.Outputs
{

    [OutputType]
    public sealed class RecoveryPlanActionResponse
    {
        /// <summary>
        /// The action name.
        /// </summary>
        public readonly string ActionName;
        /// <summary>
        /// The custom details.
        /// </summary>
        public readonly Outputs.RecoveryPlanActionDetailsResponse CustomDetails;
        /// <summary>
        /// The list of failover directions.
        /// </summary>
        public readonly ImmutableArray<string> FailoverDirections;
        /// <summary>
        /// The list of failover types.
        /// </summary>
        public readonly ImmutableArray<string> FailoverTypes;

        [OutputConstructor]
        private RecoveryPlanActionResponse(
            string actionName,

            Outputs.RecoveryPlanActionDetailsResponse customDetails,

            ImmutableArray<string> failoverDirections,

            ImmutableArray<string> failoverTypes)
        {
            ActionName = actionName;
            CustomDetails = customDetails;
            FailoverDirections = failoverDirections;
            FailoverTypes = failoverTypes;
        }
    }
}