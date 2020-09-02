// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The Media Graph.
 */
export class MediaGraph extends pulumi.CustomResource {
    /**
     * Get an existing MediaGraph resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MediaGraph {
        return new MediaGraph(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:media/v20190901preview:MediaGraph';

    /**
     * Returns true if the given object is an instance of MediaGraph.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MediaGraph {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MediaGraph.__pulumiType;
    }

    /**
     * Date the Media Graph was created
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Media Graph  description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Date the Media Graph was last modified
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Media Graph sinks
     */
    public readonly sinks!: pulumi.Output<outputs.media.v20190901preview.MediaGraphSinkResponse[]>;
    /**
     * Media Graph sources
     */
    public readonly sources!: pulumi.Output<outputs.media.v20190901preview.MediaGraphSourceResponse[]>;
    /**
     * Media Graph state
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a MediaGraph resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MediaGraphArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MediaGraphArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as MediaGraphArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.mediaGraphName === undefined) {
                throw new Error("Missing required property 'mediaGraphName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sinks === undefined) {
                throw new Error("Missing required property 'sinks'");
            }
            if (!args || args.sources === undefined) {
                throw new Error("Missing required property 'sources'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["mediaGraphName"] = args ? args.mediaGraphName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sinks"] = args ? args.sinks : undefined;
            inputs["sources"] = args ? args.sources : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:media/v20200201preview:MediaGraph" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(MediaGraph.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a MediaGraph resource.
 */
export interface MediaGraphArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Media Graph  description
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Media Graph name.
     */
    readonly mediaGraphName: pulumi.Input<string>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Media Graph sinks
     */
    readonly sinks: pulumi.Input<pulumi.Input<inputs.media.v20190901preview.MediaGraphSink>[]>;
    /**
     * Media Graph sources
     */
    readonly sources: pulumi.Input<pulumi.Input<inputs.media.v20190901preview.MediaGraphSource>[]>;
}
