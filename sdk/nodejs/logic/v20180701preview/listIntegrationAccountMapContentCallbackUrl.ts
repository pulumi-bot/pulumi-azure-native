// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listIntegrationAccountMapContentCallbackUrl(args: ListIntegrationAccountMapContentCallbackUrlArgs, opts?: pulumi.InvokeOptions): Promise<ListIntegrationAccountMapContentCallbackUrlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:logic/v20180701preview:listIntegrationAccountMapContentCallbackUrl", {
        "integrationAccountName": args.integrationAccountName,
        "keyType": args.keyType,
        "mapName": args.mapName,
        "notAfter": args.notAfter,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListIntegrationAccountMapContentCallbackUrlArgs {
    /**
     * The integration account name.
     */
    readonly integrationAccountName: string;
    /**
     * The key type.
     */
    readonly keyType?: string;
    /**
     * The integration account map name.
     */
    readonly mapName: string;
    /**
     * The expiry time.
     */
    readonly notAfter?: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * The workflow trigger callback URL.
 */
export interface ListIntegrationAccountMapContentCallbackUrlResult {
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
    readonly queries?: outputs.logic.v20180701preview.WorkflowTriggerListCallbackUrlQueriesResponse;
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
