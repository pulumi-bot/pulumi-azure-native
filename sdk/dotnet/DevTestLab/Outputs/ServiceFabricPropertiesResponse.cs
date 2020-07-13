// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.Outputs
{

    [OutputType]
    public sealed class ServiceFabricPropertiesResponse
    {
        /// <summary>
        /// The applicable schedule for the virtual machine.
        /// </summary>
        public readonly Outputs.ApplicableScheduleResponse ApplicableSchedule;
        /// <summary>
        /// The resource id of the environment under which the service fabric resource is present
        /// </summary>
        public readonly string? EnvironmentId;
        /// <summary>
        /// The backing service fabric resource's id
        /// </summary>
        public readonly string? ExternalServiceFabricId;
        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        public readonly string UniqueIdentifier;

        [OutputConstructor]
        private ServiceFabricPropertiesResponse(
            Outputs.ApplicableScheduleResponse applicableSchedule,

            string? environmentId,

            string? externalServiceFabricId,

            string provisioningState,

            string uniqueIdentifier)
        {
            ApplicableSchedule = applicableSchedule;
            EnvironmentId = environmentId;
            ExternalServiceFabricId = externalServiceFabricId;
            ProvisioningState = provisioningState;
            UniqueIdentifier = uniqueIdentifier;
        }
    }
}