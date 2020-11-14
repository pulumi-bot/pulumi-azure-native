// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getVirtualMachineScaleSet(args: GetVirtualMachineScaleSetArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineScaleSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:compute/v20190701:getVirtualMachineScaleSet", {
        "resourceGroupName": args.resourceGroupName,
        "vmScaleSetName": args.vmScaleSetName,
    }, opts);
}

export interface GetVirtualMachineScaleSetArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the VM scale set.
     */
    readonly vmScaleSetName: string;
}

/**
 * Describes a Virtual Machine Scale Set.
 */
export interface GetVirtualMachineScaleSetResult {
    /**
     * Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
     */
    readonly additionalCapabilities?: outputs.compute.v20190701.AdditionalCapabilitiesResponse;
    /**
     * Policy for automatic repairs.
     */
    readonly automaticRepairsPolicy?: outputs.compute.v20190701.AutomaticRepairsPolicyResponse;
    /**
     * When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
     */
    readonly doNotRunExtensionsOnOverprovisionedVMs?: boolean;
    /**
     * The identity of the virtual machine scale set, if configured.
     */
    readonly identity?: outputs.compute.v20190701.VirtualMachineScaleSetIdentityResponse;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Specifies whether the Virtual Machine Scale Set should be overprovisioned.
     */
    readonly overprovision?: boolean;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    readonly plan?: outputs.compute.v20190701.PlanResponse;
    /**
     * Fault Domain count for each placement group.
     */
    readonly platformFaultDomainCount?: number;
    /**
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
     */
    readonly proximityPlacementGroup?: outputs.compute.v20190701.SubResourceResponse;
    /**
     * Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
     */
    readonly scaleInPolicy?: outputs.compute.v20190701.ScaleInPolicyResponse;
    /**
     * When true this limits the scale set to a single placement group, of max size 100 virtual machines.
     */
    readonly singlePlacementGroup?: boolean;
    /**
     * The virtual machine scale set sku.
     */
    readonly sku?: outputs.compute.v20190701.SkuResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Specifies the ID which uniquely identifies a Virtual Machine Scale Set.
     */
    readonly uniqueId: string;
    /**
     * The upgrade policy.
     */
    readonly upgradePolicy?: outputs.compute.v20190701.UpgradePolicyResponse;
    /**
     * The virtual machine profile.
     */
    readonly virtualMachineProfile?: outputs.compute.v20190701.VirtualMachineScaleSetVMProfileResponse;
    /**
     * Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
     */
    readonly zoneBalance?: boolean;
    /**
     * The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
     */
    readonly zones?: string[];
}
