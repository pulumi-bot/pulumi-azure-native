// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Defines web application firewall policy for Azure CDN.
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cdn/v20200415:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Describes custom rules inside the policy.
     */
    public readonly customRules!: pulumi.Output<outputs.cdn.v20200415.CustomRuleListResponse | undefined>;
    /**
     * Describes Azure CDN endpoints associated with this Web Application Firewall policy.
     */
    public /*out*/ readonly endpointLinks!: pulumi.Output<outputs.cdn.v20200415.CdnEndpointResponse[]>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Describes managed rules inside the policy.
     */
    public readonly managedRules!: pulumi.Output<outputs.cdn.v20200415.ManagedRuleSetListResponse | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Describes  policySettings for policy
     */
    public readonly policySettings!: pulumi.Output<outputs.cdn.v20200415.PolicySettingsResponse | undefined>;
    /**
     * Provisioning state of the WebApplicationFirewallPolicy.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Describes rate limit rules inside the policy.
     */
    public readonly rateLimitRules!: pulumi.Output<outputs.cdn.v20200415.RateLimitRuleListResponse | undefined>;
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
     */
    public readonly sku!: pulumi.Output<outputs.cdn.v20200415.SkuResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PolicyArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.policyName === undefined) {
                throw new Error("Missing required property 'policyName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["customRules"] = args ? args.customRules : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedRules"] = args ? args.managedRules : undefined;
            inputs["policyName"] = args ? args.policyName : undefined;
            inputs["policySettings"] = args ? args.policySettings : undefined;
            inputs["rateLimitRules"] = args ? args.rateLimitRules : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["endpointLinks"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:cdn/latest:Policy" }, { type: "azurerm:cdn/v20190615:Policy" }, { type: "azurerm:cdn/v20190615preview:Policy" }, { type: "azurerm:cdn/v20200331:Policy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Describes custom rules inside the policy.
     */
    readonly customRules?: pulumi.Input<inputs.cdn.v20200415.CustomRuleList>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Describes managed rules inside the policy.
     */
    readonly managedRules?: pulumi.Input<inputs.cdn.v20200415.ManagedRuleSetList>;
    /**
     * The name of the CdnWebApplicationFirewallPolicy.
     */
    readonly policyName: pulumi.Input<string>;
    /**
     * Describes  policySettings for policy
     */
    readonly policySettings?: pulumi.Input<inputs.cdn.v20200415.PolicySettings>;
    /**
     * Describes rate limit rules inside the policy.
     */
    readonly rateLimitRules?: pulumi.Input<inputs.cdn.v20200415.RateLimitRuleList>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
     */
    readonly sku: pulumi.Input<inputs.cdn.v20200415.Sku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
