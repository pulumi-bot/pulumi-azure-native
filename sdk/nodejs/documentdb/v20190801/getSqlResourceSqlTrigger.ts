// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getSqlResourceSqlTrigger(args: GetSqlResourceSqlTriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetSqlResourceSqlTriggerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:documentdb/v20190801:getSqlResourceSqlTrigger", {
        "accountName": args.accountName,
        "containerName": args.containerName,
        "databaseName": args.databaseName,
        "resourceGroupName": args.resourceGroupName,
        "triggerName": args.triggerName,
    }, opts);
}

export interface GetSqlResourceSqlTriggerArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: string;
    /**
     * Cosmos DB container name.
     */
    readonly containerName: string;
    /**
     * Cosmos DB database name.
     */
    readonly databaseName: string;
    /**
     * Name of an Azure resource group.
     */
    readonly resourceGroupName: string;
    /**
     * Cosmos DB trigger name.
     */
    readonly triggerName: string;
}

/**
 * An Azure Cosmos DB trigger.
 */
export interface GetSqlResourceSqlTriggerResult {
    /**
     * The location of the resource group to which the resource belongs.
     */
    readonly location?: string;
    /**
     * The name of the ARM resource.
     */
    readonly name: string;
    readonly resource?: outputs.documentdb.v20190801.SqlTriggerGetPropertiesResponseResource;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of Azure resource.
     */
    readonly type: string;
}
