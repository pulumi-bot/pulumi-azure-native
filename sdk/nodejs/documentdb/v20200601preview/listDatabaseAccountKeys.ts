// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listDatabaseAccountKeys(args: ListDatabaseAccountKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListDatabaseAccountKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:documentdb/v20200601preview:listDatabaseAccountKeys", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListDatabaseAccountKeysArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * The access keys for the given database account.
 */
export interface ListDatabaseAccountKeysResult {
    /**
     * Base 64 encoded value of the primary read-write key.
     */
    readonly primaryMasterKey: string;
    /**
     * Base 64 encoded value of the primary read-only key.
     */
    readonly primaryReadonlyMasterKey: string;
    /**
     * Base 64 encoded value of the secondary read-write key.
     */
    readonly secondaryMasterKey: string;
    /**
     * Base 64 encoded value of the secondary read-only key.
     */
    readonly secondaryReadonlyMasterKey: string;
}
