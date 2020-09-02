// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * A virtual network rule.
 */
export class VirtualNetworkRule extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetworkRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualNetworkRule {
        return new VirtualNetworkRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:dbforpostgresql/v20171201preview:VirtualNetworkRule';

    /**
     * Returns true if the given object is an instance of VirtualNetworkRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetworkRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetworkRule.__pulumiType;
    }

    /**
     * Create firewall rule before the virtual network has vnet service endpoint enabled.
     */
    public readonly ignoreMissingVnetServiceEndpoint!: pulumi.Output<boolean | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Virtual Network Rule State
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The ARM resource id of the virtual network subnet.
     */
    public readonly virtualNetworkSubnetId!: pulumi.Output<string>;

    /**
     * Create a VirtualNetworkRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VirtualNetworkRuleArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            if (!args || args.virtualNetworkRuleName === undefined) {
                throw new Error("Missing required property 'virtualNetworkRuleName'");
            }
            if (!args || args.virtualNetworkSubnetId === undefined) {
                throw new Error("Missing required property 'virtualNetworkSubnetId'");
            }
            inputs["ignoreMissingVnetServiceEndpoint"] = args ? args.ignoreMissingVnetServiceEndpoint : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["virtualNetworkRuleName"] = args ? args.virtualNetworkRuleName : undefined;
            inputs["virtualNetworkSubnetId"] = args ? args.virtualNetworkSubnetId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:dbforpostgresql/latest:VirtualNetworkRule" }, { type: "azurerm:dbforpostgresql/v20171201:VirtualNetworkRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualNetworkRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualNetworkRule resource.
 */
export interface VirtualNetworkRuleArgs {
    /**
     * Create firewall rule before the virtual network has vnet service endpoint enabled.
     */
    readonly ignoreMissingVnetServiceEndpoint?: pulumi.Input<boolean>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The name of the virtual network rule.
     */
    readonly virtualNetworkRuleName: pulumi.Input<string>;
    /**
     * The ARM resource id of the virtual network subnet.
     */
    readonly virtualNetworkSubnetId: pulumi.Input<string>;
}
