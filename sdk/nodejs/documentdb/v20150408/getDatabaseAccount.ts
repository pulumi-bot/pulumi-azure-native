// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getDatabaseAccount(args: GetDatabaseAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:documentdb/v20150408:getDatabaseAccount", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDatabaseAccountArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: string;
    /**
     * Name of an Azure resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * An Azure Cosmos DB database account.
 */
export interface GetDatabaseAccountResult {
    /**
     * List of Cosmos DB capabilities for the account
     */
    readonly capabilities?: outputs.documentdb.v20150408.CapabilityResponse[];
    /**
     * The cassandra connector offer type for the Cosmos DB database C* account.
     */
    readonly connectorOffer?: string;
    /**
     * The consistency policy for the Cosmos DB database account.
     */
    readonly consistencyPolicy?: outputs.documentdb.v20150408.ConsistencyPolicyResponse;
    /**
     * The offer type for the Cosmos DB database account. Default value: Standard.
     */
    readonly databaseAccountOfferType: string;
    /**
     * The connection endpoint for the Cosmos DB database account.
     */
    readonly documentEndpoint: string;
    /**
     * Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
     */
    readonly enableAutomaticFailover?: boolean;
    /**
     * Enables the cassandra connector on the Cosmos DB C* account
     */
    readonly enableCassandraConnector?: boolean;
    /**
     * Enables the account to write in multiple locations
     */
    readonly enableMultipleWriteLocations?: boolean;
    /**
     * An array that contains the regions ordered by their failover priorities.
     */
    readonly failoverPolicies: outputs.documentdb.v20150408.FailoverPolicyResponse[];
    /**
     * Cosmos DB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IPs for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
     */
    readonly ipRangeFilter?: string;
    /**
     * Flag to indicate whether to enable/disable Virtual Network ACL rules.
     */
    readonly isVirtualNetworkFilterEnabled?: boolean;
    /**
     * Indicates the type of database account. This can only be set at database account creation.
     */
    readonly kind?: string;
    /**
     * The location of the resource group to which the resource belongs.
     */
    readonly location?: string;
    /**
     * The name of the database account.
     */
    readonly name: string;
    /**
     * The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'Offline' - the Cosmos DB account is not active. 'DeletionFailed' – the Cosmos DB account deletion failed.
     */
    readonly provisioningState: string;
    /**
     * An array that contains of the read locations enabled for the Cosmos DB account.
     */
    readonly readLocations: outputs.documentdb.v20150408.LocationResponse[];
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of Azure resource.
     */
    readonly type: string;
    /**
     * List of Virtual Network ACL rules configured for the Cosmos DB account.
     */
    readonly virtualNetworkRules?: outputs.documentdb.v20150408.VirtualNetworkRuleResponse[];
    /**
     * An array that contains the write location for the Cosmos DB account.
     */
    readonly writeLocations: outputs.documentdb.v20150408.LocationResponse[];
}
