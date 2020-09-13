// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A Streaming Locator resource
 */
export class StreamingLocator extends pulumi.CustomResource {
    /**
     * Get an existing StreamingLocator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StreamingLocator {
        return new StreamingLocator(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:media/v20180701:StreamingLocator';

    /**
     * Returns true if the given object is an instance of StreamingLocator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamingLocator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamingLocator.__pulumiType;
    }

    /**
     * Alternative Media ID of this Streaming Locator
     */
    public readonly alternativeMediaId!: pulumi.Output<string | undefined>;
    /**
     * Asset Name
     */
    public readonly assetName!: pulumi.Output<string>;
    /**
     * The ContentKeys used by this Streaming Locator.
     */
    public readonly contentKeys!: pulumi.Output<outputs.media.v20180701.StreamingLocatorContentKeyResponse[] | undefined>;
    /**
     * The creation time of the Streaming Locator.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Name of the default ContentKeyPolicy used by this Streaming Locator.
     */
    public readonly defaultContentKeyPolicyName!: pulumi.Output<string | undefined>;
    /**
     * The end time of the Streaming Locator.
     */
    public readonly endTime!: pulumi.Output<string | undefined>;
    /**
     * A list of asset or account filters which apply to this streaming locator
     */
    public readonly filters!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The start time of the Streaming Locator.
     */
    public readonly startTime!: pulumi.Output<string | undefined>;
    /**
     * The StreamingLocatorId of the Streaming Locator.
     */
    public readonly streamingLocatorId!: pulumi.Output<string | undefined>;
    /**
     * Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
     */
    public readonly streamingPolicyName!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StreamingLocator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamingLocatorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.assetName === undefined) {
                throw new Error("Missing required property 'assetName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.streamingLocatorName === undefined) {
                throw new Error("Missing required property 'streamingLocatorName'");
            }
            if (!args || args.streamingPolicyName === undefined) {
                throw new Error("Missing required property 'streamingPolicyName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["alternativeMediaId"] = args ? args.alternativeMediaId : undefined;
            inputs["assetName"] = args ? args.assetName : undefined;
            inputs["contentKeys"] = args ? args.contentKeys : undefined;
            inputs["defaultContentKeyPolicyName"] = args ? args.defaultContentKeyPolicyName : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["filters"] = args ? args.filters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["streamingLocatorId"] = args ? args.streamingLocatorId : undefined;
            inputs["streamingLocatorName"] = args ? args.streamingLocatorName : undefined;
            inputs["streamingPolicyName"] = args ? args.streamingPolicyName : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["alternativeMediaId"] = undefined /*out*/;
            inputs["assetName"] = undefined /*out*/;
            inputs["contentKeys"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
            inputs["defaultContentKeyPolicyName"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["filters"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["streamingLocatorId"] = undefined /*out*/;
            inputs["streamingPolicyName"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:media/latest:StreamingLocator" }, { type: "azurerm:media/v20180330preview:StreamingLocator" }, { type: "azurerm:media/v20180601preview:StreamingLocator" }, { type: "azurerm:media/v20200501:StreamingLocator" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(StreamingLocator.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StreamingLocator resource.
 */
export interface StreamingLocatorArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Alternative Media ID of this Streaming Locator
     */
    readonly alternativeMediaId?: pulumi.Input<string>;
    /**
     * Asset Name
     */
    readonly assetName: pulumi.Input<string>;
    /**
     * The ContentKeys used by this Streaming Locator.
     */
    readonly contentKeys?: pulumi.Input<pulumi.Input<inputs.media.v20180701.StreamingLocatorContentKey>[]>;
    /**
     * Name of the default ContentKeyPolicy used by this Streaming Locator.
     */
    readonly defaultContentKeyPolicyName?: pulumi.Input<string>;
    /**
     * The end time of the Streaming Locator.
     */
    readonly endTime?: pulumi.Input<string>;
    /**
     * A list of asset or account filters which apply to this streaming locator
     */
    readonly filters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The start time of the Streaming Locator.
     */
    readonly startTime?: pulumi.Input<string>;
    /**
     * The StreamingLocatorId of the Streaming Locator.
     */
    readonly streamingLocatorId?: pulumi.Input<string>;
    /**
     * The Streaming Locator name.
     */
    readonly streamingLocatorName: pulumi.Input<string>;
    /**
     * Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
     */
    readonly streamingPolicyName: pulumi.Input<string>;
}
