// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Cache details.
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
    public static readonly __pulumiType = 'azurerm:apimanagement/v20191201preview:Cache';

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
     * Runtime connection string to cache
     */
    public readonly connectionString!: pulumi.Output<string>;
    /**
     * Cache description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Original uri of entity in external system cache points to
     */
    public readonly resourceId!: pulumi.Output<string | undefined>;
    /**
     * Resource type for API Management resource.
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
            if (!args || args.cacheId === undefined) {
                throw new Error("Missing required property 'cacheId'");
            }
            if (!args || args.connectionString === undefined) {
                throw new Error("Missing required property 'connectionString'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["cacheId"] = args ? args.cacheId : undefined;
            inputs["connectionString"] = args ? args.connectionString : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:Cache" }, { type: "azurerm:apimanagement/v20180601preview:Cache" }, { type: "azurerm:apimanagement/v20190101:Cache" }, { type: "azurerm:apimanagement/v20191201:Cache" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Cache.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cache resource.
 */
export interface CacheArgs {
    /**
     * Identifier of the Cache entity. Cache identifier (should be either 'default' or valid Azure region identifier).
     */
    readonly cacheId: pulumi.Input<string>;
    /**
     * Runtime connection string to cache
     */
    readonly connectionString: pulumi.Input<string>;
    /**
     * Cache description
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Original uri of entity in external system cache points to
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
