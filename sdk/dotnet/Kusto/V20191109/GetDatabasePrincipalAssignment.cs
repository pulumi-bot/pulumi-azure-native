// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Kusto.V20191109
{
    public static class GetDatabasePrincipalAssignment
    {
        public static Task<GetDatabasePrincipalAssignmentResult> InvokeAsync(GetDatabasePrincipalAssignmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabasePrincipalAssignmentResult>("azure-nextgen:kusto/v20191109:getDatabasePrincipalAssignment", args ?? new GetDatabasePrincipalAssignmentArgs(), options.WithVersion());
    }


    public sealed class GetDatabasePrincipalAssignmentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the database in the Kusto cluster.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the Kusto principalAssignment.
        /// </summary>
        [Input("principalAssignmentName", required: true)]
        public string PrincipalAssignmentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDatabasePrincipalAssignmentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabasePrincipalAssignmentResult
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The principal name
        /// </summary>
        public readonly string PrincipalName;
        /// <summary>
        /// Principal type.
        /// </summary>
        public readonly string PrincipalType;
        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Database principal role.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// The tenant id of the principal
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The tenant name of the principal
        /// </summary>
        public readonly string TenantName;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDatabasePrincipalAssignmentResult(
            string name,

            string principalId,

            string principalName,

            string principalType,

            string provisioningState,

            string role,

            string? tenantId,

            string tenantName,

            string type)
        {
            Name = name;
            PrincipalId = principalId;
            PrincipalName = principalName;
            PrincipalType = principalType;
            ProvisioningState = provisioningState;
            Role = role;
            TenantId = tenantId;
            TenantName = tenantName;
            Type = type;
        }
    }
}
