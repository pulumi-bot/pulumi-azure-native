// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Single item in a List or Get VirtualNetworkRules operation
 */
export class NamespaceVirtualNetworkRule extends pulumi.CustomResource {
    /**
     * Get an existing NamespaceVirtualNetworkRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NamespaceVirtualNetworkRule {
        return new NamespaceVirtualNetworkRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicebus/v20180101preview:NamespaceVirtualNetworkRule';

    /**
     * Returns true if the given object is an instance of NamespaceVirtualNetworkRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NamespaceVirtualNetworkRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NamespaceVirtualNetworkRule.__pulumiType;
    }

    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Resource ID of Virtual Network Subnet
     */
    public readonly virtualNetworkSubnetId!: pulumi.Output<string | undefined>;

    /**
     * Create a NamespaceVirtualNetworkRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceVirtualNetworkRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualNetworkRuleName === undefined) {
                throw new Error("Missing required property 'virtualNetworkRuleName'");
            }
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["virtualNetworkRuleName"] = args ? args.virtualNetworkRuleName : undefined;
            inputs["virtualNetworkSubnetId"] = args ? args.virtualNetworkSubnetId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworkSubnetId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:servicebus/latest:NamespaceVirtualNetworkRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NamespaceVirtualNetworkRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NamespaceVirtualNetworkRule resource.
 */
export interface NamespaceVirtualNetworkRuleArgs {
    /**
     * The namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Virtual Network Rule name.
     */
    readonly virtualNetworkRuleName: pulumi.Input<string>;
    /**
     * Resource ID of Virtual Network Subnet
     */
    readonly virtualNetworkSubnetId?: pulumi.Input<string>;
}
