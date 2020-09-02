// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getRegistrationAssignment(args: GetRegistrationAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistrationAssignmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:managedservices/v20180601preview:getRegistrationAssignment", {
        "expandRegistrationDefinition": args.expandRegistrationDefinition,
        "registrationAssignmentId": args.registrationAssignmentId,
        "scope": args.scope,
    }, opts);
}

export interface GetRegistrationAssignmentArgs {
    /**
     * Tells whether to return registration definition details also along with registration assignment details.
     */
    readonly expandRegistrationDefinition?: boolean;
    /**
     * Guid of the registration assignment.
     */
    readonly registrationAssignmentId: string;
    /**
     * Scope of the resource.
     */
    readonly scope: string;
}

/**
 * Registration assignment.
 */
export interface GetRegistrationAssignmentResult {
    /**
     * Name of the registration assignment.
     */
    readonly name: string;
    /**
     * Properties of a registration assignment.
     */
    readonly properties: outputs.managedservices.v20180601preview.RegistrationAssignmentPropertiesResponse;
    /**
     * Type of the resource.
     */
    readonly type: string;
}
