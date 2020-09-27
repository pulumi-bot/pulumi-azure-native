// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDatabaseAccountSqlContainer(args: GetDatabaseAccountSqlContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseAccountSqlContainerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:documentdb/v20160319:getDatabaseAccountSqlContainer", {
        "accountName": args.accountName,
        "containerName": args.containerName,
        "databaseName": args.databaseName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDatabaseAccountSqlContainerArgs {
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
}

/**
 * An Azure Cosmos DB container.
 */
export interface GetDatabaseAccountSqlContainerResult {
    /**
     * The conflict resolution policy for the container.
     */
    readonly conflictResolutionPolicy?: outputs.documentdb.v20160319.ConflictResolutionPolicyResponse;
    /**
     * Default time to live
     */
    readonly defaultTtl?: number;
    /**
     * A system generated property representing the resource etag required for optimistic concurrency control.
     */
    readonly etag?: string;
    /**
     * The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the container
     */
    readonly indexingPolicy?: outputs.documentdb.v20160319.IndexingPolicyResponse;
    /**
     * The location of the resource group to which the resource belongs.
     */
    readonly location?: string;
    /**
     * The name of the database account.
     */
    readonly name: string;
    /**
     * The configuration of the partition key to be used for partitioning data into multiple partitions
     */
    readonly partitionKey?: outputs.documentdb.v20160319.ContainerPartitionKeyResponse;
    /**
     * A system generated property. A unique identifier.
     */
    readonly rid?: string;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of Azure resource.
     */
    readonly type: string;
    /**
     * The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure Cosmos DB service.
     */
    readonly uniqueKeyPolicy?: outputs.documentdb.v20160319.UniqueKeyPolicyResponse;
}
