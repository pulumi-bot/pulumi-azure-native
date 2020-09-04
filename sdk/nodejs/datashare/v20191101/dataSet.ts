// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * A DataSet data transfer object.
 *
 * ## Example Usage
 * ### DataSets_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSet = new azurerm.datashare.v20191101.DataSet("dataSet", {
 *     accountName: "Account1",
 *     dataSetName: "Dataset1",
 *     kind: "Blob",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareName: "Share1",
 * });
 *
 * ```
 * ### DataSets_KustoCluster_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSet = new azurerm.datashare.v20191101.DataSet("dataSet", {
 *     accountName: "Account1",
 *     dataSetName: "Dataset1",
 *     kind: "KustoCluster",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareName: "Share1",
 * });
 *
 * ```
 * ### DataSets_KustoDatabase_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSet = new azurerm.datashare.v20191101.DataSet("dataSet", {
 *     accountName: "Account1",
 *     dataSetName: "Dataset1",
 *     kind: "KustoDatabase",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareName: "Share1",
 * });
 *
 * ```
 * ### DataSets_SqlDBTable_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSet = new azurerm.datashare.v20191101.DataSet("dataSet", {
 *     accountName: "Account1",
 *     dataSetName: "Dataset1",
 *     kind: "SqlDBTable",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareName: "Share1",
 * });
 *
 * ```
 * ### DataSets_SqlDWTable_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSet = new azurerm.datashare.v20191101.DataSet("dataSet", {
 *     accountName: "Account1",
 *     dataSetName: "Dataset1",
 *     kind: "SqlDWTable",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareName: "Share1",
 * });
 *
 * ```
 */
export class DataSet extends pulumi.CustomResource {
    /**
     * Get an existing DataSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataSet {
        return new DataSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datashare/v20191101:DataSet';

    /**
     * Returns true if the given object is an instance of DataSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSet.__pulumiType;
    }

    /**
     * Kind of data set.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Name of the azure resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Type of the azure resource
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DataSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.dataSetName === undefined) {
                throw new Error("Missing required property 'dataSetName'");
            }
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.shareName === undefined) {
                throw new Error("Missing required property 'shareName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["dataSetName"] = args ? args.dataSetName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["shareName"] = args ? args.shareName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:datashare/latest:DataSet" }, { type: "azurerm:datashare/v20181101preview:DataSet" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DataSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataSet resource.
 */
export interface DataSetArgs {
    /**
     * The name of the share account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the dataSet.
     */
    readonly dataSetName: pulumi.Input<string>;
    /**
     * Kind of data set.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the share to add the data set to.
     */
    readonly shareName: pulumi.Input<string>;
}
