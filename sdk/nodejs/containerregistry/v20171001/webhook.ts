// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * An object that represents a webhook for a container registry.
 *
 * ## Example Usage
 * ### WebhookCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const webhook = new azurerm.containerregistry.v20171001.Webhook("webhook", {
 *     actions: ["push"],
 *     customHeaders: {
 *         Authorization: "Basic 000000000000000000000000000000000000000000000000000",
 *     },
 *     location: "westus",
 *     registryName: "myRegistry",
 *     resourceGroupName: "myResourceGroup",
 *     scope: "myRepository",
 *     serviceUri: "http://myservice.com",
 *     status: "enabled",
 *     tags: {
 *         key: "value",
 *     },
 *     webhookName: "myWebhook",
 * });
 *
 * ```
 */
export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Webhook {
        return new Webhook(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerregistry/v20171001:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    /**
     * The list of actions that trigger the webhook to post notifications.
     */
    public readonly actions!: pulumi.Output<string[]>;
    /**
     * The location of the resource. This cannot be changed after the resource is created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the webhook at the time the operation was called.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The status of the webhook at the time the operation was called.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.actions === undefined) {
                throw new Error("Missing required property 'actions'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.registryName === undefined) {
                throw new Error("Missing required property 'registryName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceUri === undefined) {
                throw new Error("Missing required property 'serviceUri'");
            }
            if (!args || args.webhookName === undefined) {
                throw new Error("Missing required property 'webhookName'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["customHeaders"] = args ? args.customHeaders : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["serviceUri"] = args ? args.serviceUri : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["webhookName"] = args ? args.webhookName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["actions"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["scope"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:containerregistry/latest:Webhook" }, { type: "azurerm:containerregistry/v20170601preview:Webhook" }, { type: "azurerm:containerregistry/v20190501:Webhook" }, { type: "azurerm:containerregistry/v20191201preview:Webhook" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Webhook.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    /**
     * The list of actions that trigger the webhook to post notifications.
     */
    readonly actions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom headers that will be added to the webhook notifications.
     */
    readonly customHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of the webhook. This cannot be changed after the resource is created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the container registry.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * The service URI for the webhook to post notifications.
     */
    readonly serviceUri: pulumi.Input<string>;
    /**
     * The status of the webhook at the time the operation was called.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The tags for the webhook.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the webhook.
     */
    readonly webhookName: pulumi.Input<string>;
}
