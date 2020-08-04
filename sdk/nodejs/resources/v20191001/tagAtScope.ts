// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
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
        return new TagAtScope(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:resources/v20191001:TagAtScope';

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
    public readonly name!: pulumi.Output<string>;
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
    constructor(name: string, args: TagAtScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagAtScopeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as TagAtScopeArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TagAtScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TagAtScope resource.
 */
export interface TagAtScopeArgs {
    /**
     * The resource scope.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The set of tags.
     */
    readonly properties: pulumi.Input<inputs.resources.v20191001.Tags>;
}
