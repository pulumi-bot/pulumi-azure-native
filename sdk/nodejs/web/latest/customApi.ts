// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A custom API
 *
 * ## Example Usage
 * ### Replace a custom API
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const customApi = new azurerm.web.latest.CustomApi("customApi", {
 *     apiName: "testCustomApi",
 *     resourceGroupName: "testResourceGroup",
 * });
 *
 * ```
 */
export class CustomApi extends pulumi.CustomResource {
    /**
     * Get an existing CustomApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomApi {
        return new CustomApi(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/latest:CustomApi';

    /**
     * Returns true if the given object is an instance of CustomApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomApi.__pulumiType;
    }

    /**
     * Resource ETag
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Custom API properties
     */
    public readonly properties!: pulumi.Output<outputs.web.latest.CustomApiPropertiesDefinitionResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CustomApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomApiArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CustomApiArgs | undefined;
            if (!args || args.apiName === undefined) {
                throw new Error("Missing required property 'apiName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiName"] = args ? args.apiName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:web/v20160601:CustomApi" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(CustomApi.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomApi resource.
 */
export interface CustomApiArgs {
    /**
     * API name
     */
    readonly apiName: pulumi.Input<string>;
    /**
     * Resource ETag
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Custom API properties
     */
    readonly properties?: pulumi.Input<inputs.web.latest.CustomApiPropertiesDefinition>;
    /**
     * The resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
