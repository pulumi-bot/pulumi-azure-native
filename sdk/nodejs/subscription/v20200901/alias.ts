// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Subscription Information with the alias.
 *
 * ## Example Usage
 * ### CreateAlias
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const alias = new azurerm.subscription.v20200901.Alias("alias", {aliasName: "aliasForNewSub"});
 *
 * ```
 */
export class Alias extends pulumi.CustomResource {
    /**
     * Get an existing Alias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Alias {
        return new Alias(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:subscription/v20200901:Alias';

    /**
     * Returns true if the given object is an instance of Alias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alias.__pulumiType;
    }

    /**
     * Alias ID.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Put Alias response properties.
     */
    public readonly properties!: pulumi.Output<outputs.subscription.v20200901.PutAliasResponsePropertiesResponse>;
    /**
     * Resource type, Microsoft.Subscription/aliases.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Alias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AliasArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AliasArgs | undefined;
            if (!args || args.aliasName === undefined) {
                throw new Error("Missing required property 'aliasName'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            inputs["aliasName"] = args ? args.aliasName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:subscription/latest:Alias" }, { type: "azurerm:subscription/v20191001preview:Alias" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Alias.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Alias resource.
 */
export interface AliasArgs {
    /**
     * Alias Name
     */
    readonly aliasName: pulumi.Input<string>;
    /**
     * Put alias request properties.
     */
    readonly properties: pulumi.Input<inputs.subscription.v20200901.PutAliasRequestProperties>;
}
