// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * A firewall rule on a redis cache has a name, and describes a contiguous range of IP addresses permitted to connect
 *
 * ## Example Usage
 * ### RedisCacheFirewallRuleCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const firewallRule = new azurerm.cache.v20170201.FirewallRule("firewallRule", {
 *     cacheName: "cache1",
 *     endIP: "192.168.1.4",
 *     resourceGroupName: "rg1",
 *     ruleName: "rule1",
 *     startIP: "192.168.1.1",
 * });
 *
 * ```
 */
export class FirewallRule extends pulumi.CustomResource {
    /**
     * Get an existing FirewallRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FirewallRule {
        return new FirewallRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cache/v20170201:FirewallRule';

    /**
     * Returns true if the given object is an instance of FirewallRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallRule.__pulumiType;
    }

    /**
     * highest IP address included in the range
     */
    public readonly endIP!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * lowest IP address included in the range
     */
    public readonly startIP!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a FirewallRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as FirewallRuleArgs | undefined;
            if (!args || args.cacheName === undefined) {
                throw new Error("Missing required property 'cacheName'");
            }
            if (!args || args.endIP === undefined) {
                throw new Error("Missing required property 'endIP'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.ruleName === undefined) {
                throw new Error("Missing required property 'ruleName'");
            }
            if (!args || args.startIP === undefined) {
                throw new Error("Missing required property 'startIP'");
            }
            inputs["cacheName"] = args ? args.cacheName : undefined;
            inputs["endIP"] = args ? args.endIP : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["ruleName"] = args ? args.ruleName : undefined;
            inputs["startIP"] = args ? args.startIP : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:cache/latest:FirewallRule" }, { type: "azurerm:cache/v20160401:FirewallRule" }, { type: "azurerm:cache/v20171001:FirewallRule" }, { type: "azurerm:cache/v20180301:FirewallRule" }, { type: "azurerm:cache/v20190701:FirewallRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(FirewallRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FirewallRule resource.
 */
export interface FirewallRuleArgs {
    /**
     * The name of the Redis cache.
     */
    readonly cacheName: pulumi.Input<string>;
    /**
     * highest IP address included in the range
     */
    readonly endIP: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the firewall rule.
     */
    readonly ruleName: pulumi.Input<string>;
    /**
     * lowest IP address included in the range
     */
    readonly startIP: pulumi.Input<string>;
}
