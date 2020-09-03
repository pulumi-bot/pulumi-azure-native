// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The Live Event.
 *
 * ## Example Usage
 * ### Create a LiveEvent
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const liveEvent = new azurerm.media.latest.LiveEvent("liveEvent", {
 *     accountName: "slitestmedia10",
 *     description: "test event 1",
 *     input: {
 *         keyFrameIntervalDuration: "PT6S",
 *         streamingProtocol: "RTMP",
 *     },
 *     liveEventName: "myLiveEvent1",
 *     location: "West US",
 *     preview: {
 *         accessControl: {
 *             ip: {
 *                 allow: [{
 *                     address: "0.0.0.0",
 *                     name: "AllowAll",
 *                     subnetPrefixLength: 0,
 *                 }],
 *             },
 *         },
 *     },
 *     resourceGroupName: "mediaresources",
 *     tags: {
 *         tag1: "value1",
 *         tag2: "value2",
 *     },
 * });
 *
 * ```
 */
export class LiveEvent extends pulumi.CustomResource {
    /**
     * Get an existing LiveEvent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LiveEvent {
        return new LiveEvent(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:media/latest:LiveEvent';

    /**
     * Returns true if the given object is an instance of LiveEvent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LiveEvent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LiveEvent.__pulumiType;
    }

    /**
     * The exact time the Live Event was created.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * The Live Event access policies.
     */
    public readonly crossSiteAccessPolicies!: pulumi.Output<outputs.media.latest.CrossSiteAccessPoliciesResponse | undefined>;
    /**
     * The Live Event description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Live Event encoding.
     */
    public readonly encoding!: pulumi.Output<outputs.media.latest.LiveEventEncodingResponse | undefined>;
    /**
     * The Live Event input.
     */
    public readonly input!: pulumi.Output<outputs.media.latest.LiveEventInputResponse>;
    /**
     * The exact time the Live Event was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Live Event preview.
     */
    public readonly preview!: pulumi.Output<outputs.media.latest.LiveEventPreviewResponse | undefined>;
    /**
     * The provisioning state of the Live Event.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource state of the Live Event.
     */
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
     */
    public readonly streamOptions!: pulumi.Output<string[] | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
     */
    public readonly useStaticHostname!: pulumi.Output<boolean | undefined>;

    /**
     * Create a LiveEvent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LiveEventArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LiveEventArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as LiveEventArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.input === undefined) {
                throw new Error("Missing required property 'input'");
            }
            if (!args || args.liveEventName === undefined) {
                throw new Error("Missing required property 'liveEventName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["autoStart"] = args ? args.autoStart : undefined;
            inputs["crossSiteAccessPolicies"] = args ? args.crossSiteAccessPolicies : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encoding"] = args ? args.encoding : undefined;
            inputs["input"] = args ? args.input : undefined;
            inputs["liveEventName"] = args ? args.liveEventName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["preview"] = args ? args.preview : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["streamOptions"] = args ? args.streamOptions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["useStaticHostname"] = args ? args.useStaticHostname : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:media/v20180330preview:LiveEvent" }, { type: "azurerm:media/v20180601preview:LiveEvent" }, { type: "azurerm:media/v20180701:LiveEvent" }, { type: "azurerm:media/v20190501preview:LiveEvent" }, { type: "azurerm:media/v20200501:LiveEvent" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LiveEvent.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LiveEvent resource.
 */
export interface LiveEventArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The flag indicates if the resource should be automatically started on creation.
     */
    readonly autoStart?: pulumi.Input<boolean>;
    /**
     * The Live Event access policies.
     */
    readonly crossSiteAccessPolicies?: pulumi.Input<inputs.media.latest.CrossSiteAccessPolicies>;
    /**
     * The Live Event description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Live Event encoding.
     */
    readonly encoding?: pulumi.Input<inputs.media.latest.LiveEventEncoding>;
    /**
     * The Live Event input.
     */
    readonly input: pulumi.Input<inputs.media.latest.LiveEventInput>;
    /**
     * The name of the Live Event.
     */
    readonly liveEventName: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The Live Event preview.
     */
    readonly preview?: pulumi.Input<inputs.media.latest.LiveEventPreview>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
     */
    readonly streamOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
     */
    readonly useStaticHostname?: pulumi.Input<boolean>;
}
