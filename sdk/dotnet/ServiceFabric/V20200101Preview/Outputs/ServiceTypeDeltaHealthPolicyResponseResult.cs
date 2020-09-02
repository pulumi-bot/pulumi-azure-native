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
    public sealed class ServiceTypeDeltaHealthPolicyResponseResult
    {
        /// <summary>
        /// The maximum allowed percentage of services health degradation allowed during cluster upgrades.
        /// The delta is measured between the state of the services at the beginning of upgrade and the state of the services at the time of the health evaluation.
        /// The check is performed after every upgrade domain upgrade completion to make sure the global state of the cluster is within tolerated limits.
        /// </summary>
        public readonly int? MaxPercentDeltaUnhealthyServices;

        [OutputConstructor]
        private ServiceTypeDeltaHealthPolicyResponseResult(int? maxPercentDeltaUnhealthyServices)
        {
            MaxPercentDeltaUnhealthyServices = maxPercentDeltaUnhealthyServices;
        }
    }
}
