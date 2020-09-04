// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The IoT Central application.
 */
export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:iotcentral/v20170701privatepreview:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    /**
     * The ID of the application.
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
    /**
     * The display name of the application.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The ARM resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A valid instance SKU.
     */
    public readonly sku!: pulumi.Output<outputs.iotcentral.v20170701privatepreview.AppSkuInfoResponse>;
    /**
     * The subdomain of the application.
     */
    public readonly subdomain!: pulumi.Output<string | undefined>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
     */
    public readonly template!: pulumi.Output<string | undefined>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["subdomain"] = args ? args.subdomain : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["template"] = args ? args.template : undefined;
            inputs["applicationId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["applicationId"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["subdomain"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["template"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:iotcentral/latest:App" }, { type: "azurerm:iotcentral/v20180901:App" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(App.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * The display name of the application.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the IoT Central application.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The ARM resource name of the IoT Central application.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * A valid instance SKU.
     */
    readonly sku: pulumi.Input<inputs.iotcentral.v20170701privatepreview.AppSkuInfo>;
    /**
     * The subdomain of the application.
     */
    readonly subdomain?: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
     */
    readonly template?: pulumi.Input<string>;
}
