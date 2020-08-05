// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Premier add-on.
 */
export class WebAppPremierAddOnSlot extends pulumi.CustomResource {
    /**
     * Get an existing WebAppPremierAddOnSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppPremierAddOnSlot {
        return new WebAppPremierAddOnSlot(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20181101:WebAppPremierAddOnSlot';

    /**
     * Returns true if the given object is an instance of WebAppPremierAddOnSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppPremierAddOnSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppPremierAddOnSlot.__pulumiType;
    }

    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * PremierAddOn resource specific properties
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.web.v20181101.PremierAddOnResponseProperties>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebAppPremierAddOnSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppPremierAddOnSlotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppPremierAddOnSlotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WebAppPremierAddOnSlotArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.slot === undefined) {
                throw new Error("Missing required property 'slot'");
            }
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["marketplaceOffer"] = args ? args.marketplaceOffer : undefined;
            inputs["marketplacePublisher"] = args ? args.marketplacePublisher : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["product"] = args ? args.product : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["slot"] = args ? args.slot : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vendor"] = args ? args.vendor : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WebAppPremierAddOnSlot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppPremierAddOnSlot resource.
 */
export interface WebAppPremierAddOnSlotArgs {
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Premier add on Marketplace offer.
     */
    readonly marketplaceOffer?: pulumi.Input<string>;
    /**
     * Premier add on Marketplace publisher.
     */
    readonly marketplacePublisher?: pulumi.Input<string>;
    /**
     * Add-on name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Premier add on Product.
     */
    readonly product?: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Premier add on SKU.
     */
    readonly sku?: pulumi.Input<string>;
    /**
     * Name of the deployment slot. If a slot is not specified, the API will update the named add-on for the production slot.
     */
    readonly slot: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Premier add on Vendor.
     */
    readonly vendor?: pulumi.Input<string>;
}
