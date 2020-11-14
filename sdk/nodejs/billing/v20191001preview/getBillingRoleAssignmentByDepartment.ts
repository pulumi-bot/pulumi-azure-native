// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getBillingRoleAssignmentByDepartment(args: GetBillingRoleAssignmentByDepartmentArgs, opts?: pulumi.InvokeOptions): Promise<GetBillingRoleAssignmentByDepartmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:billing/v20191001preview:getBillingRoleAssignmentByDepartment", {
        "billingAccountName": args.billingAccountName,
        "billingRoleAssignmentName": args.billingRoleAssignmentName,
        "departmentName": args.departmentName,
    }, opts);
}

export interface GetBillingRoleAssignmentByDepartmentArgs {
    /**
     * The ID that uniquely identifies a billing account.
     */
    readonly billingAccountName: string;
    /**
     * The ID that uniquely identifies a role assignment.
     */
    readonly billingRoleAssignmentName: string;
    /**
     * The ID that uniquely identifies a department.
     */
    readonly departmentName: string;
}

/**
 * The role assignment
 */
export interface GetBillingRoleAssignmentByDepartmentResult {
    /**
     * The principal Id of the user who created the role assignment.
     */
    readonly createdByPrincipalId: string;
    /**
     * The tenant Id of the user who created the role assignment.
     */
    readonly createdByPrincipalTenantId: string;
    /**
     * The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
     */
    readonly createdByUserEmailAddress: string;
    /**
     * The date the role assignment was created.
     */
    readonly createdOn: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The principal id of the user to whom the role was assigned.
     */
    readonly principalId?: string;
    /**
     * The principal tenant id of the user to whom the role was assigned.
     */
    readonly principalTenantId?: string;
    /**
     * The ID of the role definition.
     */
    readonly roleDefinitionId?: string;
    /**
     * The scope at which the role was assigned.
     */
    readonly scope: string;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
     */
    readonly userAuthenticationType?: string;
    /**
     * The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
     */
    readonly userEmailAddress?: string;
}
