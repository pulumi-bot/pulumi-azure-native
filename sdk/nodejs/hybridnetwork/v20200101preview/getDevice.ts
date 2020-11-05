// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDevice(args: GetDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:hybridnetwork/v20200101preview:getDevice", {
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDeviceArgs {
    /**
     * The name of the device resource.
     */
    readonly deviceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Device resource.
 */
export interface GetDeviceResult {
    /**
     * The reference to the Azure stack edge device.
     */
    readonly azureStackEdge?: outputs.hybridnetwork.v20200101preview.SubResourceResponse;
    /**
     * The type of the device.
     */
    readonly deviceType: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The list of network functions deployed on the device.
     */
    readonly networkFunctions: outputs.hybridnetwork.v20200101preview.SubResourceResponse[];
    /**
     * The provisioning state of the device resource.
     */
    readonly provisioningState: string;
    /**
     * The current device status.
     */
    readonly status: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
