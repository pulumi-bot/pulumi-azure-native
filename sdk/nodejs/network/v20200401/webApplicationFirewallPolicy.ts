// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Defines web application firewall policy.
 */
export class WebApplicationFirewallPolicy extends pulumi.CustomResource {
    /**
     * Get an existing WebApplicationFirewallPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebApplicationFirewallPolicy {
        return new WebApplicationFirewallPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200401:WebApplicationFirewallPolicy';

    /**
     * Returns true if the given object is an instance of WebApplicationFirewallPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebApplicationFirewallPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebApplicationFirewallPolicy.__pulumiType;
    }

    /**
     * A collection of references to application gateways.
     */
    public /*out*/ readonly applicationGateways!: pulumi.Output<outputs.network.v20200401.ApplicationGatewayResponse[]>;
    /**
     * The custom rules inside the policy.
     */
    public readonly customRules!: pulumi.Output<outputs.network.v20200401.WebApplicationFirewallCustomRuleResponse[] | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A collection of references to application gateway http listeners.
     */
    public /*out*/ readonly httpListeners!: pulumi.Output<outputs.network.v20200401.SubResourceResponse[]>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Describes the managedRules structure.
     */
    public readonly managedRules!: pulumi.Output<outputs.network.v20200401.ManagedRulesDefinitionResponse>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A collection of references to application gateway path rules.
     */
    public /*out*/ readonly pathBasedRules!: pulumi.Output<outputs.network.v20200401.SubResourceResponse[]>;
    /**
     * The PolicySettings for policy.
     */
    public readonly policySettings!: pulumi.Output<outputs.network.v20200401.PolicySettingsResponse | undefined>;
    /**
     * The provisioning state of the web application firewall policy resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource status of the policy.
     */
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebApplicationFirewallPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebApplicationFirewallPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.managedRules === undefined) {
                throw new Error("Missing required property 'managedRules'");
            }
            if (!args || args.policyName === undefined) {
                throw new Error("Missing required property 'policyName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["customRules"] = args ? args.customRules : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedRules"] = args ? args.managedRules : undefined;
            inputs["policyName"] = args ? args.policyName : undefined;
            inputs["policySettings"] = args ? args.policySettings : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["applicationGateways"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["httpListeners"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pathBasedRules"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["applicationGateways"] = undefined /*out*/;
            inputs["customRules"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["httpListeners"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedRules"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pathBasedRules"] = undefined /*out*/;
            inputs["policySettings"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20181201:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20190201:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20190401:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20190601:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20190701:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20190801:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20190901:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20191101:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20191201:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20200301:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20200501:WebApplicationFirewallPolicy" }, { type: "azurerm:network/v20200601:WebApplicationFirewallPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WebApplicationFirewallPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebApplicationFirewallPolicy resource.
 */
export interface WebApplicationFirewallPolicyArgs {
    /**
     * The custom rules inside the policy.
     */
    readonly customRules?: pulumi.Input<pulumi.Input<inputs.network.v20200401.WebApplicationFirewallCustomRule>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Describes the managedRules structure.
     */
    readonly managedRules: pulumi.Input<inputs.network.v20200401.ManagedRulesDefinition>;
    /**
     * The name of the policy.
     */
    readonly policyName: pulumi.Input<string>;
    /**
     * The PolicySettings for policy.
     */
    readonly policySettings?: pulumi.Input<inputs.network.v20200401.PolicySettings>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
