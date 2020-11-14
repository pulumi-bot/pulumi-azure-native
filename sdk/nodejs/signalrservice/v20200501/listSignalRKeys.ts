// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listSignalRKeys(args: ListSignalRKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListSignalRKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:signalrservice/v20200501:listSignalRKeys", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface ListSignalRKeysArgs {
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the SignalR resource.
     */
    readonly resourceName: string;
}

/**
 * A class represents the access keys of SignalR service.
 */
export interface ListSignalRKeysResult {
    /**
     * SignalR connection string constructed via the primaryKey
     */
    readonly primaryConnectionString?: string;
    /**
     * The primary access key.
     */
    readonly primaryKey?: string;
    /**
     * SignalR connection string constructed via the secondaryKey
     */
    readonly secondaryConnectionString?: string;
    /**
     * The secondary access key.
     */
    readonly secondaryKey?: string;
}
