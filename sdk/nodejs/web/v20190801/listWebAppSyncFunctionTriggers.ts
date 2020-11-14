// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listWebAppSyncFunctionTriggers(args: ListWebAppSyncFunctionTriggersArgs, opts?: pulumi.InvokeOptions): Promise<ListWebAppSyncFunctionTriggersResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:web/v20190801:listWebAppSyncFunctionTriggers", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListWebAppSyncFunctionTriggersArgs {
    /**
     * Name of the app.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Function secrets.
 */
export interface ListWebAppSyncFunctionTriggersResult {
    /**
     * Secret key.
     */
    readonly key?: string;
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * Trigger URL.
     */
    readonly triggerUrl?: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
