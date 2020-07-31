// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getProfile(args: GetProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetProfileResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:customerinsights/v20170426:getProfile", {
        "hubName": args.hubName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProfileArgs {
    /**
     * The name of the hub.
     */
    readonly hubName: string;
    /**
     * The name of the profile.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The profile resource format.
 */
export interface GetProfileResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The profile type definition.
     */
    readonly properties: outputs.customerinsights.v20170426.ProfileTypeDefinitionResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}