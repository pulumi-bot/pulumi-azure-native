// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * A Cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
 * Latest API Version: 2020-10-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:storagecache:getCache'. */
export function getCache(args: GetCacheArgs, opts?: pulumi.InvokeOptions): Promise<GetCacheResult> {
    pulumi.log.warn("getCache is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:storagecache:getCache'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:storagecache/latest:getCache", {
        "cacheName": args.cacheName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetCacheArgs {
    /**
     * Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
     */
    readonly cacheName: string;
    /**
     * Target resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * A Cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
 */
export interface GetCacheResult {
    /**
     * The size of this Cache, in GB.
     */
    readonly cacheSizeGB?: number;
    /**
     * Specifies Directory Services settings of the cache.
     */
    readonly directoryServicesSettings?: outputs.storagecache.latest.CacheDirectorySettingsResponse;
    /**
     * Specifies encryption settings of the cache.
     */
    readonly encryptionSettings?: outputs.storagecache.latest.CacheEncryptionSettingsResponse;
    /**
     * Health of the Cache.
     */
    readonly health: outputs.storagecache.latest.CacheHealthResponse;
    /**
     * Resource ID of the Cache.
     */
    readonly id: string;
    /**
     * The identity of the cache, if configured.
     */
    readonly identity?: outputs.storagecache.latest.CacheIdentityResponse;
    /**
     * Region name string.
     */
    readonly location?: string;
    /**
     * Array of IP addresses that can be used by clients mounting this Cache.
     */
    readonly mountAddresses: string[];
    /**
     * Name of Cache.
     */
    readonly name: string;
    /**
     * Specifies network settings of the cache.
     */
    readonly networkSettings?: outputs.storagecache.latest.CacheNetworkSettingsResponse;
    /**
     * ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
     */
    readonly provisioningState?: string;
    /**
     * Specifies security settings of the cache.
     */
    readonly securitySettings?: outputs.storagecache.latest.CacheSecuritySettingsResponse;
    /**
     * SKU for the Cache.
     */
    readonly sku?: outputs.storagecache.latest.CacheResponseSku;
    /**
     * Subnet used for the Cache.
     */
    readonly subnet?: string;
    /**
     * The system meta data relating to this resource.
     */
    readonly systemData: outputs.storagecache.latest.SystemDataResponse;
    /**
     * ARM tags as name/value pairs.
     */
    readonly tags?: any;
    /**
     * Type of the Cache; Microsoft.StorageCache/Cache
     */
    readonly type: string;
    /**
     * Upgrade status of the Cache.
     */
    readonly upgradeStatus?: outputs.storagecache.latest.CacheUpgradeStatusResponse;
}
