// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An Azure Cosmos DB Cassandra keyspace.
 */
export class DatabaseAccountCassandraKeyspace extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseAccountCassandraKeyspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabaseAccountCassandraKeyspace {
        return new DatabaseAccountCassandraKeyspace(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:documentdb/v20150408:DatabaseAccountCassandraKeyspace';

    /**
     * Returns true if the given object is an instance of DatabaseAccountCassandraKeyspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseAccountCassandraKeyspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseAccountCassandraKeyspace.__pulumiType;
    }

    /**
     * The location of the resource group to which the resource belongs.
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the database account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of an Azure Cosmos DB Cassandra keyspace
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.documentdb.v20150408.CassandraKeyspacePropertiesResponse>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    public /*out*/ readonly tags!: pulumi.Output<outputs.documentdb.v20150408.TagsResponse | undefined>;
    /**
     * The type of Azure resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DatabaseAccountCassandraKeyspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseAccountCassandraKeyspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseAccountCassandraKeyspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DatabaseAccountCassandraKeyspaceArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
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
            inputs["name"] = args ? args.name : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["resource"] = args ? args.resource : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatabaseAccountCassandraKeyspace.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabaseAccountCassandraKeyspace resource.
 */
export interface DatabaseAccountCassandraKeyspaceArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Cosmos DB keyspace name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
     */
    readonly options: pulumi.Input<inputs.documentdb.v20150408.CreateUpdateOptions>;
    /**
     * The standard JSON format of a Cassandra keyspace
     */
    readonly resource: pulumi.Input<inputs.documentdb.v20150408.CassandraKeyspaceResource>;
    /**
     * Name of an Azure resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
