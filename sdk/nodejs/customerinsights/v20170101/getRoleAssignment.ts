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
    return pulumi.runtime.invoke("azurerm:customerinsights/v20170101:getRoleAssignment", {
        "assignmentName": args.assignmentName,
        "hubName": args.hubName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRoleAssignmentArgs {
    /**
     * The name of the role assignment.
     */
    readonly assignmentName: string;
    /**
     * The name of the hub.
     */
    readonly hubName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The Role Assignment resource format.
 */
export interface GetRoleAssignmentResult {
    /**
     * The name of the metadata object.
     */
    readonly assignmentName: string;
    /**
     * Widget types set for the assignment.
     */
    readonly conflationPolicies?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Connectors set for the assignment.
     */
    readonly connectors?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Localized description for the metadata.
     */
    readonly description?: {[key: string]: string};
    /**
     * Localized display names for the metadata.
     */
    readonly displayName?: {[key: string]: string};
    /**
     * Interactions set for the assignment.
     */
    readonly interactions?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Kpis set for the assignment.
     */
    readonly kpis?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Links set for the assignment.
     */
    readonly links?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The principals being assigned to.
     */
    readonly principals: outputs.customerinsights.v20170101.AssignmentPrincipalResponse[];
    /**
     * Profiles set for the assignment.
     */
    readonly profiles?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Provisioning state.
     */
    readonly provisioningState: string;
    /**
     * The Role assignments set for the relationship links.
     */
    readonly relationshipLinks?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * The Role assignments set for the relationships.
     */
    readonly relationships?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Type of roles.
     */
    readonly role: RoleTypes;
    /**
     * The Role assignments set for the assignment.
     */
    readonly roleAssignments?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Sas Policies set for the assignment.
     */
    readonly sasPolicies?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * The Role assignments set for the assignment.
     */
    readonly segments?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * The hub name.
     */
    readonly tenantId: string;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * Views set for the assignment.
     */
    readonly views?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
    /**
     * Widget types set for the assignment.
     */
    readonly widgetTypes?: outputs.customerinsights.v20170101.ResourceSetDescriptionResponse;
}
