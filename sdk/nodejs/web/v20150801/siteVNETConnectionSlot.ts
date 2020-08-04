// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * VNETInfo contract. This contract is public and is a stripped down version of VNETInfoInternal
 */
export class SiteVNETConnectionSlot extends pulumi.CustomResource {
    /**
     * Get an existing SiteVNETConnectionSlot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SiteVNETConnectionSlot {
        return new SiteVNETConnectionSlot(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20150801:SiteVNETConnectionSlot';

    /**
     * Returns true if the given object is an instance of SiteVNETConnectionSlot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SiteVNETConnectionSlot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SiteVNETConnectionSlot.__pulumiType;
    }

    /**
     * Kind of resource
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly properties!: pulumi.Output<outputs.web.v20150801.VnetInfoResponseProperties>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a SiteVNETConnectionSlot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SiteVNETConnectionSlotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SiteVNETConnectionSlotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SiteVNETConnectionSlotArgs | undefined;
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
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["slot"] = args ? args.slot : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SiteVNETConnectionSlot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SiteVNETConnectionSlot resource.
 */
export interface SiteVNETConnectionSlotArgs {
    /**
     * Resource Id
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Kind of resource
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the Virtual Network
     */
    readonly name: pulumi.Input<string>;
    readonly properties?: pulumi.Input<inputs.web.v20150801.VnetInfoProperties>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the slot for this web app.
     */
    readonly slot: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource type
     */
    readonly type?: pulumi.Input<string>;
}
