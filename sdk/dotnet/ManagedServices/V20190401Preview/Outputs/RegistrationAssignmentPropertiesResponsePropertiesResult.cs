// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ManagedServices.V20190401Preview.Outputs
{

    [OutputType]
    public sealed class RegistrationAssignmentPropertiesResponsePropertiesResult
    {
        /// <summary>
        /// Authorization tuple containing principal id of the user/security group or service principal and id of the build-in role.
        /// </summary>
        public readonly ImmutableArray<Outputs.AuthorizationResponseResult> Authorizations;
        /// <summary>
        /// Description of the registration definition.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Id of the managedBy tenant.
        /// </summary>
        public readonly string? ManagedByTenantId;
        /// <summary>
        /// Name of the managedBy tenant.
        /// </summary>
        public readonly string? ManagedByTenantName;
        /// <summary>
        /// Id of the home tenant.
        /// </summary>
        public readonly string? ManageeTenantId;
        /// <summary>
        /// Name of the home tenant.
        /// </summary>
        public readonly string? ManageeTenantName;
        /// <summary>
        /// Current state of the registration definition.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Name of the registration definition.
        /// </summary>
        public readonly string? RegistrationDefinitionName;

        [OutputConstructor]
        private RegistrationAssignmentPropertiesResponsePropertiesResult(
            ImmutableArray<Outputs.AuthorizationResponseResult> authorizations,

            string? description,

            string? managedByTenantId,

            string? managedByTenantName,

            string? manageeTenantId,

            string? manageeTenantName,

            string? provisioningState,

            string? registrationDefinitionName)
        {
            Authorizations = authorizations;
            Description = description;
            ManagedByTenantId = managedByTenantId;
            ManagedByTenantName = managedByTenantName;
            ManageeTenantId = manageeTenantId;
            ManageeTenantName = manageeTenantName;
            ProvisioningState = provisioningState;
            RegistrationDefinitionName = registrationDefinitionName;
        }
    }
}
