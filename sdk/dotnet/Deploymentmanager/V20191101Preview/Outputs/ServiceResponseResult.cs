// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DeploymentManager.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class ServiceResponseResult
    {
        /// <summary>
        /// Name of the service.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The detailed information about the units that make up the service.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceUnitResponseResult> ServiceUnits;
        /// <summary>
        /// The Azure location to which the resources in the service belong to or should be deployed to.
        /// </summary>
        public readonly string TargetLocation;
        /// <summary>
        /// The subscription to which the resources in the service belong to or should be deployed to.
        /// </summary>
        public readonly string TargetSubscriptionId;

        [OutputConstructor]
        private ServiceResponseResult(
            string? name,

            ImmutableArray<Outputs.ServiceUnitResponseResult> serviceUnits,

            string targetLocation,

            string targetSubscriptionId)
        {
            Name = name;
            ServiceUnits = serviceUnits;
            TargetLocation = targetLocation;
            TargetSubscriptionId = targetSubscriptionId;
        }
    }
}
