// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Rule Group resource.
 */
export class FirewallPolicyRuleGroup extends pulumi.CustomResource {
    /**
     * Get an existing FirewallPolicyRuleGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FirewallPolicyRuleGroup {
        return new FirewallPolicyRuleGroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200401:FirewallPolicyRuleGroup';

    /**
     * Returns true if the given object is an instance of FirewallPolicyRuleGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallPolicyRuleGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallPolicyRuleGroup.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The properties of the firewall policy rule group.
     */
    public readonly properties!: pulumi.Output<outputs.network.v20200401.FirewallPolicyRuleGroupPropertiesResponse>;
    /**
     * Rule Group type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a FirewallPolicyRuleGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallPolicyRuleGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallPolicyRuleGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as FirewallPolicyRuleGroupArgs | undefined;
            if (!args || args.firewallPolicyName === undefined) {
                throw new Error("Missing required property 'firewallPolicyName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["firewallPolicyName"] = args ? args.firewallPolicyName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FirewallPolicyRuleGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FirewallPolicyRuleGroup resource.
 */
export interface FirewallPolicyRuleGroupArgs {
    /**
     * The name of the Firewall Policy.
     */
    readonly firewallPolicyName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the FirewallPolicyRuleGroup.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The properties of the firewall policy rule group.
     */
    readonly properties?: pulumi.Input<inputs.network.v20200401.FirewallPolicyRuleGroupProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
