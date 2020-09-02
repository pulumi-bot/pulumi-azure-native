// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Action for alert rule.
 *
 * ## Example Usage
 * ### Creates or updates an action of alert rule.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const alertRuleAction = new azurerm.operationalinsights.latest.AlertRuleAction("alertRuleAction", {
 *     actionId: "912bec42-cb66-4c03-ac63-1761b6898c3e",
 *     etag: "\"0300bf09-0000-0000-0000-5c37296e0000\"",
 *     logicAppResourceId: "/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.Logic/workflows/MyAlerts",
 *     resourceGroupName: "myRg",
 *     ruleId: "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
 *     triggerUri: `https://prod-31.northcentralus.logic.azure.com:443/workflows/cd3765391efd48549fd7681ded1d48d7/triggers/manual/paths/invoke?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=signature`,
 *     workspaceName: "myWorkspace",
 * });
 *
 * ```
 */
export class AlertRuleAction extends pulumi.CustomResource {
    /**
     * Get an existing AlertRuleAction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AlertRuleAction {
        return new AlertRuleAction(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:operationalinsights/latest:AlertRuleAction';

    /**
     * Returns true if the given object is an instance of AlertRuleAction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertRuleAction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertRuleAction.__pulumiType;
    }

    /**
     * Etag of the action.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
     */
    public readonly logicAppResourceId!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The name of the logic app's workflow.
     */
    public /*out*/ readonly workflowId!: pulumi.Output<string | undefined>;

    /**
     * Create a AlertRuleAction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertRuleActionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertRuleActionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AlertRuleActionArgs | undefined;
            if (!args || args.actionId === undefined) {
                throw new Error("Missing required property 'actionId'");
            }
            if (!args || args.logicAppResourceId === undefined) {
                throw new Error("Missing required property 'logicAppResourceId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.ruleId === undefined) {
                throw new Error("Missing required property 'ruleId'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["actionId"] = args ? args.actionId : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["logicAppResourceId"] = args ? args.logicAppResourceId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["ruleId"] = args ? args.ruleId : undefined;
            inputs["triggerUri"] = args ? args.triggerUri : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["workflowId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:operationalinsights/v20200101:AlertRuleAction" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AlertRuleAction.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AlertRuleAction resource.
 */
export interface AlertRuleActionArgs {
    /**
     * Action ID
     */
    readonly actionId: pulumi.Input<string>;
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
     */
    readonly logicAppResourceId: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Alert rule ID
     */
    readonly ruleId: pulumi.Input<string>;
    /**
     * Logic App Callback URL for this specific workflow.
     */
    readonly triggerUri?: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
