// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An Azure Cosmos DB database account.
 */
export class DatabaseAccount extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabaseAccount {
        return new DatabaseAccount(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:documentdb/v20200401:DatabaseAccount';

    /**
     * Returns true if the given object is an instance of DatabaseAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseAccount.__pulumiType;
    }

    /**
     * Indicates the type of database account. This can only be set at database account creation.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the ARM resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties for the database account.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.documentdb.v20200401.DatabaseAccountGetPropertiesResponse>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    public readonly tags!: pulumi.Output<outputs.documentdb.v20200401.TagsResponse | undefined>;
    /**
     * The type of Azure resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DatabaseAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseAccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DatabaseAccountArgs | undefined;
            if (!args || args.databaseAccountOfferType === undefined) {
                throw new Error("Missing required property 'databaseAccountOfferType'");
            }
            if (!args || args.locations === undefined) {
                throw new Error("Missing required property 'locations'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiProperties"] = args ? args.apiProperties : undefined;
            inputs["capabilities"] = args ? args.capabilities : undefined;
            inputs["connectorOffer"] = args ? args.connectorOffer : undefined;
            inputs["consistencyPolicy"] = args ? args.consistencyPolicy : undefined;
            inputs["databaseAccountOfferType"] = args ? args.databaseAccountOfferType : undefined;
            inputs["disableKeyBasedMetadataWriteAccess"] = args ? args.disableKeyBasedMetadataWriteAccess : undefined;
            inputs["enableAnalyticalStorage"] = args ? args.enableAnalyticalStorage : undefined;
            inputs["enableAutomaticFailover"] = args ? args.enableAutomaticFailover : undefined;
            inputs["enableCassandraConnector"] = args ? args.enableCassandraConnector : undefined;
            inputs["enableFreeTier"] = args ? args.enableFreeTier : undefined;
            inputs["enableMultipleWriteLocations"] = args ? args.enableMultipleWriteLocations : undefined;
            inputs["ipRules"] = args ? args.ipRules : undefined;
            inputs["isVirtualNetworkFilterEnabled"] = args ? args.isVirtualNetworkFilterEnabled : undefined;
            inputs["keyVaultKeyUri"] = args ? args.keyVaultKeyUri : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["locations"] = args ? args.locations : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publicNetworkAccess"] = args ? args.publicNetworkAccess : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkRules"] = args ? args.virtualNetworkRules : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatabaseAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabaseAccount resource.
 */
export interface DatabaseAccountArgs {
    /**
     * API specific properties. Currently, supported only for MongoDB API.
     */
    readonly apiProperties?: pulumi.Input<inputs.documentdb.v20200401.ApiProperties>;
    /**
     * List of Cosmos DB capabilities for the account
     */
    readonly capabilities?: pulumi.Input<pulumi.Input<inputs.documentdb.v20200401.Capability>[]>;
    /**
     * The cassandra connector offer type for the Cosmos DB database C* account.
     */
    readonly connectorOffer?: pulumi.Input<string>;
    /**
     * The consistency policy for the Cosmos DB account.
     */
    readonly consistencyPolicy?: pulumi.Input<inputs.documentdb.v20200401.ConsistencyPolicy>;
    /**
     * The offer type for the database
     */
    readonly databaseAccountOfferType: pulumi.Input<string>;
    /**
     * Disable write operations on metadata resources (databases, containers, throughput) via account keys
     */
    readonly disableKeyBasedMetadataWriteAccess?: pulumi.Input<boolean>;
    /**
     * Flag to indicate whether to enable storage analytics.
     */
    readonly enableAnalyticalStorage?: pulumi.Input<boolean>;
    /**
     * Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
     */
    readonly enableAutomaticFailover?: pulumi.Input<boolean>;
    /**
     * Enables the cassandra connector on the Cosmos DB C* account
     */
    readonly enableCassandraConnector?: pulumi.Input<boolean>;
    /**
     * Flag to indicate whether Free Tier is enabled.
     */
    readonly enableFreeTier?: pulumi.Input<boolean>;
    /**
     * Enables the account to write in multiple locations
     */
    readonly enableMultipleWriteLocations?: pulumi.Input<boolean>;
    /**
     * List of IpRules.
     */
    readonly ipRules?: pulumi.Input<inputs.documentdb.v20200401.IPRules>;
    /**
     * Flag to indicate whether to enable/disable Virtual Network ACL rules.
     */
    readonly isVirtualNetworkFilterEnabled?: pulumi.Input<boolean>;
    /**
     * The URI of the key vault
     */
    readonly keyVaultKeyUri?: pulumi.Input<string>;
    /**
     * Indicates the type of database account. This can only be set at database account creation.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * An array that contains the georeplication locations enabled for the Cosmos DB account.
     */
    readonly locations: pulumi.Input<pulumi.Input<inputs.documentdb.v20200401.Location>[]>;
    /**
     * Cosmos DB database account name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Whether requests from Public Network are allowed
     */
    readonly publicNetworkAccess?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    readonly tags?: pulumi.Input<inputs.documentdb.v20200401.Tags>;
    /**
     * List of Virtual Network ACL rules configured for the Cosmos DB account.
     */
    readonly virtualNetworkRules?: pulumi.Input<pulumi.Input<inputs.documentdb.v20200401.VirtualNetworkRule>[]>;
}
