// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * A data set mapping data transfer object.
 *
 * ## Example Usage
 * ### DataSetMappings_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSetMapping = new azurerm.datashare.v20191101.DataSetMapping("dataSetMapping", {
 *     accountName: "Account1",
 *     dataSetMappingName: "DatasetMapping1",
 *     kind: "Blob",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareSubscriptionName: "ShareSubscription1",
 * });
 *
 * ```
 * ### DataSetMappings_SqlDB_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSetMapping = new azurerm.datashare.v20191101.DataSetMapping("dataSetMapping", {
 *     accountName: "Account1",
 *     dataSetMappingName: "DatasetMapping1",
 *     kind: "SqlDBTable",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareSubscriptionName: "ShareSubscription1",
 * });
 *
 * ```
 * ### DataSetMappings_SqlDWDataSetToAdlsGen2File_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSetMapping = new azurerm.datashare.v20191101.DataSetMapping("dataSetMapping", {
 *     accountName: "Account1",
 *     dataSetMappingName: "DatasetMapping1",
 *     kind: "AdlsGen2File",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareSubscriptionName: "ShareSubscription1",
 * });
 *
 * ```
 * ### DataSetMappings_SqlDW_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dataSetMapping = new azurerm.datashare.v20191101.DataSetMapping("dataSetMapping", {
 *     accountName: "Account1",
 *     dataSetMappingName: "DatasetMapping1",
 *     kind: "SqlDWTable",
 *     resourceGroupName: "SampleResourceGroup",
 *     shareSubscriptionName: "ShareSubscription1",
 * });
 *
 * ```
 */
export class DataSetMapping extends pulumi.CustomResource {
    /**
     * Get an existing DataSetMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataSetMapping {
        return new DataSetMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datashare/v20191101:DataSetMapping';

    /**
     * Returns true if the given object is an instance of DataSetMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSetMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSetMapping.__pulumiType;
    }

    /**
     * Kind of data set mapping.
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
     * Create a DataSetMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSetMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.dataSetMappingName === undefined) {
                throw new Error("Missing required property 'dataSetMappingName'");
            }
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.shareSubscriptionName === undefined) {
                throw new Error("Missing required property 'shareSubscriptionName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["dataSetMappingName"] = args ? args.dataSetMappingName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["shareSubscriptionName"] = args ? args.shareSubscriptionName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azurerm:datashare/latest:DataSetMapping" }, { type: "azurerm:datashare/v20181101preview:DataSetMapping" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DataSetMapping.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataSetMapping resource.
 */
export interface DataSetMappingArgs {
    /**
     * The name of the share account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the data set mapping to be created.
     */
    readonly dataSetMappingName: pulumi.Input<string>;
    /**
     * Kind of data set mapping.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the share subscription which will hold the data set sink.
     */
    readonly shareSubscriptionName: pulumi.Input<string>;
}
