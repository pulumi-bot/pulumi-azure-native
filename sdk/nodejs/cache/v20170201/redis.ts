// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A single Redis item in List or Get Operation.
 */
export class Redis extends pulumi.CustomResource {
    /**
     * Get an existing Redis resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Redis {
        return new Redis(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cache/v20170201:Redis';

    /**
     * Returns true if the given object is an instance of Redis.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Redis {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Redis.__pulumiType;
    }

    /**
     * The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
     */
    public /*out*/ readonly accessKeys!: pulumi.Output<outputs.cache.v20170201.RedisAccessKeysResponse>;
    /**
     * Specifies whether the non-ssl Redis server port (6379) is enabled.
     */
    public readonly enableNonSslPort!: pulumi.Output<boolean | undefined>;
    /**
     * Redis host name.
     */
    public /*out*/ readonly hostName!: pulumi.Output<string>;
    /**
     * List of the linked servers associated with the cache
     */
    public /*out*/ readonly linkedServers!: pulumi.Output<outputs.cache.v20170201.RedisLinkedServerListResponse>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Redis non-SSL port.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * Redis instance provisioning status.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
     */
    public readonly redisConfiguration!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Redis version.
     */
    public /*out*/ readonly redisVersion!: pulumi.Output<string>;
    /**
     * The number of shards to be created on a Premium Cluster Cache.
     */
    public readonly shardCount!: pulumi.Output<number | undefined>;
    /**
     * The SKU of the Redis cache to deploy.
     */
    public readonly sku!: pulumi.Output<outputs.cache.v20170201.SkuResponse | undefined>;
    /**
     * Redis SSL port.
     */
    public /*out*/ readonly sslPort!: pulumi.Output<number>;
    /**
     * Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
     */
    public readonly staticIP!: pulumi.Output<string | undefined>;
    /**
     * The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * tenantSettings
     */
    public readonly tenantSettings!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Redis resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RedisArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["enableNonSslPort"] = args ? args.enableNonSslPort : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["redisConfiguration"] = args ? args.redisConfiguration : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["shardCount"] = args ? args.shardCount : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["staticIP"] = args ? args.staticIP : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenantSettings"] = args ? args.tenantSettings : undefined;
            inputs["accessKeys"] = undefined /*out*/;
            inputs["hostName"] = undefined /*out*/;
            inputs["linkedServers"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["redisVersion"] = undefined /*out*/;
            inputs["sslPort"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["accessKeys"] = undefined /*out*/;
            inputs["enableNonSslPort"] = undefined /*out*/;
            inputs["hostName"] = undefined /*out*/;
            inputs["linkedServers"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["redisConfiguration"] = undefined /*out*/;
            inputs["redisVersion"] = undefined /*out*/;
            inputs["shardCount"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["sslPort"] = undefined /*out*/;
            inputs["staticIP"] = undefined /*out*/;
            inputs["subnetId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tenantSettings"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:cache/latest:Redis" }, { type: "azurerm:cache/v20150801:Redis" }, { type: "azurerm:cache/v20160401:Redis" }, { type: "azurerm:cache/v20171001:Redis" }, { type: "azurerm:cache/v20180301:Redis" }, { type: "azurerm:cache/v20190701:Redis" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Redis.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Redis resource.
 */
export interface RedisArgs {
    /**
     * Specifies whether the non-ssl Redis server port (6379) is enabled.
     */
    readonly enableNonSslPort?: pulumi.Input<boolean>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the Redis cache.
     */
    readonly name: pulumi.Input<string>;
    /**
     * All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
     */
    readonly redisConfiguration?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The number of shards to be created on a Premium Cluster Cache.
     */
    readonly shardCount?: pulumi.Input<number>;
    /**
     * The SKU of the Redis cache to deploy.
     */
    readonly sku: pulumi.Input<inputs.cache.v20170201.Sku>;
    /**
     * Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
     */
    readonly staticIP?: pulumi.Input<string>;
    /**
     * The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * tenantSettings
     */
    readonly tenantSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
