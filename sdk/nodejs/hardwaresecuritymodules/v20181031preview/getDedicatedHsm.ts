// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getDedicatedHsm(args: GetDedicatedHsmArgs, opts?: pulumi.InvokeOptions): Promise<GetDedicatedHsmResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:hardwaresecuritymodules/v20181031preview:getDedicatedHsm", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDedicatedHsmArgs {
    /**
     * The name of the dedicated HSM.
     */
    readonly name: string;
    /**
     * The name of the Resource Group to which the dedicated hsm belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Resource information with extended details.
 */
export interface GetDedicatedHsmResult {
    /**
     * The supported Azure location where the dedicated HSM should be created.
     */
    readonly location: string;
    /**
     * The name of the dedicated HSM.
     */
    readonly name: string;
    /**
     * Specifies the network interfaces of the dedicated hsm.
     */
    readonly networkProfile?: outputs.hardwaresecuritymodules.v20181031preview.NetworkProfileResponse;
    /**
     * Provisioning state.
     */
    readonly provisioningState: string;
    /**
     * SKU details
     */
    readonly sku: outputs.hardwaresecuritymodules.v20181031preview.SkuResponse;
    /**
     * This field will be used when RP does not support Availability zones.
     */
    readonly stampId?: string;
    /**
     * Resource Status Message.
     */
    readonly statusMessage: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * The resource type of the dedicated HSM.
     */
    readonly type: string;
    /**
     * The Dedicated Hsm zones.
     */
    readonly zones?: string[];
}
