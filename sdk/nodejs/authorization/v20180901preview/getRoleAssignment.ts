// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getRoleAssignment(args: GetRoleAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleAssignmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:authorization/v20180901preview:getRoleAssignment", {
        "roleAssignmentName": args.roleAssignmentName,
        "scope": args.scope,
    }, opts);
}

export interface GetRoleAssignmentArgs {
    /**
     * The name of the role assignment to get.
     */
    readonly roleAssignmentName: string;
    /**
     * The scope of the role assignment.
     */
    readonly scope: string;
}

/**
 * Role Assignments
 */
export interface GetRoleAssignmentResult {
    /**
     * The Delegation flag for the role assignment
     */
    readonly canDelegate?: boolean;
    /**
     * The role assignment name.
     */
    readonly name: string;
    /**
     * The principal ID.
     */
    readonly principalId?: string;
    /**
     * The principal type of the assigned principal ID.
     */
    readonly principalType?: string;
    /**
     * The role definition ID.
     */
    readonly roleDefinitionId?: string;
    /**
     * The role assignment scope.
     */
    readonly scope?: string;
    /**
     * The role assignment type.
     */
    readonly type: string;
}
