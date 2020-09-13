// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An action group resource.
 */
export class ActionGroup extends pulumi.CustomResource {
    /**
     * Get an existing ActionGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ActionGroup {
        return new ActionGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/v20190601:ActionGroup';

    /**
     * Returns true if the given object is an instance of ActionGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionGroup.__pulumiType;
    }

    /**
     * The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
     */
    public readonly armRoleReceivers!: pulumi.Output<outputs.insights.v20190601.ArmRoleReceiverResponse[] | undefined>;
    /**
     * The list of AutomationRunbook receivers that are part of this action group.
     */
    public readonly automationRunbookReceivers!: pulumi.Output<outputs.insights.v20190601.AutomationRunbookReceiverResponse[] | undefined>;
    /**
     * The list of AzureAppPush receivers that are part of this action group.
     */
    public readonly azureAppPushReceivers!: pulumi.Output<outputs.insights.v20190601.AzureAppPushReceiverResponse[] | undefined>;
    /**
     * The list of azure function receivers that are part of this action group.
     */
    public readonly azureFunctionReceivers!: pulumi.Output<outputs.insights.v20190601.AzureFunctionReceiverResponse[] | undefined>;
    /**
     * The list of email receivers that are part of this action group.
     */
    public readonly emailReceivers!: pulumi.Output<outputs.insights.v20190601.EmailReceiverResponse[] | undefined>;
    /**
     * Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The short name of the action group. This will be used in SMS messages.
     */
    public readonly groupShortName!: pulumi.Output<string>;
    /**
     * The list of ITSM receivers that are part of this action group.
     */
    public readonly itsmReceivers!: pulumi.Output<outputs.insights.v20190601.ItsmReceiverResponse[] | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The list of logic app receivers that are part of this action group.
     */
    public readonly logicAppReceivers!: pulumi.Output<outputs.insights.v20190601.LogicAppReceiverResponse[] | undefined>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The list of SMS receivers that are part of this action group.
     */
    public readonly smsReceivers!: pulumi.Output<outputs.insights.v20190601.SmsReceiverResponse[] | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The list of voice receivers that are part of this action group.
     */
    public readonly voiceReceivers!: pulumi.Output<outputs.insights.v20190601.VoiceReceiverResponse[] | undefined>;
    /**
     * The list of webhook receivers that are part of this action group.
     */
    public readonly webhookReceivers!: pulumi.Output<outputs.insights.v20190601.WebhookReceiverResponse[] | undefined>;

    /**
     * Create a ActionGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.actionGroupName === undefined) {
                throw new Error("Missing required property 'actionGroupName'");
            }
            if (!args || args.enabled === undefined) {
                throw new Error("Missing required property 'enabled'");
            }
            if (!args || args.groupShortName === undefined) {
                throw new Error("Missing required property 'groupShortName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["actionGroupName"] = args ? args.actionGroupName : undefined;
            inputs["armRoleReceivers"] = args ? args.armRoleReceivers : undefined;
            inputs["automationRunbookReceivers"] = args ? args.automationRunbookReceivers : undefined;
            inputs["azureAppPushReceivers"] = args ? args.azureAppPushReceivers : undefined;
            inputs["azureFunctionReceivers"] = args ? args.azureFunctionReceivers : undefined;
            inputs["emailReceivers"] = args ? args.emailReceivers : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["groupShortName"] = args ? args.groupShortName : undefined;
            inputs["itsmReceivers"] = args ? args.itsmReceivers : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["logicAppReceivers"] = args ? args.logicAppReceivers : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["smsReceivers"] = args ? args.smsReceivers : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["voiceReceivers"] = args ? args.voiceReceivers : undefined;
            inputs["webhookReceivers"] = args ? args.webhookReceivers : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["armRoleReceivers"] = undefined /*out*/;
            inputs["automationRunbookReceivers"] = undefined /*out*/;
            inputs["azureAppPushReceivers"] = undefined /*out*/;
            inputs["azureFunctionReceivers"] = undefined /*out*/;
            inputs["emailReceivers"] = undefined /*out*/;
            inputs["enabled"] = undefined /*out*/;
            inputs["groupShortName"] = undefined /*out*/;
            inputs["itsmReceivers"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["logicAppReceivers"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["smsReceivers"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["voiceReceivers"] = undefined /*out*/;
            inputs["webhookReceivers"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:insights/latest:ActionGroup" }, { type: "azurerm:insights/v20170401:ActionGroup" }, { type: "azurerm:insights/v20180301:ActionGroup" }, { type: "azurerm:insights/v20180901:ActionGroup" }, { type: "azurerm:insights/v20190301:ActionGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ActionGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ActionGroup resource.
 */
export interface ActionGroupArgs {
    /**
     * The name of the action group.
     */
    readonly actionGroupName: pulumi.Input<string>;
    /**
     * The list of ARM role receivers that are part of this action group. Roles are Azure RBAC roles and only built-in roles are supported.
     */
    readonly armRoleReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.ArmRoleReceiver>[]>;
    /**
     * The list of AutomationRunbook receivers that are part of this action group.
     */
    readonly automationRunbookReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.AutomationRunbookReceiver>[]>;
    /**
     * The list of AzureAppPush receivers that are part of this action group.
     */
    readonly azureAppPushReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.AzureAppPushReceiver>[]>;
    /**
     * The list of azure function receivers that are part of this action group.
     */
    readonly azureFunctionReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.AzureFunctionReceiver>[]>;
    /**
     * The list of email receivers that are part of this action group.
     */
    readonly emailReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.EmailReceiver>[]>;
    /**
     * Indicates whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications.
     */
    readonly enabled: pulumi.Input<boolean>;
    /**
     * The short name of the action group. This will be used in SMS messages.
     */
    readonly groupShortName: pulumi.Input<string>;
    /**
     * The list of ITSM receivers that are part of this action group.
     */
    readonly itsmReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.ItsmReceiver>[]>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The list of logic app receivers that are part of this action group.
     */
    readonly logicAppReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.LogicAppReceiver>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The list of SMS receivers that are part of this action group.
     */
    readonly smsReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.SmsReceiver>[]>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of voice receivers that are part of this action group.
     */
    readonly voiceReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.VoiceReceiver>[]>;
    /**
     * The list of webhook receivers that are part of this action group.
     */
    readonly webhookReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20190601.WebhookReceiver>[]>;
}
