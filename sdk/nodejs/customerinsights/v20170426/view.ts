// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The view resource format.
 *
 * ## Example Usage
 * ### Views_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const view = new azurerm.customerinsights.v20170426.View("view", {
 *     definition: "{\\\"isProfileType\\\":false,\\\"profileTypes\\\":[],\\\"widgets\\\":[],\\\"style\\\":[]}",
 *     displayName: {
 *         en: "some name",
 *     },
 *     hubName: "sdkTestHub",
 *     resourceGroupName: "TestHubRG",
 *     userId: "testUser",
 *     viewName: "testView",
 * });
 *
 * ```
 */
export class View extends pulumi.CustomResource {
    /**
     * Get an existing View resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): View {
        return new View(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/v20170426:View';

    /**
     * Returns true if the given object is an instance of View.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is View {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === View.__pulumiType;
    }

    /**
     * Date time when view was last modified.
     */
    public /*out*/ readonly changed!: pulumi.Output<string>;
    /**
     * Date time when view was created.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * View definition.
     */
    public readonly definition!: pulumi.Output<string>;
    /**
     * Localized display name for the view.
     */
    public readonly displayName!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * the hub name.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * the user ID.
     */
    public readonly userId!: pulumi.Output<string | undefined>;
    /**
     * Name of the view.
     */
    public readonly viewName!: pulumi.Output<string>;

    /**
     * Create a View resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ViewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ViewArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ViewArgs | undefined;
            if (!args || args.definition === undefined) {
                throw new Error("Missing required property 'definition'");
            }
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.viewName === undefined) {
                throw new Error("Missing required property 'viewName'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["viewName"] = args ? args.viewName : undefined;
            inputs["changed"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:customerinsights/latest:View" }, { type: "azurerm:customerinsights/v20170101:View" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(View.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a View resource.
 */
export interface ViewArgs {
    /**
     * View definition.
     */
    readonly definition: pulumi.Input<string>;
    /**
     * Localized display name for the view.
     */
    readonly displayName?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the hub.
     */
    readonly hubName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * the user ID.
     */
    readonly userId?: pulumi.Input<string>;
    /**
     * The name of the view.
     */
    readonly viewName: pulumi.Input<string>;
}
