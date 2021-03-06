// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Friendly Rules name mapping to the any Rules or secret related information.
 * Latest API Version: 2020-09-01.
 *
 * @deprecated The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:cdn:Rule'.
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Rule {
        pulumi.log.warn("Rule is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:cdn:Rule'.")
        return new Rule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:cdn/latest:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * A list of actions that are executed when all the conditions of a rule are satisfied.
     */
    public readonly actions!: pulumi.Output<outputs.cdn.latest.DeliveryRuleCacheExpirationActionResponse | outputs.cdn.latest.DeliveryRuleCacheKeyQueryStringActionResponse | outputs.cdn.latest.DeliveryRuleRequestHeaderActionResponse | outputs.cdn.latest.DeliveryRuleResponseHeaderActionResponse | outputs.cdn.latest.OriginGroupOverrideActionResponse | outputs.cdn.latest.UrlRedirectActionResponse | outputs.cdn.latest.UrlRewriteActionResponse | outputs.cdn.latest.UrlSigningActionResponse[]>;
    /**
     * A list of conditions that must be matched for the actions to be executed
     */
    public readonly conditions!: pulumi.Output<outputs.cdn.latest.DeliveryRuleCookiesConditionResponse | outputs.cdn.latest.DeliveryRuleHttpVersionConditionResponse | outputs.cdn.latest.DeliveryRuleIsDeviceConditionResponse | outputs.cdn.latest.DeliveryRulePostArgsConditionResponse | outputs.cdn.latest.DeliveryRuleQueryStringConditionResponse | outputs.cdn.latest.DeliveryRuleRemoteAddressConditionResponse | outputs.cdn.latest.DeliveryRuleRequestBodyConditionResponse | outputs.cdn.latest.DeliveryRuleRequestHeaderConditionResponse | outputs.cdn.latest.DeliveryRuleRequestMethodConditionResponse | outputs.cdn.latest.DeliveryRuleRequestSchemeConditionResponse | outputs.cdn.latest.DeliveryRuleRequestUriConditionResponse | outputs.cdn.latest.DeliveryRuleUrlFileExtensionConditionResponse | outputs.cdn.latest.DeliveryRuleUrlFileNameConditionResponse | outputs.cdn.latest.DeliveryRuleUrlPathConditionResponse[] | undefined>;
    public /*out*/ readonly deploymentStatus!: pulumi.Output<string>;
    /**
     * If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
     */
    public readonly matchProcessingBehavior!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * Provisioning status
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Read only system data
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.cdn.latest.SystemDataResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:cdn:Rule'. */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Rule is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:cdn:Rule'.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.order === undefined) && !opts.urn) {
                throw new Error("Missing required property 'order'");
            }
            if ((!args || args.profileName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'profileName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.ruleSetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleSetName'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["conditions"] = args ? args.conditions : undefined;
            inputs["matchProcessingBehavior"] = args ? args.matchProcessingBehavior : undefined;
            inputs["order"] = args ? args.order : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["ruleName"] = args ? args.ruleName : undefined;
            inputs["ruleSetName"] = args ? args.ruleSetName : undefined;
            inputs["deploymentStatus"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["actions"] = undefined /*out*/;
            inputs["conditions"] = undefined /*out*/;
            inputs["deploymentStatus"] = undefined /*out*/;
            inputs["matchProcessingBehavior"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["order"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:cdn/latest:Rule" }, { type: "azure-native:cdn:Rule" }, { type: "azure-nextgen:cdn:Rule" }, { type: "azure-native:cdn/v20200901:Rule" }, { type: "azure-nextgen:cdn/v20200901:Rule" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Rule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * A list of actions that are executed when all the conditions of a rule are satisfied.
     */
    readonly actions: pulumi.Input<pulumi.Input<inputs.cdn.latest.DeliveryRuleCacheExpirationAction | inputs.cdn.latest.DeliveryRuleCacheKeyQueryStringAction | inputs.cdn.latest.DeliveryRuleRequestHeaderAction | inputs.cdn.latest.DeliveryRuleResponseHeaderAction | inputs.cdn.latest.OriginGroupOverrideAction | inputs.cdn.latest.UrlRedirectAction | inputs.cdn.latest.UrlRewriteAction | inputs.cdn.latest.UrlSigningAction>[]>;
    /**
     * A list of conditions that must be matched for the actions to be executed
     */
    readonly conditions?: pulumi.Input<pulumi.Input<inputs.cdn.latest.DeliveryRuleCookiesCondition | inputs.cdn.latest.DeliveryRuleHttpVersionCondition | inputs.cdn.latest.DeliveryRuleIsDeviceCondition | inputs.cdn.latest.DeliveryRulePostArgsCondition | inputs.cdn.latest.DeliveryRuleQueryStringCondition | inputs.cdn.latest.DeliveryRuleRemoteAddressCondition | inputs.cdn.latest.DeliveryRuleRequestBodyCondition | inputs.cdn.latest.DeliveryRuleRequestHeaderCondition | inputs.cdn.latest.DeliveryRuleRequestMethodCondition | inputs.cdn.latest.DeliveryRuleRequestSchemeCondition | inputs.cdn.latest.DeliveryRuleRequestUriCondition | inputs.cdn.latest.DeliveryRuleUrlFileExtensionCondition | inputs.cdn.latest.DeliveryRuleUrlFileNameCondition | inputs.cdn.latest.DeliveryRuleUrlPathCondition>[]>;
    /**
     * If this rule is a match should the rules engine continue running the remaining rules or stop. If not present, defaults to Continue.
     */
    readonly matchProcessingBehavior?: pulumi.Input<string | enums.cdn.latest.MatchProcessingBehavior>;
    /**
     * The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.
     */
    readonly order: pulumi.Input<number>;
    /**
     * Name of the CDN profile which is unique within the resource group.
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the delivery rule which is unique within the endpoint.
     */
    readonly ruleName?: pulumi.Input<string>;
    /**
     * Name of the rule set under the profile.
     */
    readonly ruleSetName: pulumi.Input<string>;
}
