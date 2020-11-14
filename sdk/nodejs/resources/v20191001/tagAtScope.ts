// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Wrapper resource for tags API requests and responses.
 */
export class TagAtScope extends pulumi.CustomResource {
    /**
     * Get an existing TagAtScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TagAtScope {
        return new TagAtScope(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:resources/v20191001:TagAtScope';

    /**
     * Returns true if the given object is an instance of TagAtScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagAtScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagAtScope.__pulumiType;
    }

    /**
     * The name of the tags wrapper resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The set of tags.
     */
    public readonly properties!: pulumi.Output<outputs.resources.v20191001.TagsResponse>;
    /**
     * The type of the tags wrapper resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a TagAtScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagAtScopeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["properties"] = args ? args.properties : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:resources/latest:TagAtScope" }, { type: "azure-nextgen:resources/v20200601:TagAtScope" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(TagAtScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TagAtScope resource.
 */
export interface TagAtScopeArgs {
    /**
     * The set of tags.
     */
    readonly properties: pulumi.Input<inputs.resources.v20191001.Tags>;
    /**
     * The resource scope.
     */
    readonly scope: pulumi.Input<string>;
}
