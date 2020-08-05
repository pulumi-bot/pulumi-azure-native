// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A firewall rule on a redis cache has a name, and describes a contiguous range of IP addresses permitted to connect
 */
export class RedisFirewallRule extends pulumi.CustomResource {
    /**
     * Get an existing RedisFirewallRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RedisFirewallRule {
        return new RedisFirewallRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cache/v20160401:RedisFirewallRule';

    /**
     * Returns true if the given object is an instance of RedisFirewallRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RedisFirewallRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RedisFirewallRule.__pulumiType;
    }

    /**
     * name of the firewall rule
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * redis cache firewall rule properties
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.cache.v20160401.RedisFirewallRulePropertiesResponse>;
    /**
     * type (of the firewall rule resource = 'Microsoft.Cache/redis/firewallRule')
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RedisFirewallRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RedisFirewallRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RedisFirewallRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RedisFirewallRuleArgs | undefined;
            if (!args || args.cacheName === undefined) {
                throw new Error("Missing required property 'cacheName'");
            }
            if (!args || args.endIP === undefined) {
                throw new Error("Missing required property 'endIP'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.startIP === undefined) {
                throw new Error("Missing required property 'startIP'");
            }
            inputs["cacheName"] = args ? args.cacheName : undefined;
            inputs["endIP"] = args ? args.endIP : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["startIP"] = args ? args.startIP : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RedisFirewallRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RedisFirewallRule resource.
 */
export interface RedisFirewallRuleArgs {
    /**
     * The name of the Redis cache.
     */
    readonly cacheName: pulumi.Input<string>;
    /**
     * highest IP address included in the range
     */
    readonly endIP: pulumi.Input<string>;
    /**
     * The name of the firewall rule.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * lowest IP address included in the range
     */
    readonly startIP: pulumi.Input<string>;
}
