// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Description of a namespace authorization rule.
 *
 * ## Example Usage
 * ### RelayAuthorizationRuleCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const wcfRelayAuthorizationRule = new azurerm.relay.v20170401.WCFRelayAuthorizationRule("wcfRelayAuthorizationRule", {
 *     authorizationRuleName: "sdk-RelayAuthRules-01",
 *     namespaceName: "sdk-RelayNamespace-01",
 *     relayName: "sdk-Relay-wcf-01",
 *     resourceGroupName: "RG-eg",
 *     rights: [
 *         "Listen",
 *         "Send",
 *     ],
 * });
 *
 * ```
 */
export class WCFRelayAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing WCFRelayAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WCFRelayAuthorizationRule {
        return new WCFRelayAuthorizationRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:relay/v20170401:WCFRelayAuthorizationRule';

    /**
     * Returns true if the given object is an instance of WCFRelayAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WCFRelayAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WCFRelayAuthorizationRule.__pulumiType;
    }

    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The rights associated with the rule.
     */
    public readonly rights!: pulumi.Output<string[]>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WCFRelayAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WCFRelayAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WCFRelayAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WCFRelayAuthorizationRuleArgs | undefined;
            if (!args || args.authorizationRuleName === undefined) {
                throw new Error("Missing required property 'authorizationRuleName'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.relayName === undefined) {
                throw new Error("Missing required property 'relayName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.rights === undefined) {
                throw new Error("Missing required property 'rights'");
            }
            inputs["authorizationRuleName"] = args ? args.authorizationRuleName : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["relayName"] = args ? args.relayName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["rights"] = args ? args.rights : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:relay/latest:WCFRelayAuthorizationRule" }, { type: "azurerm:relay/v20160701:WCFRelayAuthorizationRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WCFRelayAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WCFRelayAuthorizationRule resource.
 */
export interface WCFRelayAuthorizationRuleArgs {
    /**
     * The authorization rule name.
     */
    readonly authorizationRuleName: pulumi.Input<string>;
    /**
     * The namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * The relay name.
     */
    readonly relayName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The rights associated with the rule.
     */
    readonly rights: pulumi.Input<pulumi.Input<string>[]>;
}
