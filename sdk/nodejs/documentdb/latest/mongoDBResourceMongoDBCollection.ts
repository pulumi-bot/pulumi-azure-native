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
 * const mongoDBResourceMongoDBCollection = new azurerm.documentdb.latest.MongoDBResourceMongoDBCollection("mongoDBResourceMongoDBCollection", {
 *     accountName: "ddb1",
 *     collectionName: "collectionName",
 *     databaseName: "databaseName",
 *     location: "West US",
 *     options: {},
 *     resource: {
 *         analyticalStorageTtl: 500,
 *         id: "collectionName",
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
 *     tags: {},
 * });
 *
 * ```
 */
export class MongoDBResourceMongoDBCollection extends pulumi.CustomResource {
    /**
     * Get an existing MongoDBResourceMongoDBCollection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MongoDBResourceMongoDBCollection {
        return new MongoDBResourceMongoDBCollection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:documentdb/latest:MongoDBResourceMongoDBCollection';

    /**
     * Returns true if the given object is an instance of MongoDBResourceMongoDBCollection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MongoDBResourceMongoDBCollection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MongoDBResourceMongoDBCollection.__pulumiType;
    }

    /**
     * Identity for the resource.
     */
    public readonly identity!: pulumi.Output<outputs.documentdb.latest.ManagedServiceIdentityResponse | undefined>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the ARM resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly options!: pulumi.Output<outputs.documentdb.latest.MongoDBCollectionGetPropertiesResponseOptions | undefined>;
    public readonly resource!: pulumi.Output<outputs.documentdb.latest.MongoDBCollectionGetPropertiesResponseResource | undefined>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of Azure resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a MongoDBResourceMongoDBCollection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MongoDBResourceMongoDBCollectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
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
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["resource"] = args ? args.resource : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["options"] = undefined /*out*/;
            inputs["resource"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:documentdb/v20190801:MongoDBResourceMongoDBCollection" }, { type: "azurerm:documentdb/v20191212:MongoDBResourceMongoDBCollection" }, { type: "azurerm:documentdb/v20200301:MongoDBResourceMongoDBCollection" }, { type: "azurerm:documentdb/v20200401:MongoDBResourceMongoDBCollection" }, { type: "azurerm:documentdb/v20200601preview:MongoDBResourceMongoDBCollection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(MongoDBResourceMongoDBCollection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a MongoDBResourceMongoDBCollection resource.
 */
export interface MongoDBResourceMongoDBCollectionArgs {
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
     * Identity for the resource.
     */
    readonly identity?: pulumi.Input<inputs.documentdb.latest.ManagedServiceIdentity>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
     */
    readonly options: pulumi.Input<inputs.documentdb.latest.CreateUpdateOptions>;
    /**
     * The standard JSON format of a MongoDB collection
     */
    readonly resource: pulumi.Input<inputs.documentdb.latest.MongoDBCollectionResource>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
