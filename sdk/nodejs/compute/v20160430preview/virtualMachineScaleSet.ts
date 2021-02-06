// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
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
        return new VirtualMachineScaleSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:compute/v20160430preview:VirtualMachineScaleSet';

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
    public readonly identity!: pulumi.Output<outputs.compute.v20160430preview.VirtualMachineScaleSetIdentityResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies whether the Virtual Machine Scale Set should be overprovisioned.
     */
    public readonly overProvision!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    public readonly plan!: pulumi.Output<outputs.compute.v20160430preview.PlanResponse | undefined>;
    /**
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * When true this limits the scale set to a single placement group, of max size 100 virtual machines.
     */
    public readonly singlePlacementGroup!: pulumi.Output<boolean | undefined>;
    /**
     * The virtual machine scale set sku.
     */
    public readonly sku!: pulumi.Output<outputs.compute.v20160430preview.SkuResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The upgrade policy.
     */
    public readonly upgradePolicy!: pulumi.Output<outputs.compute.v20160430preview.UpgradePolicyResponse | undefined>;
    /**
     * The virtual machine profile.
     */
    public readonly virtualMachineProfile!: pulumi.Output<outputs.compute.v20160430preview.VirtualMachineScaleSetVMProfileResponse | undefined>;

    /**
     * Create a VirtualMachineScaleSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineScaleSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["overProvision"] = args ? args.overProvision : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["singlePlacementGroup"] = args ? args.singlePlacementGroup : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["upgradePolicy"] = args ? args.upgradePolicy : undefined;
            inputs["virtualMachineProfile"] = args ? args.virtualMachineProfile : undefined;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["overProvision"] = undefined /*out*/;
            inputs["plan"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["singlePlacementGroup"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["upgradePolicy"] = undefined /*out*/;
            inputs["virtualMachineProfile"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:compute/latest:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20150615:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20160330:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20170330:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20171201:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20180401:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20180601:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20181001:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20190301:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20190701:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20191201:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20200601:VirtualMachineScaleSet" }, { type: "azure-nextgen:compute/v20201201:VirtualMachineScaleSet" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
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
    readonly identity?: pulumi.Input<inputs.compute.v20160430preview.VirtualMachineScaleSetIdentity>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the VM scale set to create or update.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies whether the Virtual Machine Scale Set should be overprovisioned.
     */
    readonly overProvision?: pulumi.Input<boolean>;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    readonly plan?: pulumi.Input<inputs.compute.v20160430preview.Plan>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * When true this limits the scale set to a single placement group, of max size 100 virtual machines.
     */
    readonly singlePlacementGroup?: pulumi.Input<boolean>;
    /**
     * The virtual machine scale set sku.
     */
    readonly sku?: pulumi.Input<inputs.compute.v20160430preview.Sku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The upgrade policy.
     */
    readonly upgradePolicy?: pulumi.Input<inputs.compute.v20160430preview.UpgradePolicy>;
    /**
     * The virtual machine profile.
     */
    readonly virtualMachineProfile?: pulumi.Input<inputs.compute.v20160430preview.VirtualMachineScaleSetVMProfile>;
}
