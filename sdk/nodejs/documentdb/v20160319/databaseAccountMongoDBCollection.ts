// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An Azure Cosmos DB MongoDB collection.
 *
 * ## Example Usage
 * ### CosmosDBMongoDBCollectionCreateUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const databaseAccountMongoDBCollection = new azurerm.documentdb.v20160319.DatabaseAccountMongoDBCollection("databaseAccountMongoDBCollection", {
 *     accountName: "ddb1",
 *     collectionName: "collectionName",
 *     databaseName: "databaseName",
 *     options: {},
 *     resource: {
 *         id: "testcoll",
 *         indexes: [{
 *             key: {
 *                 keys: ["testKey"],
 *             },
 *             options: {
 *                 expireAfterSeconds: 100,
 *                 unique: true,
 *             },
 *         }],
 *         shardKey: {
 *             testKey: "Hash",
 *         },
 *     },
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 */
export class DatabaseAccountMongoDBCollection extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseAccountMongoDBCollection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabaseAccountMongoDBCollection {
        return new DatabaseAccountMongoDBCollection(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:documentdb/v20160319:DatabaseAccountMongoDBCollection';

    /**
     * Returns true if the given object is an instance of DatabaseAccountMongoDBCollection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseAccountMongoDBCollection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseAccountMongoDBCollection.__pulumiType;
    }

    /**
     * List of index keys
     */
    public /*out*/ readonly indexes!: pulumi.Output<outputs.documentdb.v20160319.MongoIndexResponse[] | undefined>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the database account.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A key-value pair of shard keys to be applied for the request.
     */
    public /*out*/ readonly shardKey!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of Azure resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DatabaseAccountMongoDBCollection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseAccountMongoDBCollectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseAccountMongoDBCollectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DatabaseAccountMongoDBCollectionArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.collectionName === undefined) {
                throw new Error("Missing required property 'collectionName'");
            }
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.options === undefined) {
                throw new Error("Missing required property 'options'");
            }
            if (!args || args.resource === undefined) {
                throw new Error("Missing required property 'resource'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["collectionName"] = args ? args.collectionName : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["resource"] = args ? args.resource : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["indexes"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["shardKey"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:documentdb/latest:DatabaseAccountMongoDBCollection" }, { type: "azurerm:documentdb/v20150401:DatabaseAccountMongoDBCollection" }, { type: "azurerm:documentdb/v20150408:DatabaseAccountMongoDBCollection" }, { type: "azurerm:documentdb/v20151106:DatabaseAccountMongoDBCollection" }, { type: "azurerm:documentdb/v20160331:DatabaseAccountMongoDBCollection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DatabaseAccountMongoDBCollection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabaseAccountMongoDBCollection resource.
 */
export interface DatabaseAccountMongoDBCollectionArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Cosmos DB collection name.
     */
    readonly collectionName: pulumi.Input<string>;
    /**
     * Cosmos DB database name.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
     */
    readonly options: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The standard JSON format of a MongoDB collection
     */
    readonly resource: pulumi.Input<inputs.documentdb.v20160319.MongoDBCollectionResource>;
    /**
     * Name of an Azure resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
