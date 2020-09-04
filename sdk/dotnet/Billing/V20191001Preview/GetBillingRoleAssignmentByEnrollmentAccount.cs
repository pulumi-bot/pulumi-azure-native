// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Billing.V20191001Preview
{
    public static class GetBillingRoleAssignmentByEnrollmentAccount
    {
        public static Task<GetBillingRoleAssignmentByEnrollmentAccountResult> InvokeAsync(GetBillingRoleAssignmentByEnrollmentAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBillingRoleAssignmentByEnrollmentAccountResult>("azurerm:billing/v20191001preview:getBillingRoleAssignmentByEnrollmentAccount", args ?? new GetBillingRoleAssignmentByEnrollmentAccountArgs(), options.WithVersion());
    }


    public sealed class GetBillingRoleAssignmentByEnrollmentAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID that uniquely identifies a billing account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public string BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies a role assignment.
        /// </summary>
        [Input("billingRoleAssignmentName", required: true)]
        public string BillingRoleAssignmentName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies an enrollment account.
        /// </summary>
        [Input("enrollmentAccountName", required: true)]
        public string EnrollmentAccountName { get; set; } = null!;

        public GetBillingRoleAssignmentByEnrollmentAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBillingRoleAssignmentByEnrollmentAccountResult
    {
        /// <summary>
        /// The principal Id of the user who created the role assignment.
        /// </summary>
        public readonly string CreatedByPrincipalId;
        /// <summary>
        /// The tenant Id of the user who created the role assignment.
        /// </summary>
        public readonly string CreatedByPrincipalTenantId;
        /// <summary>
        /// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
        /// </summary>
        public readonly string CreatedByUserEmailAddress;
        /// <summary>
        /// The date the role assignment was created.
        /// </summary>
        public readonly string CreatedOn;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The principal id of the user to whom the role was assigned.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// The principal tenant id of the user to whom the role was assigned.
        /// </summary>
        public readonly string? PrincipalTenantId;
        /// <summary>
        /// The ID of the role definition.
        /// </summary>
        public readonly string? RoleDefinitionId;
        /// <summary>
        /// The scope at which the role was assigned.
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        /// </summary>
        public readonly string? UserAuthenticationType;
        /// <summary>
        /// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        /// </summary>
        public readonly string? UserEmailAddress;

        [OutputConstructor]
        private GetBillingRoleAssignmentByEnrollmentAccountResult(
            string createdByPrincipalId,

            string createdByPrincipalTenantId,

            string createdByUserEmailAddress,

            string createdOn,

            string name,

            string? principalId,

            string? principalTenantId,

            string? roleDefinitionId,

            string scope,

            string type,

            string? userAuthenticationType,

            string? userEmailAddress)
        {
            CreatedByPrincipalId = createdByPrincipalId;
            CreatedByPrincipalTenantId = createdByPrincipalTenantId;
            CreatedByUserEmailAddress = createdByUserEmailAddress;
            CreatedOn = createdOn;
            Name = name;
            PrincipalId = principalId;
            PrincipalTenantId = principalTenantId;
            RoleDefinitionId = roleDefinitionId;
            Scope = scope;
            Type = type;
            UserAuthenticationType = userAuthenticationType;
            UserEmailAddress = userEmailAddress;
        }
    }
}
