// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An Azure Cosmos DB Gremlin graph.
 *
 * ## Example Usage
 * ### CosmosDBGremlinGraphCreateUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const gremlinResourceGremlinGraph = new azurerm.documentdb.v20200301.GremlinResourceGremlinGraph("gremlinResourceGremlinGraph", {
 *     accountName: "ddb1",
 *     databaseName: "databaseName",
 *     graphName: "graphName",
 *     location: "West US",
 *     options: {},
 *     resource: {
 *         conflictResolutionPolicy: {
 *             conflictResolutionPath: "/path",
 *             mode: "LastWriterWins",
 *         },
 *         defaultTtl: 100,
 *         id: "graphName",
 *         indexingPolicy: {
 *             automatic: true,
 *             excludedPaths: [],
 *             includedPaths: [{
 *                 indexes: [
 *                     {
 *                         dataType: "String",
 *                         kind: "Range",
 *                         precision: -1,
 *                     },
 *                     {
 *                         dataType: "Number",
 *                         kind: "Range",
 *                         precision: -1,
 *                     },
 *                 ],
 *                 path: "/*",
 *             }],
 *             indexingMode: "Consistent",
 *         },
 *         partitionKey: {
 *             kind: "Hash",
 *             paths: ["/AccountNumber"],
 *         },
 *         uniqueKeyPolicy: {
 *             uniqueKeys: [{
 *                 paths: ["/testPath"],
 *             }],
 *         },
 *     },
 *     resourceGroupName: "rg1",
 *     tags: {},
 * });
 *
 * ```
 */
export class GremlinResourceGremlinGraph extends pulumi.CustomResource {
    /**
     * Get an existing GremlinResourceGremlinGraph resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GremlinResourceGremlinGraph {
        return new GremlinResourceGremlinGraph(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:documentdb/v20200301:GremlinResourceGremlinGraph';

    /**
     * Returns true if the given object is an instance of GremlinResourceGremlinGraph.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GremlinResourceGremlinGraph {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GremlinResourceGremlinGraph.__pulumiType;
    }

    /**
     * The location of the resource group to which the resource belongs.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the ARM resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly options!: pulumi.Output<outputs.documentdb.v20200301.GremlinGraphGetPropertiesResponseOptions | undefined>;
    public readonly resource!: pulumi.Output<outputs.documentdb.v20200301.GremlinGraphGetPropertiesResponseResource | undefined>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of Azure resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GremlinResourceGremlinGraph resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GremlinResourceGremlinGraphArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GremlinResourceGremlinGraphArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as GremlinResourceGremlinGraphArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.graphName === undefined) {
                throw new Error("Missing required property 'graphName'");
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
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["graphName"] = args ? args.graphName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["resource"] = args ? args.resource : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:documentdb/latest:GremlinResourceGremlinGraph" }, { type: "azurerm:documentdb/v20190801:GremlinResourceGremlinGraph" }, { type: "azurerm:documentdb/v20191212:GremlinResourceGremlinGraph" }, { type: "azurerm:documentdb/v20200401:GremlinResourceGremlinGraph" }, { type: "azurerm:documentdb/v20200601preview:GremlinResourceGremlinGraph" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(GremlinResourceGremlinGraph.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GremlinResourceGremlinGraph resource.
 */
export interface GremlinResourceGremlinGraphArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Cosmos DB database name.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * Cosmos DB graph name.
     */
    readonly graphName: pulumi.Input<string>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
     */
    readonly options: pulumi.Input<inputs.documentdb.v20200301.CreateUpdateOptions>;
    /**
     * The standard JSON format of a Gremlin graph
     */
    readonly resource: pulumi.Input<inputs.documentdb.v20200301.GremlinGraphResource>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
