// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Public certificate object
 */
export class WebAppPublicCertificateSlot extends pulumi.CustomResource {
    /**
     * Get an existing WebAppPublicCertificateSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppPublicCertificateSlot {
        return new WebAppPublicCertificateSlot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/latest:WebAppPublicCertificateSlot';

    /**
     * Returns true if the given object is an instance of WebAppPublicCertificateSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppPublicCertificateSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppPublicCertificateSlot.__pulumiType;
    }

    /**
     * Public Certificate byte array
     */
    public readonly blob!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Public Certificate Location
     */
    public readonly publicCertificateLocation!: pulumi.Output<string | undefined>;
    /**
     * Certificate Thumbprint
     */
    public /*out*/ readonly thumbprint!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebAppPublicCertificateSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppPublicCertificateSlotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.publicCertificateName === undefined) {
                throw new Error("Missing required property 'publicCertificateName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.slot === undefined) {
                throw new Error("Missing required property 'slot'");
            }
            inputs["blob"] = args ? args.blob : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publicCertificateLocation"] = args ? args.publicCertificateLocation : undefined;
            inputs["publicCertificateName"] = args ? args.publicCertificateName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["slot"] = args ? args.slot : undefined;
            inputs["thumbprint"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["blob"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["publicCertificateLocation"] = undefined /*out*/;
            inputs["thumbprint"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:web/v20160801:WebAppPublicCertificateSlot" }, { type: "azurerm:web/v20180201:WebAppPublicCertificateSlot" }, { type: "azurerm:web/v20181101:WebAppPublicCertificateSlot" }, { type: "azurerm:web/v20190801:WebAppPublicCertificateSlot" }, { type: "azurerm:web/v20200601:WebAppPublicCertificateSlot" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WebAppPublicCertificateSlot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppPublicCertificateSlot resource.
 */
export interface WebAppPublicCertificateSlotArgs {
    /**
     * Public Certificate byte array
     */
    readonly blob?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the app.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Public Certificate Location
     */
    readonly publicCertificateLocation?: pulumi.Input<string>;
    /**
     * Public certificate name.
     */
    readonly publicCertificateName: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the deployment slot. If a slot is not specified, the API will create a binding for the production slot.
     */
    readonly slot: pulumi.Input<string>;
}
