// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A notification.
 */
export class NotificationChannel extends pulumi.CustomResource {
    /**
     * Get an existing NotificationChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NotificationChannel {
        return new NotificationChannel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devtestlab/latest:NotificationChannel';

    /**
     * Returns true if the given object is an instance of NotificationChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationChannel.__pulumiType;
    }

    /**
     * The creation date of the notification channel.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * Description of notification.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
     */
    public readonly emailRecipient!: pulumi.Output<string | undefined>;
    /**
     * The list of event for which this notification is enabled.
     */
    public readonly events!: pulumi.Output<outputs.devtestlab.latest.EventResponse[] | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The locale to use when sending a notification (fallback for unsupported languages is EN).
     */
    public readonly notificationLocale!: pulumi.Output<string | undefined>;
    /**
     * The provisioning status of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    public /*out*/ readonly uniqueIdentifier!: pulumi.Output<string>;
    /**
     * The webhook URL to send notifications to.
     */
    public readonly webHookUrl!: pulumi.Output<string | undefined>;

    /**
     * Create a NotificationChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationChannelArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["emailRecipient"] = args ? args.emailRecipient : undefined;
            inputs["events"] = args ? args.events : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationLocale"] = args ? args.notificationLocale : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["webHookUrl"] = args ? args.webHookUrl : undefined;
            inputs["createdDate"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
        } else {
            inputs["createdDate"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["emailRecipient"] = undefined /*out*/;
            inputs["events"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notificationLocale"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
            inputs["webHookUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:devtestlab/v20160515:NotificationChannel" }, { type: "azurerm:devtestlab/v20180915:NotificationChannel" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NotificationChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NotificationChannel resource.
 */
export interface NotificationChannelArgs {
    /**
     * Description of notification.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
     */
    readonly emailRecipient?: pulumi.Input<string>;
    /**
     * The list of event for which this notification is enabled.
     */
    readonly events?: pulumi.Input<pulumi.Input<inputs.devtestlab.latest.Event>[]>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the notification channel.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The locale to use when sending a notification (fallback for unsupported languages is EN).
     */
    readonly notificationLocale?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The webhook URL to send notifications to.
     */
    readonly webHookUrl?: pulumi.Input<string>;
}
