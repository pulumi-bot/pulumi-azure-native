// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The DataManager resource.
 */
export class DataManager extends pulumi.CustomResource {
    /**
     * Get an existing DataManager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataManager {
        return new DataManager(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:hybriddata/v20160601:DataManager';

    /**
     * Returns true if the given object is an instance of DataManager.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataManager {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataManager.__pulumiType;
    }

    /**
     * Etag of the Resource.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
     * US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
     * region is specified on update the request will succeed.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The Resource Name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The sku type.
     */
    public readonly sku!: pulumi.Output<outputs.hybriddata.v20160601.SkuResponse | undefined>;
    /**
     * The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
     * (across resource groups).
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DataManager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataManagerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.dataManagerName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dataManagerName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dataManagerName"] = args ? args.dataManagerName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:hybriddata/latest:DataManager" }, { type: "azure-nextgen:hybriddata/v20190601:DataManager" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DataManager.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataManager resource.
 */
export interface DataManagerArgs {
    /**
     * The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
     */
    readonly dataManagerName: pulumi.Input<string>;
    /**
     * Etag of the Resource.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East
     * US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo
     * region is specified on update the request will succeed.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Resource Group Name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku type.
     */
    readonly sku?: pulumi.Input<inputs.hybriddata.v20160601.Sku>;
    /**
     * The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource
     * (across resource groups).
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
