// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVendor(args: GetVendorArgs, opts?: pulumi.InvokeOptions): Promise<GetVendorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:hybridnetwork/v20200101preview:getVendor", {
        "vendorName": args.vendorName,
    }, opts);
}

export interface GetVendorArgs {
    /**
     * The name of vendor.
     */
    readonly vendorName: string;
}

/**
 * Vendor resource.
 */
export interface GetVendorResult {
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the vendor resource.
     */
    readonly provisioningState: string;
    /**
     * A list of ids of the vendor skus offered by the vendor.
     */
    readonly skus: outputs.hybridnetwork.v20200101preview.SubResourceResponse[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
