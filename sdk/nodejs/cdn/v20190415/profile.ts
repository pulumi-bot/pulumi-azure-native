// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * CDN profile is a logical grouping of endpoints that share the same settings, such as CDN provider and pricing tier.
 *
 * ## Example Usage
 * ### Profiles_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const profile = new azurerm.cdn.v20190415.Profile("profile", {
 *     location: "WestCentralUs",
 *     profileName: "profile1",
 *     resourceGroupName: "RG",
 *     sku: {
 *         name: "Standard_Verizon",
 *     },
 * });
 *
 * ```
 */
export class Profile extends pulumi.CustomResource {
    /**
     * Get an existing Profile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Profile {
        return new Profile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cdn/v20190415:Profile';

    /**
     * Returns true if the given object is an instance of Profile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profile.__pulumiType;
    }

    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning status of the profile.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource status of the profile.
     */
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
     */
    public readonly sku!: pulumi.Output<outputs.cdn.v20190415.SkuResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Profile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.profileName === undefined) {
                throw new Error("Missing required property 'profileName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:cdn/latest:Profile" }, { type: "azurerm:cdn/v20150601:Profile" }, { type: "azurerm:cdn/v20160402:Profile" }, { type: "azurerm:cdn/v20161002:Profile" }, { type: "azurerm:cdn/v20170402:Profile" }, { type: "azurerm:cdn/v20171012:Profile" }, { type: "azurerm:cdn/v20190615:Profile" }, { type: "azurerm:cdn/v20190615preview:Profile" }, { type: "azurerm:cdn/v20191231:Profile" }, { type: "azurerm:cdn/v20200331:Profile" }, { type: "azurerm:cdn/v20200415:Profile" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Profile.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Profile resource.
 */
export interface ProfileArgs {
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the CDN profile which is unique within the resource group.
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
     */
    readonly sku: pulumi.Input<inputs.cdn.v20190415.Sku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
