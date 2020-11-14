// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getReplicationProtectedItem(args: GetReplicationProtectedItemArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationProtectedItemResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:recoveryservices/v20160810:getReplicationProtectedItem", {
        "fabricName": args.fabricName,
        "protectionContainerName": args.protectionContainerName,
        "replicatedProtectedItemName": args.replicatedProtectedItemName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetReplicationProtectedItemArgs {
    /**
     * Fabric unique name.
     */
    readonly fabricName: string;
    /**
     * Protection container name.
     */
    readonly protectionContainerName: string;
    /**
     * Replication protected item name.
     */
    readonly replicatedProtectedItemName: string;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the recovery services vault.
     */
    readonly resourceName: string;
}

/**
 * Replication protected item.
 */
export interface GetReplicationProtectedItemResult {
    /**
     * Resource Location
     */
    readonly location?: string;
    /**
     * Resource Name
     */
    readonly name: string;
    /**
     * The custom data.
     */
    readonly properties: outputs.recoveryservices.v20160810.ReplicationProtectedItemPropertiesResponse;
    /**
     * Resource Type
     */
    readonly type: string;
}
