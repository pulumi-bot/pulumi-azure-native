// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDatabaseAccountSqlDatabaseContainer(args: GetDatabaseAccountSqlDatabaseContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseAccountSqlDatabaseContainerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:documentdb:getDatabaseAccountSqlDatabaseContainer", {
        "accountName": args.accountName,
        "databaseName": args.databaseName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDatabaseAccountSqlDatabaseContainerArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: string;
    /**
     * Cosmos DB database name.
     */
    readonly databaseName: string;
    /**
     * Cosmos DB container name.
     */
    readonly name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * An Azure Cosmos DB container.
 */
export interface GetDatabaseAccountSqlDatabaseContainerResult {
    /**
     * The location of the resource group to which the resource belongs.
     */
    readonly location?: string;
    /**
     * The name of the ARM resource.
     */
    readonly name: string;
    /**
     * The properties of an Azure Cosmos DB container
     */
    readonly properties: outputs.documentdb.SqlContainerGetPropertiesResponse;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    readonly tags?: outputs.documentdb.TagsResponse;
    /**
     * The type of Azure resource.
     */
    readonly type: string;
}