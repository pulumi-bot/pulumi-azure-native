// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Describes a Virtual Machine Scale Set.
 */
export class VirtualMachineScaleSet extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachineScaleSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachineScaleSet {
        return new VirtualMachineScaleSet(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20191201:VirtualMachineScaleSet';

    /**
     * Returns true if the given object is an instance of VirtualMachineScaleSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachineScaleSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachineScaleSet.__pulumiType;
    }

    /**
     * The identity of the virtual machine scale set, if configured.
     */
    public readonly identity!: pulumi.Output<outputs.compute.v20191201.VirtualMachineScaleSetIdentityResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    public readonly plan!: pulumi.Output<outputs.compute.v20191201.PlanResponse | undefined>;
    /**
     * Describes the properties of a Virtual Machine Scale Set.
     */
    public readonly properties!: pulumi.Output<outputs.compute.v20191201.VirtualMachineScaleSetPropertiesResponse>;
    /**
     * The virtual machine scale set sku.
     */
    public readonly sku!: pulumi.Output<outputs.compute.v20191201.SkuResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a VirtualMachineScaleSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineScaleSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualMachineScaleSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VirtualMachineScaleSetArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualMachineScaleSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualMachineScaleSet resource.
 */
export interface VirtualMachineScaleSetArgs {
    /**
     * The identity of the virtual machine scale set, if configured.
     */
    readonly identity?: pulumi.Input<inputs.compute.v20191201.VirtualMachineScaleSetIdentity>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the VM scale set to create or update.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    readonly plan?: pulumi.Input<inputs.compute.v20191201.Plan>;
    /**
     * Describes the properties of a Virtual Machine Scale Set.
     */
    readonly properties?: pulumi.Input<inputs.compute.v20191201.VirtualMachineScaleSetProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The virtual machine scale set sku.
     */
    readonly sku?: pulumi.Input<inputs.compute.v20191201.Sku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
