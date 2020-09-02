// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Customer subscription which can use a sku.
 */
export class VendorSkuPreview extends pulumi.CustomResource {
    /**
     * Get an existing VendorSkuPreview resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VendorSkuPreview {
        return new VendorSkuPreview(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:hybridnetwork/v20200101preview:VendorSkuPreview';

    /**
     * Returns true if the given object is an instance of VendorSkuPreview.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VendorSkuPreview {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VendorSkuPreview.__pulumiType;
    }

    /**
     * Preview subscription id
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Type of the resource
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VendorSkuPreview resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VendorSkuPreviewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VendorSkuPreviewArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VendorSkuPreviewArgs | undefined;
            if (!args || args.previewSubscription === undefined) {
                throw new Error("Missing required property 'previewSubscription'");
            }
            if (!args || args.skuName === undefined) {
                throw new Error("Missing required property 'skuName'");
            }
            if (!args || args.vendorName === undefined) {
                throw new Error("Missing required property 'vendorName'");
            }
            inputs["previewSubscription"] = args ? args.previewSubscription : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["vendorName"] = args ? args.vendorName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VendorSkuPreview.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VendorSkuPreview resource.
 */
export interface VendorSkuPreviewArgs {
    /**
     * Preview subscription id.
     */
    readonly previewSubscription: pulumi.Input<string>;
    /**
     * The name of the vendor sku.
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * The name of the vendor.
     */
    readonly vendorName: pulumi.Input<string>;
}
