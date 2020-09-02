// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The Live Output.
 */
export class LiveOutput extends pulumi.CustomResource {
    /**
     * Get an existing LiveOutput resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LiveOutput {
        return new LiveOutput(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:media/v20180601preview:LiveOutput';

    /**
     * Returns true if the given object is an instance of LiveOutput.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LiveOutput {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LiveOutput.__pulumiType;
    }

    /**
     * ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
     */
    public readonly archiveWindowLength!: pulumi.Output<string>;
    /**
     * The asset name.
     */
    public readonly assetName!: pulumi.Output<string>;
    /**
     * The exact time the Live Output was created.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * The description of the Live Output.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The HLS configuration.
     */
    public readonly hls!: pulumi.Output<outputs.media.v20180601preview.HlsResponse | undefined>;
    /**
     * The exact time the Live Output was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The manifest file name.
     */
    public readonly manifestName!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The output snapshot time.
     */
    public readonly outputSnapTime!: pulumi.Output<number | undefined>;
    /**
     * The provisioning state of the Live Output.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource state of the Live Output.
     */
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LiveOutput resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LiveOutputArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LiveOutputArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as LiveOutputArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.archiveWindowLength === undefined) {
                throw new Error("Missing required property 'archiveWindowLength'");
            }
            if (!args || args.assetName === undefined) {
                throw new Error("Missing required property 'assetName'");
            }
            if (!args || args.liveEventName === undefined) {
                throw new Error("Missing required property 'liveEventName'");
            }
            if (!args || args.liveOutputName === undefined) {
                throw new Error("Missing required property 'liveOutputName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["archiveWindowLength"] = args ? args.archiveWindowLength : undefined;
            inputs["assetName"] = args ? args.assetName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["hls"] = args ? args.hls : undefined;
            inputs["liveEventName"] = args ? args.liveEventName : undefined;
            inputs["liveOutputName"] = args ? args.liveOutputName : undefined;
            inputs["manifestName"] = args ? args.manifestName : undefined;
            inputs["outputSnapTime"] = args ? args.outputSnapTime : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azurerm:media/latest:LiveOutput" }, { type: "azurerm:media/v20180330preview:LiveOutput" }, { type: "azurerm:media/v20180701:LiveOutput" }, { type: "azurerm:media/v20190501preview:LiveOutput" }, { type: "azurerm:media/v20200501:LiveOutput" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LiveOutput.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LiveOutput resource.
 */
export interface LiveOutputArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
     */
    readonly archiveWindowLength: pulumi.Input<string>;
    /**
     * The asset name.
     */
    readonly assetName: pulumi.Input<string>;
    /**
     * The description of the Live Output.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The HLS configuration.
     */
    readonly hls?: pulumi.Input<inputs.media.v20180601preview.Hls>;
    /**
     * The name of the Live Event.
     */
    readonly liveEventName: pulumi.Input<string>;
    /**
     * The name of the Live Output.
     */
    readonly liveOutputName: pulumi.Input<string>;
    /**
     * The manifest file name.
     */
    readonly manifestName?: pulumi.Input<string>;
    /**
     * The output snapshot time.
     */
    readonly outputSnapTime?: pulumi.Input<number>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
