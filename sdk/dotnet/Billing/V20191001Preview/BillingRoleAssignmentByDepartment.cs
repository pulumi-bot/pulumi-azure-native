// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Billing.V20191001Preview
{
    /// <summary>
    /// The role assignment
    /// </summary>
    public partial class BillingRoleAssignmentByDepartment : Pulumi.CustomResource
    {
        /// <summary>
        /// The principal Id of the user who created the role assignment.
        /// </summary>
        [Output("createdByPrincipalId")]
        public Output<string> CreatedByPrincipalId { get; private set; } = null!;

        /// <summary>
        /// The tenant Id of the user who created the role assignment.
        /// </summary>
        [Output("createdByPrincipalTenantId")]
        public Output<string> CreatedByPrincipalTenantId { get; private set; } = null!;

        /// <summary>
        /// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
        /// </summary>
        [Output("createdByUserEmailAddress")]
        public Output<string> CreatedByUserEmailAddress { get; private set; } = null!;

        /// <summary>
        /// The date the role assignment was created.
        /// </summary>
        [Output("createdOn")]
        public Output<string> CreatedOn { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The principal id of the user to whom the role was assigned.
        /// </summary>
        [Output("principalId")]
        public Output<string?> PrincipalId { get; private set; } = null!;

        /// <summary>
        /// The principal tenant id of the user to whom the role was assigned.
        /// </summary>
        [Output("principalTenantId")]
        public Output<string?> PrincipalTenantId { get; private set; } = null!;

        /// <summary>
        /// The ID of the role definition.
        /// </summary>
        [Output("roleDefinitionId")]
        public Output<string?> RoleDefinitionId { get; private set; } = null!;

        /// <summary>
        /// The scope at which the role was assigned.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        /// </summary>
        [Output("userAuthenticationType")]
        public Output<string?> UserAuthenticationType { get; private set; } = null!;

        /// <summary>
        /// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        /// </summary>
        [Output("userEmailAddress")]
        public Output<string?> UserEmailAddress { get; private set; } = null!;


        /// <summary>
        /// Create a BillingRoleAssignmentByDepartment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BillingRoleAssignmentByDepartment(string name, BillingRoleAssignmentByDepartmentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:billing/v20191001preview:BillingRoleAssignmentByDepartment", name, args ?? new BillingRoleAssignmentByDepartmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BillingRoleAssignmentByDepartment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:billing/v20191001preview:BillingRoleAssignmentByDepartment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BillingRoleAssignmentByDepartment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BillingRoleAssignmentByDepartment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BillingRoleAssignmentByDepartment(name, id, options);
        }
    }

    public sealed class BillingRoleAssignmentByDepartmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID that uniquely identifies a billing account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public Input<string> BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies a role assignment.
        /// </summary>
        [Input("billingRoleAssignmentName", required: true)]
        public Input<string> BillingRoleAssignmentName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies a department.
        /// </summary>
        [Input("departmentName", required: true)]
        public Input<string> DepartmentName { get; set; } = null!;

        /// <summary>
        /// The principal id of the user to whom the role was assigned.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The principal tenant id of the user to whom the role was assigned.
        /// </summary>
        [Input("principalTenantId")]
        public Input<string>? PrincipalTenantId { get; set; }

        /// <summary>
        /// The ID of the role definition.
        /// </summary>
        [Input("roleDefinitionId")]
        public Input<string>? RoleDefinitionId { get; set; }

        /// <summary>
        /// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        /// </summary>
        [Input("userAuthenticationType")]
        public Input<string>? UserAuthenticationType { get; set; }

        /// <summary>
        /// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
        /// </summary>
        [Input("userEmailAddress")]
        public Input<string>? UserEmailAddress { get; set; }

        public BillingRoleAssignmentByDepartmentArgs()
        {
        }
    }
}
