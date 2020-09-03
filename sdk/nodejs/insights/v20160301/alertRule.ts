// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The alert rule resource.
 *
 * ## Example Usage
 * ### Create or update an alert rule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const alertRule = new azurerm.insights.v20160301.AlertRule("alertRule", {
 *     actions: [],
 *     condition: {
 *         dataSource: {
 *             odataType: "Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource",
 *             resourceUri: "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest",
 *         },
 *         odataType: "Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition",
 *     },
 *     description: "Pura Vida",
 *     isEnabled: true,
 *     location: "West US",
 *     name: "chiricutin",
 *     resourceGroupName: "Rac46PostSwapRG",
 *     ruleName: "chiricutin",
 *     tags: {},
 * });
 *
 * ```
 */
export class AlertRule extends pulumi.CustomResource {
    /**
     * Get an existing AlertRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AlertRule {
        return new AlertRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/v20160301:AlertRule';

    /**
     * Returns true if the given object is an instance of AlertRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertRule.__pulumiType;
    }

    /**
     * the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
     */
    public readonly actions!: pulumi.Output<outputs.insights.v20160301.RuleActionResponse[] | undefined>;
    /**
     * the condition that results in the alert rule being activated.
     */
    public readonly condition!: pulumi.Output<outputs.insights.v20160301.RuleConditionResponse>;
    /**
     * the description of the alert rule that will be included in the alert email.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * the flag that indicates whether the alert rule is enabled.
     */
    public readonly isEnabled!: pulumi.Output<boolean>;
    /**
     * Last time the rule was updated in ISO8601 format.
     */
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AlertRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AlertRuleArgs | undefined;
            if (!args || args.condition === undefined) {
                throw new Error("Missing required property 'condition'");
            }
            if (!args || args.isEnabled === undefined) {
                throw new Error("Missing required property 'isEnabled'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.ruleName === undefined) {
                throw new Error("Missing required property 'ruleName'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["isEnabled"] = args ? args.isEnabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["ruleName"] = args ? args.ruleName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["lastUpdatedTime"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:insights/latest:AlertRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AlertRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AlertRule resource.
 */
export interface AlertRuleArgs {
    /**
     * the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
     */
    readonly actions?: pulumi.Input<pulumi.Input<inputs.insights.v20160301.RuleAction>[]>;
    /**
     * the condition that results in the alert rule being activated.
     */
    readonly condition: pulumi.Input<inputs.insights.v20160301.RuleCondition>;
    /**
     * the description of the alert rule that will be included in the alert email.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * the flag that indicates whether the alert rule is enabled.
     */
    readonly isEnabled: pulumi.Input<boolean>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * the name of the alert rule.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the rule.
     */
    readonly ruleName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
