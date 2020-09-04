// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * FirewallPolicy Resource.
 *
 * ## Example Usage
 * ### Create FirewallPolicy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const firewallPolicy = new azurerm.network.v20191101.FirewallPolicy("firewallPolicy", {
 *     firewallPolicyName: "firewallPolicy",
 *     location: "West US",
 *     resourceGroupName: "rg1",
 *     tags: {
 *         key1: "value1",
 *     },
 *     threatIntelMode: "Alert",
 * });
 *
 * ```
 */
export class FirewallPolicy extends pulumi.CustomResource {
    /**
     * Get an existing FirewallPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FirewallPolicy {
        return new FirewallPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20191101:FirewallPolicy';

    /**
     * Returns true if the given object is an instance of FirewallPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallPolicy.__pulumiType;
    }

    /**
     * The parent firewall policy from which rules are inherited.
     */
    public readonly basePolicy!: pulumi.Output<outputs.network.v20191101.SubResourceResponse | undefined>;
    /**
     * List of references to Child Firewall Policies.
     */
    public /*out*/ readonly childPolicies!: pulumi.Output<outputs.network.v20191101.SubResourceResponse[]>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * List of references to Azure Firewalls that this Firewall Policy is associated with.
     */
    public /*out*/ readonly firewalls!: pulumi.Output<outputs.network.v20191101.SubResourceResponse[]>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the firewall policy resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * List of references to FirewallPolicyRuleGroups.
     */
    public /*out*/ readonly ruleGroups!: pulumi.Output<outputs.network.v20191101.SubResourceResponse[]>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The operation mode for Threat Intelligence.
     */
    public readonly threatIntelMode!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a FirewallPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.firewallPolicyName === undefined) {
                throw new Error("Missing required property 'firewallPolicyName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["basePolicy"] = args ? args.basePolicy : undefined;
            inputs["firewallPolicyName"] = args ? args.firewallPolicyName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["threatIntelMode"] = args ? args.threatIntelMode : undefined;
            inputs["childPolicies"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["firewalls"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["ruleGroups"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["basePolicy"] = undefined /*out*/;
            inputs["childPolicies"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["firewalls"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["ruleGroups"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["threatIntelMode"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:FirewallPolicy" }, { type: "azurerm:network/v20190601:FirewallPolicy" }, { type: "azurerm:network/v20190701:FirewallPolicy" }, { type: "azurerm:network/v20190801:FirewallPolicy" }, { type: "azurerm:network/v20190901:FirewallPolicy" }, { type: "azurerm:network/v20191201:FirewallPolicy" }, { type: "azurerm:network/v20200301:FirewallPolicy" }, { type: "azurerm:network/v20200401:FirewallPolicy" }, { type: "azurerm:network/v20200501:FirewallPolicy" }, { type: "azurerm:network/v20200601:FirewallPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(FirewallPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FirewallPolicy resource.
 */
export interface FirewallPolicyArgs {
    /**
     * The parent firewall policy from which rules are inherited.
     */
    readonly basePolicy?: pulumi.Input<inputs.network.v20191101.SubResource>;
    /**
     * The name of the Firewall Policy.
     */
    readonly firewallPolicyName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The operation mode for Threat Intelligence.
     */
    readonly threatIntelMode?: pulumi.Input<string>;
}
