// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getIotDpsResource(args: GetIotDpsResourceArgs, opts?: pulumi.InvokeOptions): Promise<GetIotDpsResourceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:devices/v20200101:getIotDpsResource", {
        "provisioningServiceName": args.provisioningServiceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetIotDpsResourceArgs {
    /**
     * Name of the provisioning service to retrieve.
     */
    readonly provisioningServiceName: string;
    /**
     * Resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * The description of the provisioning service.
 */
export interface GetIotDpsResourceResult {
    /**
     * The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
     */
    readonly etag?: string;
    /**
     * The resource location.
     */
    readonly location: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * Service specific properties for a provisioning service
     */
    readonly properties: outputs.devices.v20200101.IotDpsPropertiesDescriptionResponse;
    /**
     * Sku info for a provisioning Service.
     */
    readonly sku: outputs.devices.v20200101.IotDpsSkuInfoResponse;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The resource type.
     */
    readonly type: string;
}
