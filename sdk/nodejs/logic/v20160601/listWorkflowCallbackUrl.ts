// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listWorkflowCallbackUrl(args: ListWorkflowCallbackUrlArgs, opts?: pulumi.InvokeOptions): Promise<ListWorkflowCallbackUrlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:logic/v20160601:listWorkflowCallbackUrl", {
        "keyType": args.keyType,
        "notAfter": args.notAfter,
        "resourceGroupName": args.resourceGroupName,
        "workflowName": args.workflowName,
    }, opts);
}

export interface ListWorkflowCallbackUrlArgs {
    /**
     * The key type.
     */
    readonly keyType?: string;
    /**
     * The expiry time.
     */
    readonly notAfter?: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The workflow name.
     */
    readonly workflowName: string;
}

/**
 * The workflow trigger callback URL.
 */
export interface ListWorkflowCallbackUrlResult {
    /**
     * Gets the workflow trigger callback URL base path.
     */
    readonly basePath: string;
    /**
     * Gets the workflow trigger callback URL HTTP method.
     */
    readonly method: string;
    /**
     * Gets the workflow trigger callback URL query parameters.
     */
    readonly queries?: outputs.logic.v20160601.WorkflowTriggerListCallbackUrlQueriesResponse;
    /**
     * Gets the workflow trigger callback URL relative path.
     */
    readonly relativePath: string;
    /**
     * Gets the workflow trigger callback URL relative path parameters.
     */
    readonly relativePathParameters?: string[];
    /**
     * Gets the workflow trigger callback URL.
     */
    readonly value: string;
}
