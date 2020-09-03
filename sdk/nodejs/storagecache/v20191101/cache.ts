// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A Cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
 *
 * ## Example Usage
 * ### Caches_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const cache = new azurerm.storagecache.v20191101.Cache("cache", {
 *     cacheName: "sc1",
 *     cacheSizeGB: 3072,
 *     location: "westus",
 *     resourceGroupName: "scgroup",
 *     sku: {
 *         name: "Standard_2G",
 *     },
 *     subnet: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/sub1",
 *     tags: {
 *         Dept: "Initech",
 *     },
 * });
 *
 * ```
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
     * The size of this Cache, in GB.
     */
    public readonly cacheSizeGB!: pulumi.Output<number | undefined>;
    /**
     * Health of the Cache.
     */
    public /*out*/ readonly health!: pulumi.Output<outputs.storagecache.v20191101.CacheHealthResponse>;
    /**
     * Region name string.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Array of IP addresses that can be used by clients mounting this Cache.
     */
    public /*out*/ readonly mountAddresses!: pulumi.Output<string[]>;
    /**
     * Name of Cache.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * SKU for the Cache.
     */
    public readonly sku!: pulumi.Output<outputs.storagecache.v20191101.CacheResponseSku | undefined>;
    /**
     * Subnet used for the Cache.
     */
    public readonly subnet!: pulumi.Output<string | undefined>;
    /**
     * ARM tags as name/value pairs.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Type of the Cache; Microsoft.StorageCache/Cache
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Upgrade status of the Cache.
     */
    public /*out*/ readonly upgradeStatus!: pulumi.Output<outputs.storagecache.v20191101.CacheUpgradeStatusResponse | undefined>;

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
            if (!args || args.cacheName === undefined) {
                throw new Error("Missing required property 'cacheName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["cacheName"] = args ? args.cacheName : undefined;
            inputs["cacheSizeGB"] = args ? args.cacheSizeGB : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["health"] = undefined /*out*/;
            inputs["mountAddresses"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["upgradeStatus"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storagecache/latest:Cache" }, { type: "azurerm:storagecache/v20190801preview:Cache" }, { type: "azurerm:storagecache/v20200301:Cache" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Cache.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cache resource.
 */
export interface CacheArgs {
    /**
     * Name of Cache.
     */
    readonly cacheName: pulumi.Input<string>;
    /**
     * The size of this Cache, in GB.
     */
    readonly cacheSizeGB?: pulumi.Input<number>;
    /**
     * Region name string.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Target resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * SKU for the Cache.
     */
    readonly sku?: pulumi.Input<inputs.storagecache.v20191101.CacheSku>;
    /**
     * Subnet used for the Cache.
     */
    readonly subnet?: pulumi.Input<string>;
    /**
     * ARM tags as name/value pairs.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
