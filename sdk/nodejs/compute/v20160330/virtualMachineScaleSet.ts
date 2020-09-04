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
        return new VirtualMachineScaleSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20160330:VirtualMachineScaleSet';

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
    public readonly identity!: pulumi.Output<outputs.compute.v20160330.VirtualMachineScaleSetIdentityResponse | undefined>;
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
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The virtual machine scale set sku.
     */
    public readonly sku!: pulumi.Output<outputs.compute.v20160330.SkuResponse | undefined>;
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
    public readonly upgradePolicy!: pulumi.Output<outputs.compute.v20160330.UpgradePolicyResponse | undefined>;
    /**
     * The virtual machine profile.
     */
    public readonly virtualMachineProfile!: pulumi.Output<outputs.compute.v20160330.VirtualMachineScaleSetVMProfileResponse | undefined>;

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
            inputs["overProvision"] = args ? args.overProvision : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
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
            inputs["provisioningState"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20150615:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20160430preview:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20170330:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20171201:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20180401:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20180601:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20181001:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20190301:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20190701:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20191201:VirtualMachineScaleSet" }, { type: "azurerm:compute/v20200601:VirtualMachineScaleSet" }] };
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
    readonly identity?: pulumi.Input<inputs.compute.v20160330.VirtualMachineScaleSetIdentity>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the VM scale set to create or update.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies whether the Virtual Machine Scale Set should be overprovisioned.
     */
    readonly overProvision?: pulumi.Input<boolean>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The virtual machine scale set sku.
     */
    readonly sku?: pulumi.Input<inputs.compute.v20160330.Sku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The upgrade policy.
     */
    readonly upgradePolicy?: pulumi.Input<inputs.compute.v20160330.UpgradePolicy>;
    /**
     * The virtual machine profile.
     */
    readonly virtualMachineProfile?: pulumi.Input<inputs.compute.v20160330.VirtualMachineScaleSetVMProfile>;
}
