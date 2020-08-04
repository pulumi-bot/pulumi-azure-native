// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A Cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
 */
export class Cache extends pulumi.CustomResource {
    /**
     * Get an existing Cache resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cache {
        return new Cache(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storagecache/v20191101:Cache';

    /**
     * Returns true if the given object is an instance of Cache.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cache {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cache.__pulumiType;
    }

    /**
     * Region name string.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Name of Cache.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the Cache.
     */
    public readonly properties!: pulumi.Output<outputs.storagecache.v20191101.CacheResponseProperties>;
    /**
     * SKU for the Cache.
     */
    public readonly sku!: pulumi.Output<outputs.storagecache.v20191101.CacheResponseProperties | undefined>;
    /**
     * ARM tags as name/value pairs.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Type of the Cache; Microsoft.StorageCache/Cache
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Cache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CacheArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CacheArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CacheArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cache.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cache resource.
 */
export interface CacheArgs {
    /**
     * Region name string.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Name of Cache.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Properties of the Cache.
     */
    readonly properties?: pulumi.Input<inputs.storagecache.v20191101.CacheProperties>;
    /**
     * Target resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * SKU for the Cache.
     */
    readonly sku?: pulumi.Input<inputs.storagecache.v20191101.CacheProperties>;
    /**
     * ARM tags as name/value pairs.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
