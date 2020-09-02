// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getPrivateLinkHub(args: GetPrivateLinkHubArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateLinkHubResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:synapse/v20190601preview:getPrivateLinkHub", {
        "privateLinkHubName": args.privateLinkHubName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateLinkHubArgs {
    /**
     * The name of the privateLinkHub
     */
    readonly privateLinkHubName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * A privateLinkHub
 */
export interface GetPrivateLinkHubResult {
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * PrivateLinkHub provisioning state
     */
    readonly provisioningState: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
