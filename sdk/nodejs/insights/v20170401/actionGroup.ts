// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An action group resource.
 *
 * ## Example Usage
 * ### Create or update an action group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const actionGroup = new azurerm.insights.v20170401.ActionGroup("actionGroup", {
 *     actionGroupName: "SampleActionGroup",
 *     automationRunbookReceivers: [{
 *         automationAccountId: "/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest",
 *         isGlobalRunbook: false,
 *         name: "testRunbook",
 *         runbookName: "Sample runbook",
 *         serviceUri: `https://s13events.azure-automation.net/webhooks?token=iimE%2fD19Eg%2bvDy22yUMecIZY6Uiz%2bHfuQ67r8r1wY%2fI%3d`,
 *         webhookResourceId: "/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest/webhooks/Alert1510184037084",
 *     }],
 *     azureAppPushReceivers: [{
 *         emailAddress: "johndoe@email.com",
 *         name: "Sample azureAppPush",
 *     }],
 *     emailReceivers: [
 *         {
 *             emailAddress: "johndoe@email.com",
 *             name: "John Doe's email",
 *         },
 *         {
 *             emailAddress: "janesmith@email.com",
 *             name: "Jane Smith's email",
 *         },
 *     ],
 *     enabled: true,
 *     groupShortName: "sample",
 *     itsmReceivers: [{
 *         connectionId: "a3b9076c-ce8e-434e-85b4-aff10cb3c8f1",
 *         name: "Sample itsm",
 *         region: "westcentralus",
 *         ticketConfiguration: "{\"PayloadRevision\":0,\"WorkItemType\":\"Incident\",\"UseTemplate\":false,\"WorkItemData\":\"{}\",\"CreateOneWIPerCI\":false}",
 *         workspaceId: "5def922a-3ed4-49c1-b9fd-05ec533819a3|55dfd1f8-7e59-4f89-bf56-4c82f5ace23c",
 *     }],
 *     location: "Global",
 *     resourceGroupName: "Default-NotificationRules",
 *     smsReceivers: [
 *         {
 *             countryCode: "1",
 *             name: "John Doe's mobile",
 *             phoneNumber: "1234567890",
 *         },
 *         {
 *             countryCode: "1",
 *             name: "Jane Smith's mobile",
 *             phoneNumber: "0987654321",
 *         },
 *     ],
 *     tags: {},
 *     webhookReceivers: [{
 *         name: "Sample webhook",
 *         serviceUri: "http://www.example.com/webhook",
 *     }],
 * });
 *
 * ```
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
        return new ActionGroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/v20170401:ActionGroup';

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
     * The list of AutomationRunbook receivers that are part of this action group.
     */
    public readonly automationRunbookReceivers!: pulumi.Output<outputs.insights.v20170401.AutomationRunbookReceiverResponse[] | undefined>;
    /**
     * The list of AzureAppPush receivers that are part of this action group.
     */
    public readonly azureAppPushReceivers!: pulumi.Output<outputs.insights.v20170401.AzureAppPushReceiverResponse[] | undefined>;
    /**
     * The list of email receivers that are part of this action group.
     */
    public readonly emailReceivers!: pulumi.Output<outputs.insights.v20170401.EmailReceiverResponse[] | undefined>;
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
    public readonly itsmReceivers!: pulumi.Output<outputs.insights.v20170401.ItsmReceiverResponse[] | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The list of SMS receivers that are part of this action group.
     */
    public readonly smsReceivers!: pulumi.Output<outputs.insights.v20170401.SmsReceiverResponse[] | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The list of webhook receivers that are part of this action group.
     */
    public readonly webhookReceivers!: pulumi.Output<outputs.insights.v20170401.WebhookReceiverResponse[] | undefined>;

    /**
     * Create a ActionGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ActionGroupArgs | undefined;
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
            inputs["automationRunbookReceivers"] = args ? args.automationRunbookReceivers : undefined;
            inputs["azureAppPushReceivers"] = args ? args.azureAppPushReceivers : undefined;
            inputs["emailReceivers"] = args ? args.emailReceivers : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["groupShortName"] = args ? args.groupShortName : undefined;
            inputs["itsmReceivers"] = args ? args.itsmReceivers : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["smsReceivers"] = args ? args.smsReceivers : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["webhookReceivers"] = args ? args.webhookReceivers : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:insights/latest:ActionGroup" }, { type: "azurerm:insights/v20180301:ActionGroup" }, { type: "azurerm:insights/v20180901:ActionGroup" }, { type: "azurerm:insights/v20190301:ActionGroup" }, { type: "azurerm:insights/v20190601:ActionGroup" }] };
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
     * The list of AutomationRunbook receivers that are part of this action group.
     */
    readonly automationRunbookReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20170401.AutomationRunbookReceiver>[]>;
    /**
     * The list of AzureAppPush receivers that are part of this action group.
     */
    readonly azureAppPushReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20170401.AzureAppPushReceiver>[]>;
    /**
     * The list of email receivers that are part of this action group.
     */
    readonly emailReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20170401.EmailReceiver>[]>;
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
    readonly itsmReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20170401.ItsmReceiver>[]>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The list of SMS receivers that are part of this action group.
     */
    readonly smsReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20170401.SmsReceiver>[]>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of webhook receivers that are part of this action group.
     */
    readonly webhookReceivers?: pulumi.Input<pulumi.Input<inputs.insights.v20170401.WebhookReceiver>[]>;
}
