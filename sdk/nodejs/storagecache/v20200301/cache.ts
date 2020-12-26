// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
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
        return new Cache(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:storagecache/v20200301:Cache';

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
     * Specifies encryption settings of the cache.
     */
    public readonly encryptionSettings!: pulumi.Output<outputs.storagecache.v20200301.CacheEncryptionSettingsResponse | undefined>;
    /**
     * Health of the Cache.
     */
    public /*out*/ readonly health!: pulumi.Output<outputs.storagecache.v20200301.CacheHealthResponse>;
    /**
     * The identity of the cache, if configured.
     */
    public readonly identity!: pulumi.Output<outputs.storagecache.v20200301.CacheIdentityResponse | undefined>;
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
     * Specifies network settings of the cache.
     */
    public readonly networkSettings!: pulumi.Output<outputs.storagecache.v20200301.CacheNetworkSettingsResponse | undefined>;
    /**
     * ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Specifies security settings of the cache.
     */
    public readonly securitySettings!: pulumi.Output<outputs.storagecache.v20200301.CacheSecuritySettingsResponse | undefined>;
    /**
     * SKU for the Cache.
     */
    public readonly sku!: pulumi.Output<outputs.storagecache.v20200301.CacheResponseSku | undefined>;
    /**
     * Subnet used for the Cache.
     */
    public readonly subnet!: pulumi.Output<string | undefined>;
    /**
     * The system meta data relating to this resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.storagecache.v20200301.SystemDataResponse>;
    /**
     * ARM tags as name/value pairs.
     */
    public readonly tags!: pulumi.Output<any | undefined>;
    /**
     * Type of the Cache; Microsoft.StorageCache/Cache
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Upgrade status of the Cache.
     */
    public /*out*/ readonly upgradeStatus!: pulumi.Output<outputs.storagecache.v20200301.CacheUpgradeStatusResponse | undefined>;

    /**
     * Create a Cache resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CacheArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.cacheName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'cacheName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["cacheName"] = args ? args.cacheName : undefined;
            inputs["cacheSizeGB"] = args ? args.cacheSizeGB : undefined;
            inputs["encryptionSettings"] = args ? args.encryptionSettings : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkSettings"] = args ? args.networkSettings : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["securitySettings"] = args ? args.securitySettings : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["health"] = undefined /*out*/;
            inputs["mountAddresses"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["upgradeStatus"] = undefined /*out*/;
        } else {
            inputs["cacheSizeGB"] = undefined /*out*/;
            inputs["encryptionSettings"] = undefined /*out*/;
            inputs["health"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["mountAddresses"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkSettings"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["securitySettings"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["subnet"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["upgradeStatus"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:storagecache/latest:Cache" }, { type: "azure-nextgen:storagecache/v20190801preview:Cache" }, { type: "azure-nextgen:storagecache/v20191101:Cache" }, { type: "azure-nextgen:storagecache/v20201001:Cache" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Cache.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cache resource.
 */
export interface CacheArgs {
    /**
     * Name of Cache. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
     */
    readonly cacheName: pulumi.Input<string>;
    /**
     * The size of this Cache, in GB.
     */
    readonly cacheSizeGB?: pulumi.Input<number>;
    /**
     * Specifies encryption settings of the cache.
     */
    readonly encryptionSettings?: pulumi.Input<inputs.storagecache.v20200301.CacheEncryptionSettings>;
    /**
     * The identity of the cache, if configured.
     */
    readonly identity?: pulumi.Input<inputs.storagecache.v20200301.CacheIdentity>;
    /**
     * Region name string.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies network settings of the cache.
     */
    readonly networkSettings?: pulumi.Input<inputs.storagecache.v20200301.CacheNetworkSettings>;
    /**
     * ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
     */
    readonly provisioningState?: pulumi.Input<string | enums.storagecache.v20200301.ProvisioningStateType>;
    /**
     * Target resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies security settings of the cache.
     */
    readonly securitySettings?: pulumi.Input<inputs.storagecache.v20200301.CacheSecuritySettings>;
    /**
     * SKU for the Cache.
     */
    readonly sku?: pulumi.Input<inputs.storagecache.v20200301.CacheSku>;
    /**
     * Subnet used for the Cache.
     */
    readonly subnet?: pulumi.Input<string>;
    /**
     * ARM tags as name/value pairs.
     */
    readonly tags?: any;
}
