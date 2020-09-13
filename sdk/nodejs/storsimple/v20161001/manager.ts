// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The StorSimple Manager
 */
export class Manager extends pulumi.CustomResource {
    /**
     * Get an existing Manager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Manager {
        return new Manager(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storsimple/v20161001:Manager';

    /**
     * Returns true if the given object is an instance of Manager.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Manager {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Manager.__pulumiType;
    }

    /**
     * Specifies if the Manager is Garda or Helsinki
     */
    public readonly cisIntrinsicSettings!: pulumi.Output<outputs.storsimple.v20161001.ManagerIntrinsicSettingsResponse | undefined>;
    /**
     * ETag of the Manager
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The Geo location of the Manager
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The Resource Name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Specifies the Sku
     */
    public readonly sku!: pulumi.Output<outputs.storsimple.v20161001.ManagerSkuResponse | undefined>;
    /**
     * Tags attached to the Manager
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Manager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["cisIntrinsicSettings"] = args ? args.cisIntrinsicSettings : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["cisIntrinsicSettings"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storsimple/latest:Manager" }, { type: "azurerm:storsimple/v20170601:Manager" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Manager.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Manager resource.
 */
export interface ManagerArgs {
    /**
     * Specifies if the Manager is Garda or Helsinki
     */
    readonly cisIntrinsicSettings?: pulumi.Input<inputs.storsimple.v20161001.ManagerIntrinsicSettings>;
    /**
     * ETag of the Manager
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The Geo location of the Manager
     */
    readonly location: pulumi.Input<string>;
    /**
     * The manager name
     */
    readonly managerName: pulumi.Input<string>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the Sku
     */
    readonly sku?: pulumi.Input<inputs.storsimple.v20161001.ManagerSku>;
    /**
     * Tags attached to the Manager
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
