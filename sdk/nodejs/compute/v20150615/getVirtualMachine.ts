// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getVirtualMachine(args: GetVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:compute/v20150615:getVirtualMachine", {
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "vmName": args.vmName,
    }, opts);
}

export interface GetVirtualMachineArgs {
    /**
     * The expand expression to apply on the operation.
     */
    readonly expand?: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the virtual machine.
     */
    readonly vmName: string;
}

/**
 * Describes a Virtual Machine.
 */
export interface GetVirtualMachineResult {
    /**
     * Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
     */
    readonly availabilitySet?: outputs.compute.v20150615.SubResourceResponse;
    /**
     * Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
     */
    readonly diagnosticsProfile?: outputs.compute.v20150615.DiagnosticsProfileResponse;
    /**
     * Specifies the hardware settings for the virtual machine.
     */
    readonly hardwareProfile?: outputs.compute.v20150615.HardwareProfileResponse;
    /**
     * The virtual machine instance view.
     */
    readonly instanceView: outputs.compute.v20150615.VirtualMachineInstanceViewResponse;
    /**
     * Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
     */
    readonly licenseType?: string;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Specifies the network interfaces of the virtual machine.
     */
    readonly networkProfile?: outputs.compute.v20150615.NetworkProfileResponse;
    /**
     * Specifies the operating system settings for the virtual machine.
     */
    readonly osProfile?: outputs.compute.v20150615.OSProfileResponse;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    readonly plan?: outputs.compute.v20150615.PlanResponse;
    /**
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * The virtual machine child extension resources.
     */
    readonly resources: outputs.compute.v20150615.VirtualMachineExtensionResponse[];
    /**
     * Specifies the storage settings for the virtual machine disks.
     */
    readonly storageProfile?: outputs.compute.v20150615.StorageProfileResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
