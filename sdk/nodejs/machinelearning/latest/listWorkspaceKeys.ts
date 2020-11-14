// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listWorkspaceKeys(args: ListWorkspaceKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListWorkspaceKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:machinelearning/latest:listWorkspaceKeys", {
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListWorkspaceKeysArgs {
    /**
     * The name of the resource group to which the machine learning workspace belongs.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the machine learning workspace.
     */
    readonly workspaceName: string;
}

/**
 * Workspace authorization keys for a workspace.
 */
export interface ListWorkspaceKeysResult {
    /**
     * Primary authorization key for this workspace.
     */
    readonly primaryToken?: string;
    /**
     * Secondary authorization key for this workspace.
     */
    readonly secondaryToken?: string;
}
