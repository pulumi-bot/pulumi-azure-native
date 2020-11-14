// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getAppResourceUploadUrl(args: GetAppResourceUploadUrlArgs, opts?: pulumi.InvokeOptions): Promise<GetAppResourceUploadUrlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:appplatform/v20190501preview:getAppResourceUploadUrl", {
        "appName": args.appName,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetAppResourceUploadUrlArgs {
    /**
     * The name of the App resource.
     */
    readonly appName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Service resource.
     */
    readonly serviceName: string;
}

/**
 * Resource upload definition payload
 */
export interface GetAppResourceUploadUrlResult {
    /**
     * Source relative path
     */
    readonly relativePath?: string;
    /**
     * Upload URL
     */
    readonly uploadUrl?: string;
}
