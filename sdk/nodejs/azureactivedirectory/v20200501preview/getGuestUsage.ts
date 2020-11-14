// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getGuestUsage(args: GetGuestUsageArgs, opts?: pulumi.InvokeOptions): Promise<GetGuestUsageResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:azureactivedirectory/v20200501preview:getGuestUsage", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetGuestUsageArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The initial domain name of the AAD tenant.
     */
    readonly resourceName: string;
}

/**
 * Guest Usages Resource
 */
export interface GetGuestUsageResult {
    /**
     * Location of the Guest Usages resource.
     */
    readonly location?: string;
    /**
     * The name of the Guest Usages resource.
     */
    readonly name: string;
    /**
     * Key-value pairs of additional resource provisioning properties.
     */
    readonly tags?: {[key: string]: string};
    /**
     * An identifier for the tenant for which the resource is being created
     */
    readonly tenantId?: string;
    /**
     * The type of the Guest Usages resource.
     */
    readonly type: string;
}
