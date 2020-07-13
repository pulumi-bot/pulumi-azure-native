// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Bot channel resource definition
 */
export class BotServiceChannel extends pulumi.CustomResource {
    /**
     * Get an existing BotServiceChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BotServiceChannel {
        return new BotServiceChannel(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:botservice:BotServiceChannel';

    /**
     * Returns true if the given object is an instance of BotServiceChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BotServiceChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BotServiceChannel.__pulumiType;
    }

    /**
     * Entity Tag
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Required. Gets or sets the Kind of the resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Specifies the location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The set of properties specific to bot channel resource
     */
    public readonly properties!: pulumi.Output<outputs.botservice.ChannelResponse>;
    /**
     * Gets or sets the SKU of the resource.
     */
    public readonly sku!: pulumi.Output<outputs.botservice.SkuResponse | undefined>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a BotServiceChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BotServiceChannelArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
        inputs["etag"] = args ? args.etag : undefined;
        inputs["kind"] = args ? args.kind : undefined;
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["resourceName"] = args ? args.resourceName : undefined;
        inputs["sku"] = args ? args.sku : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BotServiceChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BotServiceChannel resource.
 */
export interface BotServiceChannelArgs {
    /**
     * Entity Tag
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Required. Gets or sets the Kind of the resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Specifies the location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Channel resource.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The set of properties specific to bot channel resource
     */
    readonly properties?: pulumi.Input<inputs.botservice.Channel>;
    /**
     * The name of the Bot resource group in the user subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Bot resource.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * Gets or sets the SKU of the resource.
     */
    readonly sku?: pulumi.Input<inputs.botservice.Sku>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
