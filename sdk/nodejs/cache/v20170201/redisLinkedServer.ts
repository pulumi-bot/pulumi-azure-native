// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Response to put/get linked server (with properties) for Redis cache.
 */
export class RedisLinkedServer extends pulumi.CustomResource {
    /**
     * Get an existing RedisLinkedServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RedisLinkedServer {
        return new RedisLinkedServer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cache/v20170201:RedisLinkedServer';

    /**
     * Returns true if the given object is an instance of RedisLinkedServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RedisLinkedServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RedisLinkedServer.__pulumiType;
    }

    /**
     * Fully qualified resourceId of the linked redis cache.
     */
    public readonly linkedRedisCacheId!: pulumi.Output<string>;
    /**
     * Location of the linked redis cache.
     */
    public readonly linkedRedisCacheLocation!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Terminal state of the link between primary and secondary redis cache.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Role of the linked server.
     */
    public readonly serverRole!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RedisLinkedServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RedisLinkedServerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.linkedRedisCacheId === undefined) {
                throw new Error("Missing required property 'linkedRedisCacheId'");
            }
            if (!args || args.linkedRedisCacheLocation === undefined) {
                throw new Error("Missing required property 'linkedRedisCacheLocation'");
            }
            if (!args || args.linkedServerName === undefined) {
                throw new Error("Missing required property 'linkedServerName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverRole === undefined) {
                throw new Error("Missing required property 'serverRole'");
            }
            inputs["linkedRedisCacheId"] = args ? args.linkedRedisCacheId : undefined;
            inputs["linkedRedisCacheLocation"] = args ? args.linkedRedisCacheLocation : undefined;
            inputs["linkedServerName"] = args ? args.linkedServerName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverRole"] = args ? args.serverRole : undefined;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["linkedRedisCacheId"] = undefined /*out*/;
            inputs["linkedRedisCacheLocation"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["serverRole"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:cache/latest:RedisLinkedServer" }, { type: "azurerm:cache/v20171001:RedisLinkedServer" }, { type: "azurerm:cache/v20180301:RedisLinkedServer" }, { type: "azurerm:cache/v20190701:RedisLinkedServer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RedisLinkedServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RedisLinkedServer resource.
 */
export interface RedisLinkedServerArgs {
    /**
     * Fully qualified resourceId of the linked redis cache.
     */
    readonly linkedRedisCacheId: pulumi.Input<string>;
    /**
     * Location of the linked redis cache.
     */
    readonly linkedRedisCacheLocation: pulumi.Input<string>;
    /**
     * The name of the linked server that is being added to the Redis cache.
     */
    readonly linkedServerName: pulumi.Input<string>;
    /**
     * The name of the Redis cache.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Role of the linked server.
     */
    readonly serverRole: pulumi.Input<string>;
}
