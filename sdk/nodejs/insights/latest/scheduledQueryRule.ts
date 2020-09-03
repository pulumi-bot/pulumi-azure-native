// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The Log Search Rule resource.
 *
 * ## Example Usage
 * ### Create or Update rule - AlertingAction
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const scheduledQueryRule = new azurerm.insights.latest.ScheduledQueryRule("scheduledQueryRule", {
 *     action: {
 *         odataType: "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction",
 *     },
 *     description: "log alert description",
 *     enabled: "true",
 *     location: "eastus",
 *     resourceGroupName: "Rac46PostSwapRG",
 *     ruleName: "logalertfoo",
 *     schedule: {
 *         frequencyInMinutes: 15,
 *         timeWindowInMinutes: 15,
 *     },
 *     source: {
 *         dataSourceId: "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.OperationalInsights/workspaces/sampleWorkspace",
 *         query: "Heartbeat | summarize AggregatedValue = count() by bin(TimeGenerated, 5m)",
 *         queryType: "ResultCount",
 *     },
 *     tags: {},
 * });
 *
 * ```
 * ### Create or Update rule - AlertingAction with Cross-Resource
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const scheduledQueryRule = new azurerm.insights.latest.ScheduledQueryRule("scheduledQueryRule", {
 *     action: {
 *         odataType: "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.AlertingAction",
 *     },
 *     description: "Sample Cross Resource alert",
 *     enabled: "true",
 *     location: "eastus",
 *     resourceGroupName: "Rac46PostSwapRG",
 *     ruleName: "SampleCrossResourceAlert",
 *     schedule: {
 *         frequencyInMinutes: 60,
 *         timeWindowInMinutes: 60,
 *     },
 *     source: {
 *         authorizedResources: [
 *             "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.OperationalInsights/workspaces/sampleWorkspace",
 *             "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/components/sampleAI",
 *         ],
 *         dataSourceId: "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/microsoft.insights/components/sampleAI",
 *         query: "union requests, workspace(\"sampleWorkspace\").Update",
 *         queryType: "ResultCount",
 *     },
 *     tags: {},
 * });
 *
 * ```
 * ### Create or Update rule - LogToMetricAction
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const scheduledQueryRule = new azurerm.insights.latest.ScheduledQueryRule("scheduledQueryRule", {
 *     action: {
 *         odataType: "Microsoft.WindowsAzure.Management.Monitoring.Alerts.Models.Microsoft.AppInsights.Nexus.DataContracts.Resources.ScheduledQueryRules.LogToMetricAction",
 *     },
 *     description: "log to metric description",
 *     enabled: "true",
 *     location: "West Europe",
 *     resourceGroupName: "alertsweu",
 *     ruleName: "logtometricfoo",
 *     source: {
 *         dataSourceId: "/subscriptions/af52d502-a447-4bc6-8cb7-4780fbb00490/resourceGroups/alertsweu/providers/Microsoft.OperationalInsights/workspaces/alertsweu",
 *     },
 *     tags: {},
 * });
 *
 * ```
 */
export class ScheduledQueryRule extends pulumi.CustomResource {
    /**
     * Get an existing ScheduledQueryRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ScheduledQueryRule {
        return new ScheduledQueryRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/latest:ScheduledQueryRule';

    /**
     * Returns true if the given object is an instance of ScheduledQueryRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScheduledQueryRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScheduledQueryRule.__pulumiType;
    }

    /**
     * Action needs to be taken on rule execution.
     */
    public readonly action!: pulumi.Output<outputs.insights.latest.ActionResponse>;
    /**
     * The description of the Log Search rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The flag which indicates whether the Log Search rule is enabled. Value should be true or false
     */
    public readonly enabled!: pulumi.Output<string | undefined>;
    /**
     * Last time the rule was updated in IS08601 format.
     */
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the scheduled query rule
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Schedule (Frequency, Time Window) for rule. Required for action type - AlertingAction
     */
    public readonly schedule!: pulumi.Output<outputs.insights.latest.ScheduleResponse | undefined>;
    /**
     * Data Source against which rule will Query Data
     */
    public readonly source!: pulumi.Output<outputs.insights.latest.SourceResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ScheduledQueryRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduledQueryRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduledQueryRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ScheduledQueryRuleArgs | undefined;
            if (!args || args.action === undefined) {
                throw new Error("Missing required property 'action'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.ruleName === undefined) {
                throw new Error("Missing required property 'ruleName'");
            }
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["ruleName"] = args ? args.ruleName : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["lastUpdatedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:insights/v20180416:ScheduledQueryRule" }, { type: "azurerm:insights/v20200501preview:ScheduledQueryRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ScheduledQueryRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScheduledQueryRule resource.
 */
export interface ScheduledQueryRuleArgs {
    /**
     * Action needs to be taken on rule execution.
     */
    readonly action: pulumi.Input<inputs.insights.latest.Action>;
    /**
     * The description of the Log Search rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The flag which indicates whether the Log Search rule is enabled. Value should be true or false
     */
    readonly enabled?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the rule.
     */
    readonly ruleName: pulumi.Input<string>;
    /**
     * Schedule (Frequency, Time Window) for rule. Required for action type - AlertingAction
     */
    readonly schedule?: pulumi.Input<inputs.insights.latest.Schedule>;
    /**
     * Data Source against which rule will Query Data
     */
    readonly source: pulumi.Input<inputs.insights.latest.Source>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
