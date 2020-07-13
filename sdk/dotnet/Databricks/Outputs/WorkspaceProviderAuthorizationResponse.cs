// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Databricks.Outputs
{

    [OutputType]
    public sealed class WorkspaceProviderAuthorizationResponse
    {
        /// <summary>
        /// The provider's principal identifier. This is the identity that the provider will use to call ARM to manage the workspace resources.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The provider's role definition identifier. This role will define all the permissions that the provider must have on the workspace's container resource group. This role definition cannot have permission to delete the resource group.
        /// </summary>
        public readonly string RoleDefinitionId;

        [OutputConstructor]
        private WorkspaceProviderAuthorizationResponse(
            string principalId,

            string roleDefinitionId)
        {
            PrincipalId = principalId;
            RoleDefinitionId = roleDefinitionId;
        }
    }
}