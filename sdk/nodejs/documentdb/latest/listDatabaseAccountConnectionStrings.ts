// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listDatabaseAccountConnectionStrings(args: ListDatabaseAccountConnectionStringsArgs, opts?: pulumi.InvokeOptions): Promise<ListDatabaseAccountConnectionStringsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:documentdb/latest:listDatabaseAccountConnectionStrings", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListDatabaseAccountConnectionStringsArgs {
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
 * The connection strings for the given database account.
 */
export interface ListDatabaseAccountConnectionStringsResult {
    /**
     * An array that contains the connection strings for the Cosmos DB account.
     */
    readonly connectionStrings?: outputs.documentdb.latest.DatabaseAccountConnectionStringResponse[];
}
