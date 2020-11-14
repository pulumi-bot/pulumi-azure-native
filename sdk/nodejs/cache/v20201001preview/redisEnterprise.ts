// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Describes the RedisEnterprise cluster
 */
export class RedisEnterprise extends pulumi.CustomResource {
    /**
     * Get an existing RedisEnterprise resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RedisEnterprise {
        return new RedisEnterprise(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:cache/v20201001preview:RedisEnterprise';

    /**
     * Returns true if the given object is an instance of RedisEnterprise.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RedisEnterprise {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RedisEnterprise.__pulumiType;
    }

    /**
     * DNS name of the cluster endpoint
     */
    public /*out*/ readonly hostName!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The minimum TLS version for the cluster to support, e.g. '1.2'
     */
    public readonly minimumTlsVersion!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of private endpoint connections associated with the specified RedisEnterprise cluster
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.cache.v20201001preview.PrivateEndpointConnectionResponse[]>;
    /**
     * Current provisioning status of the cluster
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Version of redis the cluster supports, e.g. '6'
     */
    public /*out*/ readonly redisVersion!: pulumi.Output<string>;
    /**
     * Current resource status of the cluster
     */
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * The SKU to create, which affects price, performance, and features.
     */
    public readonly sku!: pulumi.Output<outputs.cache.v20201001preview.SkuResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The zones where this cluster will be deployed.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a RedisEnterprise resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RedisEnterpriseArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["minimumTlsVersion"] = args ? args.minimumTlsVersion : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["hostName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["redisVersion"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["hostName"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["minimumTlsVersion"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["redisVersion"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["zones"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RedisEnterprise.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RedisEnterprise resource.
 */
export interface RedisEnterpriseArgs {
    /**
     * The name of the RedisEnterprise cluster.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The minimum TLS version for the cluster to support, e.g. '1.2'
     */
    readonly minimumTlsVersion?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU to create, which affects price, performance, and features.
     */
    readonly sku: pulumi.Input<inputs.cache.v20201001preview.Sku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The zones where this cluster will be deployed.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
